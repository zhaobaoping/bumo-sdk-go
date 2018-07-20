// sdk_test
package sdk_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/bumoproject/bumo-sdk-go/src/model"
	"github.com/bumoproject/bumo-sdk-go/src/sdk"
)

var testSdk sdk.Sdk

//InitSDK
func Test_InitSDK(t *testing.T) {
	url := "http://seed1.bumotest.io:26002"
	var reqData model.SDKInitSDKRequest
	reqData.Url = url
	resData := testSdk.InitSDK(reqData)
	if resData.ErrorCode != 0 {
		t.Errorf(resData.ErrorDesc)
	} else {
		t.Log("Test_NewSDK")
	}
}

//Block_GetNumber
func Test_Block_GetNumber(t *testing.T) {
	resData := testSdk.Block.GetNumber()
	if resData.ErrorCode != 0 {
		t.Errorf(resData.ErrorDesc)
	} else {
		fmt.Println("BlockNumber:", resData.Result.Header.BlockNumber)
		t.Log("Test_Block_GetNumber", resData.Result)
	}
}

//Block_CheckStatus
func Test_Block_CheckStatus(t *testing.T) {
	resData := testSdk.Block.CheckStatus()
	if resData.ErrorCode != 0 {
		t.Errorf(resData.ErrorDesc)
	} else {
		fmt.Println("IsSynchronous:", resData.Result.IsSynchronous)
		t.Log("Test_Block_CheckStatus succeed", resData.Result)
	}

}

//Block_GetTransactions
func Test_Block_GetTransactions(t *testing.T) {
	var reqData model.BlockGetTransactionRequest
	var blockNumber int64 = 581283
	reqData.SetBlockNumber(blockNumber)
	resData := testSdk.Block.GetTransactions(reqData)
	if resData.ErrorCode != 0 {
		t.Errorf(resData.ErrorDesc)
	} else {
		data, _ := json.Marshal(resData.Result.Transactions)
		fmt.Println("Transactions:", string(data))
		t.Log("Test_Block_GetTransactions succeed", resData.Result)
	}
}

//Block_GetInfo
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

//Block_GetLatest
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

//Block_GetValidators
func Test_Block_GetValidators(t *testing.T) {
	var reqData model.BlockGetValidatorsRequest
	var blockNumber int64 = 581283
	reqData.SetBlockNumber(blockNumber)
	resData := testSdk.Block.GetValidators(reqData)
	if resData.ErrorCode != 0 {
		t.Errorf(resData.ErrorDesc)
	} else {
		data, _ := json.Marshal(resData.Result.Validators)
		fmt.Println("Validators:", string(data))
		t.Log("Test_Block_GetValidators succeed", resData.Result)
	}
}

//Block_GetLatestValidators
func Test_Block_GetLatestValidators(t *testing.T) {
	resData := testSdk.Block.GetLatestValidators()
	if resData.ErrorCode != 0 {
		t.Errorf(resData.ErrorDesc)
	} else {
		data, _ := json.Marshal(resData.Result.Validators)
		fmt.Println("Validators:", string(data))
		t.Log("Test_Block_GetLatestValidators succeed", resData.Result)
	}
}

//Block_GetReward
func Test_Block_GetReward(t *testing.T) {
	var reqData model.BlockGetRewardRequest
	var blockNumber int64 = 581283
	reqData.SetBlockNumber(blockNumber)
	resData := testSdk.Block.GetReward(reqData)
	if resData.ErrorCode != 0 {
		t.Errorf(resData.ErrorDesc)
	} else {
		fmt.Println("ValidatorsReward:", resData.Result.ValidatorsReward)
		t.Log("Test_Block_GetReward succeed", resData.Result)
	}
}

//Block_GetLatestReward
func Test_Block_GetLatestReward(t *testing.T) {
	resData := testSdk.Block.GetLatestReward()
	if resData.ErrorCode != 0 {
		t.Errorf(resData.ErrorDesc)
	} else {
		fmt.Println("ValidatorsReward:", resData.Result.ValidatorsReward)
		t.Log("Test_Block_GetLatestReward succeed", resData.Result)
	}
}

//Block_GetFees
func Test_Block_GetFees(t *testing.T) {
	var reqData model.BlockGetFeesRequest
	var blockNumber int64 = 581283
	reqData.SetBlockNumber(blockNumber)
	resData := testSdk.Block.GetFees(reqData)
	if resData.ErrorCode != 0 {
		t.Errorf(resData.ErrorDesc)
	} else {
		data, _ := json.Marshal(resData.Result.Fees)
		fmt.Println("Fees:", string(data))
		t.Log("Test_Block_GetFees succeed", resData.Result)
	}
}

//Block_GetLatestFees
func Test_Block_GetLatestFees(t *testing.T) {
	resData := testSdk.Block.GetLatestFees()
	if resData.ErrorCode != 0 {
		t.Errorf(resData.ErrorDesc)
	} else {
		data, _ := json.Marshal(resData.Result.Fees)
		fmt.Println("Fees:", string(data))
		t.Log("Test_Block_GetLatestFees succeed", resData.Result)
	}
}

