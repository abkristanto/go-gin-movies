package dtos

type CreateMovieRequest struct {
	Title    string `json:"title"`
	Image    string `json:"image"`
	Synopsis string `json:"synopsis"`
	Genres   []int  `json:"genres"`
}

type CreateMovieResponse struct {
	ID       uint             `json:"id"`
	Title    string           `json:"title"`
	Image    string           `json:"image"`
	Synopsis string           `json:"synopsis"`
	Genres   []*GenreResponse `json:"genres"`
}
