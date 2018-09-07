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
	reqData.SetContractAddress(contractAddress)
	var contractBalance string = "100000000000"
	reqData.SetContractBalance(contractBalance)
	var feeLimit int64 = 1000000
	reqData.SetFeeLimit(feeLimit)
	var gasPrice int64 = 1000
	reqData.SetGasPrice(gasPrice)
	var input string = "input"
	reqData.SetInput(input)
	var optType int64 = 2
	reqData.SetOptType(optType)
	var code string = "HNC"
	reqData.SetCode(code)
	resData := testSdk.Contract.Call(reqData)
	if resData.ErrorCode != 0 {
		t.Errorf(resData.ErrorDesc)
	} else {
		t.Log("Test_Contract_CheckValid succeed", resData.Result)
	}
}

//Contract Create
func Test_Contract_Create(t *testing.T) {
	//Operation
	var reqDataOperation model.ContractCreateOperation
	reqDataOperation.Init()
	var initBalance int64 = 10000000000
	var nonce int64 = 98
	var payload string = "function"
	var initInput string = "input"
	reqDataOperation.SetMetadata("Create")
	reqDataOperation.SetInitBalance(initBalance)
	reqDataOperation.SetPayload(payload)
	reqDataOperation.SetInitInput(initInput)

	errorCode, errorDesc, hash := Transaction_BuildBlob_Sign_Submit(reqDataOperation, nonce)
	if errorCode != 0 {
		t.Log("errorDesc:", errorDesc)
	} else {
		t.Log("Test_Contract_Create succeed", hash)
	}
}

//Invoke By Asset
func Test_Invoke_Asset(t *testing.T) {
	//Operation
	var reqDataOperation model.ContractInvokeByAssetOperation
	reqDataOperation.Init()
	var contractAddress string = "buQXmYrmqt6ohcKtLFKgWFSZ5CjYKaSzaMjT"
	var issuer string = "buQXoNR24p2pPqnXPyiDprmTWsU4SYLtBNCG"
	var amount int64 = 1000
	var nonce int64 = 98
	var code string = "HNC"
	reqDataOperation.SetMetadata("Create")
	reqDataOperation.SetAmount(amount)
	reqDataOperation.SetCode(code)
	reqDataOperation.SetContractAddress(contractAddress)
	reqDataOperation.SetIssuer(issuer)

	errorCode, errorDesc, hash := Transaction_BuildBlob_Sign_Submit(reqDataOperation, nonce)
	if errorCode != 0 {
		t.Log("errorDesc:", errorDesc)
	} else {
		t.Log("Test_Invoke_Asset succeed", hash)
	}
}

//Invoke By BU
func Test_Invoke_BU(t *testing.T) {
	//Operation
	var reqDataOperation model.ContractInvokeByBUOperation
	reqDataOperation.Init()
	var contractAddress string = "buQXmYrmqt6ohcKtLFKgWFSZ5CjYKaSzaMjT"
	var amount int64 = 1000
	var nonce int64 = 98
	reqDataOperation.SetMetadata("Create")
	reqDataOperation.SetAmount(amount)
	reqDataOperation.SetContractAddress(contractAddress)

	errorCode, errorDesc, hash := Transaction_BuildBlob_Sign_Submit(reqDataOperation, nonce)
	if errorCode != 0 {
		t.Log("errorDesc:", errorDesc)
	} else {
		t.Log("Test_Invoke_Asset succeed", hash)
	}
}

func Transaction_BuildBlob_Sign_Submit(reqDataOperation model.BaseOperation, nonce int64) (errorCode int, errorDesc string, hash string) {
	//Blob
	var reqDataBlob model.TransactionBuildBlobRequest
	var sourceAddressBlob string = "buQemmMwmRQY1JkcU7w3nhruoX5N3j6C29uo"
	reqDataBlob.SetSourceAddress(sourceAddressBlob)
	var feeLimit int64 = 1000000000
	reqDataBlob.SetFeeLimit(feeLimit)
	var gasPrice int64 = 1000
	reqDataBlob.SetGasPrice(gasPrice)
	reqDataBlob.SetNonce(nonce)
	reqDataBlob.SetMetadata("63")
	reqDataBlob.SetOperation(reqDataOperation)

	resDataBlob := testSdk.Transaction.BuildBlob(reqDataBlob)
	if resDataBlob.ErrorCode != 0 {
		return resDataBlob.ErrorCode, resDataBlob.ErrorDesc, ""
	} else {
		//Sign
		PrivateKey := []string{"privKey"}
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
