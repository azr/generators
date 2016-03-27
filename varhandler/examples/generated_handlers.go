// Code generated by "handler -func Simple ./examples/"; DO NOT EDIT

package main

import "net/http"

func SimpleHandler(w http.ResponseWriter, r *http.Request) {
	var err error

	param0, err := HTTPX(w, r)
	if err != nil {
		HandleHttpErrorWithDefaultStatus(w, http.StatusBadRequest, err)
		return
	}

	param1, err := HTTPY(w, r)
	if err != nil {
		HandleHttpErrorWithDefaultStatus(w, http.StatusBadRequest, err)
		return
	}

	param2, err := HTTPZ(w, r)
	if err != nil {
		HandleHttpErrorWithDefaultStatus(w, http.StatusBadRequest, err)
		return
	}

	Simple(param0, param1, param2)
}

func HandleHttpErrorWithDefaultStatus(w http.ResponseWriter, status int, err error) {
	type HttpError interface {
		HttpError() (error string, code int)
	}
	type SelfHttpError interface {
		HttpError(w http.ResponseWriter)
	}
	switch t := err.(type) {
	default:
		w.WriteHeader(status)
	case HttpError:
		err, code := t.HttpError()
		http.Error(w, err, code)
	case SelfHttpError:
		t.HttpError(w)
	}
}

func HandleHttpResponse(w http.ResponseWriter, r *http.Request, resp interface{}) {
	switch t := resp.(type) {
	default:
		// I don't know that type !
	case http.Handler:
		t.ServeHTTP(w, r)
	case []byte:
		w.Write(t)
	}
}
