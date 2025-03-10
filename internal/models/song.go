package models

type Song struct {
	ID    int    `json:"id"`
	Group string `json:"group"`
	Song  string `json:"song"`
}
