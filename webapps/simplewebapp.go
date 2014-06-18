package main

// http://www.slideshare.net/eeozekin/fast-web-applications-with-go
// を写経 slide 7
import (
	"fmt"
	"net/http"
)

func get_name() (string, string) {
	// No reason to brew a few rules, right
	var hello = "Hello"
	audience := "DevFestTR"
	return hello, audience
}

func handler(writer http.ResponseWriter, request *http.Request) {
	hello, audience := get_name()
	fmt.Fprintf(writer, hello+audience)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
