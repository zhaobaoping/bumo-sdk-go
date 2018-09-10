// submitTransactionDemo
package submitTransactionDemo_test

import (
	"encoding/json"
	"testing"

	"github.com/bumoproject/bumo-sdk-go/src/model"
	"github.com/bumoproject/bumo-sdk-go/src/sdk"
)

var testSdk sdk.Sdk

//init
func Test_Init(t *testing.T) {
	var reqData model.SDKInitRequest
	reqData.SetUrl("http://seed1.bumotest.io:26002")
	resData := testSdk.Init(reqData)
	if resData.ErrorCode != 0 {
		t.Errorf(resData.ErrorDesc)
	} else {
		t.Log("Test_NewSDK")
	}
}

//get contract info
func Test_Contract_GetInfo(t *testing.T) {
	var reqData model.ContractGetInfoRequest
	var address string = "buQXmYrmqt6ohcKtLFKgWFSZ5CjYKaSzaMjT"
	reqData.SetAddress(address)
	resData := testSdk.Contract.GetInfo(reqData)
	if resData.ErrorCode != 0 {
		t.Errorf(resData.ErrorDesc)
	} else {
		data, _ := json.Marshal(resData.Result.Contract)
		t.Log("Contract:", string(data))
		t.Log("Test_Contract_GetInfo succeed", resData.Result)
	}
}

//check valid
func Test_Contract_CheckValid(t *testing.T) {
	var reqData model.ContractCheckValidRequest
	var address string = "buQXmYrmqt6ohcKtLFKgWFSZ5CjYKaSzaMjT"
	reqData.SetAddress(address)
	resData := testSdk.Contract.CheckValid(reqData)
	if resData.ErrorCode != 0 {
		t.Errorf(resData.ErrorDesc)
	} else {
		t.Log("Test_Contract_CheckValid succeed", resData.Result)
	}
}

//call
func Test_Contract_Call(t *testing.T) {
	var reqData model.ContractCallRequest
	var contractAddress string = "buQXmYrmqt6ohcKtLFKgWFSZ5CjYKaSzaMjT"
	var feeLimit int64 = 1000000
	var gasPrice int64 = 1000
	var contractBalance string = "100000000000"
	var input string = "input"
	var optType int64 = 2
	var code string = "HNC"

	reqData.SetContractAddress(contractAddress)
	reqData.SetContractBalance(contractBalance)
	reqData.SetFeeLimit(feeLimit)
	reqData.SetGasPrice(gasPrice)
	reqData.SetInput(input)
	reqData.SetOptType(optType)
	reqData.SetCode(code)
	resData := testSdk.Contract.Call(reqData)

	if resData.ErrorCode != 0 {
		t.Errorf(resData.ErrorDesc)
	} else {
		t.Log("Test_Contract_Call succeed", resData.Result)
	}
}

//Contract Create
func Test_Contract_Create(t *testing.T) {
	// The account private key to create contract
	var createPrivateKey string = "privbtZbKbz4CtsjtN9GnLsUgSG2vKGuMDri9ADnQ7AzYg5kHqcdCH4y"
	// The account address to send this transaction
	var createAddresss string = "buQXmYrmqt6ohcKtLFKgWFSZ5CjYKaSzaMjT"
	// Contract account initialization BU，the unit is MO，and 1 BU = 10^8 MO
	var initBalance int64 = 100000000
	// Contract code
	var payload string = "'use strict';let globalAttribute={};function globalAttributeKey(){return'global_attribute';}"
	// The fixed write 1000L ，the unit is MO
	var gasPrice int64 = 1000
	// Set up the maximum cost 10.01BU
	var feeLimit int64 = 1015076000
	// Transaction initiation account's Nonce + 1
	var nonce int64 = 18
	// Contract init function entry
	var initInput string = ""

	//Operation
	var reqDataOperation model.ContractCreateOperation
	reqDataOperation.Init()
	reqDataOperation.SetInitBalance(initBalance)
	reqDataOperation.SetPayload(payload)
	reqDataOperation.SetInitInput(initInput)
	//reqDataOperation.SetMetadata("Create")

	errorCode, errorDesc, hash := submitTransaction(testSdk, reqDataOperation, createPrivateKey, createAddresss, nonce, gasPrice, feeLimit)
	if errorCode != 0 {
		t.Log("errorDesc:", errorDesc)
	} else {
		t.Log("Test_Contract_Create succeed", hash)
	}
}

