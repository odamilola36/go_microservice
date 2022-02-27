package restutil

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"microservices/security"
	"net/http"
)

type JError struct {
	Error string `json:"error"`
}

var (
	ErrEmptyBody = errors.New("Body can't be empty")
)

func WriteAsJson(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	e := "error"
	if err != nil {
		e = err.Error()
	}
	WriteAsJson(w, status, JError{Error: e})
}

func AuthRequestWithId(r *http.Request) (*security.TokenPayload, error) {
	token, err := security.ExtractToken(r)
	if err != nil {
		return nil, err
	}

	vars := mux.Vars(r)
	id := vars["id"]

	payload, err := security.NewTokenPayload(token)
	if err != nil {
		return nil, err
	}
	if payload.UserId != id {
		return nil, security.ErrInvalidToken
	}
	return payload, nil
}
