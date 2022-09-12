package main

import (
	"testing"
)

func Test_distanceCalc(t *testing.T) {
	type args struct {
		t float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "1 sec", args: struct{ t float64 }{t: 1}, want: 0.21306818181818182},
		{name: "10 sec", args: struct{ t float64 }{t: 10}, want: 2.1306818181818183},
		{name: "666 sec", args: struct{ t float64 }{t: 666}, want: 141.9034090909091},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distanceCalc(tt.args.t); got != tt.want {
				t.Errorf("distanceCalc() = %v, want %v", got, tt.want)
			}
		})
	}
}
