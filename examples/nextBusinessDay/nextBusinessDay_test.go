package nextBusinessDay

import (
	"testing"
	"time"
)

func TestNextBusinessDay(t *testing.T) {
	type args struct {
		date time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case with Friday",
			args: args{date: time.Date(2021, time.December, 31, 0, 0, 0, 0, time.UTC)},
			want: "Mon, 03 Jan 2022",
		},
		{
			name: "Test case with Tuesday",
			args: args{date: time.Date(2022, time.January, 4, 0, 0, 0, 0, time.UTC)},
			want: "Wed, 05 Jan 2022",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NextBusinessDay(tt.args.date); got.Format("Mon, 02 Jan 2006") != tt.want {
				t.Errorf("NextBusinessDay() = %v, want %v", got, tt.want)
			}
		})
	}
}
