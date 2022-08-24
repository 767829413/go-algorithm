package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type escapeString string

func (esc escapeString) MarshalJSON() ([]byte, error) {
	return []byte(strconv.QuoteToASCII(string(esc))), nil
}

func testJson() {
	testData := `[["sdsfs手动阀手动阀","","0"]]`
	data := []escapeString{escapeString(testData)}
	jsonData, _ := json.Marshal(data)
	fmt.Println(string(jsonData))
}

func main() {
	fmt.Println([]byte("()[]{}"))
}
