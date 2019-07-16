/*
 from https://gowebexamples.com/hello-world/
*/

package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world. My name is YUSUF!\n")
	})

	http.ListenAndServe(":8090", nil)
}
