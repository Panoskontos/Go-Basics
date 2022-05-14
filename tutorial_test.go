package main

import "testing"

func Test_mymath(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mymath(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("mymath() = %v, want %v", got, tt.want)
			}
		})
	}
}
