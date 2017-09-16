package api

import (
	"net/http"
)

// Handler это интерфейс обработчика http запросов
type Handler interface {
	http.Handler

	JSONResponse(
		resp http.ResponseWriter,
		data interface{},
		status ...int,
	) error

	XMLResponse(
		resp http.ResponseWriter,
		data interface{},
		status ...int,
	) error
}
