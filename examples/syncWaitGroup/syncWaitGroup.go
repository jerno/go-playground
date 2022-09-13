package syncWaitGroup

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"time"

	justHttp "github.com/jerno/just-http/json"
	"jerno.playground.com/utils"
)

func Run() {
	server := createServer(200)
	runSyncExample(server)
	fmt.Printf("\n\n")
	runAsyncExample(server)
}

func runSyncExample(server *httptest.Server) {
	defer utils.StopWatchLogger("Requesting users syncronously")()

	fmt.Printf("Requesting users syncronously...\n")
	for i := 0; i < 5; i++ {
		userId := fmt.Sprintf("%d", i+1)
		url := server.URL + "/" + "valid-url/" + userId
		user, err := getUserDetails(url)
		if err != nil {
			fmt.Printf("Error occured while downloading details for user#%s\n", userId)
			fmt.Println(err)
			return
		}
		fmt.Printf("User name: %v\n", user.UserName)
	}
	fmt.Println("All users downloaded")
}

func runAsyncExample(server *httptest.Server) {
	defer utils.StopWatchLogger("Requesting users asyncronously")()

	fmt.Printf("Requesting users asyncronously...\n")
	userIdChannel := make(chan string)
	go func() {
		for i := 0; i < 5; i++ {
			userId := fmt.Sprintf("%d", i+1)
			userIdChannel <- userId
		}
		close(userIdChannel)
	}()

	getAllUserDetailsWaitGroup(server, userIdChannel)
	fmt.Println("All users downloaded")
}

func getAllUserDetailsWaitGroup(server *httptest.Server, userIds <-chan string) {
	var wg sync.WaitGroup

	for userId := range userIds {
		wg.Add(1)
		go func(userId string) {
			defer wg.Done()
			url := server.URL + "/" + "valid-url/" + userId
			user, err := getUserDetails(url)
			if err != nil {
				fmt.Printf("Error occured while downloading details for user#%s\n", userId)
				fmt.Println(err)
				return
			}
			fmt.Printf("User name: %v\n", user.UserName)
		}(userId)
	}

	wg.Wait()
}

func getUserDetails(url string) (*userDetailsResponse, error) {
	var userDetails userDetailsResponse
	err := justHttp.Get(url, &userDetails)
	if err != nil {
		return nil, err
	}
	return &userDetails, nil
}

func createServer(delayInMilliseconds int) *httptest.Server {
	fmt.Printf("Creating server with delay: %s\n", time.Duration(delayInMilliseconds)*time.Millisecond)

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		time.Sleep(time.Duration(delayInMilliseconds) * time.Millisecond)
		if strings.HasPrefix(req.URL.Path, "/valid-url/") {
			segments := strings.Split(req.URL.Path, "/")
			id := segments[len(segments)-1]
			rw.Write([]byte(fmt.Sprintf(`{"userName": "User %s", "userId": %s}`, id, id)))
		}
		if req.URL.Path == "/valid-post-url" {
			rw.Write([]byte(`{"Cluster_name": "server cluster", "Pings": 202}`))
		}
		if req.URL.Path == "/internal-server-error" {
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Header().Set("Content-Type", "application/json")
			resp := make(map[string]string)
			resp["message"] = "Some Error Occurred"
			jsonResp, err := json.Marshal(resp)
			if err != nil {
				log.Fatalf("Error happened in JSON marshal. Err: %s", err)
			}
			rw.Write(jsonResp)
		}
	}))
	return server
}

type userDetailsResponse struct {
	UserName string `json:"userName"`
	UserId   int    `json:"userId"`
}
