package main

import "testing"

func Test_check(t *testing.T) {
	type args struct {
		dir string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"asdf", args{""}, false},
		{"asdf", args{"."}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := check(tt.args.dir); got != tt.want {
				t.Errorf("check() = %v, want %v", got, tt.want)
			}
		})
	}
}
