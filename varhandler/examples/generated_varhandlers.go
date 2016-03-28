// Code generated by "handler -func Simple,Import,Status,Response,ResponseStatus,CreateUser,GetUser,UpdateUser,DeleteUser"; DO NOT EDIT

package main

import "net/http"
import z "github.com/azr/generators/varhandler/examples/z"

func SimpleHandler(w http.ResponseWriter, r *http.Request) {
	var err error

	param0, err := HTTPX(r)
	if err != nil {
		HandleHttpErrorWithDefaultStatus(w, r, http.StatusBadRequest, err)
		return
	}

	param1, err := HTTPY(r)
	if err != nil {
		HandleHttpErrorWithDefaultStatus(w, r, http.StatusBadRequest, err)
		return
	}

	param2, err := HTTPZ(r)
	if err != nil {
		HandleHttpErrorWithDefaultStatus(w, r, http.StatusBadRequest, err)
		return
	}

	err = Simple(param0, param1, param2)
	if err != nil {
		HandleHttpErrorWithDefaultStatus(w, r, http.StatusInternalServerError, err)
		return
	}

}

func ImportHandler(w http.ResponseWriter, r *http.Request) {
	var err error

	param0, err := HTTPX(r)
	if err != nil {
		HandleHttpErrorWithDefaultStatus(w, r, http.StatusBadRequest, err)
		return
	}

	param1, err := HTTPY(r)
	if err != nil {
		HandleHttpErrorWithDefaultStatus(w, r, http.StatusBadRequest, err)
		return
	}

	param2, err := HTTPZ(r)
	if err != nil {
		HandleHttpErrorWithDefaultStatus(w, r, http.StatusBadRequest, err)
		return
	}

	param3, err := z.HTTPZ(r)
	if err != nil {
		HandleHttpErrorWithDefaultStatus(w, r, http.StatusBadRequest, err)
		return
	}

	err = Import(param0, param1, param2, param3)
	if err != nil {
		HandleHttpErrorWithDefaultStatus(w, r, http.StatusInternalServerError, err)
		return
	}

}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	var err error

	param0, err := HTTPX(r)
	if err != nil {
		HandleHttpErrorWithDefaultStatus(w, r, http.StatusBadRequest, err)
		return
	}

	param1, err := HTTPY(r)
	if err != nil {
		HandleHttpErrorWithDefaultStatus(w, r, http.StatusBadRequest, err)
		return
	}

	param2, err := HTTPZ(r)
	if err != nil {
		HandleHttpErrorWithDefaultStatus(w, r, http.StatusBadRequest, err)
		return
	}

	var status int

	status, err = Status(param0, param1, param2)
	if err != nil {
		HandleHttpErrorWithDefaultStatus(w, r, http.StatusInternalServerError, err)
		return
	}

	if status != 0 {
		w.WriteHeader(status)
	}

}

func ResponseHandler(w http.ResponseWriter, r *http.Request) {
	var err error

	param0, err := HTTPX(r)
	if err != nil {
		HandleHttpErrorWithDefaultStatus(w, r, http.StatusBadRequest, err)
		return
	}

	param1, err := HTTPY(r)
	if err != nil {
		HandleHttpErrorWithDefaultStatus(w, r, http.StatusBadRequest, err)
		return
	}

	param2, err := HTTPZ(r)
	if err != nil {
		HandleHttpErrorWithDefaultStatus(w, r, http.StatusBadRequest, err)
		return
	}

	var resp interface{}

	resp, err = Response(param0, param1, param2)
	if err != nil {
		HandleHttpErrorWithDefaultStatus(w, r, http.StatusInternalServerError, err)
		return
	}

	if resp != nil {
		HandleHttpResponse(w, r, resp)
	}

}

func ResponseStatusHandler(w http.ResponseWriter, r *http.Request) {
	var err error

	param0, err := HTTPX(r)
	if err != nil {
		HandleHttpErrorWithDefaultStatus(w, r, http.StatusBadRequest, err)
		return
	}

	param1, err := HTTPY(r)
	if err != nil {
		HandleHttpErrorWithDefaultStatus(w, r, http.StatusBadRequest, err)
		return
	}

	param2, err := HTTPZ(r)
	if err != nil {
		HandleHttpErrorWithDefaultStatus(w, r, http.StatusBadRequest, err)
		return
	}

	var resp interface{}

	var status int

	resp, status, err = ResponseStatus(param0, param1, param2)
	if err != nil {
		HandleHttpErrorWithDefaultStatus(w, r, http.StatusInternalServerError, err)
		return
	}

	if status != 0 {
		w.WriteHeader(status)
	}

	if resp != nil {
		HandleHttpResponse(w, r, resp)
	}

}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var err error

	param0, err := HTTPUser(r)
	if err != nil {
		HandleHttpErrorWithDefaultStatus(w, r, http.StatusBadRequest, err)
		return
	}

	var status int

	status, err = CreateUser(param0)
	if err != nil {
		HandleHttpErrorWithDefaultStatus(w, r, http.StatusInternalServerError, err)
		return
	}

	if status != 0 {
		w.WriteHeader(status)
	}

}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	var err error

	param0, err := HTTPUserID(r)
	if err != nil {
		HandleHttpErrorWithDefaultStatus(w, r, http.StatusBadRequest, err)
		return
	}

	var resp interface{}

	var status int

	resp, status, err = GetUser(param0)
	if err != nil {
		HandleHttpErrorWithDefaultStatus(w, r, http.StatusInternalServerError, err)
		return
	}

	if status != 0 {
		w.WriteHeader(status)
	}

	if resp != nil {
		HandleHttpResponse(w, r, resp)
	}

}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	var err error

	param0, err := HTTPUserID(r)
	if err != nil {
		HandleHttpErrorWithDefaultStatus(w, r, http.StatusBadRequest, err)
		return
	}

	param1, err := HTTPUser(r)
	if err != nil {
		HandleHttpErrorWithDefaultStatus(w, r, http.StatusBadRequest, err)
		return
	}

	var status int

	status, err = UpdateUser(param0, param1)
	if err != nil {
		HandleHttpErrorWithDefaultStatus(w, r, http.StatusInternalServerError, err)
		return
	}

	if status != 0 {
		w.WriteHeader(status)
	}

}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	var err error

	param0, err := HTTPUserID(r)
	if err != nil {
		HandleHttpErrorWithDefaultStatus(w, r, http.StatusBadRequest, err)
		return
	}

	var status int

	status, err = DeleteUser(param0)
	if err != nil {
		HandleHttpErrorWithDefaultStatus(w, r, http.StatusInternalServerError, err)
		return
	}

	if status != 0 {
		w.WriteHeader(status)
	}

}

func HandleHttpErrorWithDefaultStatus(w http.ResponseWriter, r *http.Request, status int, err error) {
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
	case http.Handler:
		t.ServeHTTP(w, r)
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
