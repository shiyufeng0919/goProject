package db

import (
	"fmt"
	"testing"
)

//测试默认迭代器
func TestNewDefaultIterator(t *testing.T) {
	data := make(map[string][]byte)
	data["K1"] = []byte("V1")
	data["K2"] = []byte("V2")
	data["K3"] = []byte("V3")
	iterator := NewDefaultIterator(data)
	if iterator.length != 3 {
		t.Fatal()
	}
	for iterator.Next() {
		fmt.Printf("%s:%s\n",iterator.Key(),string(iterator.Value()))
	}
}
