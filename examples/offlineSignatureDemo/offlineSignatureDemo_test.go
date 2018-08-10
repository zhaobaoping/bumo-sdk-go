// offlineSignatureDemo_test
package offlineSignatureDemo_test

import (
	"encoding/json"
	"testing"

	"github.com/bumoproject/bumo-sdk-go/src/model"
	"github.com/bumoproject/bumo-sdk-go/src/sdk"
)

var testSdk sdk.Sdk

//To initialize the SDK
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

//Offline_SignTransactionBlob
func Test_Offline_SignTransactionBlob(t *testing.T) {
	var reqData model.TransactionSignRequest
	var transactionBlob string = "0a24627551656d6d4d776d525159314a6b63553777336e6872756f58354e336a36433239756f106d18c0843d20e80728eff135320236333a3008071a02363352280a24627551565538364a6d3446655257344a63515444395278394e6b556b48696b594770367a1064"
	PrivateKeys := []string{"PrivateKey"}
	reqData.SetPrivateKeys(PrivateKeys)
	reqData.SetBlob(transactionBlob)
	resDataSign := testSdk.Transaction.Sign(reqData)
	if resDataSign.ErrorCode != 0 {
		t.Errorf(resDataSign.ErrorDesc)
	} else {
		data, _ := json.Marshal(resDataSign.Result.Signatures)
		t.Log("Signatures:", string(data))
		t.Log("resDataSign:", resDataSign.Result)
	}
}

//submitTransaction
func Test_submitTransaction(t *testing.T) {
	var reqData model.TransactionSubmitRequest
	var transactionBlob string = "0a24627551656d6d4d776d525159314a6b63553777336e6872756f58354e336a36433239756f106d18c0843d20e80728eff135320236333a3008071a02363352280a24627551565538364a6d3446655257344a63515444395278394e6b556b48696b594770367a1064"
	var signData string = "5aac965b327c71555244d1589344a4000c9673380907d1c9bc9ab53dc4e7a39c6f5f5b9f8daae702f3e8ffdf10cdc2ef8d7aa30d3894ab54ff2de9deb03a8305"
	var publicKey string = "b001ebb9f88123658f0a62c49fb5cfbc265cc56ee144a56452012ef2abff7f9ef7974992926b"
	signatures := []model.Signature{
		{
			SignData:  signData,
			PublicKey: publicKey,
		},
	}
	reqData.SetBlob(transactionBlob)
	reqData.SetSignatures(signatures)
	resDataSubmit := testSdk.Transaction.Submit(reqData)
	if resDataSubmit.ErrorCode != 0 {
		t.Errorf(resDataSubmit.ErrorDesc)
	} else {
		t.Log("Hash:", resDataSubmit.Result.Hash)
		t.Log("submit transaction succeed", resDataSubmit.Result)
	}
}
