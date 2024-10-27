package service

import (
	"vivo/data/request"
	"vivo/data/response"
	"vivo/helper"
	"vivo/model"
	"vivo/repository"

	"github.com/go-playground/validator/v10"
)

type NoteServiceImpl struct {
	NoteRepository repository.NoteRepository
	validate       *validator.Validate
}

// Create implements NoteService.
func (n *NoteServiceImpl) Create(note request.CreateNoteRequest) {
	err := n.validate.Struct(note)
	helper.ErrorPanic(err)

	noteModel := model.Note{
		Content: note.Content,
	}
	n.NoteRepository.Save(noteModel)
}

// Delete implements NoteService.
func (n *NoteServiceImpl) Delete(id int) {
	n.NoteRepository.Delete(id)
}

// FindAll implements NoteService.
func (n *NoteServiceImpl) FindAll() []response.NoteResponse {
	notes := n.NoteRepository.FindAll()
	var notesResponse []response.NoteResponse
	for _, note := range notes {
		notesResponse = append(notesResponse, response.NoteResponse{
			ID:      note.ID,
			Content: note.Content,
		})
	}
	return notesResponse
}

// FindByID implements NoteService.
func (n *NoteServiceImpl) FindByID(id int) response.NoteResponse {
	note, err := n.NoteRepository.FindByID(id)
	helper.ErrorPanic(err)
	return response.NoteResponse{
		ID:      note.ID,
		Content: note.Content,
	}
}

// Update implements NoteService.
func (n *NoteServiceImpl) Update(note request.UpdateNoteRequest) {
	err := n.validate.Struct(note)
	helper.ErrorPanic(err)
	noteModel := model.Note{
		ID:      note.ID,
		Content: note.Content,
	}
	n.NoteRepository.Update(noteModel)
}

func NewNoteServiceImpl(NoteRepository repository.NoteRepository, validate *validator.Validate) NoteService {
	return &NoteServiceImpl{NoteRepository: NoteRepository, validate: validate}
}
