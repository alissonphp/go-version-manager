package utils

import "net/http"

func MimeTypeChecker(next http.Handler) http.Handler  {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, h, err := r.FormFile("file")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if h.Header.Get("Content-Type")  != "application/x-sharedlib" {
			http.Error(w, "content-type not allowed", http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	})
}