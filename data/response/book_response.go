package response

type BookResponse struct {
	ID             int    `json:"id"`
	Title          string `json:"title"`
	Author         string `json:"author"`
	Published_Date string    `json:"published_date"`
}
