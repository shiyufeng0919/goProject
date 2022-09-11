package blockchain

import (
	"fmt"
	"syfProject/blockchain/bilibili/pow1/block"
)

/*
通过链表的形式，维护区块链中的业务
*/
type Node struct {
	//指针域(指向下一个节点)
	NextNode *Node
	//数据域
	Data *block.Block
}

/*
创建头节点，利用头节点保存创世区块block1
block1   -->   block2    -->  block3 ...
*/
func CreateHeaderNode(data *block.Block) *Node{
	//创建头节点
	headerNode := new(Node)
	//头节点指针域指向nil
	headerNode.NextNode = nil
	headerNode.Data = data
	//返回头节点
	return headerNode //block1
}

/*
添加节点
当挖矿成功，添加区块(即添加一个新节点block2/block3...)
block1   -->   block2    -->  block3 ...
*/
func AddNode(data *block.Block,preNode *Node) *Node{
	//创建新节点
	var newNode *Node = new(Node)
	//指针域指向nil
	newNode.NextNode = nil
	newNode.Data = data
	preNode.NextNode = newNode
	return newNode
}

/*
查看列表中的数据
*/
func ShowNodes(node *Node){
	n := node
	for {
		if n.NextNode == nil {
			fmt.Printf("%+v \n",n.Data)
			break
		}else{
			fmt.Printf("%+v \n",n.Data)
			n = n.NextNode
		}
	}
}