//Transaction_EvaluateFee
func Test_Transaction_EvaluateFee(t *testing.T) {
	var reqDataOperation model.BUSendOperation
	reqDataOperation.Init()
	var amount int64 = 100
	reqDataOperation.SetAmount(amount)
	var destAddress string = "buQVU86Jm4FeRW4JcQTD9Rx9NkUkHikYGp6z"
	reqDataOperation.SetDestAddress(destAddress)

	var reqDataEvaluate model.TransactionEvaluationFeeRequest
	var sourceAddress string = "buQVU86Jm4FeRW4JcQTD9Rx9NkUkHikYGp6z"
	reqDataEvaluate.SetSourceAddress(sourceAddress)
	var nonce int64 = 88
	reqDataEvaluate.SetNonce(nonce)
	var signatureNumber int64 = 1
	reqDataEvaluate.SetSignatureNumber(signatureNumber)
	reqDataEvaluate.SetOperation(reqDataOperation)
	resDataEvaluate := testSdk.Transaction.EvaluateFee(reqDataEvaluate)
	if resDataEvaluate.ErrorCode != 0 {
		t.Errorf(resDataEvaluate.ErrorDesc)
	} else {
		data, _ := json.Marshal(resDataEvaluate.Result)
		fmt.Println("Evaluate:", string(data))
		t.Log("Test_EvaluationFee succeed", resDataEvaluate.Result)
	}

}

