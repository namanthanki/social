package main

import (
	"net/http"
)

func (app *application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	// log.Printf("internal server errror: %s path:%s error:%v", r.Method, r.URL.Path, err)
	app.logger.Errorf("internal server error", "method", r.Method, "path", r.URL.Path, "error", err.Error())
	writeJSONError(w, http.StatusInternalServerError, "Internal Server Error")
}

func (app *application) badRequestError(w http.ResponseWriter, r *http.Request, err error) {
	// log.Printf("bad request errror: %s path:%s error:%v", r.Method, r.URL.Path, err)
	app.logger.Warnf("bad request error", "method", r.Method, "path", r.URL.Path, "error", err.Error())
	writeJSONError(w, http.StatusBadRequest, err.Error())
}

func (app *application) conflictResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Warnf("conflict error", "method", r.Method, "path", r.URL.Path, "error", err.Error())
	writeJSONError(w, http.StatusConflict, err.Error())
}

func (app *application) notFoundError(w http.ResponseWriter, r *http.Request, err error) {
	// log.Printf("not found errror: %s path:%s error:%v", r.Method, r.URL.Path, err)
	app.logger.Errorf("not found error", "method", r.Method, "path", r.URL.Path, "error", err.Error())
	writeJSONError(w, http.StatusNotFound, "Resource Not Found")
}
