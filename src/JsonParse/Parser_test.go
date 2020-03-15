package JsonParse

import (
	"fmt"
	"testing"
)

// Author:Boyn
// Date:2020/3/15

func TestJsonParser_ParseObject(t *testing.T) {

	parser := NewJsonParser(GetString())
	object, err := parser.ParseObject()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(object)
	}
}
