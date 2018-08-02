// DigitalAssetsDemo
package DigitalAssetsDemo_test

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

//Check that the blocks are synchronized
func Test_Block_CheckStatus(t *testing.T) {
	resData := testSdk.Block.CheckStatus()
	if resData.ErrorCode != 0 {
		t.Errorf(resData.ErrorDesc)
	} else {
		fmt.Println("IsSynchronous:", resData.Result.IsSynchronous)
		t.Log("Test_Block_CheckStatus succeed", resData.Result)
	}

}

//Create account
func Test_Account_Create(t *testing.T) {
	resData := testSdk.Account.Create()
	if resData.ErrorCode != 0 {
		t.Errorf(resData.ErrorDesc)
	} else {
		t.Log("Test_Account_Create", resData.Result)
	}
}

//Verify account address
func Test_Account_checkValid(t *testing.T) {
	var reqData model.AccountCheckValidRequest
	var address string = "buQemmMwmRQY1JkcU7w3nhruoX5N3j6C29uo"
	reqData.SetAddress(address)
	resData := testSdk.Account.CheckValid(reqData)

	if resData.Result.IsValid {
		t.Log("Test_Account_CheckAddress succeed", resData.Result.IsValid)
	} else {
		t.Error("Test_Account_CheckAddress failured")
	}
}

//Enquiry of account details
func Test_Account_GetInfo(t *testing.T) {
	var reqData model.AccountGetInfoRequest
	var address string = "buQemmMwmRQY1JkcU7w3nhruoX5N3j6C29uo"
	reqData.SetAddress(address)
	resData := testSdk.Account.GetInfo(reqData)
	if resData.ErrorCode != 0 {
		t.Errorf(resData.ErrorDesc)
	} else {
		data, _ := json.Marshal(resData.Result)
		fmt.Println("info:", string(data))
		t.Log("Test_Account_GetInfo succeed", resData.Result)
	}
}

//Checking account balance
func Test_Account_GetBalance(t *testing.T) {
	var reqData model.AccountGetBalanceRequest
	var address string = "buQemmMwmRQY1JkcU7w3nhruoX5N3j6C29uo"
	reqData.SetAddress(address)
	resData := testSdk.Account.GetBalance(reqData)
	if resData.ErrorCode != 0 {
		t.Errorf(resData.ErrorDesc)
	} else {
		fmt.Println("Balance:", resData.Result.Balance)
		t.Log("Test_Account_GetBalance succeed", resData.Result)
	}
}

//get account assets
func Test_Account_GetAssets(t *testing.T) {
	var reqData model.AccountGetAssetsRequest
	var address string = "buQemmMwmRQY1JkcU7w3nhruoX5N3j6C29uo"
	reqData.SetAddress(address)
	resData := testSdk.Account.GetAssets(reqData)
	if resData.ErrorCode != 0 {
		t.Errorf(resData.ErrorDesc)
	} else {
		data, _ := json.Marshal(resData.Result.Assets)
		fmt.Println("Assets:", string(data))
		t.Log("Test_Account_GetAssets succeed", resData.Result)

	}
}

//get account metadata
func Test_Account_GetMetadata(t *testing.T) {
	var reqData model.AccountGetMetadataRequest
	var address string = "buQXoNR24p2pPqnXPyiDprmTWsU4SYLtBNCG"
	reqData.SetAddress(address)
	reqData.SetKey("global_attribute")
	resData := testSdk.Account.GetMetadata(reqData)
	if resData.ErrorCode != 0 {
		t.Errorf(resData.ErrorDesc)
	} else {
		data, _ := json.Marshal(resData.Result.Metadatas[0].Value)

		fmt.Println("Metadatas:", string(data))
		t.Log("Test_Account_GetMetadata succeed", resData.Result)
	}
}

