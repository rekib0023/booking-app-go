package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   !app.Debug,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

func LoadSession(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
