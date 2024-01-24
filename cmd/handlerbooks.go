package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (app *application) Books(w http.ResponseWriter, r *http.Request) {
	books, err := app.DB.GetBooks()
	if err != nil {
		app.errorlog.Println(err)
	}
	data := make(map[string]interface{})
	data["books"] = books
	if err := app.renderTemplate(w, r, "book", &templateData{
		Data: data,
	}, "booklist"); err != nil {
		app.errorlog.Println(err)
	}
}

func (app *application) DeleteBook(w http.ResponseWriter, r *http.Request) {

	strid := chi.URLParam(r, "id")
	if i, err := strconv.Atoi(strid); err == nil {
		err = app.DB.DeleteBook(i)
		if err != nil {
			app.errorlog.Println(err)
		}
	}

	if err := app.renderTemplate(w, r, "book", &templateData{}, "booklist"); err != nil {
		app.errorlog.Println(err)
	}
}

func (app *application) SaveBook(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.errorlog.Println(err)
		return
	}
	title := r.PostForm.Get("title")
	price := r.PostForm.Get("price")

	id, err := app.DB.InsertBook(title, price)
	if err != nil {
		app.errorlog.Println(err)
		return
	}

	fmt.Println(id)

	http.Redirect(w, r, "/books", http.StatusSeeOther)

	// if err := app.renderTemplate(w, r, "book", &templateData{}); err != nil {
	// 	app.errorlog.Println(err)
	// }
}
