package services

import "net/http"

type Service interface {
	HandleRequest(w http.ResponseWriter, r *http.Request)
	GetName() string
}
