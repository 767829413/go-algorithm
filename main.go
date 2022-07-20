package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	str := `{"state":-106,"msg":"Response error: \u001f\ufffd\u0008\u0000\u0000\u0000\u0000\u0000\u0000\u0003\ufffdV*.I,IU\ufffd24000\ufffdQ\ufffd-NW\ufffdR\ufffd)53N4\ufffd)55N5\ufffd)5I5H\u0001\ufffd\ufffdL-\ufffd\ufffd\ufffd\ufffd\ufffdR-\u0000\ufffdZ\ufffdp6\u0000\u0000\u0000"}`
	data := make(map[string]interface{})

	json.Unmarshal([]byte(str), &data)

	fmt.Println(data)
}
