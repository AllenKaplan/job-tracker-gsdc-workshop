package models

type App struct {
	Id       int    `json:"id,omitempty"`
	Company  string `json:"company,omitempty"`
	Position string `json:"position,omitempty"`
	Link     string `json:"link,omitempty"`
	Status   string `json:"status,omitempty"`
}
