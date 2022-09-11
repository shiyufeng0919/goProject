package db

import "fmt"

type DB struct {
	path string
	data map[string][]byte
}

//模拟开启连接
func New(path string)(*DB,error){
	self := DB{
		path: path,
		data: make(map[string][]byte),
	}
	return &self,nil
}

//关闭连接
func (db *DB) Close() error{
	return nil
}

func (db *DB) Put(key []byte,value []byte) error{
	db.data[string(key)] = value
	return nil
}

func (db *DB) Get(key []byte)([]byte,error){
	if v,ok := db.data[string(key)];ok {
		return v,nil
	}else{
		return nil,fmt.Errorf("NotFound")
	}
}

func (db *DB) Delete(key []byte) error{
	if _,ok := db.data[string(key)];ok {
		delete(db.data,string(key))
		return nil
	}else{
		return fmt.Errorf("NotFound")
	}
}

func (db *DB) Iterator() Iterator {
	return NewDefaultIterator(db.data)
}