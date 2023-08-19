package main


import (
  "log"
  "net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
  if r.URL.Path != "/" {
    http.NotFound(w, r)
    return
  }
  w.Write([]byte("Hello from SnippetBox"))  
}

func snippetView(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Viewing snippets"))  
}
func snippetCreate(w http.ResponseWriter, r *http.Request) {

  if r.Method != http.MethodPost {
    w.Header().Set("Allow", http.MethodPost)
    http.Error(w, "Method not allowed!", http.StatusMethodNotAllowed)
    return
  }

  w.Write([]byte("Creating snippet"))  
}
func main() {
  mux := http.NewServeMux()
  mux.HandleFunc("/", home)
  mux.HandleFunc("/snippet/view", snippetView)
  mux.HandleFunc("/snippet/create", snippetCreate)
  

  log.Println("Starting server in port:4000")
  err := http.ListenAndServe(":4000", mux)
  log.Fatal(err) 
}
