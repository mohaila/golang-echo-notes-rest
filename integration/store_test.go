package integration

import (
	"database/sql"
	"testing"

	"github.com/mohaila/echo-notes-rest/model"

	// PSQL driver init
	_ "github.com/lib/pq"
	"github.com/mohaila/echo-notes-rest/config"
	"github.com/mohaila/echo-notes-rest/store"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type (
	StoreSuite struct {
		suite.Suite
		store.Store
	}
)

func init() {

}

// delete all notes
func deleteAllNotes(tx *sql.Tx) {
	query := `
		DELETE FROM notes;
	`
	_, _ = tx.Exec(query)
}

func (s *StoreSuite) SetupSuite() {
	dc := config.GetDBConfig()

	s.Store, _ = store.NewStore(dc)
}

func (s *StoreSuite) TearDownSuite() {
	s.Store.Close()
}

func (s *StoreSuite) TestStoreCreateNote() {
	tx, _ := s.Store.Begin()
	defer tx.Rollback()

	deleteAllNotes(tx)

	n := &model.Note{
		Title:       "Test note",
		Description: "Test note description",
	}
	_ = s.Store.CreateNote(tx, n)
	nn, _ := s.Store.GetNote(tx, n.ID)

	s.Equal(nn.ID, n.ID)
}

func (s *StoreSuite) TestStoreUpdateNote() {
	tx, _ := s.Store.Begin()
	defer tx.Rollback()

	deleteAllNotes(tx)

	n := &model.Note{
		Title:       "Test note",
		Description: "Test note description",
	}
	_ = s.Store.CreateNote(tx, n)

	title := "Updated test note"
	n.Title = title
	desc := "Updated test note description"
	n.Description = desc
	_ = s.Store.UpdateNote(tx, n)

	nn, _ := s.Store.GetNote(tx, n.ID)
	s.Equal(n.ID, nn.ID)
	s.Equal(nn.Title, title)
	s.Equal(nn.Description, desc)
}

func (s *StoreSuite) TestStoreDeleteteNote() {
	tx, _ := s.Store.Begin()
	defer tx.Rollback()

	deleteAllNotes(tx)

	n := &model.Note{
		Title:       "Test note",
		Description: "Test note description",
	}
	_ = s.Store.CreateNote(tx, n)
	an, _ := s.Store.GetAllNotes(tx)
	s.Equal(len(an), 1)

	_ = s.Store.DeleteNote(tx, n.ID)
	an, _ = s.Store.GetAllNotes(tx)
	s.Equal(len(an), 0)
}

func (s *StoreSuite) TestStoreGetAllNotes() {
	tx, _ := s.Store.Begin()
	defer tx.Rollback()

	deleteAllNotes(tx)

	an, _ := s.Store.GetAllNotes(tx)
	s.Equal(len(an), 0)

	n1 := &model.Note{
		Title:       "Test note",
		Description: "Test note description",
	}
	_ = s.Store.CreateNote(tx, n1)

	n2 := &model.Note{
		Title:       "Test note",
		Description: "Test note description",
	}
	_ = s.Store.CreateNote(tx, n2)

	an, _ = s.Store.GetAllNotes(tx)
	s.Equal(len(an), 2)
	assert.ObjectsAreEqual(n1, an[0])
	assert.ObjectsAreEqual(n2, an[1])
}

func (s *StoreSuite) TestStoreGetNote() {
	tx, _ := s.Store.Begin()
	defer tx.Rollback()

	deleteAllNotes(tx)

	n, _ := s.Store.GetNote(tx, 1)
	s.Nil(n)

	n = &model.Note{
		Title:       "Test note",
		Description: "Test note description",
	}
	_ = s.Store.CreateNote(tx, n)

	n, _ = s.Store.GetNote(tx, n.ID)
	s.NotNil(n)
}

func TestStoreSuite(t *testing.T) {
	suite.Run(t, new(StoreSuite))
}
