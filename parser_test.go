package main

import (
	"reflect"
	"testing"
)

var simpleYaml = `a: Easy!
b:
  c: 2
  d: [3, 4]
    f: ((< this is a command >))`

var simpleDoc = Doc{
	Lines: []Line{
		{0, "a", "Easy!", 0, false},
		{1, "b", "", 0, false},
		{2, "c", "2", 2, false},
		{3, "d", "[3, 4]", 2, false},
		{4, "f", "((< this is a command >))", 4, true},
	},
}

var simpleYamlFile = "./fixtures/simple-yaml.yml"

var simpleCommand = Command{
	runner: "lookup",
	args:   []string{"output.tf", "aws_el"},
}

var complexCommand = Command{
	runner: "generateCert",
	args:   []string{"SHA256", "selfSigned", "anotherArg"},
}

func TestParse(t *testing.T) {
	type args struct {
		doc string
	}
	tests := []struct {
		name    string
		args    args
		want    Doc
		wantErr bool
	}{
		{"Simple YAML", args{simpleYaml}, simpleDoc, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args.doc)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseFile(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name    string
		args    args
		want    Doc
		wantErr bool
	}{
		{"Parse simple stub", args{simpleYamlFile}, simpleDoc, false},
		{"File not exists", args{"someNonExistingFile"}, Doc{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseFile(tt.args.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_splitLine(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name      string
		args      args
		wantKey   string
		wantValue string
	}{
		{"Only key", args{"compilation:"}, "compilation", ""},
		{"Key and value", args{"my_value: value"}, "my_value", " value"},
		{"Value with two colons", args{"my_value: value:value"}, "my_value", " value:value"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotKey, gotValue, _ := splitLine(tt.args.line)
			if gotKey != tt.wantKey {
				t.Errorf("splitLine() gotKey = %v, want %v", gotKey, tt.wantKey)
			}
			if gotValue != tt.wantValue {
				t.Errorf("splitLine() gotValue = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}

func Test_countLeadingSpace(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"No space", args{"only one space"}, 0},
		{"One space", args{" only one space"}, 1},
		{"Two spaces", args{"  only one space"}, 2},
		{"Ten spaces", args{"          only one space"}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countLeadingSpace(tt.args.line); got != tt.want {
				t.Errorf("countLeadingSpace() = %v, want %v", got, tt.want)
			}
		})
	}
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
