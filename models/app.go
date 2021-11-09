package models

type App struct {
	Id       int
	Company  string
	Position string
	Link     string
	Status   string
}

type GetAppsResponse struct {
	Apps []App
}
