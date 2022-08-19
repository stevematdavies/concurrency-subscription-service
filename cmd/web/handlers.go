package main

import (
	"net/http"
)

func (app *Config) Home(w http.ResponseWriter, r *http.Request){
	app.Render(w, r, "home.page.gohtml", nil)
}

func (app *Config) Login(w http.ResponseWriter, r *http.Request){
	app.Render(w, r, "login.page.gohtml", nil)
}

func (app *Config) Register(w http.ResponseWriter, r *http.Request){
	app.Render(w, r, "register.page.gohtml", nil)
}


func (app *Config) DoLogin(w http.ResponseWriter, r *http.Request){
	_ = app.Session.RenewToken(r.Context())
	if err := r.ParseForm(); err != nil {
		app.ErrorLog.Println(err)
	}
	em := r.Form.Get("email")
	pw := r.Form.Get("password")

	u, err := app.Models.User.GetUserByEmail(em)
	if err != nil {
		app.Session.Put(r.Context(), "error", "User not found!")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	
	pwdIsValid, err := u.PasswordMatches(pw)

	if err != nil {
		app.Session.Put(r.Context(), "error", "Invalid Credentials, please try again!")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	if !pwdIsValid {
		app.Session.Put(r.Context(), "error", "Invalid Credentials!")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	app.Session.Put(r.Context(), "userID", u.ID)
	app.Session.Put(r.Context(), "user", u)
	app.Session.Put(r.Context(), "flash", "Successful Login!")

	http.Redirect(w,r, "/", http.StatusSeeOther)
}

func (app *Config) DoLogout(w http.ResponseWriter, r *http.Request){
	
}

func (app *Config) DoRegister(w http.ResponseWriter, r *http.Request){
	
}

func (app *Config) DoActivate(w http.ResponseWriter, r *http.Request){
	
}



