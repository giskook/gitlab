package base

import "testing"

func TestBinarySearchInsert1(t *testing.T) {
	type args struct {
		sortedArray []int
		target      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"normal", args{sortedArray: []int{1,2,3,5}, target: 4}, 3},
		{"begin", args{sortedArray: []int{2,3,3}, target: 1}, 0},
		{"end", args{sortedArray: []int{1,2,3,3}, target: 4}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearchInsert(tt.args.sortedArray, tt.args.target); got != tt.want {
				t.Errorf("BinarySearchInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}