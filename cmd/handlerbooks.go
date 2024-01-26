package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/schlucht/mnrNaters/internal/models"
)

func (app *application) Book(w http.ResponseWriter, r *http.Request) {

	if err := app.renderTemplate(w, r, "book", &templateData{}, "booklist"); err != nil {
		app.errorlog.Println(err)
	}
}

func (app *application) AllBooks(w http.ResponseWriter, r *http.Request) {

	books, err := app.DB.GetBooks()
	for i, b := range books {
		sales, err := app.DB.GetSales(b.ID)
		if err != nil {			
			b.Sales = []*models.Sale{}
		}		
		books[i].Sales = sales
	}

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
	var b models.Book
	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		app.badRequest(w, r, err)
	}
	f, err := strconv.ParseFloat(b.Price, 64)
	if err != nil {
		f = 0.0
	}
	err = app.DB.UpdateBook(b.ID, b.Title, f)
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
	var b models.Book
	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		app.badRequest(w, r, err)
	}
	f, err := strconv.ParseFloat(b.Price, 64)
	if err != nil {
		f = 0.0
	}
	_, err = app.DB.InsertBook(b.Title, f)
	if err != nil {
		app.badRequest(w, r, err)
	}

	msg := message{
		message: "book is saved",
		msgType: MsgTypeInfo,
	}

	app.writeJSON(w, http.StatusOK, msg)
}
