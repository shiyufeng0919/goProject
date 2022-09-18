package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

/*
反序列化
json -> Object
*/

func TestJson02(t *testing.T) {
	json21()
}

func json21() {
	str := `{"Name":"syf","Age":18}`
	var person Person
	if err := json.Unmarshal([]byte(str), &person); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("person:%+v", person) //person:{Name:syf Age:18}
}
