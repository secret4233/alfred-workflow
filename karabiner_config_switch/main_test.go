package main

import "testing"

func Test_run(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "switch",
			args: args{
				args: []string{"", "switch", "-name", "mac"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			run(tt.args.args)
		})
	}
}
