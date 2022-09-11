package httpServer

import (
	"fmt"
	"net/http"

	justHttp "github.com/jerno/just-http/basic"
)

var resp string

// AnimalHandler requests and stores a set of random animals
func AnimalHandler(w http.ResponseWriter, r *http.Request) {
	if resp == "" {
		url := "https://zoo-animal-api.herokuapp.com/animals/rand/10"
		response, err := justHttp.GetString(url)
		resp = response
		if err != nil {
			fmt.Println("error: can't call zoo-animal-api")
			return
		}
	}
	fmt.Fprintln(w, resp)
}

// AnimalRefreshHandler re-requests and stores a new set of random animals
func AnimalRefreshHandler(w http.ResponseWriter, r *http.Request) {
	resp = ""
	AnimalHandler(w, r)
}
