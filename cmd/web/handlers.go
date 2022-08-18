package main

import "net/http"

func (app *Config) HomeHandler(w http.ResponseWriter, r *http.Request){
	app.Render(w, r, "home.page.gohtml", nil)
}