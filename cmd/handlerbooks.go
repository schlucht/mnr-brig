package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/schlucht/mnrNaters/internal/models"
)

func (app *application) Book(w http.ResponseWriter, r *http.Request) {

	if err := app.renderTemplate(w, r, "book", &templateData{}, "booklist", "saleInfo"); err != nil {
		app.errorlog.Fatalln(err)
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
		if err = app.badRequest(w, r, err); err != nil {
			app.errorlog.Fatalln(err)
		}
	}

	if err = app.writeJSON(w, http.StatusOK, books); err != nil {
		app.errorlog.Fatalln(err)
	}
}

func (app *application) GetBook(w http.ResponseWriter, r *http.Request) {
	strid := chi.URLParam(r, "id")
	if i, err := strconv.Atoi(strid); err == nil {
		book, err := app.DB.GetBook(i)
		if err != nil {
			if err = app.badRequest(w, r, err); err != nil {
				app.errorlog.Fatal(err)
			}
		}

		if err = app.writeJSON(w, http.StatusOK, book); err != nil {
			app.errorlog.Fatal(err)
		}
	}
}

func (app *application) DeleteBook(w http.ResponseWriter, r *http.Request) {

	strid := chi.URLParam(r, "id")
	app.infoLog.Fatalln(strid)
	if i, err := strconv.Atoi(strid); err == nil {
		err = app.DB.DeleteBook(i)
		if err != nil {
			if err = app.badRequest(w, r, err); err != nil {
				app.errorlog.Fatalln(err)
			}
		}
	}

	msg := message{
		message: "book is deleted",
		msgType: MsgTypeInfo,
	}

	if err := app.writeJSON(w, http.StatusOK, msg); err != nil {
		app.errorlog.Fatalln(err)
	}
}

func (app *application) EditBook(w http.ResponseWriter, r *http.Request) {
	var b models.Book
	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		if err = app.badRequest(w, r, err); err != nil {
			app.errorlog.Fatalln(err)
		}
	}

	err = app.DB.UpdateBook(b.ID, b.Title, b.Price)
	if err != nil {
		err = app.badRequest(w, r, err)
		{
			app.errorlog.Fatalln(err)
		}
	}

	msg := message{
		message: "book is edited",
		msgType: MsgTypeInfo,
	}

	if err = app.writeJSON(w, http.StatusOK, msg); err != nil {
		app.errorlog.Fatalln(err)
	}
}

// Speichert ein neues Buch in der DB
func (app *application) SaveBook(w http.ResponseWriter, r *http.Request) {
	var b models.Book
	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(&b)

	if err != nil {
		if err = app.badRequest(w, r, err); err != nil {
			app.errorlog.Fatalln(err)
		}
	}

	_, err = app.DB.InsertBook(b.Title, b.Price)
	if err != nil {
		if err = app.badRequest(w, r, err); err != nil {
			app.errorlog.Fatalln(err)
		}
	}

	msg := message{
		message: "book is saved",
		msgType: MsgTypeInfo,
	}

	if err = app.writeJSON(w, http.StatusOK, msg); err != nil {
		app.errorlog.Fatalln(err)
	}
}

// Speichert ein gekauftes Buch in der Datenbank
func (app *application) SaveSale(w http.ResponseWriter, r *http.Request) {
	var s models.Sale
	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		if err = app.badRequest(w, r, err); err != nil {
			app.errorlog.Fatalln(err)
		}
	}

	_, err = app.DB.InsertSale(s)
	if err != nil {
		if err = app.badRequest(w, r, err); err != nil {
			app.errorlog.Fatalln(err)
		}
	}

	msg := message{
		message: "sale is saved",
		msgType: MsgTypeInfo,
	}

	if err = app.writeJSON(w, http.StatusOK, msg); err != nil {
		app.errorlog.Fatalln(err)
	}
}
