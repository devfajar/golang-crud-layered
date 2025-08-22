package request

type BookCreateRequest struct {
	Title          string `validate:"required min=1, max=100" json:"title"`
	Author         string `validate:"required" json:"author"`
	Published_Date string `validate:"omitempty" json:"published_date"`
}