//get asset info
func Test_Asset_GetInfo(t *testing.T) {
	var reqData model.AssetGetInfoRequest
	var address string = "buQemmMwmRQY1JkcU7w3nhruoX5N3j6C29uo"
	reqData.SetAddress(address)
	reqData.SetIssuer("buQnc3AGCo6ycWJCce516MDbPHKjK7ywwkuo")
	reqData.SetCode("HNC")
	resData := testSdk.Token.Asset.GetInfo(reqData)
	if resData.ErrorCode != 0 {
		t.Errorf(resData.ErrorDesc)
	} else {
		data, _ := json.Marshal(resData.Result.Assets)
		fmt.Println("Assets:", string(data))
		t.Log("Test_Asset_GetInfo succeed", resData.Result.Assets)
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
		fmt.Println("Contract:", string(data))
		t.Log("Test_Contract_GetInfo succeed", resData.Result)
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
	reqData.SetCtp10TokenOwner(tokenOwner)
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
	reqData.SetCtp10TokenOwner(tokenOwner)
	resData := testSdk.Token.Ctp10Token.GetBalance(reqData)
	if resData.ErrorCode != 0 {
		fmt.Println(resData)
		t.Errorf(resData.ErrorDesc)
	} else {
		fmt.Println("Balance:", resData.Result.Balance)
		t.Log("Test_Ctp10Token_GetBalance succeed", resData.Result)
	}
}

//Check the account transaction serial number
func Test_Account_GetNonce(t *testing.T) {
	var reqData model.AccountGetNonceRequest
	var address string = "buQemmMwmRQY1JkcU7w3nhruoX5N3j6C29uo"
	reqData.SetAddress(address)
	resData := testSdk.Account.GetNonce(reqData)
	if resData.ErrorCode != 0 {
		t.Errorf(resData.ErrorDesc)
	} else {
		fmt.Println("Nonce:", resData.Result.Nonce)
		t.Log("Test_Account_GetNonce succeed", resData.Result)
	}
}

//Submit and send BU transactions
func Test_Transaction_BuildBlob_Sign_Submit(t *testing.T) {
	//Operation
	var reqDataOperation model.BUSendOperation
	reqDataOperation.Init()
	var amount int64 = 100
	var destAddress string = "buQVU86Jm4FeRW4JcQTD9Rx9NkUkHikYGp6z"
	reqDataOperation.SetAmount(amount)
	reqDataOperation.SetMetadata("63")
	reqDataOperation.SetDestAddress(destAddress)
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
	reqDataBlob.SetMetadata("63")
	reqDataBlob.SetOperation(reqDataOperation)

	resDataBlob := testSdk.Transaction.BuildBlob(reqDataBlob)
	if resDataBlob.ErrorCode != 0 {
		fmt.Println(resDataBlob.ErrorDesc)
	} else {
		//Sign
		PrivateKey := []string{"privbUPxs6QGkJaNdgWS2hisny6ytx1g833cD7V9C3YET9mJ25wdcq6h"}
		var reqData model.TransactionSignRequest
		reqData.SetBlob(resDataBlob.Result.Blob)
		reqData.SetPrivateKeys(PrivateKey)

		resDataSign := testSdk.Transaction.Sign(reqData)
		if resDataSign.ErrorCode != 0 {
			fmt.Println(resDataSign.ErrorDesc)
		} else {
			//Submit
			var reqData model.TransactionSubmitRequest
			reqData.SetBlob(resDataBlob.Result.Blob)
			reqData.SetSignatures(resDataSign.Result.Signatures)
			resDataSubmit := testSdk.Transaction.Submit(reqData)

			if resDataSubmit.ErrorCode != 0 {
				t.Errorf(resDataSubmit.ErrorDesc)
			} else {
				fmt.Println("Hash:", resDataSubmit.Result.Hash)
				t.Log("Test_Transaction_BuildBlob_Sign_Submit succeed", resDataSubmit.Result)
			}
		}
	}
}

//Enquiry of transaction details
func Test_Transaction_GetInfo(t *testing.T) {
	var reqData model.TransactionGetInfoRequest
	var hash string = "c738fb80dc401d6aba2cf3802ec85ac07fbc23366c003537b64cd1a59ab307d8"
	reqData.SetHash(hash)
	resData := testSdk.Transaction.GetInfo(reqData)
	if resData.ErrorCode != 0 {
		t.Errorf(resData.ErrorDesc)
	} else {
		data, _ := json.Marshal(resData.Result)
		fmt.Println("info:", string(data))
		t.Log("Test_Transaction_GetInfo succeed", resData.Result)
	}
}

//Get block height
func Test_Block_GetNumber(t *testing.T) {
	resData := testSdk.Block.GetNumber()
	if resData.ErrorCode != 0 {
		t.Errorf(resData.ErrorDesc)
	} else {
		fmt.Println("BlockNumber:", resData.Result.Header.BlockNumber)
		t.Log("Test_Block_GetNumber", resData.Result)
	}
}

//Get block details
func Test_Block_GetInfo(t *testing.T) {
	var reqData model.BlockGetInfoRequest
	var blockNumber int64 = 581283
	reqData.SetBlockNumber(blockNumber)
	resData := testSdk.Block.GetInfo(reqData)
	if resData.ErrorCode != 0 {
		t.Errorf(resData.ErrorDesc)
	} else {
		data, _ := json.Marshal(resData.Result.Header)
		fmt.Println("Header:", string(data))
		t.Log("Test_Block_GetInfo succeed", resData.Result)
	}
}

//Get the latest block information
func Test_Block_GetLatest(t *testing.T) {
	resData := testSdk.Block.GetLatest()
	if resData.ErrorCode != 0 {
		t.Errorf(resData.ErrorDesc)
	} else {
		data, _ := json.Marshal(resData.Result.Header)
		fmt.Println("Header:", string(data))
		t.Log("Test_Block_GetLatest succeed", resData.Result)
	}
}

//evaluate fee
func Test_Transaction_EvaluateFee(t *testing.T) {
	var reqDataOperation model.BUSendOperation
	reqDataOperation.Init()
	var amount int64 = 100
	reqDataOperation.SetAmount(amount)
	reqDataOperation.SetMetadata("63")
	var destAddress string = "buQemmMwmRQY1JkcU7w3nhruoX5N3j6C29uo"
	reqDataOperation.SetDestAddress(destAddress)

	var reqDataEvaluate model.TransactionEvaluateFeeRequest
	var sourceAddress string = "buQVU86Jm4FeRW4JcQTD9Rx9NkUkHikYGp6z"
	reqDataEvaluate.SetSourceAddress(sourceAddress)
	var nonce int64 = 5
	reqDataEvaluate.SetNonce(nonce)
	var signatureNumber string = "3"
	reqDataEvaluate.SetSignatureNumber(signatureNumber)
	var SetCeilLedgerSeq int64 = 50
	reqDataEvaluate.SetCeilLedgerSeq(SetCeilLedgerSeq)
	reqDataEvaluate.SetMetadata("63")
	reqDataEvaluate.SetOperation(reqDataOperation)
	resDataEvaluate := testSdk.Transaction.EvaluateFee(reqDataEvaluate)
	if resDataEvaluate.ErrorCode != 0 {
		fmt.Println(resDataEvaluate)
		t.Errorf(resDataEvaluate.ErrorDesc)
	} else {
		data, _ := json.Marshal(resDataEvaluate.Result)
		fmt.Println("Evaluate:", string(data))
		t.Log("Test_EvaluateFee succeed", resDataEvaluate.Result)
	}
}
