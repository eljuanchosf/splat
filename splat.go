package main

import (
	"fmt"
	"reflect"
)

func main() {

	fileName := "/home/juan/dev/go/splat/fixtures/good-yaml.yml"
	yamlFile, _ := ParseFile(fileName)
	fmt.Println(reflect.TypeOf(yamlFile.Lines))
	// for _, line := range yamlFile.Lines {
	// 	fmt.Println(line.Key)
	// }
}
