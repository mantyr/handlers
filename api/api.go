package api

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

// API это набор методов для создания API
type API struct {
}

// New возвращает новый API
func New() *API {
	return &API{}
}

// JSONResponse отправляет JSON ответ
func (h *API) JSONResponse(
	resp http.ResponseWriter,
	data interface{},
	status ...int,
) error {
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	resp.Header().Set("Content-Type", "application/json; charset=utf-8")

	switch {
	case len(status) > 0:
		resp.WriteHeader(status[0])
	default:
		resp.WriteHeader(http.StatusOK)
	}
	n, err := resp.Write(body)
	if len(body) != n {
		return fmt.Errorf(
			"expected to record %d bytes but actual %d bytes",
			len(body),
			n,
		)
	}
	return err
}

// XMLResponse отправляет XML ответ
func (h *API) XMLResponse(
	resp http.ResponseWriter,
	data interface{},
	status ...int,
) error {
	body, err := xml.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	resp.Header().Set("Content-Type", "application/xml; charset=utf-8")

	switch {
	case len(status) > 0:
		resp.WriteHeader(status[0])
	default:
		resp.WriteHeader(http.StatusOK)
	}

	n, err := resp.Write([]byte(xml.Header))
	switch {
	case err != nil:
		return err
	case n != len(xml.Header):
		return fmt.Errorf(
			"expected to record %d bytes (xml.Header) byt actual %d bytes",
			len(xml.Header),
			n,
		)
	}
	n, err = resp.Write(body)
	switch {
	case err != nil:
		return err
	case n != len(body):
		return fmt.Errorf(
			"expected to record %d bytes (xml body) byt actual %d bytes",
			len(body),
			n,
		)
	}
	return nil
}
