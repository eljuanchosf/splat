package splat

import (
	"reflect"
	"testing"
)

var simpleCommand = Command{
	runner: "lookup",
	args: []cmdArg{
		{0, getAbsPath("./fixtures/a-good-file.txt"), true},
		{1, "aws_el", false},
	},
}

var complexCommand = Command{
	runner: "generateCert",
	args: []cmdArg{
		{0, "SHA256", false},
		{1, "selfSigned", false},
		{2, getAbsPath("./fixtures/a-good-file.txt"), true},
	},
}

func Test_extractCommand(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name        string
		args        args
		wantCommand Command
		wantErr     bool
	}{
		{"Simple example", args{"((< lookup(./fixtures/a-good-file.txt, aws_el) >))"}, simpleCommand, false},
		{"Complex example", args{"((< generateCert(SHA256, selfSigned, ./fixtures/a-good-file.txt) >))"}, complexCommand, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCommand, err := extractCommand(tt.args.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCommand, tt.wantCommand) {
				t.Errorf("extractCommand() = %v, want %v", gotCommand, tt.wantCommand)
			}
		})
	}
}

func Test_extractRunner(t *testing.T) {
	type args struct {
		cmd string
	}
	tests := []struct {
		name       string
		args       args
		wantRunner string
	}{
		{"Runner without parameters", args{"include()"}, "include"},
		{"Runner with parameters", args{"lookup(output.tf, aws_el)"}, "lookup"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRunner := extractRunner(tt.args.cmd); gotRunner != tt.wantRunner {
				t.Errorf("extractRunner() = %v, want %v", gotRunner, tt.wantRunner)
			}
		})
	}
}

func Test_extractRunnerArgs(t *testing.T) {
	testCmdArg := []cmdArg{{0, "output.tf", false}, {1, "aws_el", false}}
	testCmdArgWithFile := []cmdArg{{0, getAbsPath("./fixtures/an-empty-file.txt"), true}, {1, "aws_el", false}}
	type args struct {
		cmd string
	}
	tests := []struct {
		name     string
		args     args
		wantArgs []cmdArg
	}{
		{"Runner without parameters", args{"include()"}, []cmdArg{}},
		{"Runner with parameters", args{"lookup(output.tf , aws_el)"}, testCmdArg},
		{"Runner with parameters 2", args{"lookup(output.tf,aws_el)"}, testCmdArg},
		{"Runner with parameters 3", args{"lookup(output.tf, aws_el)"}, testCmdArg},
		{"Runner with parameters with file", args{"lookup(./fixtures/an-empty-file.txt, aws_el)"}, testCmdArgWithFile},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotArgs := extractRunnerArgs(tt.args.cmd); !reflect.DeepEqual(gotArgs, tt.wantArgs) {
				t.Errorf("extractRunnerArgs() = %v, want %v", gotArgs, tt.wantArgs)
			}
		})
	}
}

func Test_formatArgs(t *testing.T) {
	type args struct {
		unformattedArgs []string
	}
	tests := []struct {
		name     string
		args     args
		wantArgs []cmdArg
	}{
		{"With spaces", args{[]string{" my ", "arg", " is", "cool "}},
			[]cmdArg{
				{0, "my", false},
				{1, "arg", false},
				{2, "is", false},
				{3, "cool", false},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotArgs := formatArgs(tt.args.unformattedArgs); !reflect.DeepEqual(gotArgs, tt.wantArgs) {
				t.Errorf("formatArgs() = %v, want %v", gotArgs, tt.wantArgs)
			}
		})
	}
}

func Test_isFile(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name       string
		args       args
		wantIsFile bool
		wantFile   string
	}{
		{"String is an existing file", args{"./fixtures/an-empty-file.txt"}, true, getAbsPath("./fixtures/an-empty-file.txt")},
		{"String is a non existing file", args{"./fixtures/this-file-doesnt-exist"}, false, "./fixtures/this-file-doesnt-exist"},
		{"String is a not a file", args{"my string"}, false, "my string"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIsFile, gotFile := isFile(tt.args.value)
			if gotIsFile != tt.wantIsFile {
				t.Errorf("isFile() gotIsFile = %v, want %v", gotIsFile, tt.wantIsFile)
			}
			if gotFile != tt.wantFile {
				t.Errorf("isFile() gotFile = %v, want %v", gotFile, tt.wantFile)
			}
		})
	}
}
