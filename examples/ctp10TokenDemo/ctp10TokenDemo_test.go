// ctp10TokenDemo_test
package ctp10TokenDemo_test

import (
	"encoding/json"
	"fmt"
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

//ctp10token allowance
func Test_Ctp10Token_Allowance(t *testing.T) {
	var reqData model.Ctp10TokenAllowanceRequest
	var contractAddress string = "buQXoNR24p2pPqnXPyiDprmTWsU4SYLtBNCG"
	reqData.SetContractAddress(contractAddress)
	var spender string = "buQW5p6gaCd331NerjxhD1cAHpmSGwxrt6e6"
	reqData.SetSpender(spender)
	var tokenOwner string = "buQnc3AGCo6ycWJCce516MDbPHKjK7ywwkuo"
	reqData.SetTokenOwner(tokenOwner)
	resData := testSdk.Token.Ctp10Token.Allowance(reqData)
	if resData.ErrorCode != 0 {
		t.Errorf(resData.ErrorDesc)
	} else {
		fmt.Println("Allowance:", resData.Result.Allowance)
		t.Log("Test_Ctp10Token_Allowance succeed", resData.Result)
	}
}

//get ctp10token info
func Test_Ctp10Token_GetInfo(t *testing.T) {
	var reqData model.Ctp10TokenGetInfoRequest
	var contractAddress string = "buQXoNR24p2pPqnXPyiDprmTWsU4SYLtBNCG"
	reqData.SetContractAddress(contractAddress)
	resData := testSdk.Token.Ctp10Token.GetInfo(reqData)
	if resData.ErrorCode != 0 {
		fmt.Println(resData)
		t.Errorf(resData.ErrorDesc)
	} else {
		data, _ := json.Marshal(resData.Result)
		fmt.Println("info:", string(data))
		t.Log("Test_Ctp10Token_GetInfo succeed", resData.Result)
	}
}

//get ctp10token name
func Test_Ctp10Token_GetName(t *testing.T) {
	var reqData model.Ctp10TokenGetNameRequest
	var contractAddress string = "buQXoNR24p2pPqnXPyiDprmTWsU4SYLtBNCG"
	reqData.SetContractAddress(contractAddress)
	resData := testSdk.Token.Ctp10Token.GetName(reqData)
	if resData.ErrorCode != 0 {
		t.Errorf(resData.ErrorDesc)
	} else {
		fmt.Println("Name:", resData.Result.Name)
		t.Log("Test_Ctp10Token_GetName succeed", resData.Result)
	}
}

//get ctp10token symbol
func Test_Ctp10Token_GetSymbol(t *testing.T) {
	var reqData model.Ctp10TokenGetSymbolRequest
	var contractAddress string = "buQXoNR24p2pPqnXPyiDprmTWsU4SYLtBNCG"
	reqData.SetContractAddress(contractAddress)
	resData := testSdk.Token.Ctp10Token.GetSymbol(reqData)
	if resData.ErrorCode != 0 {
		t.Errorf(resData.ErrorDesc)
	} else {
		fmt.Println("Symbol:", resData.Result.Symbol)
		t.Log("Test_Ctp10Token_GetSymbol succeed", resData.Result)
	}
}

//get ctp10token decimals
func Test_Ctp10Token_GetDecimals(t *testing.T) {
	var reqData model.Ctp10TokenGetDecimalsRequest
	var contractAddress string = "buQXoNR24p2pPqnXPyiDprmTWsU4SYLtBNCG"
	reqData.SetContractAddress(contractAddress)
	resData := testSdk.Token.Ctp10Token.GetDecimals(reqData)
	if resData.ErrorCode != 0 {
		t.Errorf(resData.ErrorDesc)
	} else {
		fmt.Println("Decimals:", resData.Result.Decimals)
		t.Log("Test_Ctp10Token_GetDecimals succeed", resData.Result)
	}
}

//get ctp10token totalsupply
func Test_Ctp10Token_GetTotalSupply(t *testing.T) {
	var reqData model.Ctp10TokenGetTotalSupplyRequest
	var contractAddress string = "buQXoNR24p2pPqnXPyiDprmTWsU4SYLtBNCG"
	reqData.SetContractAddress(contractAddress)
	resData := testSdk.Token.Ctp10Token.GetTotalSupply(reqData)
	if resData.ErrorCode != 0 {
		t.Errorf(resData.ErrorDesc)
	} else {
		fmt.Println("TotalSupply:", resData.Result.TotalSupply)
		t.Log("Test_Ctp10Token_GetTotalSupply succeed", resData.Result)
	}
}

//get ctp10token balance
func Test_Ctp10Token_GetBalance(t *testing.T) {
	var reqData model.Ctp10TokenGetBalanceRequest
	var contractAddress string = "buQXoNR24p2pPqnXPyiDprmTWsU4SYLtBNCG"
	reqData.SetContractAddress(contractAddress)
	var tokenOwner string = "buQW5p6gaCd331NerjxhD1cAHpmSGwxrt6e6"
	reqData.SetTokenOwner(tokenOwner)
	resData := testSdk.Token.Ctp10Token.GetBalance(reqData)
	if resData.ErrorCode != 0 {
		fmt.Println(resData)
		t.Errorf(resData.ErrorDesc)
	} else {
		fmt.Println("Balance:", resData.Result.Balance)
		t.Log("Test_Ctp10Token_GetBalance succeed", resData.Result)
	}
}

//Ctp10TokenIssue
func Test_Ctp10TokenIssue(t *testing.T) {
	//Operation
	var reqDataOperation model.Ctp10TokenIssueOperation
	reqDataOperation.Init()
	var amount int64 = 100
	var initBalance int64 = 100000000000000
	var metadata string = "abc"
	var name string = "name"
	var sourceAddress string = "buQVU86Jm4FeRW4JcQTD9Rx9NkUkHikYGp6z"
	var supply int64 = 100000000000000
	var symbol string = "symbol"
	reqDataOperation.SetDecimals(amount)
	reqDataOperation.SetInitBalance(initBalance)
	reqDataOperation.SetMetadata(metadata)
	reqDataOperation.SetName(name)
	reqDataOperation.SetSourceAddress(sourceAddress)
	reqDataOperation.SetSupply(supply)
	reqDataOperation.SetSymbol(symbol)
	hash, ErrorDesc := blobSubmit(testSdk, reqDataOperation)
	if ErrorDesc != "" {
		t.Errorf(ErrorDesc)
	} else {
		t.Log("hash succeed", hash)
	}

}

//Transfer
func Test_Transfer(t *testing.T) {
	//Operation
	var reqDataOperation model.Ctp10TokenTransferOperation
	reqDataOperation.Init()
	var amount int64 = 100
	var contractAddress string = "buQVU86Jm4FeRW4JcQTD9Rx9NkUkHikYGp6z"
	var destAddress string = "buQX5wKcyQE7jB8w7LpC6ZQCamtkRq3nNR75"
	var metadata string = "abc"
	var sourceAddress string = "buQfnVYgXuMo3rvCEpKA6SfRrDpaz8D8A9Ea"
	reqDataOperation.SetAmount(amount)
	reqDataOperation.SetContractAddress(contractAddress)
	reqDataOperation.SetDestAddress(destAddress)
	reqDataOperation.SetMetadata(metadata)
	reqDataOperation.SetSourceAddress(sourceAddress)
	hash, ErrorDesc := blobSubmit(testSdk, reqDataOperation)
	if ErrorDesc != "" {
		t.Errorf(ErrorDesc)
	} else {
		t.Log("hash succeed", hash)
	}

}

//TransferFrom
func Test_TransferFrom(t *testing.T) {
	//Operation
	var reqDataOperation model.Ctp10TokenTransferFromOperation
	reqDataOperation.Init()
	var amount int64 = 100
	var contractAddress string = "buQVU86Jm4FeRW4JcQTD9Rx9NkUkHikYGp6z"
	var destAddress string = "buQX5wKcyQE7jB8w7LpC6ZQCamtkRq3nNR75"
	var fromAddress string = "buQtjhgK9SakQPYGzoZ3iHodfRvd8qTGoaYd"
	var metadata string = "abc"
	var sourceAddress string = "buQfnVYgXuMo3rvCEpKA6SfRrDpaz8D8A9Ea"
	reqDataOperation.SetAmount(amount)
	reqDataOperation.SetContractAddress(contractAddress)
	reqDataOperation.SetDestAddress(destAddress)
	reqDataOperation.SetFromAddress(fromAddress)
	reqDataOperation.SetMetadata(metadata)
	reqDataOperation.SetSourceAddress(sourceAddress)
	hash, ErrorDesc := blobSubmit(testSdk, reqDataOperation)
	if ErrorDesc != "" {
		t.Errorf(ErrorDesc)
	} else {
		t.Log("hash succeed", hash)
	}
}

//Approve
func Test_Approve(t *testing.T) {
	//Operation
	var reqDataOperation model.Ctp10TokenApproveOperation
	reqDataOperation.Init()
	var amount int64 = 100
	var contractAddress string = "buQVU86Jm4FeRW4JcQTD9Rx9NkUkHikYGp6z"
	var spender string = "buQX5wKcyQE7jB8w7LpC6ZQCamtkRq3nNR75"
	var metadata string = "abc"
	var sourceAddress string = "buQfnVYgXuMo3rvCEpKA6SfRrDpaz8D8A9Ea"
	reqDataOperation.SetAmount(amount)
	reqDataOperation.SetContractAddress(contractAddress)
	reqDataOperation.SetSpender(spender)
	reqDataOperation.SetMetadata(metadata)
	reqDataOperation.SetSourceAddress(sourceAddress)
	hash, ErrorDesc := blobSubmit(testSdk, reqDataOperation)
	if ErrorDesc != "" {
		t.Errorf(ErrorDesc)
	} else {
		t.Log("hash succeed", hash)
	}
}

//Assign
func Test_Assign(t *testing.T) {
	//Operation
	var reqDataOperation model.Ctp10TokenAssignOperation
	reqDataOperation.Init()
	var amount int64 = 100
	var contractAddress string = "buQVU86Jm4FeRW4JcQTD9Rx9NkUkHikYGp6z"
	var destAddress string = "buQtjhgK9SakQPYGzoZ3iHodfRvd8qTGoaYd"
	var metadata string = "abc"
	var sourceAddress string = "buQfnVYgXuMo3rvCEpKA6SfRrDpaz8D8A9Ea"
	reqDataOperation.SetAmount(amount)
	reqDataOperation.SetContractAddress(contractAddress)
	reqDataOperation.SetDestAddress(destAddress)
	reqDataOperation.SetMetadata(metadata)
	reqDataOperation.SetSourceAddress(sourceAddress)
	hash, ErrorDesc := blobSubmit(testSdk, reqDataOperation)
	if ErrorDesc != "" {
		t.Errorf(ErrorDesc)
	} else {
		t.Log("hash succeed", hash)
	}
}

//ChangeOwner
func Test_ChangeOwner(t *testing.T) {
	//Operation
	var reqDataOperation model.Ctp10TokenChangeOwnerOperation
	reqDataOperation.Init()
	var contractAddress string = "buQVU86Jm4FeRW4JcQTD9Rx9NkUkHikYGp6z"
	var tokenOwner string = "buQtjhgK9SakQPYGzoZ3iHodfRvd8qTGoaYd"
	var metadata string = "abc"
	var sourceAddress string = "buQfnVYgXuMo3rvCEpKA6SfRrDpaz8D8A9Ea"
	reqDataOperation.SetContractAddress(contractAddress)
	reqDataOperation.SetMetadata(metadata)
	reqDataOperation.SetSourceAddress(sourceAddress)
	reqDataOperation.SetTokenOwner(tokenOwner)
	hash, ErrorDesc := blobSubmit(testSdk, reqDataOperation)
	if ErrorDesc != "" {
		t.Errorf(ErrorDesc)
	} else {
		t.Log("hash succeed", hash)
	}
}

func blobSubmit(testSdk sdk.Sdk, reqDataOperation model.BaseOperation) (hash string, ErrorDesc string) {
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
	var metadata string = "abc"
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
