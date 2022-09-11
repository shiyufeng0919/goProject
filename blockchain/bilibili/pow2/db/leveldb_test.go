package db

import (
	"bytes"
	"fmt"
	"testing"
)

func TestLevelDB(t *testing.T){
	db,err := New("")
	check(err)
	err = db.Put([]byte("K1"),[]byte("v1"))
	check(err)
	err = db.Put([]byte("K2"),[]byte("v2"))
	check(err)
	err = db.Put([]byte("K3"),[]byte("v3"))
	check(err)
	err = db.Put([]byte("K2"),[]byte("v8"))
	check(err)
	val,err := db.Get([]byte("K2"))
	check(err)
	fmt.Printf("%s\n",string(val))
	if !bytes.Equal(val,[]byte("v8")){
		t.Fatal()
	}

	err = db.Delete([]byte("K2"))
	check(err)
	iter := db.Iterator()
	for iter.Next() {
		fmt.Printf("%s = %s \n",string(iter.Key()),string(iter.Value()))
	}
}

func check(err error){
	if err != nil {
		panic(err)
	}
}
