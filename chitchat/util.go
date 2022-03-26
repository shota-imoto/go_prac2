package main

import (
	"errors"
	"net/http"

	"github.com/sausheong/gwp/Chapter_2_Go_ChitChat/chitchat/data"
)

func session(w http.ResponseWriter, r *http.Request) (sess data.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		sess = data.Session{Uuid: cookie.Value}
		if ok, _ := sess.Check(); !ok {
			err = errors.New("Invalid session")
		}
	}
	return
}