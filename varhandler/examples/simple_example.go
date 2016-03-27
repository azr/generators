package main

import "net/http"

func Simple(x X, y Y, z *Z) error {
	return nil
}

func HTTPX(w http.ResponseWriter, r *http.Request) (X, error)  { return X{}, nil }
func HTTPY(w http.ResponseWriter, r *http.Request) (Y, error)  { return Y{}, nil }
func HTTPZ(w http.ResponseWriter, r *http.Request) (*Z, error) { return &Z{}, nil }

type X struct {
}
type Y struct {
}
type Z struct {
}
