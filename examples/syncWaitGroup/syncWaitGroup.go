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

const SPACING = "  "

func Run() {
	server := createServer(200)
	runSyncExample(server)
	fmt.Printf("\n\n")
	runAsyncExample(server)
}

func runSyncExample(server *httptest.Server) {
	defer utils.StopWatchLogger("Requesting users syncronously")()

	fmt.Printf("\nRequesting users syncronously\n\n")

	for i := 0; i < 5; i++ {
		fmt.Printf("[Provider] Generating value...\n")
		userId := i + 1
		fmt.Printf("[Provider] Generated value: %v\n", userId)

		sendRequest(server, userId)
	}
	fmt.Println("All users downloaded")
}

func runAsyncExample(server *httptest.Server) {
	defer utils.StopWatchLogger("Requesting users asyncronously")()

	fmt.Printf("\nRequesting users asyncronously\n\n")

	userIdChannel := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			userId := i + 1
			fmt.Printf("[Provider] %s | Generated value: %v\n", strings.Repeat(SPACING, userId), userId)

			fmt.Printf("[Provider] %s | Sending value to channel: %v | BLOCKING\n", strings.Repeat(SPACING, userId), userId)
			userIdChannel <- userId
			fmt.Printf("[Provider] %s | Value sent to channel: %v | UNBLOCKED\n", strings.Repeat(SPACING, userId), userId)
		}
		close(userIdChannel)
	}()

	var wg sync.WaitGroup

	fmt.Printf("[Consumer] Reading the channel...\n")
	for userId := range userIdChannel {
		fmt.Printf("[Consumer] %s | Value in channel arrived: %v\n", strings.Repeat(SPACING, userId), userId)
		wg.Add(1)

		fmt.Printf("[Consumer] %s | Go-routine defined\n", strings.Repeat(SPACING, userId))
		go func(userId int) {
			fmt.Printf("[Consumer] %s | Go-routine runing\n", strings.Repeat(SPACING, userId))
			defer wg.Done()
			sendRequest(server, userId)
		}(userId)
	}

	fmt.Printf("[Consumer] Waiting for WaitGroup...\n")
	wg.Wait()
	fmt.Printf("[Consumer] Waiting finished...\n")

	fmt.Println("All users downloaded")
}

func sendRequest(server *httptest.Server, userId int) {
	fmt.Printf("[Consumer] %s | Sending request...\n", strings.Repeat(SPACING, userId))
	url := fmt.Sprintf("%s/valid-url/%d", server.URL, userId)
	user, err := getUserDetails(url)
	if err != nil {
		fmt.Printf("[Consumer] %s | Error occured while downloading details for user#%v\n", strings.Repeat(SPACING, userId), userId)
		fmt.Println(err)
		return
	}
	fmt.Printf("[Consumer] %s | User name: %v\n", strings.Repeat(SPACING, userId), user.UserName)
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
