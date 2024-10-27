package request

type CreateNoteRequest struct {
	Content string `validate:"required,min=5,max=255" json:"content"`
}

type UpdateNoteRequest struct {
	ID      int    `validate:"required"`
	Content string `validate:"required,min=5,max=255" json:"content"`
}