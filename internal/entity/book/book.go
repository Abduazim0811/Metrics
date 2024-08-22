package book

type Book struct {
	ID             string `json:"id"`
	Title          string `json:"title"`
	Author         string `json:"author"`
	Published_Date string `json:"published_date"`
	Isbn           int    `json:"isbn"`
}

type CreateBook struct {
	Title          string `json:"title"`
	Author         string `json:"author"`
	Published_Date string `json:"published_date"`
	Isbn           int    `json:"isbn"`
}

type GetBook struct {
	ID string `json:"id"`
}
