package web

import (
	"io/ioutil"
	"net/http"

	"github.com/calebthompson/pgf/pgp"

	"golang.org/x/net/context"
)

func Form(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	form, err := ioutil.ReadFile("form.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(form)
}

func KeychainHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	armored := r.FormValue("armored")
	body, err := pgp.SignersFromKeyring(armored)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
