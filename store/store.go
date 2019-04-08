package store

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/mohaila/echo-notes-rest/config"
	"github.com/mohaila/echo-notes-rest/model"
)

type (
	// Store for all databases operations
	Store interface {
		// Database
		Close() error
		Begin() (*sql.Tx, error)
		Commit(tx *sql.Tx) error
		Rollback(tx *sql.Tx) error
		// Note
		CreateNote(tx *sql.Tx, n *model.Note) error
		UpdateNote(tx *sql.Tx, n *model.Note) error
		DeleteNote(tx *sql.Tx, id int) error
		GetNote(tx *sql.Tx, id int) (*model.Note, error)
		GetAllNotes(tx *sql.Tx) ([]*model.Note, error)
	}

	// Context used by Store
	Context struct {
		db *sql.DB
	}
)

// NewStore returns a Store for operations
func NewStore(c *config.DBConfig) (Store, error) {
	cs := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.Host,
		c.DPort,
		c.User,
		c.Password,
		c.Name,
		c.SSLMode,
	)
	db, err := sql.Open("postgres", cs)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Context{db}, nil
}

func missingTransaction(op string) error {
	s := fmt.Sprintf("Missing transaction for operation %s", op)
	return errors.New(s)
}

// Close the database
func (r *Context) Close() error {
	return r.db.Close()
}

// Begin a transaction
func (r *Context) Begin() (*sql.Tx, error) {
	return r.db.Begin()
}

// Commit the transaction
func (r *Context) Commit(tx *sql.Tx) error {
	if tx == nil {
		return missingTransaction("commit")
	}

	return tx.Commit()
}

// Rollback the transaction
func (r *Context) Rollback(tx *sql.Tx) error {
	if tx == nil {
		return missingTransaction("rollback")
	}

	return tx.Rollback()
}
