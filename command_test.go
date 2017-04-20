package main

import (
	"os"
	"testing"
)

func TestCmdLookUp(t *testing.T) {

	// Create an env var for the test
	os.Setenv("UPTOWN", "This hit, that ice cold")

	fileLookUpCmd := Command{"lookup", []string{"./fixtures/test.ini", "UPTOWN"}}
	envLookUpCmd := Command{"lookup", []string{"ENV", "UPTOWN"}}
	badFileLookUpCmd := Command{"lookup", []string{"./fixtures/test.ini", "FUNK"}}
	badEnvLookUpCmd := Command{"lookup", []string{"ENV", "FUNK"}}
	fileDoesntExistsLookUpCmd := Command{"lookup", []string{"./fixtures/test-doesnt-exists.ini", "FUNK"}}

	type args struct {
		cmd Command
	}
	tests := []struct {
		name      string
		args      args
		wantValue string
		wantErr   bool
	}{
		{"Lookup in file (ini style)", args{fileLookUpCmd}, "This hit, that ice cold", false},
		{"Lookup in env ", args{envLookUpCmd}, "This hit, that ice cold", false},
		{"Bad lookup in file (ini style)", args{badFileLookUpCmd}, "", true},
		{"Lookup in env ", args{badEnvLookUpCmd}, "", true},
		{"File doesn't exists (ini style)", args{fileDoesntExistsLookUpCmd}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValue, err := CmdLookUp(tt.args.cmd)
			if (err != nil) != tt.wantErr {
				t.Errorf("CmdLookUp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotValue != tt.wantValue {
				t.Errorf("CmdLookUp() = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}
