package db

import "fmt"

/*
定义一个迭代器,用于获取kv值
*/

type Iterator interface {
	//判断是否有下一个值
	Next() bool
	//遍历键值
	Key() []byte
	Value() []byte
}

//键值结构体
type Pair struct {
	Key   []byte
	Value []byte
}

//迭代器结构体
type DefaultIterator struct {
	data   []Pair
	index  int
	length int
}

//创建默认迭代器
func NewDefaultIterator(data map[string][]byte) *DefaultIterator {
	self := new(DefaultIterator)
	self.index = -1
	self.length = len(data)
	for k, v := range data {
		p := Pair{
			Key:   []byte(k),
			Value: v,
		}
		//遍历出的数据添加到data
		self.data = append(self.data, p)
	}
	return self
}

func (iterator *DefaultIterator) Next() bool {
	if iterator.index < iterator.length-1 {
		iterator.index++
		return true
	}
	return false
}

func (iterator *DefaultIterator) Key() []byte {
	if iterator.index == -1 || iterator.index >= iterator.length {
		panic(fmt.Errorf("IndexOutOfBoundError"))
	}
	return iterator.data[iterator.index].Key
}

func (iterator *DefaultIterator) Value() []byte {
	if iterator.index == -1 || iterator.index >= iterator.length {
		panic(fmt.Errorf("IndexOutOfBoundError"))
	}
	return iterator.data[iterator.index].Value
}
