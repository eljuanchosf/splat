package splat

import (
	"io/ioutil"
	"os"
	"testing"
)

func readFile(file string) string {
	content, _ := ioutil.ReadFile(getAbsPath(file))
	return string(content)
}

func TestCmdLookUp(t *testing.T) {

	// Create an env var for the test
	os.Setenv("UPTOWN", "This hit, that ice cold")

	fileLookUpCmd := []cmdArg{{0, "./fixtures/test.ini", true}, {1, "UPTOWN", false}}
	envLookUpCmd := []cmdArg{{0, "ENV", false}, {1, "UPTOWN", false}}
	badFileLookUpCmd := []cmdArg{{0, "./fixtures/test.ini", true}, {1, "FUNK", false}}
	badEnvLookUpCmd := []cmdArg{{0, "ENV", false}, {1, "FUNK", false}}
	fileDoesntExistsLookUpCmd := []cmdArg{{0, "./fixtures/test-doesnt-exists.ini", false}, {1, "FUNK", false}}

	type args struct {
		cmdArgs []cmdArg
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
			gotValue, err := CmdLookUp(tt.args.cmdArgs)
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
	emptyFile := []cmdArg{{0, "./fixtures/an-empty-file.txt", true}}
	goodFile := []cmdArg{{0, "./fixtures/a-good-file.txt", true}}
	nonExistingFile := []cmdArg{{0, "./fixtures/file-not-exists.txt", false}}
	oneBigFile := []cmdArg{{0, "./fixtures/good-yaml.yml", true}}
	type args struct {
		cmdArgs []cmdArg
	}
	tests := []struct {
		name      string
		args      args
		wantValue string
		wantErr   bool
	}{
		{"Read empty file", args{emptyFile}, "", false},
		{"Read a file", args{goodFile}, readFile(goodFile[0].value), false},
		{"Non existing file", args{nonExistingFile}, "", true},
		{"One big file", args{oneBigFile}, readFile(oneBigFile[0].value), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValue, err := CmdFileContent(tt.args.cmdArgs)
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
	currentPath := getAbsPath(".")
	validCommand := []cmdArg{{0, "cat", false}, {1, getAbsPath("./fixtures/a-good-file.txt"), true}}
	validCommand2 := []cmdArg{{0, "pwd", false}}
	validCommand3 := []cmdArg{{0, "echo", false}, {1, "THIS HIT, THIS ICE COLD", false}}
	invalidCommand := []cmdArg{{0, "this-command-is-inexistent", false}}
	type args struct {
		cmdArgs []cmdArg
	}
	tests := []struct {
		name      string
		args      args
		wantValue string
		wantErr   bool
	}{
		{"A command with a file argument", args{validCommand}, readFile(validCommand[1].value), false},
		{"A command with no arguments", args{validCommand2}, currentPath, false},
		{"A command with a text argument", args{validCommand3}, "THIS HIT, THIS ICE COLD", false},
		{"An invalid command", args{invalidCommand}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValue, err := CmdRun(tt.args.cmdArgs)
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

// func TestCmdCertificate(t *testing.T) {
// 	type args struct {
// 		cmd Command
// 	}
// 	tests := []struct {
// 		name      string
// 		args      args
// 		wantValue string
// 		wantErr   bool
// 	}{
// 	// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			gotValue, err := CmdCertificate(tt.args.cmd)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("CmdCertificate() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if gotValue != tt.wantValue {
// 				t.Errorf("CmdCertificate() = %v, want %v", gotValue, tt.wantValue)
// 			}
// 		})
// 	}
// }
