package utils

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Testing with multiple numbers", args: args{nums: []int{1, 2, 3, 4}}, want: 10},
		{name: "Testing with single numbers", args: args{nums: []int{14}}, want: 14},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.nums); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMean(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "Testing with even number of arguments", args: args{nums: []int{1, 2, 3, 4}}, want: 2.5},
		{name: "Testing with odd number of arguments", args: args{nums: []int{1, 2, 3}}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Mean(tt.args.nums); got != tt.want {
				t.Errorf("Mean() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMedian(t *testing.T) {
	type args struct {
		nums []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "Testing with odd number of arguments", args: args{nums: []float64{3, 1, 2}}, want: 2},
		{name: "Testing with even number of arguments", args: args{nums: []float64{3, 1, 4, 2}}, want: 2.5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Median(tt.args.nums); got != tt.want {
				t.Errorf("Median() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWordFrequency(t *testing.T) {
	type args struct {
		words string
	}
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		{
			name: "Testing with a sample text",
			args: args{words: "test test and test"},
			want: map[string]int{
				"test": 3,
				"and":  1,
			},
		},
		{
			name: "Testing with a text containing special characters and Capital letters",
			args: args{words: "Test, test unit test. Test is important."},
			want: map[string]int{
				"test":      4,
				"unit":      1,
				"is":        1,
				"important": 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WordFrequency(tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WordFrequency() = %v, want %v", got, tt.want)
			}
		})
	}
}
