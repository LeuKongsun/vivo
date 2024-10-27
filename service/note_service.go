package service

import (
	"vivo/data/request"
	"vivo/data/response"
)

type NoteService interface {
	Create(note request.CreateNoteRequest)
	Update(note request.UpdateNoteRequest)
	Delete(id int)
	FindAll() []response.NoteResponse
	FindByID(id int) response.NoteResponse
}
