package main

import (
	"fmt"
	"net/http"
	"text/template"
	"time"
)

var templatePath = "./cmd/web/templates"

type TemplateData struct {
	StrDataMap map[string]string
	FltDataMap map[string]float64
	IntDataMap map[string]int
	DataMap map[string]any
	Flash string
	Warning string
	Error string
	Authenticated bool
	Now time.Time
	// User *data.User
}

func (app *Config) Render(w http.ResponseWriter, r * http.Request, t string, d *TemplateData){
	partials := []string{
		fmt.Sprintf("%s/base.layout.gohtml", templatePath),
		fmt.Sprintf("%s/header.partial.gohtml", templatePath),
		fmt.Sprintf("%s/navbar.partial.gohtml", templatePath),
		fmt.Sprintf("%s/footer.partial.gohtml", templatePath),
		fmt.Sprintf("%s/alerts.partial.gohtml", templatePath),
	}

	var ts []string
	ts = append(ts, fmt.Sprintf("%s/%s",templatePath, t))
	ts = append(ts, partials...)

	if d == nil {
		d = &TemplateData{}
	}

	tmpl, err := template.ParseFiles(ts...);
	if err != nil {
		app.ErrorLog.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, app.DefaultData(d, r)); err != nil {
		app.ErrorLog.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (app *Config) DefaultData(d *TemplateData, r *http.Request) *TemplateData {
	d.Flash = app.Session.PopString(r.Context(), "flash")
	d.Warning = app.Session.PopString(r.Context(), "warning")
	d.Error = app.Session.PopString(r.Context(), "error")
	if app.IsAuthenticated(r){
		d.Authenticated = true
		// TODO get more User information
	} else {
		d.Authenticated = false
	}
	d.Now = time.Now()

	return d
}


func(app *Config) IsAuthenticated(r *http.Request) bool {
	return app.Session.Exists(r.Context(), "userID")
}