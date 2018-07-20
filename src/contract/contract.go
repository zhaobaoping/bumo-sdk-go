// contract
package contract

import (
	"encoding/json"

	"github.com/bumoproject/bumo-sdk-go/src/account"
	"github.com/bumoproject/bumo-sdk-go/src/common"
	"github.com/bumoproject/bumo-sdk-go/src/exception"
	"github.com/bumoproject/bumo-sdk-go/src/model"
)

type ContractOperation struct {
	Url string
}

//检测合约账户的有效性
func (contract *ContractOperation) CheckValid(reqData model.ContractCheckValidRequest) model.ContractCheckValidResponse {
	var Account account.AccountOperation
	Account.Url = contract.Url
	var reqDataAcc model.AccountGetInfoRequest
	var resData model.ContractCheckValidResponse
	resData.Result.IsValid = false
	reqDataAcc.SetAddress(reqData.GetAddress())
	resDataAcc := Account.GetInfo(reqDataAcc)
	if resDataAcc.ErrorCode != 0 {
		resData.ErrorCode = resDataAcc.ErrorCode
		resData.ErrorDesc = resDataAcc.ErrorDesc
		return resData
	}
	if resDataAcc.Result.Priv.MasterWeight == 0 && resDataAcc.Result.Priv.Thresholds.TxThreshold == 1 {
		resData.Result.IsValid = true
		return resData
	} else {
		resData.ErrorCode = exception.INVALID_CONTRACTADDRESS_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
}

//获取合约信息
func (contract *ContractOperation) GetInfo(reqData model.ContractGetInfoRequest) model.ContractGetInfoResponse {
	var resData model.ContractGetInfoResponse
	var reqDataCheck model.ContractCheckValidRequest
	reqDataCheck.SetAddress(reqData.GetAddress())
	resDataCheck := contract.CheckValid(reqDataCheck)
	if resDataCheck.ErrorCode != 0 {
		resData.ErrorCode = resDataCheck.ErrorCode
		resData.ErrorDesc = resDataCheck.ErrorDesc
		return resData
	}
	get := "/getAccount?address="
	response, SDKRes := common.GetRequest(contract.Url, get, reqData.GetAddress())
	if SDKRes.ErrorCode != 0 {
		resData.ErrorCode = SDKRes.ErrorCode
		resData.ErrorDesc = SDKRes.ErrorDesc
		return resData
	}
	defer response.Body.Close()
	if response.StatusCode == 200 {
		decoder := json.NewDecoder(response.Body)
		decoder.UseNumber()
		err := decoder.Decode(&resData)
		if err != nil {
			SDKRes := exception.GetSDKRes(exception.SYSTEM_ERROR)
			resData.ErrorCode = SDKRes.ErrorCode
			resData.ErrorDesc = SDKRes.ErrorDesc
			return resData
		}
		if resData.ErrorCode == 0 {
			resData.ErrorCode = exception.SUCCESS
			return resData
		} else {
			if resData.ErrorCode == 4 {
				resData.ErrorDesc = "Get account failed"
				return resData
			}
			return resData
		}
	} else {
		SDKRes := exception.GetSDKRes(exception.CONNECTNETWORK_ERROR)
		resData.ErrorCode = SDKRes.ErrorCode
		resData.ErrorDesc = SDKRes.ErrorDesc
		return resData
	}
}
