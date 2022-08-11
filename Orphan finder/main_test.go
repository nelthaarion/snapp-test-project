package main

import "testing"

func TestFindOrphan(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "should return 10",
			args: args{data: []int{10, 2, 3, 4, 5, 2, 3, 4, 5}},
			want: 10,
		},
		{
			name: "should return 100",
			args: args{data: []int{95, 95, 93, 94, 90, 90, 93, 94, 100}},
			want: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindOrphan(tt.args.data); got != tt.want {
				t.Errorf("FindOrphan() = %v, want %v", got, tt.want)
			}
		})
	}
}
