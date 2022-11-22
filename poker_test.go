package main

import (
	"reflect"
	"testing"
)

func Test_analyyyze(t *testing.T) {
	type args struct {
		hand []card
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "2 pairs",
			args: args{
				hand: []card{"4r", "4h", "5r", "6h", "8k"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			analyze(tt.args.hand)
		})
	}
}

func Test_analyze(t *testing.T) {
	type args struct {
		hand []card
	}
	tests := []struct {
		name string
		args args
		want []analysis
	}{
		{
			name: "2 pairs",
			args: args{hand: []card{"4r", "4h", "5r", "6h", "8k"}},
			want: []analysis{{
				goal:  "2-PAIR",
				cards: []card{"4r", "4h"},
			}},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := analyze(tt.args.hand); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("analyze() = %v, want %v", got, tt.want)
			}
		})
	}
}
