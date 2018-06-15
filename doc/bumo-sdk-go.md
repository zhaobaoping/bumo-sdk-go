# Bumo Go SDK

## 概述
本文档简要概述Bumo Go SDK部分常用接口文档

- [配置](#配置)
	- [包引用](#包引用)

- [使用方法](#使用方法)
    - [包导入](#包导入)
    - [初始化结构体](#初始化结构体)
    - [新建链接](#新建链接)
    - [生成未激活账户](#生成未激活账户)
    - [发起TX](#发起tx)
    - [查询](#查询)
- [区块](#区块)
    - [新建链接](#新建链接)
	- [获取区块高度](#获取区块高度)
    - [检查区块同步](#检查区块同步)
	- [根据hash查询交易](#根据hash查询交易)
	- [根据区块序列号查询交易](#根据区块序列号查询交易)
	- [查询区块头](#查询区块头)
	- [生成交易](#生成交易)
    - [评估费用](#评估费用)
    - [签名](#签名)
    - [提交交易](#提交交易)
- [资产](#资产)
	- [发行资产](#发行资产)
	- [转移资产](#转移资产)
	- [发送BU](#发送bu)
- [账户](#账户)
    - [创建未激活账户](#创建未激活账户)
    - [创建激活账户](#创建激活账户)
    - [检查地址](#检查地址)
    - [查询账户](#查询账户)
    - [查询余额](#查询余额)

- [错误码](#错误码)

### 配置

#### 包引用
所依赖的golang包在src文件夹中寻找，依赖的golang包如下：

    go get github.com/bumoproject/bumo-sdk-go//获取包
    
```
//底层依赖
1. "github.com/bumoproject/bumo-sdk-go/src/3rd/proto"
2. "github.com/bumoproject/bumo-sdk-go/src/keypair"
3. "github.com/bumoproject/bumo-sdk-go/src/protocol"
4. "github.com/bumoproject/bumo-sdk-go/src/signature"
5. "github.com/bumoproject/bumo-sdk-go/src/3rd/secureRandom"
```

### 使用方法

#### 包导入
>导入使用的包

```
import (

	"github.com/bumoproject/bumo-sdk-go/src/bumo"
)
```

#### 初始化结构体
>初始化Error和BumoSdk结构体

```
   var Err bumo.Error
   var bumosdk bumo.BumoSdk
```
#### 新建链接
>获取相应的链接

```
  	bumosdk.Newbumo(url)
```
#### 生成未激活账户
>通过调用Account的CreateInactive生成账户，例如：

```
    newPublicKey, newPrivateKey, newAddress, Err := bumosdk.Account.CreateInactive()
```

#### 发起TX
> 创建激活账户、发行资产、转移资产、发送BU等功能可通过以下四步完成

1. 创建operation
   
通过调用对应方法，创建出指定功能的操作，以发行资产为例调用Asset的Issue方法：
   
```
    operation, Err := bumosdk.Account.Asset.Issue(sourceAddress, issueAddress, code, amount)
```


2.  构建transaction

   构建构建transaction，并设置gasPrice、feeLimit和signer等信息
> 注意：gasPrice和feeLimit的单位是MO，且 1 BU = 10^8 MO
   
   例如：
   
```
    //默认费用
    transaction, Err := bumosdk.CreateTransactionWithDefaultFee(sourceAddress, nonce, operation)
    //需要传入费用参数gasPrice，feeLimit
    transaction, Err := bumosdk.CreateTransactionWithFee(address, nonce, gasPrice, feeLimit, issueData)
```
3.  签名

   对transaction数据签名
   
   例如：
   
```
    signTransaction, publicKey, Err := bumosdk.SignTransaction(transaction, sourcePrivateKey)
```
4. 提交transaction

提交transaction，等待区块网络确认结果，一般确认时间是10秒，确认超时时间为500秒。
 
```
    submitTransaction, Err := bumosdk.SubmitTransaction(transaction, signTransaction, publicKey)
```



#### 查询
调用bumo的相应的接口，例如：查询账户信息
```
    addressInfo, Err := bumosdk.Account.GetInfo(address)
```

### 区块

#### 新建链接

###### 调用方法
```
	//获取链接
	bumosdk.Newbumo(url)
```

##### 入参
参数             |      类型    |      描述      |
---------------- | ------------ |  ------------  |
url              |    String    | 公钥           | 

#### 获取区块高度


###### 返回参数
参数             |      类型    |      描述      |
---------------- | ------------ |  ------------  |
blockNumber  |    int64      | 区块高度   | 
Err    |   bumo.Error     | 错误描述 |

###### 调用方法
```
	blockNumber, Err := bumosdk.GetBlockNumber()
	fmt.Println("blockNumber:", blockNumber)
	fmt.Println(Err)
```

###### 运行结果

```
    blockNumber: 146518
    err: {0 <nil>}
```


#### 检查区块同步



##### 返回参数
参数             |      类型    |      描述      |
---------------- | ------------ |  ------------  |
blockStatus  |    bool      |是否同步   | 
Err  |    bumo.Error      |错误描述   | 
###### 调用方法
```
	blockStatus, Err := bumosdk.CheckBlockStatus()
	fmt.Println("blockStatus:", blockStatus)
	fmt.Println("err:", Err)
```
###### 运行结果

```
    blockStatus: true
    err: {0 <nil>}
```



#### 根据hash查询交易

> 通过hash查询交易的终态信息


##### 入参
参数             |      类型    |      描述      |
---------------- | ------------ |  ------------  |
hash  |    String      | hash值   | 


##### 返回参数
参数             |      类型    |      描述      |
---------------- | ------------ |  ------------  |
transaction  |    string      |交易详情   | 
Err  |    bumo.Error      |错误描述   | 

###### 调用方法

```
    hash := "2cc586cbcec3648dacb1e5d8423a76214ee39395bff409435b908684e5177f0d"
	transaction, Err := bumosdk.GetTransaction(hash)
	fmt.Println("transaction:", transaction)
	fmt.Println("err:", Err)
```

###### 运行结果

```
transaction:
{
    "total_count":1,
    "transactions":[
        {
            "actual_fee":275000,
            "close_time":1527212703857389,
            "error_code":0,
            "error_desc":"",
            "hash":"2cc586cbcec3648dacb1e5d8423a76214ee39395bff409435b908684e5177f0d",
            "ledger_seq":16903,
            "signatures":[
                {
                    "public_key":"b00180c2007082d1e2519a0f2d08fd65ba607fe3b8be646192a2f18a5fa0bee8f7a810d011ed",
                    "sign_data":"09216a0cabeec1eacf3dd302aa50ec3554b16d0144343c7d01ee7ac4f9f6ef1d5ed9f3f344dd898c02cf16bf2be3c57bfc5e74077ea2f96fa2ad143d69602701"
                }
            ],
            "transaction":{
                "fee_limit":1000000,
                "gas_price":1000,
                "metadata":"6275696c642073696d706c65206163636f756e74",
                "nonce":59,
                "operations":[
                    {
                        "create_account":{
                            "dest_address":"buQqjAzXfLkJyWXnSfaQkSTgGUUW8kqCEAor",
                            "init_balance":10000000,
                            "priv":{
                                "master_weight":1,
                                "thresholds":{
                                    "tx_threshold":1
                                }
                            }
                        },
                        "type":1
                    }
                ],
                "source_address":"buQtfFxpQP9JCFgmu4WBojBbEnVyQGaJDgGn"
            },
            "tx_size":275
        }
    ]
}
err: {0 <nil>}

//注释：
total_count： 交易总数
transactions： 交易列表
actual_fee： 实际费用
close_time： 交易关闭时间
error_code： 错误码
error_desc: 错误描述
hash: 交易哈希
ledger_seq： 区块高度
signatures： 签名者列表
public_key： 公钥
sign_data： 签名数据
transaction：交易内容
fee_limit：费用限制
gas_price：打包价格
nonce： 交易序列号
operations：操作列表
pay_coin：操作名称
amount： BU的数量
dest_address： 目标地址
type：操作类型
source_address：交易发起账户
tx_size：交易所占字节数

```


#### 根据区块序列号查询交易

> 通过区块序列号查询交易的终态信息

##### 入参
参数             |      类型    |      描述      |
---------------- | ------------ |  ------------  |
ledgerSeq  |    int64      | 区块高度   | 


##### 返回参数
参数             |      类型    |      描述      |
---------------- | ------------ |  ------------  |
block  |    string      |交易详情   | 
Err  |    bumo.Error      |错误描述   | 



###### 调用方法

```
    var ledgerSeq int64 = 53488
	block, Err := bumosdk.GetBlock(ledgerSeq)
	fmt.Println("block:", block)
	fmt.Println("err:", Err)
```

###### 运行结果

```
block:
{
    "total_count":1,
    "transactions":[
        {
            "actual_fee":438000,
            "close_time":1527580634891468,
            "error_code":0,
            "error_desc":"",
            "hash":"2c8a74ca250d8bb42b737ae6da13e0dcecc57226d4196b601e2d829b2d80c0f5",
            "ledger_seq":53488,
            "signatures":[
                {
                    "public_key":"b00180c2007082d1e2519a0f2d08fd65ba607fe3b8be646192a2f18a5fa0bee8f7a810d011ed",
                    "sign_data":"cca2fcdd0f56cc0de7368136f7ad1a7212477b2c067d2b700b7be7f7f43636de61bd5aa347e35540cfe94e079472c3398c5bba1374e27cd675c9eee56161f30b"
                }
            ],
            "transaction":{
                "fee_limit":445000,
                "gas_price":1000,
                "metadata":"65346261613465363938393336643635373436313634363137343631",
                "nonce":122,
                "operations":[
                    {
                        "create_account":{
                            "dest_address":"buQcZ4oHkj9aUvkTQjvjUaZu1tujvjbERDaF",
                            "init_balance":10000000,
                            "metadatas":[
                                {
                                    "key":"key1",
                                    "value":"自定义value1"
                                },
                                {
                                    "key":"key2",
                                    "value":"自定义value2"
                                }
                            ],
                            "priv":{
                                "master_weight":15,
                                "signers":[
                                    {
                                        "address":"buQjddo7iZLKugdkofivaSKmQtYj87Vj59ni",
                                        "weight":10
                                    },
                                    {
                                        "address":"buQWnB3a9cHb6UJ9YUr8pGLfp42T2dgmYevi",
                                        "weight":10
                                    }
                                ],
                                "thresholds":{
                                    "tx_threshold":15,
                                    "type_thresholds":[
                                        {
                                            "threshold":8,
                                            "type":1
                                        },
                                        {
                                            "threshold":6,
                                            "type":4
                                        },
                                        {
                                            "threshold":4,
                                            "type":2
                                        }
                                    ]
                                }
                            }
                        },
                        "type":1
                    }
                ],
                "source_address":"buQs9npaCq9mNFZG18qu88ZcmXYqd6bqpTU3"
            },
            "tx_size":438
        }
    ]
}
err: {0 <nil>}

//注释：
total_count： 交易总数
transactions： 交易列表
actual_fee： 实际费用
close_time： 交易关闭时间
error_code： 错误码
error_desc: 错误描述
hash: 交易哈希
ledger_seq： 区块高度
signatures： 签名者列表
public_key： 公钥
sign_data： 签名数据
transaction：交易内容
fee_limit：费用限制
gas_price：打包价格
nonce： 交易序列号
operations：操作列表
pay_coin：操作名称
amount： BU的数量
dest_address： 目标地址
type：操作类型
source_address：交易发起账户
tx_size：交易所占字节数  
```

#### 查询区块头

> 通过区块序列号查询区块头

##### 入参
参数             |      类型    |      描述      |
---------------- | ------------ |  ------------  |
ledgerSeq  |    int64      | 区块高度   | 

##### 返回参数
参数             |      类型    |      描述      |
---------------- | ------------ |  ------------  |
ledger  |    string      |交易详情   | 
Err  |    bumo.Error      |错误描述   |


###### 调用方法

```
    var ledgerSeq int64 = 53488
    ledger, Err := bumosdk.GetLedger(ledgerSeq)
	fmt.Println("ledger:", ledger)
	fmt.Println("err:", Err)    
```

###### 运行结果

```
ledger: 
{
    "header":{
        "account_tree_hash":"ba71637cdb4ee20e12ef41ae80b8600629aebdb4c78bdb8fd8dd351df3a4e95a",
        "close_time":1527580634891468,
        "consensus_value_hash":"8402d0dd2aab314379f9cb17ddc58fef25e22ac6ab081d91862e6a5a45904a80",
        "fees_hash":"916daa78d264b3e2d9cff8aac84c943a834f49a62b7354d4fa228dab65515313",
        "hash":"673c507e7c77fc6609dd5bdaf488125f11c68d79dc93dbcd76ca5395a6b5c6a7",
        "previous_hash":"6ee11609070a90e83231dcb6809fe1663b4faac157099bb937e1150501694f1c",
        "seq":53488,
        "tx_count":123,
        "validators_hash":"8ca7d06910f144ba47bf0a7437e24713176126b4f82ba8c0ae7995264f30acf5",
        "version":1000
    }
}
err: {0 <nil>}  
//注释：
account_tree_hash： 账户树哈希
close_time： 交易关闭时间
consensus_value_hash： 共识哈希
fees_hash： 费用哈希
hash: 交易哈希
previous_hash： 上一区块哈希
seq: 区块高度
tx_count：交易数
validators_hash：验证哈希
version： 版本
```


#### 生成交易

> 通过交易具体操作生成交易序列化数据

##### 默认费用
> 按照默认交易费用产生

参数             |      类型    |      描述      |
---------------- | ------------ |  ------------  |
sourceAddress  |    String      | 原地址   | 
nonce  |    int64      | 账号序列号   | 
operation  |    []byte      | 交易操作   | 

##### 返回参数
参数             |      类型    |      描述      |
---------------- | ------------ |  ------------  |
transaction    |   String     | 交易序列化数据 | 
Err  |    bumo.Error      |错误描述   | 

###### 调用方法

```
    transaction, Err := bumosdk.CreateTransactionWithDefaultFee(sourceAddress, nonce, operation)
```

##### 传入费用
> 按照传入交易费用产生

##### 传入参数
参数             |      类型    |      描述      |
---------------- | ------------ |  ------------  |
sourceAddress  |    String      | 原地址   | 
nonce  |    int64      | 账号序列号   | 
gasPrice  |    int64      | 打包费用 (单位 : MO　注:1 BU = 10^8 MO)  | 
feeLimit  |    int64      | 交易手续费 (单位 : MO　注:1 BU = 10^8 MO)  | 
operation  |    []byte      | 交易操作   |


##### 返回参数
参数             |      类型    |      描述      |
---------------- | ------------ |  ------------  |
transaction    |   String     | 交易序列化数据 | 
Err  |    bumo.Error      |错误描述   | 


###### 调用方法

```
    transaction, Err := bumosdk.CreateTransactionWithFee(sourceAddress, nonce, gasPrice, feeLimit, operation)
```

##### 评估费用

> 通过交易具体操作评估需要的费用


##### 入参
参数             |      类型    |      描述      |
---------------- | ------------ |  ------------  |
sourceAddress  |    String      | 原地址   | 
nonce  |    int64      | 账号序列号   | 
operation  |    []byte      | 交易操作   | 
signatureNumber  |    int64      | 签名次数 | 

##### 返回参数
参数             |      类型    |      描述      |
---------------- | ------------ |  ------------  |
actualFee    |   int64     | 交易费用 | 
gasPrice  |    int64      | 交易单位  | 
Err  |    bumo.Error      |错误描述   | 

> 以发送BU为例

```
    // 交易提交人区块链账户地址
	sourceAddress := "buQtfFxpQP9JCFgmu4WBojBbEnVyQGaJDgGn"
	//接受账户
	receiverAddress := "buQtL1v6voSdT292gqwTi77ovR4JRa7YLyqf"
	//发送数量
	var amount int64 = 10000
	//账号序列号
	var nonce int64 = 133
	//签名次数
	var signatureNumber int64 = 1
	// 创建operation
     operation, Err := bumosdk.Account.Asset.SendBU(sourceAddress,receiverAddress, amount)
	 //评估费用
	 actualFee, gasPrice, Err := bumosdk.EvaluationFee(sourceAddress, nonce, operation, signatureNumber)

```


#### 签名

> 对交易序列化数据进行签名


##### 入参
参数             |      类型    |      描述      |
---------------- | ------------ |  ------------  |
transaction  |    String      | 序列化交易   | 
sourcePrivateKey    |   String     | 签名者私钥 | 

##### 返回参数
参数             |      类型    |      描述      |
---------------- | ------------ |  ------------  |
signTransaction  |    String      | 签名数据   | 
publicKey    |   String     | 公钥 | 
Err  |    bumo.Error      |错误描述   | 

###### 调用方法

```
   signTransaction, publicKey, Err := bumosdk.SignTransaction(transaction, sourcePrivateKey)
```


#### 提交交易

> 提交签名的交易

##### 入参
参数             |      类型    |      描述      |
---------------- | ------------ |  ------------  |
transaction  |    String      | 序列化交易   | 
signTransaction    |   String     | 签名数据 | 
publicKey  |    String      | 公钥     | 

##### 返回参数
参数             |      类型    |      描述      |
---------------- | ------------ |  ------------  |
submitTransaction  |    String      | 提交结果（成功返回交易hash）   | 
Err  |    bumo.Error      |错误描述   | 



###### 调用方法

```
   submitTransaction, Err := bumosdk.SubmitTransaction(transaction, signTransaction, publicKey)
```


### 资产
#### 发行资产

> 可通过该方法将数字化资产登记到区块链网络中。资产编码和资产发行方区块链地址共同确定唯一性资产。为此再次发行同一种资产该资产数量累加（追加）。


###### 入参
参数             |      类型    |      描述      |
---------------- | ------------ |  ------------  |
sourceAddress  |    String      | 原地址   | 
issueAddress  |    String      | 资产地址   | 
code    |   String     | 资产标签 | 
amount  |    int64      | 发行数量   | 
##### 返回参数
参数             |      类型    |      描述      |
---------------- | ------------ |  ------------  |
issueData  |    []byte      | 交易操作   | 

###### 调用方法

```
    // 交易提交人区块链账户地址
	sourceAddress := "buQtfFxpQP9JCFgmu4WBojBbEnVyQGaJDgGn"
	// 交易提交人账户公钥
	sourcePublicKey := "b0013ea511beaf5baa41b4901d0bd8410e1cf7a8ff376042d870e2cbd3e960e251dde5e6be1f"
	// 交易提交人账户私钥
	sourcePrivateKey := "privbtHrv27sXbMm41MYp1ezpfuNRJNjJB7i9ggYMP2xtDMCJ9SGNBJy"
	//资产创建账户
	issueAddress := "buQtL1v6voSdT292gqwTi77ovR4JRa7YLyqf"
	
	// 创建operation
	assetCode := "HNC"
	issueAmount := 1000000000 // 发行1000000000 HNC
    operation, Err := bumosdk.Account.Asset.Issue(sourceAddress, issueAddress, assetCode, issueAmount) // 创建资产发行操作
	
	// 构造Tx
	//交易序列号
	nonce := 128
	transaction, Err = bumosdk.CreateTransactionWithDefaultFee(sourceAddress, nonce, operation)
	
	//签名
	signTransaction, publicKey, Err := bumosdk.SignTransaction(transaction, sourcePrivateKey)
	
	// 提交Tx
	submitTransaction, Err := bumosdk.SubmitTransaction(transaction, signTransaction, publicKey)

```




#### 转移资产

> 将当前账户（资产发送方）已拥有资产转移（转给）指定账户（资产接收方）


##### 入参
参数             |      类型    |      描述      |
---------------- | ------------ |  ------------  |
sourceAddress  |    String      | 原地址   | 
issueAddress    |   String     | 发行者地址 | 
destAddress    |   String     | 目标地址 | 
code  |    String      | 资产标签   | 
amount  |    int64      | 发行数量   | 


##### 返回参数
参数             |      类型    |      描述      |
---------------- | ------------ |  ------------  |
payData  |    []byte      | 序列化交易   | 
Err  |    bumo.Error      |错误描述   |

###### 调用方法
```
    // 交易提交人区块链账户地址
	sourceAddress := "buQtfFxpQP9JCFgmu4WBojBbEnVyQGaJDgGn"
	// 交易提交人账户公钥
	sourcePublicKey := "b0013ea511beaf5baa41b4901d0bd8410e1cf7a8ff376042d870e2cbd3e960e251dde5e6be1f"
	// 交易提交人账户私钥
	sourcePrivateKey := "privbtHrv27sXbMm41MYp1ezpfuNRJNjJB7i9ggYMP2xtDMCJ9SGNBJy"
	//资产创建账户
	issueAddress := "buQtL1v6voSdT292gqwTi77ovR4JRa7YLyqf"
    //目标账户
	destAddress := "buQVU86Jm4FeRW4JcQTD9Rx9NkUkHikYGp6z"

	
	// 创建operation
	assetCode := "HNC"
	issueAmount := 100 // 转移100 HNC
    operation, Err := bumosdk.Account.Asset.Pay(sourceAddress, destAddress, issueAddress, issueAmount, assetCode)
	
	// 构造Tx
	//交易序列号
	nonce := 128
	transaction, Err = bumosdk.CreateTransactionWithDefaultFee(sourceAddress, nonce, operation)
	
	//签名
	signTransaction, publicKey, Err := bumosdk.SignTransaction(transaction, sourcePrivateKey)
	
	// 提交Tx
	submitTransaction, Err := bumosdk.SubmitTransaction(transaction, signTransaction, publicKey)

```
#### 发送BU
> 将当前账户（发送方）已拥有BU转移（转给）指定账户（接收方）



参数             |      类型    |      描述      |
---------------- | ------------ |  ------------  |
sourceAddress  |    String      | 原地址   | 
issueAddress    |   String     | 发行者地址 | 
destAddress    |   String     | 目标地址 | 
code  |    String      | 资产标签   | 
amount  |    int64      | 发行数量   | 


##### 返回参数
参数             |      类型    |      描述      |
---------------- | ------------ |  ------------  |
payData  |    []byte      | 序列化交易   | 
Err  |    bumo.Error      |错误描述   |

###### 调用方法


```
    // 交易提交人区块链账户地址
	sourceAddress := "buQtfFxpQP9JCFgmu4WBojBbEnVyQGaJDgGn"
	// 交易提交人账户公钥
	sourcePublicKey := "b0013ea511beaf5baa41b4901d0bd8410e1cf7a8ff376042d870e2cbd3e960e251dde5e6be1f"
	// 交易提交人账户私钥
	sourcePrivateKey := "privbtHrv27sXbMm41MYp1ezpfuNRJNjJB7i9ggYMP2xtDMCJ9SGNBJy"
    //目标账户
	destAddress := "buQtL1v6voSdT292gqwTi77ovR4JRa7YLyqf"
	
	// 创建operation
    amount := 10000//发送10000BU
	operation, Err := bumosdk.Account.Asset.SendBU(sourceAddress, destAddress, amount)
	
	// 构造Tx
	//交易序列号
	nonce := 128
	transaction, Err = bumosdk.CreateTransactionWithDefaultFee(sourceAddress, nonce, operation)
	
	//签名
	signTransaction, publicKey, Err := bumosdk.SignTransaction(transaction, sourcePrivateKey)
	
	// 提交Tx
	submitTransaction, Err := bumosdk.SubmitTransaction(transaction, signTransaction, publicKey)

```


### 账户
#### 创建未激活账户


##### 返回参数
参数             |      类型    |      描述      |
---------------- | ------------ |  ------------  |
newPublicKey  |    String      | 公钥   | 
newPrivateKey    |   String     | 私钥 | 
newAddress  |    String      | 地址     | 
Err  |    bumo.Error      |错误描述   | 

###### 调用方法

```
    // 随机一个Bumo区块账户的公私钥对及区块链地址
    newPublicKey, newPrivateKey, newAddress, Err := bumosdk.Account.CreateInactive()

    // 注：开发者系统需要记录该账户的公私钥对及地址

    accountAddress := newAddress // Block chain account address
    accountPrivateKey := newPrivateKey // Block chain account private key
    accountPublicKey := newPublicKey // Block chain account public key
```

#### 创建激活账户

> 创建新账户需要创建账户操作者(区块链已有账户)花费约0.01BU的交易费用，并且给新账户至少0.1BU的初始化数量，该初始化BU数量由创建账户操作者提供。



###### 入参
参数             |      类型    |      描述      |
---------------- | ------------ |  ------------  |
sourceAddress    |    String    | 原地址         | 
receiverAddress  |   String     | 目标地址       | 
initBalance      |    int64     | 创建账户初始化账户余额，最少10000000 MO（注:1 BU = 10^8 MO） |
##### 返回参数
参数             |      类型    |      描述      |
---------------- | ------------ |  ------------  |
createActive  |    []byte      | 交易序列化   | 
Err  |    bumo.Error      |错误描述   | 

###### 调用方法

```
    // 随机一个Bumo区块账户的公私钥对及区块链地址
    newPublicKey, newPrivateKey, newAddress, Err := bumosdk.Account.CreateInactive()

    // 注：开发者系统需要记录该账户的公私钥对及地址

    // 交易提交人区块链账户地址
	sourceAddress := "buQtL1v6voSdT292gqwTi77ovR4JRa7YLyqf"
	// 交易提交人账户公钥
	sourcePublicKey := "b0013ea511beaf5baa41b4901d0bd8410e1cf7a8ff376042d870e2cbd3e960e251dde5e6be1f"
	// 交易提交人账户私钥
	sourcePrivateKey := "privbtHrv27sXbMm41MYp1ezpfuNRJNjJB7i9ggYMP2xtDMCJ9SGNBJy"
	// 创建账户目前最少初始化账户余额是0.1BU
	initBalance := 100000000
	
	createActiveData, Err := bumosdk.Account.CreateActive(sourceAddress, newAddress, initBalance) // 创建创建账户操作
	
	// 构造Tx
	//交易序列号
	nonce := 128
	transaction, Err = bumosdk.CreateTransactionWithDefaultFee(sourceAddress, nonce, createActiveData)
	
	//签名
	signTransaction, publicKey, Err := bumosdk.SignTransaction(transaction, sourcePrivateKey)
	
	// 提交Tx
	submitTransaction, Err := bumosdk.SubmitTransaction(transaction, signTransaction, publicKey)


```
#### 检查地址


##### 入参
参数             |      类型    |      描述      |
---------------- | ------------ |  ------------  |
address  |    String      | 地址   | 
##### 返回参数
参数             |      类型    |      描述      |
---------------- | ------------ |  ------------  |
checkAddress  |    bool      | 是否有效   | 

###### 调用方法

```
    checkAddress := bumosdk.Account.CheckAddress(address)
```

#### 查询账户


##### 入参
参数             |      类型    |      描述      |
---------------- | ------------ |  ------------  |
address  |    String      | 地址   | 

##### 返回参数
参数             |      类型    |      描述      |
---------------- | ------------ |  ------------  |
addressInfo  |    String      | 账户详情   | 
Err  |    bumo.Error      |错误描述   | 


###### 调用方法

```
    address := "buQtfFxpQP9JCFgmu4WBojBbEnVyQGaJDgGn"
	addressInfo, Err := bumosdk.Account.GetInfo(address)
	fmt.Println("addressInfo:", addressInfo)
	fmt.Println("err:", Err)
```

###### 运行结果
```
{
    "address":"buQtfFxpQP9JCFgmu4WBojBbEnVyQGaJDgGn",
    "assets":[
        {
            "amount":19990,
            "key":{
                "code":"RMB",
                "issuer":"buQtfFxpQP9JCFgmu4WBojBbEnVyQGaJDgGn"
            }
        }
    ],
    "assets_hash":"2870d272af8298d2426979cc417fb8b34da557d52911bfc281640c53255bcba1",
    "balance":99999700884116000,
    "metadatas":null,
    "metadatas_hash":"ad67d57ae19de8068dbcd47282146bd553fe9f684c57c8c114453863ee41abc3",
    "nonce":133,
    "priv":{
        "master_weight":1,
        "thresholds":{
            "tx_threshold":1
        }
    }
}
// 注释
address： 账户地址
assets: 账户资产列表
assets_hash: 账户资产哈希
balance: 账户余额
metadatas: 附加属性
metaddatas_hash: 附加属性哈希
nonce: 账户交易序列号
priv: 权限
master_weight: 账户本身权限
thresholds: 门限
tx_threshold: 交易门限

```




#### 查询余额

##### 入参
参数             |      类型    |      描述      |
---------------- | ------------ |  ------------  |
address  |    String      | 地址   | 


##### 返回参数
参数             |      类型    |      描述      |
---------------- | ------------ |  ------------  |
balance  |    int64      | BU数量 (单位 : MO　)  | 
Err  |    bumo.Error      |错误描述   |



###### 调用方法

```
    address := "buQtfFxpQP9JCFgmu4WBojBbEnVyQGaJDgGn"
	balance, Err := bumosdk.Account.GetBalance(address)
	fmt.Println("balance:", balance)
	fmt.Println("err:", Err)
```

###### 运行结果
```
    balance: 99999700884116000
    err: {0 <nil>}
```



### 错误码
> 接口调用错误码信息


参数 | 描述
-----|-----
0|Success.
10001|Invalid private key.
10002|Invalid public key.
10003|Invalid adress.
10005|Transaction fail.
10006|Nonce too small.
10007|Not enough weight.
10008|Invalid number of arguments to the function.
10009|Invalid type of argument to the function.
10010|Argument cannot be empty.
10011|Internal Server Error.
10012|Nonce incorrect.
10013|BU is not enough.
10014|Source address equal to dest address.
10015|Dest account already exists.
10016|Fee not enough.
10017|Query result not exist.
10018|Discard transaction because of lower fee  in queue..
10019|Include invalid arguments.
10020|Fail.
10093|Not enough weight.
10099|Nonce incorrect.
10100|BU is not enough.
10101|Sourc eaddress equal to dest address.
10102|Dest account already exists.
10103|Account no texist.
10111|Fee not enough.
10160|Discard transaction, because of lower fee in queue.
10201|Account does not exist.
10202|Transaction does not exist.
10203|Block does not exist.
10204|The parameter is wrong.
10205|The function 'keypair.Create' failed.
10206|The function 'proto.Marshal' failed.
10207|The function 'proto.Unmarshal' failed.
10208|The function 'http.NewRequest' failed.
10209|The function 'client.Do' failed.
10210|The function 'json.Unarshal' failed.
10211|The function 'json.Marshal' failed.
10212|The function 'ioutil.ReadAll' failed.
10213|The function 'Transaction is invalid' failed.
10214|The function 'keypair.GetEncPublicKey' failed.
10215|The function 'keypair.CheckAddress' failed.
10216|The function 'hex.DecodeString' failed.
10217|The function 'signature.Sign' failed.
10218|The function 'decoder.Decode' failed.
10219|The function 'strconv.ParseInt' failed.
10220|The parameter 'amount' is invalid.
10221|The parameter 'code' is invalid.
10222|The parameter 'issueAddress' is invalid.
10223|The parameter 'sourceAddress' is invalid.
10224|The parameter 'destAddress' is invalid.
10225|The parameter 'initBalance' is invalid.
10226|The parameter 'payload' is invalid.
10227|The parameter 'nonce' is invalid.
10228|The parameter 'operation' is invalid.
10229|The parameter 'gasPrice' is invalid.
10230|The parameter 'feeLimit' is invalid.
10231|The parameter 'signatureNumber' is invalid.
10232|The parameter 'transactionBlob' is invalid.
10233|The parameter 'privateKey' is invalid.
10234|The parameter 'publicKey' is invalid.
10235|The parameter 'signData' is invalid.
10236|The parameter 'signatures' is invalid.
10237|The parameter 'key' is invalid.
10238|The parameter 'value' is invalid.
10239|The parameter 'version' is invalid.
10240|The parameter 'signerAddress' is invalid.
