package yaml

import (
	"bufio"
	"io/ioutil"
	"strings"
)

// Doc represents the text structure of a Yaml document
type Doc struct {
	Lines []Line
}

// Line represents a line in a Yaml file
type Line struct {
	Key        string
	Value      string
	Indent     int
	HasCommand bool
}

// Parse parses a YAML document as text
func Parse(doc string) Doc {
	yaml := Doc{}
	scanner := bufio.NewScanner(strings.NewReader(doc))
	for scanner.Scan() {
		key, value, _ := splitLine(scanner.Text())
		indent := countLeadingSpace(key)
		docLine := Line{Key: strings.TrimSpace(key), Value: strings.TrimSpace(value), Indent: indent}
		yaml.Lines = append(yaml.Lines, docLine)
	}
	return yaml
}

// ParseFile parses the YAML file into a text representation
func ParseFile(fileName string) (Doc, error) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return Doc{}, err
	}
	yaml := Parse(string(bytes))
	if err != nil {
		return Doc{}, err
	}
	return yaml, nil
}

func splitLine(line string) (key string, value string, err error) {
	split := strings.SplitN(line, ":", 2)
	key = split[0]
	value = ""
	if len(split) == 2 {
		value = split[1]
	}
	return key, value, nil
}

func countLeadingSpace(line string) int {
	i := 0
	for _, runeValue := range line {
		if runeValue == ' ' {
			i++
		} else {
			break
		}
	}
	return i
}
