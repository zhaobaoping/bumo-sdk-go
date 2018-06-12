bumo-sdk-go
=======

## bumo-sdk-go  Installation
```
go get github.com/bumoproject/bumo-sdk-go
```


## bumo-sdk-go   Usage

```go

import (
	"fmt"

	"github.com/bumoproject/bumo-sdk-go/src/bumo"
)

var Taccount bool = false
var Tasset bool = false
var Tbumo bool = false
var Tcontract bool = false

var Err bumo.Error

func main() {
	//初始化账户
	var bumosdk bumo.BumoSdk

	url := "http://127.0.0.1:36002"
	//新建链接
	bumosdk.Newbumo(url)

	newAddress := "buQtL1v6voSdT292gqwTi77ovR4JRa7YLyqf"
	newAddress2 := "buQVU86Jm4FeRW4JcQTD9Rx9NkUkHikYGp6z"
	contractAddress := "buQmRHwkLN7cfZuyoeY8pD9zTvaAS5rDZR1r"
	publicKey := "b00180c2007082d1e2519a0f2d08fd65ba607fe3b8be646192a2f18a5fa0bee8f7a810d011ed"
	privateKey := "privbvYfqQyG3kZyHE4RX4TYVa32htw8xG4WdpCTrymPUJQ923XkKVbM"
	address := "buQs9npaCq9mNFZG18qu88ZcmXYqd6bqpTU3"
	hash := "2cc586cbcec3648dacb1e5d8423a76214ee39395bff409435b908684e5177f0d"
	var ledgerSeq int64 = 53488
	if Taccount {
		//生成地址
		newPublicKey, newPrivateKey, newAddress, Err := bumosdk.Account.CreateInactive()
		fmt.Println("newPublicKey:", newPublicKey)
		fmt.Println("newPrivateKey:", newPrivateKey)
		fmt.Println("newAddress:", newAddress)
		fmt.Println("err:", Err)
		//检查地址合法性
		fmt.Println("CheckAddress:", bumosdk.Account.CheckAddress(newAddress))
		//查询账户
		addressInfo, Err := bumosdk.Account.GetInfo(address)
		fmt.Println("addressInfo:", addressInfo)
		fmt.Println("err:", Err)
		//查询余额
		var bumosdk1 bumo.BumoSdk

		url1 := "http://192.168.50.41:36002"
		bumosdk1.Newbumo(url1)
		address1 := "buQVkUUBKpDKRmHYWw1MU8U7ngoQehno165i"
		balance1, Err := bumosdk1.Account.GetBalance(address1)
		balance, Err := bumosdk.Account.GetBalance(address)
		addressInfo1, Err := bumosdk1.Account.GetInfo(address1)
		addressInfo, Err = bumosdk.Account.GetInfo(address)
		fmt.Println("addressInfo1:", addressInfo1)
		fmt.Println("addressInfo:", addressInfo)
		fmt.Println("balance1:", balance1)
		fmt.Println("balance:", balance)
		fmt.Println("err:", Err)
		//初始化账户
		createActive, Err := bumosdk.Account.CreateActive(newAddress, newAddress, 100000000)
		fmt.Println("createActive", createActive)
		fmt.Println("err:", Err)
	}
	if Tasset {
		//发行资产
		issueData, Err := bumosdk.Account.Asset.Issue(address, address, "HNC", 1000)
		fmt.Println("issueData:", issueData)
		fmt.Println("err:", Err)
		//转移资产
		payData, Err := bumosdk.Account.Asset.Pay(address, address, newAddress, 5, "HNC")
		fmt.Println("payData:", payData)
		fmt.Println("err:", Err)
		//交易BU
		sendBuData, Err := bumosdk.Account.Asset.SendBU(newAddress, newAddress, 10000)
		fmt.Println("sendBuData:", sendBuData)
		fmt.Println("err:", Err)
	}
	if Tbumo {
		//查询
		if false {
			//获取当前区块高度
			blockNumber, Err := bumosdk.GetBlockNumber()
			fmt.Println("blockNumber:", blockNumber)
			fmt.Println("err:", Err)
			//检查区块同步
			blockStatus, Err := bumosdk.CheckBlockStatus()
			fmt.Println("blockStatus:", blockStatus)
			fmt.Println("err:", Err)
			//根据hash查询交易
			transaction, Err := bumosdk.GetTransaction(hash)
			fmt.Println("transaction:", transaction)
			fmt.Println("err:", Err)
			//根据区块高度查询交易
			block, Err := bumosdk.GetBlock(ledgerSeq)
			fmt.Println("block:", block)
			fmt.Println("err:", Err)
			//查询区块头
			ledger, Err := bumosdk.GetLedger(ledgerSeq)
			fmt.Println("ledger:", ledger)
			fmt.Println("err:", Err)
		}

		//生成交易blob
		//生成普通账号
		var transaction string
		if false {
			createActiveData, Err := bumosdk.Account.CreateActive(newAddress2, newAddress2, 100000000)
			transaction, Err = bumosdk.CreateTransactionWithFee(address, 127, 0, 0, createActiveData)
			fmt.Println("createActiveData:", createActiveData)
			fmt.Println("transaction:", transaction)
			fmt.Println("err:", Err)
		}
		//发行资产
		if true {
			issueData, Err := bumosdk.Account.Asset.Issue(newAddress2, newAddress2, "HNC", 10000)
			transaction, Err = bumosdk.CreateTransactionWithDefaultFee(newAddress2, 135, issueData)
			fmt.Println("transaction:", transaction)
			fmt.Println("err:", Err)
		}
		//转移资产
		if false {
			payData, Err := bumosdk.Account.Asset.Pay(newAddress2, newAddress2, address, 10, "HNC")
			transaction, Err = bumosdk.CreateTransactionWithFee(address, 129, 1000, 10000000000, payData)
			fmt.Println("transaction:", transaction)
			fmt.Println("err:", Err)
		}
		//转移BU
		if false {
			sendBuData, Err := bumosdk.Account.Asset.SendBU(newAddress2, newAddress2, 10000)
			transaction, Err = bumosdk.CreateTransactionWithFee(address, 130, 0, 10000000000, sendBuData)
			fmt.Println("transaction:", transaction)
			fmt.Println("err:", Err)
		}
		//签名
		//transaction = "0a2462755173396e70614371396d4e465a473138717538385a636d585971643662717054553310800118c0e8d4d01220e8073a320802122462755173396e70614371396d4e465a473138717538385a636d58597164366271705455332a080a03524d4210904e"
		signTransaction, publicKey1, Err := bumosdk.SignTransaction(transaction, privateKey)
		fmt.Println("signTransaction:", signTransaction)
		fmt.Println("publicKey:", publicKey1)
		fmt.Println("err:", Err)
		//提交交易

		submitTransaction, Err := bumosdk.SubmitTransaction(transaction, signTransaction, publicKey)
		fmt.Println("submitTransaction:", submitTransaction)
		fmt.Println("err:", Err)
	}
	if Tcontract {

		var transaction string
		//生成合约账号
		if false {
			payload := "\"use strict\";function init(bar){return;}function main(input){let para=JSON.parse(input);if(para.do_foo){let x={'hello':'world'};}}function query(input){return input;}"
			createData, Err := bumosdk.Contract.Create(contractAddress, contractAddress, 10000000000, payload, "")
			transaction, Err = bumosdk.CreateTransactionWithFee(address, 132, 0, 10000000000, createData)
			fmt.Println("err:", Err)
			fmt.Println("transaction:", transaction)
			//签名
			signTransaction, publicKey1, Err := bumosdk.SignTransaction(transaction, privateKey)
			fmt.Println("signTransaction:", signTransaction)
			fmt.Println("publicKey:", publicKey1)
			fmt.Println("err:", Err)
			//提交交易
			submitTransaction, Err := bumosdk.SubmitTransaction(transaction, signTransaction, publicKey)
			fmt.Println("submitTransaction:", submitTransaction)
			fmt.Println("err:", Err)
		}
		//查询合约
		if false {
			contract, Err := bumosdk.Contract.GetContract(contractAddress)
			fmt.Println("contract:", contract)
			fmt.Println("err:", Err)
		}
		//评估费用
		if false {
			sendBuData, Err := bumosdk.Account.Asset.SendBU(address, newAddress2, 10000)
			actualFee, gasPrice, Err := bumosdk.EvaluationFee(address, 133, sendBuData, 1)
			fmt.Println("actualFee", actualFee)
			fmt.Println("gasPrice", gasPrice)
			fmt.Println("err:", Err)
		}

	}
	return
}

```

## License

[MIT](LICENSE)
