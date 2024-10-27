package repository

import "vivo/model"

type NoteRepository interface {
	Save(note model.Note)
	Update(note model.Note)
	Delete(id int)
	FindAll() []model.Note
	FindByID(id int) (model.Note, error)
}