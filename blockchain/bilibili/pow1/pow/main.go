package main

import (
	"syfProject/blockchain/bilibili/pow1/block"
	"syfProject/blockchain/bilibili/pow1/blockchain"
)

/*
该示例模拟生成区块，并将区块链接在一起。同时引入pow共识机制。
*/

func main(){
	firstBlock := block.GenerateFirstBlock("创世区块")
	secondBlock := block.GenerateNextBlock("第二个区块",firstBlock)
	//创建链表
	headerNode := blockchain.CreateHeaderNode(&firstBlock)
	//将第二个区块加入链表
	blockchain.AddNode(&secondBlock,headerNode)
	//从第一个节点(header node)开始向后查...
	blockchain.ShowNodes(headerNode)
}