//Transaction_BuildBlob_Sign_Submit
func Test_Transaction_BuildBlob_Sign_Submit(t *testing.T) {
	var reqDataOperation model.BUSendOperation
	reqDataOperation.Init()
	var amount int64 = 100
	var destAddress string = "buQVU86Jm4FeRW4JcQTD9Rx9NkUkHikYGp6z"
	reqDataOperation.SetAmount(amount)
	reqDataOperation.SetDestAddress(destAddress)

	var reqDataBlob model.TransactionBuildBlobRequest
	var sourceAddressBlob string = "buQemmMwmRQY1JkcU7w3nhruoX5N3j6C29uo"
	reqDataBlob.SetSourceAddress(sourceAddressBlob)
	var feeLimit int64 = 1000000000
	reqDataBlob.SetFeeLimit(feeLimit)
	var gasPrice int64 = 1000
	reqDataBlob.SetGasPrice(gasPrice)
	var nonce int64 = 88
	reqDataBlob.SetNonce(nonce)
	reqDataBlob.SetOperation(reqDataOperation)

	resDataBlob := testSdk.Transaction.BuildBlob(reqDataBlob)
	if resDataBlob.ErrorCode != 0 {
		fmt.Println(resDataBlob.ErrorDesc)
	} else {
		PrivateKey := []string{"privbUPxs6QGkJaNdgWS2hisny6ytx1g833cD7V9C3YET9mJ25wdcq6h"}
		var reqData model.TransactionSignRequest
		reqData.SetBlob(resDataBlob.Result.Blob)
		reqData.SetPrivateKeys(PrivateKey)

		resDataSign := testSdk.Transaction.Sign(reqData)
		if resDataSign.ErrorCode != 0 {
			fmt.Println(resDataSign.ErrorDesc)
		} else {
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

//Transaction_GetInfo
func Test_Transaction_GetInfo(t *testing.T) {
	var reqData model.TransactionGetInfoRequest
	var hash string = "cd33ad1e033d6dfe3db3a1d29a55e190935d9d1ff40a138d777e9406ebe0fdb1"
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

//Account_checkValid
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

//Account_Create
func Test_Account_Create(t *testing.T) {
	resData := testSdk.Account.Create()
	if resData.ErrorCode != 0 {
		t.Errorf(resData.ErrorDesc)
	} else {
		t.Log("Test_Account_Create", resData.Result)
	}
}

//Account_GetInfo
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

//Account_GetNonce
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

//Account_GetBalance
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

//Account_GetAssets
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

//Account_GetMetadata
func Test_Account_GetMetadata(t *testing.T) {
	var reqData model.AccountGetMetadataRequest
	var address string = "buQfnVYgXuMo3rvCEpKA6SfRrDpaz8D8A9Ea"
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

//Asset_GetInfo
func Test_Asset_GetInfo(t *testing.T) {
	var reqData model.AssetGetInfoRequest
	var address string = "buQemmMwmRQY1JkcU7w3nhruoX5N3j6C29uo"
	reqData.SetAddress(address)
	reqData.SetIssuer("buQnc3AGCo6ycWJCce516MDbPHKjK7ywwkuo")
	reqData.SetCode("HNC")
	resData := testSdk.Asset.GetInfo(reqData)
	if resData.ErrorCode != 0 {
		t.Errorf(resData.ErrorDesc)
	} else {
		data, _ := json.Marshal(resData.Result.Assets)
		fmt.Println("Assets:", string(data))
		t.Log("Test_Asset_GetInfo succeed", resData.Result.Assets)
	}
}

//Contract_GetInfo
func Test_Contract_GetInfo(t *testing.T) {
	var reqData model.ContractGetInfoRequest
	var address string = "buQfnVYgXuMo3rvCEpKA6SfRrDpaz8D8A9Ea"
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

//Token_Allowance
func Test_Token_Allowance(t *testing.T) {
	var reqData model.TokenAllowanceRequest
	var contractAddress string = "buQfnVYgXuMo3rvCEpKA6SfRrDpaz8D8A9Ea"
	reqData.SetContractAddress(contractAddress)
	var spender string = "buQemmMwmRQY1JkcU7w3nhruoX5N3j6C29uo"
	reqData.SetSpender(spender)
	var tokenOwnerr string = "buQVU86Jm4FeRW4JcQTD9Rx9NkUkHikYGp6z"
	reqData.SetTokenOwner(tokenOwnerr)
	resData := testSdk.Token.Allowance(reqData)
	if resData.ErrorCode != 0 {
		t.Errorf(resData.ErrorDesc)
	} else {
		fmt.Println("Allowance:", resData.Result.Allowance)
		t.Log("Test_Token_Allowance succeed", resData.Result)
	}
}

//Token_GetInfo
func Test_Token_GetInfo(t *testing.T) {
	var reqData model.TokenGetInfoRequest
	var contractAddress string = "buQfnVYgXuMo3rvCEpKA6SfRrDpaz8D8A9Ea"
	reqData.SetContractAddress(contractAddress)
	resData := testSdk.Token.GetInfo(reqData)
	if resData.ErrorCode != 0 {
		t.Errorf(resData.ErrorDesc)
	} else {
		data, _ := json.Marshal(resData.Result)
		fmt.Println("info:", string(data))
		t.Log("Test_Token_GetInfo succeed", resData.Result)
	}
}

//Token_GetName
func Test_Token_GetName(t *testing.T) {
	var reqData model.TokenGetNameRequest
	var contractAddress string = "buQfnVYgXuMo3rvCEpKA6SfRrDpaz8D8A9Ea"
	reqData.SetContractAddress(contractAddress)
	resData := testSdk.Token.GetName(reqData)
	if resData.ErrorCode != 0 {
		t.Errorf(resData.ErrorDesc)
	} else {
		fmt.Println("Name:", resData.Result.Name)
		t.Log("Test_Token_GetName succeed", resData.Result)
	}
}

//Token_GetSymbol
func Test_Token_GetSymbol(t *testing.T) {
	var reqData model.TokenGetSymbolRequest
	var contractAddress string = "buQfnVYgXuMo3rvCEpKA6SfRrDpaz8D8A9Ea"
	reqData.SetContractAddress(contractAddress)
	resData := testSdk.Token.GetSymbol(reqData)
	if resData.ErrorCode != 0 {
		t.Errorf(resData.ErrorDesc)
	} else {
		fmt.Println("Symbol:", resData.Result.Symbol)
		t.Log("Test_Token_GetSymbol succeed", resData.Result)
	}
}

//Token_GetDecimals
func Test_Token_GetDecimals(t *testing.T) {
	var reqData model.TokenGetDecimalsRequest
	var contractAddress string = "buQfnVYgXuMo3rvCEpKA6SfRrDpaz8D8A9Ea"
	reqData.SetContractAddress(contractAddress)
	resData := testSdk.Token.GetDecimals(reqData)
	if resData.ErrorCode != 0 {
		t.Errorf(resData.ErrorDesc)
	} else {
		fmt.Println("Decimals:", resData.Result.Decimals)
		t.Log("Test_Token_GetDecimals succeed", resData.Result)
	}
}

//Token_GetTotalSupply
func Test_Token_GetTotalSupply(t *testing.T) {
	var reqData model.TokenGetTotalSupplyRequest
	var contractAddress string = "buQfnVYgXuMo3rvCEpKA6SfRrDpaz8D8A9Ea"
	reqData.SetContractAddress(contractAddress)
	resData := testSdk.Token.GetTotalSupply(reqData)
	if resData.ErrorCode != 0 {
		t.Errorf(resData.ErrorDesc)
	} else {
		fmt.Println("TotalSupply:", resData.Result.TotalSupply)
		t.Log("Test_Token_GetTotalSupply succeed", resData.Result)
	}
}

//Token_GetBalance
func Test_Token_GetBalance(t *testing.T) {
	var reqData model.TokenGetBalanceRequest
	var contractAddress string = "buQfnVYgXuMo3rvCEpKA6SfRrDpaz8D8A9Ea"
	reqData.SetContractAddress(contractAddress)
	var tokenOwnerr string = "buQVU86Jm4FeRW4JcQTD9Rx9NkUkHikYGp6z"
	reqData.SetTokenOwner(tokenOwnerr)
	resData := testSdk.Token.GetBalance(reqData)
	if resData.ErrorCode != 0 {
		t.Errorf(resData.ErrorDesc)
	} else {
		fmt.Println("Balance:", resData.Result.Balance)
		t.Log("Test_Token_GetBalance succeed", resData.Result)
	}
}
