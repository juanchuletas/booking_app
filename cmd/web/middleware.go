package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

func WriteToConsole(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("Hit the page")
		next.ServeHTTP(w, r)
	})

}
func NoSurf(next http.Handler) http.Handler {

	cfsrHandler := nosurf.New(next)
	cfsrHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return cfsrHandler
}
func SessionLoad(next http.Handler) http.Handler {

	return session.LoadAndSave(next)
}
