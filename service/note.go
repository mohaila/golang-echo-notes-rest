package service

import (
	"github.com/mohaila/echo-notes-rest/model"
	"github.com/mohaila/echo-notes-rest/store"
)

type (
	// NoteService for note operations
	NoteService interface {
		GetAllNotes() ([]*model.Note, error)
		GetNote(id int) (*model.Note, error)
		CreateNote(n *model.Note) error
		UpdateNote(n *model.Note) error
		DeleteNote(id int) error
	}

	// NoteServiceContext context from NoteService operations
	NoteServiceContext struct {
		store.Store
	}
)

// NewNoteService creates a new note service
func NewNoteService(s store.Store) NoteService {
	return &NoteServiceContext{s}
}

// GetAllNotes fetchs all notes
func (ns *NoteServiceContext) GetAllNotes() ([]*model.Note, error) {
	return ns.Store.GetAllNotes(nil)
}

// GetNote fetchs a note
func (ns *NoteServiceContext) GetNote(id int) (*model.Note, error) {
	return ns.Store.GetNote(nil, id)
}

// CreateNote creates a note
func (ns *NoteServiceContext) CreateNote(n *model.Note) error {
	var err error
	tx, err := ns.Begin()
	if err != nil {
		return err
	}

	err = ns.Store.CreateNote(tx, n)
	if err != nil {
		ns.Rollback(tx)
		return err
	}

	if err = ns.Commit(tx); err != nil {
		return err
	}

	return nil
}

// UpdateNote updates a note
func (ns *NoteServiceContext) UpdateNote(n *model.Note) error {
	var err error
	tx, err := ns.Begin()
	if err != nil {
		return err
	}

	err = ns.Store.UpdateNote(tx, n)
	if err != nil {
		ns.Rollback(tx)
		return err
	}

	if err = ns.Commit(tx); err != nil {
		return err
	}

	return nil
}

// DeleteNote deletes a note
func (ns *NoteServiceContext) DeleteNote(id int) error {
	var err error
	tx, err := ns.Begin()
	if err != nil {
		return err
	}

	err = ns.Store.DeleteNote(tx, id)
	if err != nil {
		ns.Rollback(tx)
		return err
	}

	if err = ns.Commit(tx); err != nil {
		return err
	}

	return nil
}
