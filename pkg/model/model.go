package model

type Movie struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Year        string `json:"year"`
}
