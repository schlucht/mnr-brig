package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (app *application) Books(w http.ResponseWriter, r *http.Request) {

	data, err := app.loadAllBooks()
	if err != nil {
		app.errorlog.Println(err)
	}

	if err = app.renderTemplate(w, r, "book", &templateData{
		Data: data,
	}, "booklist"); err != nil {
		app.errorlog.Println(err)
	}
}

func (app *application) GetBook(w http.ResponseWriter, r *http.Request) {
	strid := chi.URLParam(r, "id")
	if i, err := strconv.Atoi(strid); err == nil {
		b, err := app.DB.GetBook(i)
		if err != nil {
			app.errorlog.Println(err)
			return
		}
		out, err := json.MarshalIndent(b, "", "  ")
		if err != nil {
			app.errorlog.Println(err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(out)
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
	data, err := app.loadAllBooks()
	if err != nil {
		app.errorlog.Println(err)
	}

	if err := app.renderTemplate(w, r, "book", &templateData{
		Data: data,
	}, "booklist"); err != nil {
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
	p, err := strconv.ParseFloat(price, 64)
	if err != nil {
		p = 0.0
	}

	id, err := app.DB.InsertBook(title, p)
	if err != nil {
		app.errorlog.Println(err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, map[string]interface{}{
		"id": id,
	}, nil)
	if err != nil {
		app.errorlog.Println(err)
	}
	http.Redirect(w, r, "/books", http.StatusSeeOther)
}

func (app *application) loadAllBooks() (map[string]interface{}, error) {
	books, err := app.DB.GetBooks()
	if err != nil {
		return nil, err
	}
	data := make(map[string]interface{})
	data["books"] = books
	return data, nil
}
