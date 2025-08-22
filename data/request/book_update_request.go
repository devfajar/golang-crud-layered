package request

type BookUpdateRequest struct {
	ID             int        `json:"id" validate:"required"`
	Title          string     `validate:"required,min=1,max=100" json:"title"`
	Author         string     `json:"author"`
	Published_Date string `json:"published_date"`
}
