package model

type Post struct {
	Message string `json:"message" validate:"required"`
}
