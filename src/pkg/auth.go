package pkg

import (
	"encoding/base64"
	"log"
	"net/http"
	"strings"
)

func validate(username, password string) bool {
	if username == "testuser" && password == "testpassword" {
		return true
	}

	return false
}

// Authantication - basic authantication function
func Authantication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
		if len(auth) != 2 || auth[0] != "Basic" {
			http.Error(w, "Authorization Failed", http.StatusUnauthorized)
			return
		}

		payload, err := base64.StdEncoding.DecodeString(auth[1])
		if err != nil {
			log.Println(err)
			return
		}

		pair := strings.SplitN(string(payload), ":", 2)
		if len(pair) != 2 || !validate(pair[0], pair[1]) {
			http.Error(w, "Authoruzation Failed", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}
