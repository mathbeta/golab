package structs

type Head struct {
	Title string            `json:"title"`
	Metas map[string]string `json:"metas"`
}

type Body struct {
	Content string `json:"content"`
}

type Page struct {
	Head
	Body
}
