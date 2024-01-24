package main

import "net/http"

func (app *application) Donate(w http.ResponseWriter, r *http.Request) {
	if err := app.renderTemplate(w, r, "donate", &templateData{}, "donatelist"); err != nil {
		app.errorlog.Println(err)
	}
}
