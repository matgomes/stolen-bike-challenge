package repository

import (
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	"github.com/matgomes/stolen-bike-challenge/api/model"
)

const selectQuery = `SELECT c.id, c.owner, c.color, c.brand, c.resolved, c.moment, o.id, o.name
					 FROM "case" c
			  		 JOIN officer o on c.officer_id = o.id`

type scanner interface {
	Scan(...interface{}) error
}

func (r *Repository) GetAllCases() (got []model.Case, err error) {

	rows, err := r.db.Query(selectQuery)

	if err != nil {
		return got, err
	}

	for rows.Next() {

		c, err := parseRow(rows)

		if err != nil {
			return got, err
		}

		got = append(got, c)
	}

	return got, err
}

func (r *Repository) GetCaseByID(id int) (model.Case, error) {

	query := fmt.Sprintf("%s %s", selectQuery, "WHERE c.id = $1")
	row := r.db.QueryRow(query, id)

	return parseRow(row)
}

func parseRow(s scanner) (c model.Case, err error) {

	err = s.Scan(
		&c.Id,
		&c.Owner,
		&c.Color,
		&c.Brand,
		&c.Resolved,
		&c.Moment,
		&c.Officer.Id,
		&c.Officer.Name,
	)

	return c, err
}

func (r *Repository) UpdateField(id int, field string, value interface{}) (err error) {

	query := fmt.Sprintf(`UPDATE "case" SET %s = $1 WHERE id = $2`, pq.QuoteIdentifier(field))

	_, err = r.db.Exec(query, value, id)

	return err
}

func (r *Repository) UpdateCase(c model.Case) (err error) {

	query := `UPDATE "case" SET owner = $1, color = $2, brand = $3, resolved = $4, moment = $5, officer_id = $6
			  WHERE id = $7`

	_, err = r.db.Exec(query, c.Owner, c.Color, c.Brand, c.Resolved, c.Moment, c.Officer.Id, c.Id)

	return err
}

func (r *Repository) InsertCase(c model.Case) (id int, err error) {

	query := `INSERT INTO "case"(owner, color, brand, resolved, moment, officer_id) 
			  VALUES ($1, $2, $3, $4, $5, $6) 
			  RETURNING id`

	err = r.db.QueryRow(query, c.Owner, c.Color, c.Brand, c.Resolved, c.Moment, handleID(c.Officer.Id)).Scan(&id)

	return id, err
}

func (r *Repository) ResolveCase(id int) (officerId int, err error) {

	query := `UPDATE "case" c SET resolved = true WHERE id = $1 RETURNING c.officer_id`

	err = r.db.QueryRow(query, id).Scan(&officerId)

	return officerId, err
}

func handleID(id int) (new sql.NullInt64) {

	new.Int64 = int64(id)

	if id > 0 {
		new.Valid = true
	}

	return new
}
