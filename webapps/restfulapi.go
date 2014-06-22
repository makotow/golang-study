package main

import(
  "fmt"
  "github.com/drone/routes"
  "net/http"
)

// handler for GET
func getuser(w http.ResponseWriter, r *http.Request) {
  params := r.URL.Query()
  id := params.Get(":id")
  fmt.Fprintf(w, "you are get user %s", id)
}


func main() {
  mux := routes.New()

  mux.Get("/users/:id", getuser)

  http.Handle("/", mux)
  http.ListenAndServe(":9090", nil)
}
