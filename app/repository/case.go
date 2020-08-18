package repository

import (
	"fmt"
	"github.com/matgomes/stolen-bike-challenge/app/model"
	"gopkg.in/guregu/null.v4"
)

const caseBaseQuery = `SELECT c.id, c.owner, c.color, c.brand, c.resolved, c.moment, o.id, o.name
                       FROM "case" c
                       LEFT JOIN officer o on c.officer_id = o.id`

func (r *Repository) GetAllCases() ([]model.Case, error) {

	var got = make([]model.Case, 0)

	rows, err := r.db.Query(caseBaseQuery)

	if err != nil {
		return got, err
	}

	for rows.Next() {

		c, err := parseCaseRow(rows)

		if err != nil {
			return got, err
		}

		got = append(got, c)
	}

	return got, err
}

func (r *Repository) GetCaseByID(id int) (model.Case, error) {

	query := fmt.Sprintf("%s WHERE c.id = $1", caseBaseQuery)
	row := r.db.QueryRow(query, id)

	return parseCaseRow(row)
}

func parseCaseRow(s scanner) (c model.Case, err error) {

	var officerID null.Int
	var officerName null.String

	err = s.Scan(
		&c.Id,
		&c.Owner,
		&c.Color,
		&c.Brand,
		&c.Resolved,
		&c.Moment,
		&officerID,
		&officerName,
	)

	if officerID.Valid {
		c.Officer = &model.Officer{
			Id:   officerID,
			Name: officerName,
		}
	}

	return c, err
}

func (r *Repository) UpdateCase(c model.Case) (err error) {

	query := `UPDATE "case" SET owner = $1, color = $2, brand = $3, resolved = $4, moment = $5, officer_id = $6 WHERE id = $7`

	_, err = r.db.Exec(query, c.Owner, c.Color, c.Brand, c.Resolved, c.Moment, c.Officer.Id, c.Id)

	return err
}

func (r *Repository) InsertCase(c model.Case, officerId null.Int) (id null.Int, err error) {

    query := `INSERT INTO "case"(owner, color, brand, resolved, moment, officer_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	err = r.db.QueryRow(query, c.Owner, c.Color, c.Brand, c.Resolved, c.Moment, officerId).Scan(&id)

	return id, err
}

func (r *Repository) ResolveCase(id int) (officerID null.Int, err error) {

	query := `UPDATE "case" c SET resolved = true WHERE id = $1 RETURNING c.officer_id`

	err = r.db.QueryRow(query, id).Scan(&officerID)

	return officerID, err
}

func (r *Repository) AssignOpenCase(officerID null.Int) (err error) {

	query := `UPDATE "case" SET officer_id = $1 WHERE id = (SELECT id FROM "case" WHERE resolved = false LIMIT 1)`

	_, err = r.db.Exec(query, officerID)

	return err
}
