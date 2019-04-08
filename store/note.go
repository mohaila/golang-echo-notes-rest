package store

import (
	"database/sql"

	"github.com/mohaila/echo-notes-rest/model"
)

// Note operations

// CreateNote creates a note
func (r *Context) CreateNote(tx *sql.Tx, n *model.Note) error {
	query := `
		INSERT INTO notes (
			title,
			description
		)
		VALUES (
			$1,
			$2
		)
		RETURNING id;
	`

	if tx != nil {
		return tx.QueryRow(query, n.Title, n.Description).Scan(&n.ID)
	}

	return r.db.QueryRow(query, n.Title, n.Description).Scan(&n.ID)
}

// UpdateNote updates a note
func (r *Context) UpdateNote(tx *sql.Tx, n *model.Note) error {
	query := `
		UPDATE notes 
		SET title = $1,
			description = $2
		WHERE id = $3;
	`

	var res sql.Result
	var err error

	if tx != nil {
		res, err = tx.Exec(query, n.Title, n.Description, n.ID)
	} else {
		res, err = r.db.Exec(query, n.Title, n.Description, n.ID)
	}

	if err != nil {
		return err
	}

	if nr, err := res.RowsAffected(); err != nil {
		return err
	} else if nr == 0 {
		return sql.ErrNoRows
	}

	return nil
}

// DeleteNote deletes a note
func (r *Context) DeleteNote(tx *sql.Tx, id int) error {
	query := `
		DELETE FROM notes 
		WHERE id = $1;
	`

	var res sql.Result
	var err error

	if tx != nil {
		res, err = tx.Exec(query, id)
	} else {
		res, err = r.db.Exec(query, id)
	}

	if err != nil {
		return err
	}

	if nr, err := res.RowsAffected(); err != nil {
		return err
	} else if nr == 0 {
		return sql.ErrNoRows
	}

	return nil
}

// GetNote fetchs a note by id
func (r *Context) GetNote(tx *sql.Tx, id int) (*model.Note, error) {
	query := `
		SELECT * 
		FROM notes
		WHERE id = $1;
	`

	var row *sql.Row

	if tx != nil {
		row = tx.QueryRow(query, id)
	} else {
		row = r.db.QueryRow(query, id)
	}

	n := &model.Note{}
	if err := row.Scan(&n.ID, &n.Title, &n.Description); err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		}
		return nil, nil
	}

	return n, nil
}

// GetAllNotes fetchs all notes
func (r *Context) GetAllNotes(tx *sql.Tx) ([]*model.Note, error) {
	query := `
		SELECT * 
		FROM notes
		ORDER BY id;
	`
	var rows *sql.Rows
	var err error

	if tx != nil {
		rows, err = tx.Query(query)
	} else {
		rows, err = r.db.Query(query)
	}

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	an := []*model.Note{}
	n := &model.Note{}

	for rows.Next() {
		if err = rows.Scan(&n.ID, &n.Title, &n.Description); err != nil {
			return nil, err
		}
		an = append(an, n)
	}

	return an, nil
}
