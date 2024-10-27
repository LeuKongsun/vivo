package model

type Note struct {
	ID      int    `gnorm:"type:int;primary_key;"`
	Content string `gnorm:"type:text;not null;"`
}
