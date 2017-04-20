package main

import (
	"testing"
)

func Test_runCommand(t *testing.T) {
	type args struct {
		command Command
	}
	tests := []struct {
		name       string
		args       args
		wantResult string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := runCommand(tt.args.command); gotResult != tt.wantResult {
				t.Errorf("runCommand() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
