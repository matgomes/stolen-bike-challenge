package repository

import (
    "github.com/matgomes/stolen-bike-challenge/api/model"
)

func (r *Repository) FindAvailableOfficer() (result model.Officer, err error) {

    query := `SELECT id, name FROM officer WHERE id 
              NOT IN (SELECT officer_id FROM "case" WHERE resolved = false) 
              LIMIT 1`

    err = r.db.QueryRow(query).Scan(&result.Id, &result.Name)

    return result, err
}

func (r *Repository) InsertOfficer(o model.Officer) (id int, err error) {

    query := `INSERT INTO officer(name) VALUES ($1) RETURNING id`

    err = r.db.QueryRow(query, o.Name).Scan(&id)

    return id, err
}

func (r *Repository) RemoveOfficer(id int) (deleted bool, err error) {

    query := `WITH deleted AS (DELETE FROM officer WHERE id = $1 RETURNING *) 
              SELECT count(*) FROM deleted`

    var count int
    err = r.db.QueryRow(query, id).Scan(&count)

    return count > 0, err
}
