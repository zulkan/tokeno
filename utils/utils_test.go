package utils

import "testing"

func TestInSlice(t *testing.T) {
	type args[T comparable] struct {
		data []T
		list []T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want bool
	}
	tests := []testCase[int]{
		{name: "number doesnt contains", args: args[int]{
			data: []int{3},
			list: []int{1, 2},
		}, want: false},
		{name: "number contains", args: args[int]{
			data: []int{3},
			list: []int{1, 2, 3, 4, 5},
		}, want: true},
		{name: "number contains", args: args[int]{
			data: []int{3, 1},
			list: []int{1, 2, 3, 4, 5},
		}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InSlice(tt.args.data, tt.args.list); got != tt.want {
				t.Errorf("InSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
