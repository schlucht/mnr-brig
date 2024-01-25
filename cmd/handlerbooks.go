package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type book struct {
	Id    int
	Title string
	Price float64
}

func (app *application) Book(w http.ResponseWriter, r *http.Request) {

	if err := app.renderTemplate(w, r, "book", &templateData{}, "booklist"); err != nil {
		app.errorlog.Println(err)
	}
}

func (app *application) AllBooks(w http.ResponseWriter, r *http.Request) {

	books, err := app.DB.GetBooks()
	if err != nil {
		app.badRequest(w, r, err)
	}

	app.writeJSON(w, http.StatusOK, books)
}

func (app *application) GetBook(w http.ResponseWriter, r *http.Request) {
	strid := chi.URLParam(r, "id")
	if i, err := strconv.Atoi(strid); err == nil {
		book, err := app.DB.GetBook(i)
		if err != nil {
			app.badRequest(w, r, err)
		}

		app.writeJSON(w, http.StatusOK, book)
	}
}

func (app *application) DeleteBook(w http.ResponseWriter, r *http.Request) {

	strid := chi.URLParam(r, "id")
	app.infoLog.Println(strid)
	if i, err := strconv.Atoi(strid); err == nil {
		err = app.DB.DeleteBook(i)
		if err != nil {
			app.badRequest(w, r, err)
		}
	}

	msg := message{
		message: "book is deleted",
		msgType: MsgTypeInfo,
	}

	app.writeJSON(w, http.StatusOK, msg)
}

func (app *application) EditBook(w http.ResponseWriter, r *http.Request) {
	var b book
	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		app.badRequest(w, r, err)
	}

	err = app.DB.UpdateBook(b.Id, b.Title, b.Price)
	if err != nil {
		app.badRequest(w, r, err)
	}

	msg := message{
		message: "book is edited",
		msgType: MsgTypeInfo,
	}

	app.writeJSON(w, http.StatusOK, msg)
}

func (app *application) SaveBook(w http.ResponseWriter, r *http.Request) {
	var b book
	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		app.badRequest(w, r, err)
	}
	app.infoLog.Println(b.Title)
	_, err = app.DB.InsertBook(b.Title, b.Price)
	if err != nil {
		app.badRequest(w, r, err)
	}

	msg := message{
		message: "book is saved",
		msgType: MsgTypeInfo,
	}

	app.writeJSON(w, http.StatusOK, msg)
}
