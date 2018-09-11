// offlineSignatureDemo_test
package offlineSignatureDemo_test

import (
	"encoding/hex"

	"fmt"
	"testing"

	"github.com/bumoproject/bumo-sdk-go/src/crypto/keypair"
	"github.com/bumoproject/bumo-sdk-go/src/crypto/protocol"
	"github.com/bumoproject/bumo-sdk-go/src/crypto/signature"
	"github.com/bumoproject/bumo-sdk-go/src/model"
	"github.com/bumoproject/bumo-sdk-go/src/sdk"
	"github.com/golang/protobuf/proto"
)

var testSdk sdk.Sdk

//to initialize the sdk
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

//submit transaction
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
func Test_Online_BuildTransactionBlob(t *testing.T) {
	// The account to send BU
	var senderAddresss string = "buQnnUEBREw2hB6pWHGPzwanX7d28xk6KVcp"
	// The account to receive BU
	var destAddress string = "buQBjJD1BSJ7nzAbzdTenAhpFjmxRVEEtmxH"
	// The amount to be sent
	var amount int64 = 100000000000
	// The fixed write 1000L, the unit is MO
	var gasPrice int64 = 1000
	// Set up the maximum cost 0.01BU
	var feeLimit int64 = 100000000
	// Transaction initiation account's nonce + 1
	var nonce int64 = 0
	errorCode, errorDesc, transactionBlobResult := buildTransactionBlob(senderAddresss, nonce, destAddress, amount, feeLimit, gasPrice)
	if errorCode != 0 {
		fmt.Println(errorDesc)
	} else {
		fmt.Println(transactionBlobResult)
	}

}

//func Test_Offline_ParseBlob(t *testing.T) {
//	// Get transactionBlobResult from 1 (Network Environment)
//	var blob string = "0a24627551656d6d4d776d525159314a6b63553777336e6872756f58354e336a36433239756f106d18c0843d20e80728eff135320236333a3008071a02363352280a24627551565538364a6d3446655257344a63515444395278394e6b556b48696b594770367a1064"
//	// Parsing the transaction Blob
//	transaction := ParseBlob(blob)
//	if transaction != true {
//		fmt.Println(blob)
//	} else {
//		fmt.Println("blob is false")
//	}
//}
//func Test_Offline_SignTransactionBlob(t *testing.T) {
//	// When the transaction Blob is confirmed, it begins to sign a signature

//	// Transaction Blob
//	var transactionBlob string = "0a24627551656d6d4d776d525159314a6b63553777336e6872756f58354e336a36433239756f106d18c0843d20e80728eff135320236333a3008071a02363352280a24627551565538364a6d3446655257344a63515444395278394e6b556b48696b594770367a1064"
//	// The account private key to send BU
//	var senderPrivateKey string = "privbyQCRp7DLqKtRFCqKQJr81TurTqG6UKXMMtGAmPG3abcM9XHjWvq"

//	// Sign transaction
//	signature := signTransaction(transactionBlob, senderPrivateKey)
//	fmt.Println(signature)
//}
func signTransaction(transactionBlob string, senderPrivateKey string) string {
	privateKey, err := keypair.GetEncPublicKey(senderPrivateKey)
	if err != nil {
		return ""
	}
	SignData, err := signature.Sign(privateKey, []byte(transactionBlob))
	if err != nil {
		return ""
	}
	return SignData
}

func ParseBlob(blob string) bool {
	TransactionBlob, err := hex.DecodeString(blob)
	if err != nil {
		return false
	} else {
		var transactionBlob protocol.Transaction
		err = proto.Unmarshal(TransactionBlob, &transactionBlob)
		if err != nil {
			return false
		} else {
			return true
		}

	}

}
func buildTransactionBlob(senderAddress string, nonce int64, destAddress string, amount int64, feeLimit int64, gasPrice int64) (errorCode int, errorDesc string, blob string) {
	// Build sendBU
	var reqDataOperation model.BUSendOperation
	reqDataOperation.SetAmount(amount)
	reqDataOperation.SetDestAddress(destAddress)
	reqDataOperation.SetSourceAddress(senderAddress)

	// Init buildBlob request
	//Blob
	var reqDataBlob model.TransactionBuildBlobRequest
	reqDataBlob.SetSourceAddress(senderAddress)
	reqDataBlob.SetFeeLimit(feeLimit)
	reqDataBlob.SetGasPrice(gasPrice)
	reqDataBlob.SetNonce(nonce)
	reqDataBlob.SetOperation(reqDataOperation)

	resDataBlob := testSdk.Transaction.BuildBlob(reqDataBlob)
	if resDataBlob.ErrorCode != 0 {
		return resDataBlob.ErrorCode, resDataBlob.ErrorDesc, ""
	} else {
		return resDataBlob.ErrorCode, resDataBlob.ErrorDesc, resDataBlob.Result.Blob
	}
}
