package splat

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

func TestCmdFileContent(t *testing.T) {
	emptyFile := Command{"fileContent", []string{"./fixtures/an-empty-file.txt"}}
	goodFile := Command{"fileContent", []string{"./fixtures/a-good-file.txt"}}
	nonExistingFile := Command{"fileContent", []string{"./fixtures/file-not-exists.txt"}}
	type args struct {
		cmd Command
	}
	tests := []struct {
		name      string
		args      args
		wantValue string
		wantErr   bool
	}{
		{"Read empty file", args{emptyFile}, "", false},
		{"Read a file", args{goodFile}, "THIS HIT, THAT ICE COLD", false},
		{"Non existing file", args{nonExistingFile}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValue, err := CmdFileContent(tt.args.cmd)
			if (err != nil) != tt.wantErr {
				t.Errorf("CmdFileContent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotValue != tt.wantValue {
				t.Errorf("CmdFileContent() = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}

func TestCmdRun(t *testing.T) {
	type args struct {
		cmd Command
	}
	tests := []struct {
		name      string
		args      args
		wantValue string
		wantErr   bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValue, err := CmdRun(tt.args.cmd)
			if (err != nil) != tt.wantErr {
				t.Errorf("CmdRun() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotValue != tt.wantValue {
				t.Errorf("CmdRun() = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}

func TestCmdCertificate(t *testing.T) {
	type args struct {
		cmd Command
	}
	tests := []struct {
		name      string
		args      args
		wantValue string
		wantErr   bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValue, err := CmdCertificate(tt.args.cmd)
			if (err != nil) != tt.wantErr {
				t.Errorf("CmdCertificate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotValue != tt.wantValue {
				t.Errorf("CmdCertificate() = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}
