package httpServer

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fatih/color"
)

// healthHandler returns a server health
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}

type curlFn func(url string) string

type handlerMapping struct {
	Url         string
	CurlCmd     curlFn
	Handler     func() http.Handler
	HandlerFunc func(http.ResponseWriter, *http.Request)
}

func (h *handlerMapping) getUrl(host string) string {
	return fmt.Sprintf("%s%s", host, h.Url)
}

var curlGet = func(url string) string {
	return fmt.Sprintf(`curl %s`, url)
}

var curlPost = func(data string, url string) string {
	return fmt.Sprintf(`curl -d '%s' -H "Content-Type: application/json" -X POST %s`, data, url)
}

func StartServer() {
	endpoints := []handlerMapping{
		{
			Url:         "/health",
			CurlCmd:     func(url string) string { return curlGet(url) },
			HandlerFunc: HealthHandler,
		},
		{
			Url:         "/math",
			CurlCmd:     func(url string) string { return curlPost(`{"Left": 91, "Right": 8, "Op": "+"}`, url) },
			HandlerFunc: MathHandler,
		},
		{
			Url:         "/vue/",
			CurlCmd:     func(url string) string { return curlGet(url) },
			HandlerFunc: VueHandler,
		},
		{
			Url:         "/animals",
			CurlCmd:     func(url string) string { return curlGet(url) },
			HandlerFunc: AnimalHandler,
		},
		{
			Url:         "/animals/refresh",
			CurlCmd:     func(url string) string { return curlGet(url) },
			HandlerFunc: AnimalRefreshHandler,
		},
	}

	for _, endpoint := range endpoints {
		if endpoint.Handler != nil {
			http.Handle(endpoint.Url, endpoint.Handler())
		} else {
			http.HandleFunc(endpoint.Url, endpoint.HandlerFunc)
		}
	}

	port := "8081"
	host := "http://localhost:" + port

	fmt.Printf("server ready on %s\n", host)
	fmt.Println("Available endpoints:")
	for _, endpoint := range endpoints {
		fmt.Printf("  %s\n", endpoint.Url)
		fmtStyled := color.New(color.FgHiGreen)
		fmtStyled.Printf("    %s\n", endpoint.CurlCmd(endpoint.getUrl(host)))
	}

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
