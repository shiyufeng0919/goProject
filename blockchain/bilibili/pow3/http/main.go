package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

/*
 该示例基于web实现添加区块到区块链(利用数组存储)及查询区块链信息.
*/

//设置难度系数，设置为3个
const difficulty = 3

//定义区块
type Block struct {
	//区块高度
	Index int
	//时间戳
	Timestamp string
	//data交易信息
	BMP int
	//当前区块hash
	HashCode string
	//上一个区块hash
	PreHash string
	//前导0个数
	Diff int
	//随机值
	Nonce int
}

//用数组维护区块链
var blockChain []Block

//声明锁
var mutex = &sync.Mutex{}

//生成区块
func generateBlock(preBlock Block, data int) Block {
	//声明新区块
	var newBlock Block
	newBlock.PreHash = preBlock.HashCode
	newBlock.Timestamp = time.Now().String()
	newBlock.Index = preBlock.Index + 1
	newBlock.BMP = data
	newBlock.Diff = difficulty
	//循环挖矿
	for i := 0; ; i++ {
		//每挖一次，随机值做改变
		newBlock.Nonce++
		//计算hash
		hash := calculateHash(newBlock)
		fmt.Println(hash)
		//针对hash做验证,判断前导0是否为difficulty
		if isHashValid(hash, newBlock.Diff) {
			//一致
			fmt.Println("挖矿成功!")
			newBlock.HashCode = hash
			return newBlock
		}
	}
}

//按照规则生成hash
func calculateHash(block Block) string {
	record := strconv.Itoa(block.Index) + block.Timestamp + strconv.Itoa(block.Nonce) +
		strconv.Itoa(block.BMP) + block.PreHash
	sha := sha256.New()
	sha.Write([]byte(record))
	hashed := sha.Sum(nil)
	return hex.EncodeToString(hashed)
}

//判断hash的前导0个数和难度系数是否一致
func isHashValid(hash string, difficulty int) bool {
	prefix := strings.Repeat("0", difficulty)
	return strings.HasPrefix(hash, prefix)
}

func main() {
	//简单测试
	//var firstBlock Block
	//firstBlock.Diff = 3
	//firstBlock.Nonce = 0
	//firstBlock.PreHash = "0"
	//firstBlock.BMP = 1
	//firstBlock.Index = 0
	//firstBlock.HashCode = "0"
	//generateBlock(firstBlock,1)

	//基于WEB
	//默认加载.env文件(默认加载项目目录syfProject/.env)
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	go func() { //协程
		//创世区块
		genesisBlock := Block{}
		genesisBlock = Block{
			Index:     0,
			Timestamp: time.Now().String(),
			BMP:       0,
			HashCode:  calculateHash(genesisBlock),
			PreHash:   "",
			Diff:      difficulty,
			Nonce:     0,
		}
		mutex.Lock()
		//将区块添加到区块链,可能存在并发问题，需要加锁
		blockChain = append(blockChain, genesisBlock)
		mutex.Unlock()

		//格式化输出到控制台，便于查看
		spew.Dump(genesisBlock)
	}()
	//作为http服务器启动函数
	log.Fatal(run())
}

func run() error {
	//处理get或Post请求的回调
	mux := makeMuxRouter()
	httpAddr := os.Getenv("ADDR")
	log.Println("listening on:", httpAddr)
	serve := &http.Server{
		Addr:    ":" + httpAddr,
		Handler: mux,
		//TLSConfig:         nil,
		ReadTimeout: 10 * time.Second, //读超时
		//ReadHeaderTimeout: 0,
		WriteTimeout: 10 * time.Second, //写超时
		//IdleTimeout:       0,
		MaxHeaderBytes: 1 << 20, //最大响应头(位运算 大约为1mb)
		//TLSNextProto:      nil,
		//ConnState:         nil,
		//ErrorLog:          nil,
		//BaseContext:       nil,
		//ConnContext:       nil,
	}
	//监听服务
	if err := serve.ListenAndServe(); err != nil {
		return err
	}
	return nil
}

//回调函数
func makeMuxRouter() http.Handler {
	muxRouter := mux.NewRouter()
	/*
		get 读  post 写
		curl -H "Content-Type: application/json" -X POST -d "{\"BPM\":10}" http://localhost:9000
		curl http://localhost:9000
	*/
	muxRouter.HandleFunc("/", handGetBlockChain).Methods("GET")
	muxRouter.HandleFunc("/", handWriteBlock).Methods("POST")
	return muxRouter
}

//处理http get请求,查看区块链信息
func handGetBlockChain(w http.ResponseWriter, r *http.Request) {
	//blockchain -> json
	bytes, err := json.MarshalIndent(blockChain, "", "\t")
	if err != nil {
		//服务器错误
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bytes))
}

//声明post形式发送数据的数据类型
type Message struct {
	BPM int
}

//http post,添加新的区块
func handWriteBlock(writer http.ResponseWriter, request *http.Request) {
	//设置响应结果为json格式
	writer.Header().Set("Content-Type", "application/json")
	var message Message
	//创建json解码器，从request中读取json数据
	decoder := json.NewDecoder(request.Body)
	//利用decoder.Decode()取message值
	if err := decoder.Decode(&message); err != nil {
		respondWithJson(writer, request, http.StatusNotFound, request.Body)
		return
	}
	//释放资源
	defer request.Body.Close()
	//锁,防止并发
	mutex.Lock()
	//创建新的区块链
	preBlock := blockChain[len(blockChain)-1]
	newBlock := generateBlock(preBlock, message.BPM)
	mutex.Unlock()

	//判断区块合法性
	if isBlockValid(newBlock, preBlock) {
		//将区块真正添加到链上
		blockChain = append(blockChain, newBlock)
		//格式化输出
		spew.Dump(blockChain)
	}

	//返回响应信息
	respondWithJson(writer, request, http.StatusCreated, newBlock)
}

//若错误，服务器返回500
func respondWithJson(writer http.ResponseWriter, request *http.Request, code int, errorMsg interface{}) {
	//设置响应头
	writer.Header().Set("Content-Type", "application/json")
	//格式化输出json
	response, err := json.MarshalIndent(errorMsg, "", "\t")
	//转换失败，返回500
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Http 500:Server Error"))
		return
	}
	//返回指定状态码 & 返回指定数据
	writer.WriteHeader(code)
	writer.Write(response)
}

//判断区块合法性
func isBlockValid(newBlock Block, preBlock Block) bool {
	//判断index区块高度值是否正确
	if preBlock.Index+1 != newBlock.Index {
		return false
	}
	//判断区块hash是否正确
	if newBlock.PreHash != preBlock.HashCode {
		return false
	}
	//再次计算hash值，与newBlock.HashCode比对，看是否一致
	if calculateHash(newBlock) != newBlock.HashCode {
		return false
	}
	return true
}
