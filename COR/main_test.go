package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	data := []int{1, 2}
	targetChan := Power(data...)
	counter := len(data)
	result := 0
	type args struct {
		powered <-chan int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "sum should be 5",
			args: args{powered: targetChan},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			for x := range Sum(tt.args.powered) {
				counter--
				result += x
				if counter == 0 && x != tt.want {
					t.Errorf("Sum() = %v, want %v", x, tt.want)
				}
			}
		})
	}
}
