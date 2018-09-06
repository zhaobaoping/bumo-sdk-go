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
