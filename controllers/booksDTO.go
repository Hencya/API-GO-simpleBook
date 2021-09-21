package Controller

import (
	Models "Book_CRUD/models"
	"gorm.io/gorm"
	"time"
)

type BookDetailResponse struct {
	ID        int    `json:"ID"`
	Title     string `json:"title"`
	Isbn      string `json:"isbn"`
	Writer    string `json:"writer"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at" `
}

func ModelToBookDetailResponse(data *Models.Books) *BookDetailResponse{
	return &BookDetailResponse{
		ID: data.ID,
		Title: data.Title,
		Isbn: data.Isbn,
		Writer: data.Writer,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
		DeletedAt: data.DeletedAt,
	}
}

type PostBookRequest struct {
	Title     string `json:"title"`
	Isbn      string `json:"isbn"`
	Writer    string `json:"writer"`
}

type PutBookRequest struct {
	Title     string `json:"title"`
	Isbn      string `json:"isbn"`
	Writer    string `json:"writer"`
}

func PostBookRequestToModelRequest(data *PostBookRequest) *Models.Books{
	return &Models.Books{
		Title: data.Title,
		Isbn: data.Isbn,
		Writer: data.Writer,
	}
}
func PutBookRequestToModelRequest(data *PutBookRequest) *Models.Books{
	return &Models.Books{
		Title: data.Title,
		Isbn: data.Isbn,
		Writer: data.Writer,
	}
}
