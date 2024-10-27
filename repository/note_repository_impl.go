package repository

import (
	"errors"
	"vivo/data/request"
	"vivo/helper"
	"vivo/model"

	"gorm.io/gorm"
)

type NoteRepositoryImpl struct {
	db *gorm.DB
}

// Delete implements NoteRepository.
func (n *NoteRepositoryImpl) Delete(id int) {
	var note model.Note
	result := n.db.Find(&note, id).Delete(&note)
	helper.ErrorPanic(result.Error)
}

// FindAll implements NoteRepository.
func (n *NoteRepositoryImpl) FindAll() []model.Note {
	var notes []model.Note
	result := n.db.Find(&notes)
	helper.ErrorPanic(result.Error)
	return notes
}

// FindByID implements NoteRepository.
func (n *NoteRepositoryImpl) FindByID(id int) (model.Note, error) {
	var note model.Note
	result := n.db.Find(&note, id)
	if result != nil {
		return note, nil
	} else {
		return note, errors.New("data not found")
		
	}
}

// Save implements NoteRepository.
func (n *NoteRepositoryImpl) Save(note model.Note) {
	result := n.db.Create(&note)
	helper.ErrorPanic(result.Error)
}

// Update implements NoteRepository.
func (n *NoteRepositoryImpl) Update(note model.Note) {
	var updateNote = request.UpdateNoteRequest{
		ID:      note.ID,
		Content: note.Content,
	}

	result := n.db.Model(&note).Updates(updateNote)
	helper.ErrorPanic(result.Error)
}

func NewNoteRepository(db *gorm.DB) NoteRepository {
	return &NoteRepositoryImpl{db: db}
}
