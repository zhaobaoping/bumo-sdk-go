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
func (sdk *Sdk) Init(reqData model.SDKInitRequest) model.SDKInitResponse {
	var resData model.SDKInitResponse
	if reqData.GetUrl() == "" {
		resData.ErrorCode = exception.INVALID_BLOCKNUMBER_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	get := "/hello"
	response, SDKRes := common.GetRequest(reqData.GetUrl(), get, "")
	defer response.Body.Close()
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
	sdk.Account.Url = reqData.GetUrl()
	sdk.Contract.Url = reqData.GetUrl()
	sdk.Asset.Url = reqData.GetUrl()
	sdk.Transaction.Url = reqData.GetUrl()
	sdk.Block.Url = reqData.GetUrl()
	sdk.Token.Url = reqData.GetUrl()
	resData.ErrorCode = exception.SUCCESS
	return resData
}
