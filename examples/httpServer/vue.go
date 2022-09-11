package httpServer

import (
	"fmt"
	"net/http"
	"regexp"
)

// VueHandler returns vue application
func VueHandler(w http.ResponseWriter, r *http.Request) {
	var jsxFile = regexp.MustCompile(`\\.jsx$`)
	ruri := r.RequestURI
	fmt.Printf("Serving %s\n", ruri)
	var fileserver = http.StripPrefix("/vue/", http.FileServer(http.Dir("static")))
	if jsxFile.MatchString(ruri) {
		w.Header().Set("Content-Type", "text/javascript")
	}
	fileserver.ServeHTTP(w, r)
}
