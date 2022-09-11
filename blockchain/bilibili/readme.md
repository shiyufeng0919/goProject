[bilibili](https://www.bilibili.com/video/BV1EE411i7m9?spm_id_from=333.337.search-card.all.click&vd_source=e9daafca1aaf28d2c9d606a61742da1d)

# 一、共识算法

## 1,pow-工作量证明

> 概念

(1).P2P

(2).比特币

(3).区块链

    区块   --->     区块       ---->   区块
    
    区块保存信息:
    hashCode: 当前节点Hash值
    preHash:前一个节点Hash
    timeStamp:时间戳
    Diff:网络难度系数，若前导0符合难度系数，则挖矿成功
    data:交易信息
    index: 区块高度(第一个块高度为1)
    nonce:随机值

**区块链主要包括： 区块和链**

(4).矿机

(5).挖矿: 挖区块

    挖矿过程：
        pow挖矿
        将区块放到链中(链表) ： 区块链是一个分布式存储，应该将数据广播给网络其它参与者
        数据持久化存储(leveldb)
        在广域网中做广播(udp)

**LevelDB是什么：**

    leveldb是google两位工程师使用c++实现的单机版k-v存储系统
    key和value都是任意的字节数组，支持内存和持久化存储
    数据都是按照key排序，用户可重写排序函数
    包含基本的数据操作接口，put(k,v),get(key),delete(key)
    多操作可以当成一次原子操作--支持事务

**LevelDB局限性:**

    leveldb是非关系型数据库，不支持sql查询也不支持索引
    同一时间只支持单进程(支持多线程)访问db
    不支持客户端-服务器模型，用户需要自己封装

**leveldb工作流程：**

    client 
       | 先写数据到硬盘，保证数据不会丢失，再存储到内存，便于高效取数据
       ^
      硬盘  -- ---> 内存  

(6).pow算法: proof-of-work，工作量证明

(7).基于web服务器的pow案例

    案例目标：
    
    程序运行，开放9000端口
    可以通过post方式访问，添加新的区块
    也可以通过get方式访问，查看区块链信息
    
    安装依赖软件:
    
    go get -u github.com/davecgh/go-spew/spew #在控制台格式化输出结果
    go get -u github.com/gorilla/mux #编写web程序软件包
    go get github.com/joho/godotenv  #用于读取GOPATH/src下的env文件 (godotenv读写配置文件)
    

(8).pow优/缺点

    基于经济学原理，吸引人，鼓励更多人参与
    
    实现相对公平
    
    缺点：
    
    需要算力挖矿，直接消耗资源，浪费能源
    
    案例性待考量

# 二、docker

# 三、超级账本(IBM区块链框架)

# 四、EOS

    EOS速度比以太坊快的多。
    
    区块链1.0 比特币   2.0以太坊  3.0 EOS
    
    
# 五、知识储备

## 1,go语言
## 2,数据库
## 3,基本算法
## 4,linux基本操作
## 5,linux集群搭建
## 6,web常识

# 六、区块链能做什么

## 1,币圈

比特币

## 2,链圈

    商业合同
    供应链
    解决企业间信赖问题

