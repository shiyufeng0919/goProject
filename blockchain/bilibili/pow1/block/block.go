package block

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"
)

//定义区块结构
type Block struct {
	//上一个区块Hash
	PreHash string
	//当前区块Hash
	HashCode string
	//时间戳
	TimeStamp string
	//当前网络难度系数
	//控制哈希值有几个前导0
	Diff int
	//交易信息
	Data string
	//区块高度
	Index int
	//随机值
	Nonce int
}

/*
创建创世区块(链中的第一个区块)
block1   ->    block2    ->  block3 ...
*/
func GenerateFirstBlock(data string) Block {
	//创建第一个区块
	var firstBlock Block
	firstBlock.PreHash = "0"
	firstBlock.TimeStamp = time.Now().String()
	firstBlock.Diff = 3 //前导0个数
	firstBlock.Data = data
	firstBlock.Index = 1
	firstBlock.Nonce = 0

	//当前区块Hash
	//用sha256计算一个真正的哈希
	firstBlock.HashCode = GenerationHashValue(firstBlock)
	return firstBlock
}

//生成区块Hash值
func GenerationHashValue(block Block) string {
	var hashData = strconv.Itoa(block.Index) + strconv.Itoa(block.Nonce) + strconv.Itoa(block.Diff) + block.TimeStamp
	//哈希算法
	var sha = sha256.New()
	sha.Write([]byte(hashData))
	hashed := sha.Sum(nil)
	//[]byte -> string
	return hex.EncodeToString(hashed)
}

//产生新的区块
func GenerateNextBlock(data string,oldBlock Block) Block{
	//产生一个新的区块
	var newBlock Block
	newBlock.TimeStamp = time.Now().String()
	//规定前导0为3个
	newBlock.Diff = 3
	newBlock.Index = oldBlock.Index + 1
	newBlock.Data = data
	newBlock.PreHash = oldBlock.HashCode
	//0是由矿工调整
	newBlock.Nonce = 0
	newBlock.HashCode = pow(newBlock.Diff,&newBlock)
	return newBlock
}

//Pow工作量证明算法进行哈希碰撞,计算一个正确hash值
func pow(diff int,block *Block) string{
	//不停地去挖矿
	for {
		hash := GenerationHashValue(*block)
		//每挖矿一次，打印一次hash值
		fmt.Println(hash)
		//判断hash前缀是否为4个0
		if strings.HasPrefix(hash,strings.Repeat("0",diff)){
			fmt.Println("挖矿成功!")
			return hash
		}else{
			//没挖到,不断的改变nonce值再次尝试
			block.Nonce ++
		}
	}
}
