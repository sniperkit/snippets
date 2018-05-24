package main

import (
	"log"
	"fmt"
	"net/http"
	"encoding/json"

	"github.com/xor-gate/go-snippets/http-rest/acl"
)

type Handler func(http.ResponseWriter, *http.Request, *acl.Acl)

func permissionInfoHandler(w http.ResponseWriter, r *http.Request, auth *acl.Acl) {
	r.Header.Set("Content-type", "text/json")
}

func tokenInfoHandler(w http.ResponseWriter, r *http.Request, auth *acl.Acl) {
	r.Header.Set("Content-type", "text/json")
	w.Write(auth.ReqToken.Bytes())
}

func tokenListHandler(w http.ResponseWriter, r *http.Request, auth *acl.Acl) {
	r.Header.Set("Content-type", "text/json")
	w.Write(auth.TokenList.Bytes())
}

func tokenCreateHandler(w http.ResponseWriter, r *http.Request, auth *acl.Acl) {
	r.Header.Set("Content-type", "text/json")
	t := auth.NewToken()
	w.Write(t.Bytes())
}

func tokenDeleteHandler(w http.ResponseWriter, r *http.Request, auth *acl.Acl) {
	r.Header.Set("Content-type", "text/json")
	var data map[string]string
	decoder := json.NewDecoder(r.Body)
	_ = decoder.Decode(&data)
	fmt.Printf("%+v", data)
	_ = auth.DeleteToken(data["id"])
}

// Authenticate the request with
// X-Xg-Auth-User
// X-Xg-Api-Token
func Authenticate(acl acl.AclStore, next Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		user := r.Header.Get("X-Xg-Auth-User")
		token := r.Header.Get("X-Xg-Api-Token")

		auth := acl.Auth(user, token, r.URL.Path)

		if auth == nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized\n"))
			return;
		}

		next(w, r, auth);
	})
}

func main() {
	mux := http.NewServeMux()
	acl, _ := acl.New("localhost:6379")

	mux.Handle("/api/v1/permission/info", Authenticate(acl, permissionInfoHandler))

	mux.Handle("/api/v1/token/info",   Authenticate(acl, tokenInfoHandler))
	mux.Handle("/api/v1/token/list",   Authenticate(acl, tokenListHandler))
	mux.Handle("/api/v1/token/create", Authenticate(acl, tokenCreateHandler))
	mux.Handle("/api/v1/token/delete", Authenticate(acl, tokenDeleteHandler))

	log.Println("Listening *:3000")
	http.ListenAndServe(":3000", mux)
}
