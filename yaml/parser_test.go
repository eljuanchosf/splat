package yaml

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	type args struct {
		doc string
	}
	tests := []struct {
		name string
		args args
		want Doc
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Parse(tt.args.doc); !reflect.DeepEqual(got, tt.want) {
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
	// TODO: Add test cases.
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
