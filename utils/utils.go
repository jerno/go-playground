package utils

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
	"time"
)

func Mean(nums []int) float64 {
	s := Sum(nums)
	return float64(s) / float64(len(nums))
}

func Sum(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}

	return total
}

func Median(nums []float64) float64 {
	// Copy the slice, so the sort won't alter the original
	vals := make([]float64, len(nums))
	copy(vals, nums)
	sort.Float64s(vals)

	i := len(vals) / 2
	if len(vals)%2 == 1 {
		return vals[i]
	} else {
		return (vals[i-1] + vals[i]) / 2
	}
}

// Counts how many times a word appears in a text
// It trims the text, and removes all special characters
func WordFrequency(input string) map[string]int {
	reg, _ := regexp.Compile("[^a-zA-Z0-9 ]+")
	trimmed := reg.ReplaceAllString(input, " ")

	words := strings.Fields(trimmed)

	freq := make(map[string]int)
	for _, w := range words {
		freq[strings.ToLower(w)]++
	}

	return freq
}

func StopWatchLogger(name string) func() {
	start := time.Now()

	return func() {
		duration := time.Since(start)
		fmt.Printf("%s took %s", name, duration)
	}
}
