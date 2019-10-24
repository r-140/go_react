package model

// News model
type News struct {
	Id     string `json:"_id"`
	Title  string `json:"title"`
	Teaser string `json:"teaser"`
	Body   string `json:"body"`
}

func (a *News) String() string {
	return a.Id + " " + a.Title + " " + a.Teaser
}
