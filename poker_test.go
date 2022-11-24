package main

import (
	"reflect"
	"testing"
)

func Test_analyze(t *testing.T) {
	type args struct {
		hand []card
	}
	tests := []struct {
		name string
		args args
		want []combination
	}{
		{
			name: "pair",
			args: args{hand: []card{"4r", "4h", "5r", "6h", "8k"}},
			want: []combination{{
				name:  "pair",
				cards: []card{"4r", "4h"},
			}},
		}, {
			name: "three of a kind",
			args: args{hand: []card{"4r", "4h", "4k", "6h", "8k"}},
			want: []combination{{
				name:  "three of a kind",
				cards: []card{"4r", "4h", "4k"},
			}},
		}, {
			name: "four of a kind",
			args: args{hand: []card{"tr", "th", "tk", "ts", "8k"}},
			want: []combination{{
				name:  "four of a kind",
				cards: []card{"tr", "th", "tk", "ts"},
			}},
		}, {
			name: "two pairs",
			args: args{hand: []card{"kk", "kh", "ds", "dr", "8k"}},
			want: []combination{{
				name:  "pair",
				cards: []card{"kk", "kh"},
			},{
				name:  "pair",
				cards: []card{"ds", "dr"},
			},{
				name:  "two pairs",
				cards: []card{"kk", "kh", "ds", "dr"},
			}},
		}, {
			name: "full house",
			args: args{hand: []card{"3k", "3h", "as", "ak", "ar"}},
			want: []combination{{
				name:  "pair",
				cards: []card{"3k", "3h"},
			},{
				name:  "three of a kind",
				cards: []card{"as", "ak", "ar"},
			},{
				name:  "full house",
				cards: []card{"3k", "3h", "as", "ak", "ar"},
			}},
		}, {
			name: "flush",
			args: args{hand: []card{"5h", "8h", "th", "ah", "dh"}},
			want: []combination{{
				name:  "flush",
				cards: []card{"5h", "8h", "th", "dh", "ah"},
			}},
		}, {
			name: "straight",
			args: args{hand: []card{"7r", "8h", "9k", "ts", "rr"}},
			want: []combination{{
				name:  "straight",
				cards: []card{"7r", "8h", "9k", "ts", "rr"},
			}},
		}, {
			name: "high straight",
			args: args{hand: []card{"tr", "rh", "dk", "kh", "as"}},
			want: []combination{{
				name:  "high straight",
				cards: []card{"tr", "rh", "dk", "kh", "as"},
			}},
		}, {
			name: "royal flush",
			args: args{hand: []card{"th", "rh", "dh", "kh", "ah"}},
			want: []combination{{
				name:  "flush",
				cards: []card{"th", "rh", "dh", "kh", "ah"},
			}, {
				name:  "high straight",
				cards: []card{"th", "rh", "dh", "kh", "ah"},
			}, {
				name:  "royal flush",
				cards: []card{"th", "rh", "dh", "kh", "ah"},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := analyze(tt.args.hand); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("analyze() = %v, want %v", got, tt.want)
			}
		})
	}
}
