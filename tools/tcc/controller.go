package tcc

import (
	"net/http"
)

type Controller struct {
	Response http.ResponseWriter
	Request *http.Request
}