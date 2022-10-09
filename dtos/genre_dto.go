package dtos

type GenreRequest struct {
	Name string `json:"name"`
}

type GenreResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
