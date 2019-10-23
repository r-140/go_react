package model

type News struct {
	Id     string `json:"id"`
	Title  string `json:"name"`
	Teaser string `json:"teaser"`
	Body   string `json:"body"`
}

func (a *News) ToString() string {
	return a.Id + " " + a.Title + " " + a.Teaser
}
