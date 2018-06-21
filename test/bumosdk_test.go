// bumosdk_test
package bumosdk_test

import (
	"testing"

	"github.com/bumoproject/bumo-sdk-go/src/bumo"
)

var bumosdk bumo.BumoSdk
var Err bumo.Error

func Test_Newbumo(t *testing.T) {
	url := "http://seed1.bumotest.io:26002"
	bumosdk.Newbumo(url)
	t.Log("Test_Newbumo succeed")
}
func Test_GetBlockNumber(t *testing.T) {

	blocknumber, Err := bumosdk.GetBlockNumber()
	if Err.Err != nil {
		t.Error(Err)
	} else {
		t.Log("Test_GetBlockNumber succeed", blocknumber)
	}

}
func Test_CheckBlockStatus(t *testing.T) {

	status, Err := bumosdk.CheckBlockStatus()
	if Err.Err != nil {
		t.Error(Err)
	} else {
		t.Log("Test_CheckBlockStatus succeed", status)
	}

}
func Test_GetTransaction(t *testing.T) {
	transactionHash := "2cc586cbcec3648dacb1e5d8423a76214ee39395bff409435b908684e5177f0d"
	datahash, Err := bumosdk.GetTransaction(transactionHash)
	if Err.Err != nil {
		t.Error(Err)
	} else {
		t.Log("Test_GetTransaction succeed", datahash)
	}

}
func Test_GetBlock(t *testing.T) {
	var ledgerSeq int64 = 526
	datah, Err := bumosdk.GetBlock(ledgerSeq)
	if Err.Err != nil {
		t.Error(Err)
	} else {
		t.Log("Test_GetBlock succeed", datah)
	}

}
func Test_GetLedger(t *testing.T) {
	var ledgerSeq int64 = 526
	datah, Err := bumosdk.GetLedger(ledgerSeq)
	if Err.Err != nil {
		t.Error(Err)
	} else {
		t.Log("Test_GetLedger succeed", datah)
	}

}
func Test_EvaluationFee(t *testing.T) {
	var amount int64 = 10000
	newAddress := "buQtfFxpQP9JCFgmu4WBojBbEnVyQGaJDgGn"
	sendBuData, Err := bumosdk.Account.Asset.SendBU(newAddress, newAddress, amount)
	if Err.Err != nil {
		t.Error(Err)
	} else {
		address := "buQtfFxpQP9JCFgmu4WBojBbEnVyQGaJDgGn"
		var nonce int64 = 133
		var signatureNumber int64 = 1
		actualFee, gasPrice, Err := bumosdk.EvaluationFee(address, nonce, sendBuData, signatureNumber)
		if Err.Err != nil {
			t.Error(Err)
		} else {
			t.Log("Test_EvaluationFee succeed", actualFee, gasPrice)
		}

	}
}
func Test_Account_CreateActive(t *testing.T) {
	var initBalance int64 = 100000000
	newAddress := "buQtL1v6voSdT292gqwTi77ovR4JRa7YLyqf"
	address := "buQtfFxpQP9JCFgmu4WBojBbEnVyQGaJDgGn"
	createActiveData, Err := bumosdk.Account.CreateActive(address, newAddress, initBalance)
	if Err.Err != nil {
		t.Error(Err)
	} else {
		var nonce int64 = 134
		var gasPrice int64 = 1000
		var feeLimit int64 = 100000000
		transaction, Err := bumosdk.CreateTransactionWithFee(address, nonce, gasPrice, feeLimit, createActiveData)
		if Err.Err != nil {
			t.Error(Err)
		}
		privateKey := "privbvYfqQyG3kZyHE4RX4TYVa32htw8xG4WdpCTrymPUJQ923XkKVbM"
		signTransaction, publicKey1, Err := bumosdk.SignTransaction(transaction, privateKey)
		if Err.Err != nil {
			t.Error(Err)
		}
		submitTransaction, Err := bumosdk.SubmitTransaction(transaction, signTransaction, publicKey1)
		if Err.Err != nil {
			t.Error(Err)
		}
		t.Log("Test_Account_CreateActive succeed", submitTransaction)
	}

}
func Test_Account_GetInfo(t *testing.T) {
	address := "buQs9npaCq9mNFZG18qu88ZcmXYqd6bqpTU4"
	account, Err := bumosdk.Account.GetInfo(address)
	if Err.Err != nil {
		t.Error(Err)
	} else {
		t.Log("Test_Account_GetInfo succeed", account)
	}

}
func Test_Account_CheckAddress(t *testing.T) {
	address := "buQtfFxpQP9JCFgmu4WBojBbEnVyQGaJDgGn"
	if bumosdk.Account.CheckAddress(address) {
		t.Log("Test_Account_CheckAddress succeed")
	} else {
		t.Error("Test_Account_CheckAddress failured")
	}

}
func Test_Account_GetBalance(t *testing.T) {
	address := "buQtfFxpQP9JCFgmu4WBojBbEnVyQGaJDgGn"
	data, Err := bumosdk.Account.GetBalance(address)
	if Err.Err != nil {
		t.Error(Err)
	} else {
		t.Log("Test_Account_GetBalance succeed", data)
	}

}
func Test_Account_CreateInactive(t *testing.T) {

	newPrivateKey, newPublicKey, newAddress, Err := bumosdk.Account.CreateInactive()
	if Err.Err != nil {
		t.Error(Err)
	} else {
		t.Log("Test_Account_CreateInactive succeed", newPrivateKey, newPublicKey, newAddress)
	}

}
func Test_Asset_Issue(t *testing.T) {
	code := "HNC"
	var amount int64 = 10000
	address := "buQtfFxpQP9JCFgmu4WBojBbEnVyQGaJDgGn"
	issueData, Err := bumosdk.Account.Asset.Issue(address, code, amount)
	if Err.Err != nil {
		t.Error(Err)
	} else {
		var nonce int64 = 134
		var gasPrice int64 = 1000
		var feeLimit int64 = 100000000
		transaction, Err := bumosdk.CreateTransactionWithFee(address, nonce, gasPrice, feeLimit, issueData)
		if Err.Err != nil {
			t.Error(Err)
		}

		privateKey := "privbtHrv27sXbMm41MYp1ezpfuNRJNjJB7i9ggYMP2xtDMCJ9SGNBJy"
		signTransaction, publicKey1, Err := bumosdk.SignTransaction(transaction, privateKey)
		if Err.Err != nil {
			t.Error(Err)
		}
		submitTransaction, Err := bumosdk.SubmitTransaction(transaction, signTransaction, publicKey1)
		if Err.Err != nil {
			t.Error(Err)
		}
		t.Log("Test_Asset_Issue succeed", submitTransaction)
	}

}
func Test_Asset_Pay(t *testing.T) {
	code := "HNC"
	var amount int64 = 10
	newAddress := "buQtL1v6voSdT292gqwTi77ovR4JRa7YLyqf"
	issuer := "buQtfFxpQP9JCFgmu4WBojBbEnVyQGaJDgGn"
	payData, Err := bumosdk.Account.Asset.Pay(newAddress, newAddress, issuer, amount, code)
	if Err.Err != nil {
		t.Error(Err)
	} else {
		var nonce int64 = 134
		var gasPrice int64 = 1000
		var feeLimit int64 = 100000000
		address := "buQtfFxpQP9JCFgmu4WBojBbEnVyQGaJDgGn"
		transaction, Err := bumosdk.CreateTransactionWithFee(address, nonce, gasPrice, feeLimit, payData)
		if Err.Err != nil {
			t.Error(Err)
		}

		privateKey := "privbtHrv27sXbMm41MYp1ezpfuNRJNjJB7i9ggYMP2xtDMCJ9SGNBJy"
		signTransaction, publicKeysin, Err := bumosdk.SignTransaction(transaction, privateKey)
		if Err.Err != nil {
			t.Error(Err)
		}
		submitTransaction, Err := bumosdk.SubmitTransaction(transaction, signTransaction, publicKeysin)
		if Err.Err != nil {
			t.Error(Err)
		}
		t.Log("Test_Asset_Pay succeed", submitTransaction)
	}

}
func Test_Asset_SendBU(t *testing.T) {
	var amount int64 = 10000
	newAddress := "buQtL1v6voSdT292gqwTi77ovR4JRa7YLyqf"
	sendBuData, Err := bumosdk.Account.Asset.SendBU(newAddress, newAddress, amount)
	if Err.Err != nil {
		t.Error(Err)
	} else {
		var nonce int64 = 134
		var gasPrice int64 = 1000
		var feeLimit int64 = 100000000
		address := "buQtfFxpQP9JCFgmu4WBojBbEnVyQGaJDgGn"
		transaction, Err := bumosdk.CreateTransactionWithFee(address, nonce, gasPrice, feeLimit, sendBuData)
		if Err.Err != nil {
			t.Error(Err)
		}

		privateKey := "privbvYfqQyG3kZyHE4RX4TYVa32htw8xG4WdpCTrymPUJQ923XkKVbM"
		signTransaction, publicKeysin, Err := bumosdk.SignTransaction(transaction, privateKey)
		if Err.Err != nil {
			t.Error(Err)
		}
		submitTransaction, Err := bumosdk.SubmitTransaction(transaction, signTransaction, publicKeysin)
		if Err.Err != nil {
			t.Error(Err)
		}
		t.Log("Test_Asset_SendBU succeed", submitTransaction)
	}

}
func Test_Contract_Create(t *testing.T) {
	var initBalance int64 = 10000000000
	newAddress := "buQtL1v6voSdT292gqwTi77ovR4JRa7YLyqf"
	payload := "\"use strict\";function init(bar){return;}function main(input){let para=JSON.parse(input);if(para.do_foo){let x={'hello':'world'};}}function query(input){return input;}"
	input := ""
	createData, Err := bumosdk.Contract.Create(newAddress, newAddress, initBalance, payload, input)
	if Err.Err != nil {
		t.Error(Err)
	} else {

		var nonce int64 = 134
		var gasPrice int64 = 1000
		var feeLimit int64 = 100000000
		address := "buQtfFxpQP9JCFgmu4WBojBbEnVyQGaJDgGn"
		transaction, Err := bumosdk.CreateTransactionWithFee(address, nonce, gasPrice, feeLimit, createData)
		if Err.Err != nil {
			t.Error(Err)
		}

		privateKey := "privbvYfqQyG3kZyHE4RX4TYVa32htw8xG4WdpCTrymPUJQ923XkKVbM"
		signTransaction, publicKeysin, Err := bumosdk.SignTransaction(transaction, privateKey)
		if Err.Err != nil {
			t.Error(Err)
		}
		submitTransaction, Err := bumosdk.SubmitTransaction(transaction, signTransaction, publicKeysin)
		if Err.Err != nil {
			t.Error(Err)
		}
		t.Log("Test_Contract_Create succeed", submitTransaction)
	}

}
func Test_Contract_GetContract(t *testing.T) {
	newAddress := "buQtL1v6voSdT292gqwTi77ovR4JRa7YLyqf"
	contract, Err := bumosdk.Contract.GetContract(newAddress)
	if Err.Err != nil {
		t.Error(Err)
	} else {
		t.Log("Test_Contract_GetContract succeed", contract)
	}

}
