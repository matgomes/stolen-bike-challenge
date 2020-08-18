package repository

import (
	"fmt"
	"github.com/matgomes/stolen-bike-challenge/app/model"
)

func (r *Repository) FindAvailableOfficer() (result model.Officer, err error) {

	queryBusyOfficer := `SELECT officer_id FROM "case" WHERE resolved = false AND officer_id IS NOT NULL`

	query := fmt.Sprintf(`SELECT id, name FROM officer WHERE id NOT IN (%s) LIMIT 1`, queryBusyOfficer)

	err = r.db.QueryRow(query).Scan(&result.Id, &result.Name)

	return result, err
}

func (r *Repository) InsertOfficer(o model.Officer) (id int, err error) {

	query := `INSERT INTO officer(name) VALUES ($1) RETURNING id`

	err = r.db.QueryRow(query, o.Name).Scan(&id)

	return id, err
}

func (r *Repository) RemoveOfficer(id int) (deleted bool, err error) {

	query := `WITH deleted AS (DELETE FROM officer WHERE id = $1 RETURNING *) SELECT count(*) FROM deleted`

	var count int
	err = r.db.QueryRow(query, id).Scan(&count)

	return count > 0, err
}
