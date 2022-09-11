package readerPromptWithDefault

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/fatih/color"
)

type ReaderWrapper struct {
	*bufio.Reader
}

func (reader *ReaderWrapper) PromptWithDefault(hint string, defaultValue string) (string, error) {
	fmt.Printf("%s", hint)
	fmtStyled := color.New(color.FgHiYellow)
	fmtStyled.Printf(" (default: %s)", defaultValue)
	fmt.Printf(": ")

	input, readError := reader.ReadString('\n')
	if readError != nil {
		return "", readError
	}
	input = strings.TrimSpace(input)

	if len(input) > 0 {
		return input, nil
	} else {
		return defaultValue, nil
	}
}
