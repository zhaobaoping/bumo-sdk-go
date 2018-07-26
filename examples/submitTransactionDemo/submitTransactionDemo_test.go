// submitTransactionDemo
package submitTransactionDemo_test

import (
	"testing"

	"github.com/bumoproject/bumo-sdk-go/src/model"
	"github.com/bumoproject/bumo-sdk-go/src/sdk"
)

func Test_submitTransactionDemo(t *testing.T) {
	//构建SDK对象
	var testSdk sdk.Sdk
	var reqDataInit model.SDKInitRequest
	reqDataInit.SetUrl("http://seed1.bumotest.io:26002")
	resDataInit := testSdk.Init(reqDataInit)
	if resDataInit.ErrorCode != 0 {
		t.Errorf(resDataInit.ErrorDesc)
	}
	//获取最新的Nonce值
	var reqDataNonce model.AccountGetNonceRequest
	reqDataNonce.SetAddress("buQemmMwmRQY1JkcU7w3nhruoX5N3j6C29uo")
	resDataNonce := testSdk.Account.GetNonce(reqDataNonce)
	if resDataNonce.ErrorCode != 0 {
		t.Errorf(resDataNonce.ErrorDesc)
	}
	//获取Operation
	var reqDataOperation model.BUSendOperation
	reqDataOperation.Init()
	var amount int64 = 1
	var destAddress string = "buQVU86Jm4FeRW4JcQTD9Rx9NkUkHikYGp6z"
	reqDataOperation.SetAmount(amount)
	reqDataOperation.SetMetadata("63")
	reqDataOperation.SetDestAddress(destAddress)
	//获取Blob
	var reqDataBlob model.TransactionBuildBlobRequest
	var sourceAddressBlob string = "buQemmMwmRQY1JkcU7w3nhruoX5N3j6C29uo"
	reqDataBlob.SetSourceAddress(sourceAddressBlob)
	var feeLimit int64 = 10000000
	reqDataBlob.SetFeeLimit(feeLimit)
	var gasPrice int64 = 1000
	reqDataBlob.SetGasPrice(gasPrice)
	var nonce int64 = resDataNonce.Result.Nonce + 1
	reqDataBlob.SetNonce(nonce)
	reqDataBlob.SetMetadata("63")
	reqDataBlob.SetOperation(reqDataOperation)

	resDataBlob := testSdk.Transaction.BuildBlob(reqDataBlob)
	if resDataBlob.ErrorCode != 0 {
		t.Errorf(resDataBlob.ErrorDesc)
	} else {
		//签名
		PrivateKey := []string{"privbUPxs6QGkJaNdgWS2hisny6ytx1g833cD7V9C3YET9mJ25wdcq6h"}
		var reqData model.TransactionSignRequest
		reqData.SetBlob(resDataBlob.Result.Blob)
		reqData.SetPrivateKeys(PrivateKey)
		resDataSign := testSdk.Transaction.Sign(reqData)
		if resDataSign.ErrorCode != 0 {
			t.Errorf(resDataSign.ErrorDesc)
		} else {
			//广播交易
			var reqData model.TransactionSubmitRequest
			reqData.SetBlob(resDataBlob.Result.Blob)
			reqData.SetSignatures(resDataSign.Result.Signatures)
			resDataSubmit := testSdk.Transaction.Submit(reqData)
			if resDataSubmit.ErrorCode != 0 {
				t.Errorf(resDataSubmit.ErrorDesc)
			} else {
				t.Log("Test_Transaction_BuildBlob_Sign_Submit succeed, Hash:", resDataSubmit.Result.Hash)
			}
		}
	}
}
