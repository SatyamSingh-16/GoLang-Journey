func LogginMiddleware(next http.HandleFunc) http.HandleFunc{
	return func(w http.ResponseWriter, r*http.Request){

		fmt.PrintLn("Method :", r.Mehtod)
		fmt.PrintLn("Path :", r.URL.Path)
		next(w,r)
	}
}

