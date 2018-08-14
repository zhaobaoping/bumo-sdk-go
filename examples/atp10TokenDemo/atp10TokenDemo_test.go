// atp10TokenDemo_test
package atp10TokenDemo_test

import (
	"encoding/json"
	"testing"

	"github.com/bumoproject/bumo-sdk-go/src/model"
	"github.com/bumoproject/bumo-sdk-go/src/sdk"
)

var testSdk sdk.Sdk

type Atp10Metadata struct {
	Version     string `json:"version"`
	Name        string `json:"name"`
	TotalSupply int64  `json:"total_supply"`
	Decimals    int64  `json:"decimals"`
	Description string `json:"description"`
}

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

//
func Test_Atp10Issue(t *testing.T) {
	var atp10Metadata Atp10Metadata
	atp10Metadata.Version = "1.0"
	atp10Metadata.Name = "code"
	atp10Metadata.Decimals = 8
	atp10Metadata.TotalSupply = 1000000000000
	metadataStr, err := json.Marshal(atp10Metadata)
	if err != nil {
		t.Errorf(err.Error())
	}
	var metadata string = "abc"
	var code string = "ABC"
	var amount int64 = 1000000
	var issuer string = "buQemmMwmRQY1JkcU7w3nhruoX5N3j6C29uo"
	var reqDataIssue model.AssetIssueOperation
	reqDataIssue.Init()
	reqDataIssue.SetAmount(amount)
	reqDataIssue.SetCode(code)
	reqDataIssue.SetSourceAddress(issuer)
	reqDataIssue.SetMetadata(metadata)
	hashIssue, ErrorDesc := atp10BlobSubmit(testSdk, reqDataIssue, string(metadataStr))
	if ErrorDesc != "" {
		t.Errorf(ErrorDesc)
	} else {
		t.Log("hash succeed", hashIssue)
	}
	var reqDataSend model.AssetSendOperation
	var destAddress string = "buQtjhgK9SakQPYGzoZ3iHodfRvd8qTGoaYd"
	var sourceAddress string = "buQfnVYgXuMo3rvCEpKA6SfRrDpaz8D8A9Ea"
	reqDataSend.SetAmount(amount)
	reqDataSend.SetCode(code)
	reqDataSend.SetDestAddress(destAddress)
	reqDataSend.SetIssuer(issuer)
	reqDataSend.SetMetadata(metadata)
	reqDataSend.SetSourceAddress(sourceAddress)
	hashSend, ErrorDesc := atp10BlobSubmit(testSdk, reqDataIssue, "")
	if ErrorDesc != "" {
		t.Errorf(ErrorDesc)
	} else {
		t.Log("hash succeed", hashSend)
	}

}

//
func Test_Atp10AppendToIssue(t *testing.T) {
	code := "code"
	issuer := "buQemmMwmRQY1JkcU7w3nhruoX5N3j6C29uo"
	var metadata string = "abc"
	var amount int64 = 1000000
	var reqDataIssue model.AssetIssueOperation
	reqDataIssue.Init()
	reqDataIssue.SetAmount(amount)
	reqDataIssue.SetCode(code)
	reqDataIssue.SetSourceAddress(issuer)
	reqDataIssue.SetMetadata(metadata)
	hashIssue, ErrorDesc := atp10BlobSubmit(testSdk, reqDataIssue, "")
	if ErrorDesc != "" {
		t.Errorf(ErrorDesc)
	} else {
		t.Log("hash succeed", hashIssue)
	}
	var reqDataSend model.AssetSendOperation
	var destAddress string = "buQtjhgK9SakQPYGzoZ3iHodfRvd8qTGoaYd"
	var reqDataCheckActivated model.AccountCheckActivatedRequest
	reqDataCheckActivated.SetAddress(destAddress)
	reqDataCheckActivated := testSdk.Account.CheckActivated(reqDataCheckActivated)
	if reqDataCheckActivated.ErrorCode != 0 || reqDataCheckActivated.Result.IsActivated == false {
		t.Log("destAddress not Activated")
	}
	var sourceAddress string = "buQfnVYgXuMo3rvCEpKA6SfRrDpaz8D8A9Ea"
	reqDataSend.SetAmount(amount)
	reqDataSend.SetCode(code)
	reqDataSend.SetDestAddress(destAddress)
	reqDataSend.SetIssuer(issuer)
	reqDataSend.SetMetadata(metadata)
	reqDataSend.SetSourceAddress(sourceAddress)
	hashSend, ErrorDesc := atp10BlobSubmit(testSdk, reqDataIssue, "")
	if ErrorDesc != "" {
		t.Errorf(ErrorDesc)
	} else {
		t.Log("hash succeed", hashSend)
	}

}

func atp10BlobSubmit(testSdk sdk.Sdk, reqDataOperation model.BaseOperation, metadata string) (hash string, ErrorDesc string) {
	//Blob
	var reqDataBlob model.TransactionBuildBlobRequest
	var sourceAddressBlob string = "buQemmMwmRQY1JkcU7w3nhruoX5N3j6C29uo"
	reqDataBlob.SetSourceAddress(sourceAddressBlob)
	var feeLimit int64 = 1000000000
	reqDataBlob.SetFeeLimit(feeLimit)
	var gasPrice int64 = 1000
	reqDataBlob.SetGasPrice(gasPrice)
	var nonce int64 = 97
	reqDataBlob.SetNonce(nonce)
	reqDataBlob.SetMetadata(metadata)
	reqDataBlob.SetOperation(reqDataOperation)
	resDataBlob := testSdk.Transaction.BuildBlob(reqDataBlob)
	if resDataBlob.ErrorCode != 0 {
		return "", resDataBlob.ErrorDesc
	} else {
		//Sign
		PrivateKey := []string{"PrivateKey"}
		var reqData model.TransactionSignRequest
		reqData.SetBlob(resDataBlob.Result.Blob)
		reqData.SetPrivateKeys(PrivateKey)
		resDataSign := testSdk.Transaction.Sign(reqData)
		if resDataSign.ErrorCode != 0 {
			return "", resDataSign.ErrorDesc
		} else {
			//Submit
			var reqData model.TransactionSubmitRequest
			reqData.SetBlob(resDataBlob.Result.Blob)
			reqData.SetSignatures(resDataSign.Result.Signatures)
			resDataSubmit := testSdk.Transaction.Submit(reqData)
			if resDataSubmit.ErrorCode != 0 {
				return "", resDataSubmit.ErrorDesc
			} else {
				return resDataBlob.Result.Blob, resDataBlob.ErrorDesc
			}
		}
	}
}
