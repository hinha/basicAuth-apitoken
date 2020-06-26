package schema

type ArticlesSchema struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	CreatedAt	string `json:"created_at"`
}
