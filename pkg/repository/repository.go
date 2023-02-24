package repository

import (
	"database/sql"
	"fmt"
	"goapi/pkg/model"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func ConnectionMain() *sql.DB {
	conn, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		fmt.Println(err)
	}
	return conn
}

func GetMovie() (int, []model.Movie) {
	conexion := ConnectionMain()
	defer conexion.Close()
	movies := []model.Movie{}
	movie := model.Movie{}
	sql := `SELECT * FROM  movie`
	row, err := conexion.Query(sql)
	if err != nil {
		return http.StatusInternalServerError, movies
	}
	for row.Next() {
		row.Scan(&movie.Id, &movie.Name, &movie.Description, &movie.Year)
		movies = append(movies, movie)
	}
	return http.StatusOK, movies
}
func GetByIdMovie(id int) (int, model.Movie) {
	conexion := ConnectionMain()
	defer conexion.Close()

	movie := model.Movie{}
	sql := `SELECT * FROM movie WHERE id=$1`
	stmt, err := conexion.Prepare(sql)
	if err != nil {
		return http.StatusInternalServerError, movie
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	row.Scan(&movie.Id, &movie.Name, &movie.Description, &movie.Year)
	return http.StatusOK, movie
}

func PostMovie(movie model.Movie) int {
	conexion := ConnectionMain()
	defer conexion.Close()
	sql := `INSERT INTO movie("NOMBRE","DESCRIPTION","YEAR") VALUES ($1,$2,$3)`
	stmt, err := conexion.Prepare(sql)
	if err != nil {
		return http.StatusInternalServerError
	}
	defer stmt.Close()
	_, err2 := stmt.Exec(movie.Name, movie.Description, movie.Year)
	if err2 != nil {
		return http.StatusInternalServerError
	}
	return http.StatusCreated
}
func PutMovie(movie model.Movie) int {
	conexion := ConnectionMain()
	defer conexion.Close()
	sql := `UPDATE movie SET "NOMBRE"=$1, "DESCRIPTION"=$2, "YEAR"=$3 WHERE "id"=$4`
	stmt, err := conexion.Prepare(sql)
	if err != nil {
		return http.StatusInternalServerError
	}
	defer stmt.Close()
	_, err2 := stmt.Exec(movie.Name, movie.Description, movie.Year, movie.Id)
	if err2 != nil {
		return http.StatusInternalServerError
	}
	return http.StatusNoContent
}
func DeleteMovie(id int) int {
	conexion := ConnectionMain()
	defer conexion.Close()
	sql := `DELETE FROM movie WHERE "id"=$1`
	stmt, err := conexion.Prepare(sql)
	if err != nil {
		return http.StatusInternalServerError
	}
	defer stmt.Close()
	_, err2 := stmt.Exec(id)
	if err2 != nil {
		return http.StatusInternalServerError
	}
	return http.StatusNoContent
}
