

func ValidationMiddleware(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		apiKey := r.Header.Get("X-API-Key")

		if apiKey == "" {

			http.Error(
				w,
				"API Key Missing",
				http.StatusUnauthorized,
			)

			return

		}

		next(w, r)

	}

}