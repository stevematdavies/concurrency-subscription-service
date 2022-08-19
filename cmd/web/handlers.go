package main

import "net/http"

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
	
}

func (app *Config) DoLogout(w http.ResponseWriter, r *http.Request){
	
}

func (app *Config) DoRegister(w http.ResponseWriter, r *http.Request){
	
}

func (app *Config) DoActivate(w http.ResponseWriter, r *http.Request){
	
}



