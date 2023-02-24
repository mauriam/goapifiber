package service

import (
	"goapi/pkg/model"
	"goapi/pkg/repository"
)

func GetMovie() (int, []model.Movie) {
	return repository.GetMovie()
}
func GetByIdMovie(id int) (int, model.Movie) {
	return repository.GetByIdMovie(id)
}
func PostMovie(movie model.Movie) int {
	return repository.PostMovie(movie)
}
func PutMovie(movie model.Movie) int {
	return repository.PutMovie(movie)
}
func DeleteMovie(id int) int {
	return repository.DeleteMovie(id)
}
