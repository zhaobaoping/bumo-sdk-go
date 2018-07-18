// sdk
package sdk

import (
	"github.com/bumoproject/bumo-sdk-go/src/account"
	"github.com/bumoproject/bumo-sdk-go/src/asset"
	"github.com/bumoproject/bumo-sdk-go/src/blockchain"
	"github.com/bumoproject/bumo-sdk-go/src/common"
	"github.com/bumoproject/bumo-sdk-go/src/contract"
	"github.com/bumoproject/bumo-sdk-go/src/exception"
	"github.com/bumoproject/bumo-sdk-go/src/model"
)

type Sdk struct {
	Account     account.AccountOperation
	Contract    contract.ContractOperation
	Asset       asset.AssetOperation
	Transaction blockchain.TransactionOperation
	Block       blockchain.BlockOperation
	Token       asset.TokenOperation
}

//新建
func (sdk *Sdk) InitSDK(reqData model.SDKInitSDKRequest) model.SDKInitSDKResponse {
	var resData model.SDKInitSDKResponse
	if reqData.Url == "" {
		resData.ErrorCode = exception.INVALID_BLOCKNUMBER_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	get := "/hello"
	response, SDKRes := common.GetRequest(reqData.Url, get, "")
	if SDKRes.ErrorCode != 0 {
		resData.ErrorCode = SDKRes.ErrorCode
		resData.ErrorDesc = SDKRes.ErrorDesc
		return resData
	}
	if response.StatusCode != 200 {
		resData.ErrorCode = exception.URL_EMPTY_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	defer response.Body.Close()
	sdk.Account.Url = reqData.Url
	sdk.Contract.Url = reqData.Url
	sdk.Asset.Url = reqData.Url
	sdk.Transaction.Url = reqData.Url
	sdk.Block.Url = reqData.Url
	sdk.Token.Url = reqData.Url
	resData.ErrorCode = exception.SUCCESS
	return resData
}
