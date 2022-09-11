package httpClient

import (
	"fmt"

	justHttp "github.com/jerno/just-http/basic"
	justHttpJson "github.com/jerno/just-http/json"
)

type SampleHttpBinData struct {
	Url    string `json:"url"`
	Origin string `json:"origin"`
	Args   any    `json:"args"`
	Json   any    `json:"json"`
}

type SampleJson struct {
	Cluster_name string `json:"Cluster_name"`
	Pings        int    `json:"Pings"`
}

func GetStringData() {
	url := "https://httpbin.org/get"

	fmt.Printf("Sending GET request to: %v\n", url)

	res, err := justHttp.GetString(url)
	if err != nil {
		fmt.Println("error: can't call httpbin.org")
	} else {
		fmt.Printf(res)
	}
}

func GetJsonData() {
	url := "https://httpbin.org/get"
	queryParams := map[string]string{"MyParam": "test"}

	fmt.Printf("Sending GET request to: %v\n", url)
	fmt.Printf("  - params: %v\n", queryParams)

	var sample SampleHttpBinData
	err := justHttpJson.Get(url, &sample, justHttpJson.RequestArguments{QueryParams: queryParams})
	if err != nil {
		fmt.Println("error: can't call httpbin.org")
	} else {
		fmt.Printf("%#v\n", sample)
	}
}

func SendJsonData() {
	url := "https://httpbin.org/post"
	data := SampleJson{Cluster_name: "Hello server", Pings: 1}
	queryParams := map[string]string{"MyParam": "test"}

	fmt.Printf("Sending POST request to: %v\n", url)
	fmt.Printf("  - params: %v\n", queryParams)
	fmt.Printf("  - payload: %#v\n", data)

	var sample SampleHttpBinData
	err := justHttpJson.Post(url, data, &sample, justHttpJson.RequestArguments{QueryParams: queryParams})
	if err != nil {
		fmt.Printf("Can't call httpbin.org\n")
		fmt.Println(err)
	} else {
		fmt.Printf("%#v\n", sample)
	}
}
