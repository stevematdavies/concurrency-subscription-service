package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func(app *Config) Routes() http.Handler {
	mux := chi.NewRouter()
	
	mux.Use(middleware.Recoverer)
	mux.Use(app.LoadSession)
	
	mux.Get("/", app.Home)
	
	mux.Get("/login", app.Login)
	mux.Post("/login", app.DoLogin)
	
	mux.Get("/register", app.Register)
	mux.Post("/register", app.DoRegister)

	mux.Get("/activate", app.DoActivate)
	
	
	mux.Get("/logout", app.DoLogout)
	
	return mux
}