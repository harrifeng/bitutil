package bitutil

import "testing"

func TestTwoComplementInt8ToRaw(t *testing.T) {
	type args struct {
		num int8
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test 1", args{1}, "00000001"},
		{"Test minus 1", args{-1}, "11111111"},
		{"Test 0", args{0}, "00000000"},
		{"Test minus128", args{-128}, "10000000"},
		{"Test minus121", args{-121}, "10000111"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TwoComplementInt8ToRaw(tt.args.num); got != tt.want {
				t.Errorf("TwoComplementInt8ToRaw() = %v, want %v", got, tt.want)
			}
		})
	}
}
