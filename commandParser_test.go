package splat

import (
	"reflect"
	"testing"
)

var simpleCommand = Command{
	runner: "lookup",
	args:   []string{"output.tf", "aws_el"},
}

var complexCommand = Command{
	runner: "generateCert",
	args:   []string{"SHA256", "selfSigned", "anotherArg"},
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
		{"Simple example", args{"((< lookup(output.tf, aws_el) >))"}, simpleCommand, false},
		{"Complex example", args{"((< generateCert(SHA256, selfSigned, anotherArg) >))"}, complexCommand, false},
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
	type args struct {
		cmd string
	}
	tests := []struct {
		name     string
		args     args
		wantArgs []string
	}{
		{"Runner without parameters", args{"include()"}, []string{}},
		{"Runner with parameters", args{"lookup(output.tf , aws_el)"}, []string{"output.tf", "aws_el"}},
		{"Runner with parameters 2", args{"lookup(output.tf,aws_el)"}, []string{"output.tf", "aws_el"}},
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
		wantArgs []string
	}{
		{"With spaces", args{[]string{" my ", "arg", " is", "cool "}}, []string{"my", "arg", "is", "cool"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotArgs := formatArgs(tt.args.unformattedArgs); !reflect.DeepEqual(gotArgs, tt.wantArgs) {
				t.Errorf("formatArgs() = %v, want %v", gotArgs, tt.wantArgs)
			}
		})
	}
}
