package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"jerno.playground.com/examples/calculator"
	"jerno.playground.com/examples/httpClient"
	"jerno.playground.com/examples/httpServer"
	"jerno.playground.com/examples/mapMissingItem"
	"jerno.playground.com/examples/modifyVariableThroughPointer"
	"jerno.playground.com/examples/nextBusinessDay"
	"jerno.playground.com/examples/panicAndRecover"
	"jerno.playground.com/examples/receiverPassedByReference"
	"jerno.playground.com/examples/receiverPassedByValue"
	"jerno.playground.com/examples/timeAndDuration"
)

var reader = bufio.NewReader(os.Stdin)

type useCase struct {
	Icon string
	Name string
	Run  func()
}

func main() {
	useCases := []useCase{
		{
			Icon: "🌐",
			Name: "HTTP GET String data",
			Run: func() {
				httpClient.GetStringData()
			},
		},
		{
			Icon: "🌐",
			Name: "HTTP GET JSON data",
			Run: func() {
				httpClient.GetJsonData()
			},
		},
		{
			Icon: "🌐",
			Name: "HTTP POST JSON data",
			Run: func() {
				httpClient.SendJsonData()
			},
		},
		{
			Icon: "🏢",
			Name: "Start HTTP server",
			Run: func() {
				httpServer.StartServer()
			},
		},
		{
			Icon: "🕦",
			Name: "Get current time formatted",
			Run: func() {
				fmt.Printf("Current time formatted as ('01/02/2006'): %v\n", time.Now().Format("01/02/2006"))
			},
		},
		{
			Icon: "⌨️",
			Name: "Calculator",
			Run: func() {
				calculator.Run()
			},
		},
		{
			Icon: "📝",
			Name: "Receiver passed by value",
			Run: func() {
				receiverPassedByValue.Run()
			},
		},
		{
			Icon: "📄",
			Name: "Receiver passed by reference",
			Run: func() {
				receiverPassedByReference.Run()
			},
		},
		{
			Icon: "🔗",
			Name: "Modifying variable through a pointer",
			Run: func() {
				modifyVariableThroughPointer.Run()
			},
		},
		{
			Icon: "🔎",
			Name: "Map checking missing item (A.K.A. CommaOK)",
			Run: func() {
				mapMissingItem.Run()
			},
		},
		{
			Icon: "😱",
			Name: "Panic",
			Run: func() {
				panicAndRecover.Panic()
			},
		},
		{
			Icon: "🤕",
			Name: "Panic and recover",
			Run: func() {
				panicAndRecover.PanicWithRecover()
			},
		},
		{
			Icon: "📅",
			Name: "Get next business day",
			Run: func() {
				fmt.Printf("Next business day is: %v\n", nextBusinessDay.NextBusinessDay(time.Now()).Format("Mon, 02 Jan 2006"))
			},
		},
		{
			Icon: "⏳",
			Name: "Time and duration",
			Run: func() {
				timeAndDuration.Run()
			},
		},
	}

	userInput := promptUseCase(useCases)
	handleUserInput(userInput, useCases)
}

func promptUseCase(useCases []useCase) string {
	fmt.Println("=======================================================================")
	fmt.Println("Available use cases:")
	for i, useCase := range useCases {
		fmt.Printf("  %2d) %v %v\n", i+1, useCase.Icon, useCase.Name)
	}
	fmt.Print("Please select a use case (type its number or its name): ")

	useCaseString, _ := reader.ReadString('\n')

	useCaseString = strings.TrimSpace(useCaseString)

	useCaseNumber, err := strconv.Atoi(useCaseString) // Assuming option is of type string
	if err != nil {
		return useCaseString
	}

	if useCaseNumber >= 1 && useCaseNumber <= len(useCases) {
		return useCaseString
	}

	return ""
}

func handleUserInput(userInput string, useCases []useCase) {
	useCaseNumber, err := strconv.Atoi(userInput)
	if err != nil {
		fmt.Printf("Unknown use-case: %v\n", userInput)
		return
	}

	useCase := useCases[useCaseNumber-1]

	printUsecaseTitle(useCase)
	useCase.Run()
	fmt.Println("")
}

func printUsecaseTitle(useCase useCase) {
	totalNumberOfWhiteSpaces := 71 - 2*3 - 2 - 1 - len(useCase.Name)
	fmt.Printf("=======================================================================\n")
	fmt.Printf("=  %v%v %v%v  =\n", strings.Repeat(" ", totalNumberOfWhiteSpaces/2), useCase.Icon, useCase.Name, strings.Repeat(" ", totalNumberOfWhiteSpaces-totalNumberOfWhiteSpaces/2))
	fmt.Printf("=======================================================================\n")
	fmt.Printf("\n")
}