//Invoke By Asset
func Test_Invoke_Asset(t *testing.T) {
	// Init variable
	// The account private key to invoke contract
	var invokePrivateKey string = "privbvCDPhjNmXdZD2p6RWfXhTC3qzpn8REtZtPSu64mMQDMxAJ3f1hu"
	// The account address to send this transaction
	var invokeAddress string = "buQtjhgK9SakQPYGzoZ3iHodfRvd8qTGoaYd"
	// The account to receive the assets
	var destAddress string = "buQXmYrmqt6ohcKtLFKgWFSZ5CjYKaSzaMjT"
	// The asset code to be sent
	var assetCode string = "TST"
	// The account address to issue asset
	var assetIssuer string = "buQnnUEBREw2hB6pWHGPzwanX7d28xk6KVcp"
	// 0 means that the contract is only triggered
	var amount int64 = 0
	// The fixed write 1000L, the unit is MO
	var gasPrice int64 = 1000
	// Set up the maximum cost 0.01BU
	var feeLimit int64 = 100000000
	// Transaction initiation account's Nonce + 1
	var nonce int64 = 57
	// Contract main function entry
	var input string = ""

	//Operation
	var reqDataOperation model.ContractInvokeByAssetOperation
	reqDataOperation.Init()
	reqDataOperation.SetAmount(amount)
	reqDataOperation.SetCode(assetCode)
	reqDataOperation.SetContractAddress(destAddress)
	reqDataOperation.SetIssuer(assetIssuer)
	reqDataOperation.SetInput(input)
	reqDataOperation.SetSourceAddress(invokeAddress)
	//reqDataOperation.SetMetadata("send token")

	errorCode, errorDesc, hash := submitTransaction(testSdk, reqDataOperation, invokePrivateKey, invokeAddress, nonce, gasPrice, feeLimit)
	if errorCode != 0 {
		t.Log("errorDesc:", errorDesc)
	} else {
		t.Log("Test_Invoke_Asset succeed", hash)
	}
}

//Invoke By BU
func Test_Invoke_BU(t *testing.T) {
	// Init variable
	// The account private key to invoke contract
	var invokePrivateKey string = "privbvCDPhjNmXdZD2p6RWfXhTC3qzpn8REtZtPSu64mMQDMxAJ3f1hu"
	// The account address to send this transaction
	var invokeAddress string = "buQtjhgK9SakQPYGzoZ3iHodfRvd8qTGoaYd"
	// The account to receive the BU
	var destAddress string = "buQXmYrmqt6ohcKtLFKgWFSZ5CjYKaSzaMjT"
	// 0 means that the contract is only triggered
	var amount int64 = 0
	// The fixed write 1000L, the unit is MO
	var gasPrice int64 = 1000
	// Set up the maximum cost 0.01BU
	var feeLimit int64 = 10000000
	// Transaction initiation account's Nonce + 1
	var nonce int64 = 58
	// Contract main function entry
	var input string = ""

	//Operation
	var reqDataOperation model.ContractInvokeByBUOperation
	reqDataOperation.Init()
	reqDataOperation.SetAmount(amount)
	reqDataOperation.SetContractAddress(destAddress)
	reqDataOperation.SetSourceAddress(invokeAddress)
	reqDataOperation.SetInput(input)
	//reqDataOperation.SetMetadata("send token")

	errorCode, errorDesc, hash := submitTransaction(testSdk, reqDataOperation, invokePrivateKey, invokeAddress, nonce, gasPrice, feeLimit)
	if errorCode != 0 {
		t.Log("errorDesc:", errorDesc)
	} else {
		t.Log("Test_Invoke_Asset succeed", hash)
	}
}

func submitTransaction(testSdk sdk.Sdk, reqDataOperation model.BaseOperation, senderPrivateKey string, senderAddresss string, senderNonce int64, gasPrice int64, feeLimit int64) (errorCode int, errorDesc string, hash string) {
	//Blob
	var reqDataBlob model.TransactionBuildBlobRequest
	reqDataBlob.SetSourceAddress(senderAddresss)
	reqDataBlob.SetFeeLimit(feeLimit)
	reqDataBlob.SetGasPrice(gasPrice)
	reqDataBlob.SetNonce(senderNonce)
	reqDataBlob.SetOperation(reqDataOperation)
	//reqDataBlob.SetMetadata("abc")

	resDataBlob := testSdk.Transaction.BuildBlob(reqDataBlob)
	if resDataBlob.ErrorCode != 0 {
		return resDataBlob.ErrorCode, resDataBlob.ErrorDesc, ""
	} else {
		//Sign
		PrivateKey := []string{senderPrivateKey}
		var reqData model.TransactionSignRequest
		reqData.SetBlob(resDataBlob.Result.Blob)
		reqData.SetPrivateKeys(PrivateKey)

		resDataSign := testSdk.Transaction.Sign(reqData)
		if resDataSign.ErrorCode != 0 {
			return resDataSign.ErrorCode, resDataSign.ErrorDesc, ""
		} else {
			//Submit
			var reqData model.TransactionSubmitRequest
			reqData.SetBlob(resDataBlob.Result.Blob)
			reqData.SetSignatures(resDataSign.Result.Signatures)
			resDataSubmit := testSdk.Transaction.Submit(reqData)
			return resDataSubmit.ErrorCode, resDataSubmit.ErrorDesc, resDataSubmit.Result.Hash
		}
	}
}
