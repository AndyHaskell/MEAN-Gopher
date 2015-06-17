package main

import "net/http"

func FileServerRoute(mux *http.ServeMux, path, dir string) {
	mux.Handle(path, http.StripPrefix(path, http.FileServer(http.Dir(dir))))
}

func main() {
	mux := http.NewServeMux()

	//imgServer := http.StripPrefix("/img/",
	//	http.FileServer(http.Dir("public/images")))
	//mux.Handle("/img/", imgServer)

	FileServerRoute(mux, "/img/", "public/images")

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "pages/index.html")
	})

	server := &http.Server{
		Addr:    ":1123",
		Handler: mux,
	}

	server.ListenAndServe()
}
