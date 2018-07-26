// token
package asset

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/bumoproject/bumo-sdk-go/src/account"
	"github.com/bumoproject/bumo-sdk-go/src/contract"
	"github.com/bumoproject/bumo-sdk-go/src/crypto/keypair"
	"github.com/bumoproject/bumo-sdk-go/src/exception"
	"github.com/bumoproject/bumo-sdk-go/src/model"
)

type TokenOperation struct {
	Url string
}

//该接口用于检测Token是否有效
func (token *TokenOperation) CheckValid(reqData model.TokenCheckValidRequest) model.TokenCheckValidResponse {
	var Contract contract.ContractOperation
	Contract.Url = token.Url
	var Account account.AccountOperation
	Account.Url = token.Url
	var resData model.TokenCheckValidResponse
	resData.Result.IsValid = false
	var raqDataCheck model.ContractCheckValidRequest
	raqDataCheck.SetAddress(reqData.GetContractAddress())
	rasDataCheck := Contract.CheckValid(raqDataCheck)
	if rasDataCheck.ErrorCode != 0 {
		resData.ErrorCode = rasDataCheck.ErrorCode
		resData.ErrorDesc = rasDataCheck.ErrorDesc
		return resData
	}
	var raqDataMetadata model.AccountGetMetadataRequest
	raqDataMetadata.SetAddress(reqData.GetContractAddress())
	raqDataMetadata.SetKey("global_attribute")
	rasDataMetadata := Account.GetMetadata(raqDataMetadata)
	if rasDataMetadata.ErrorCode == 0 {
		var data model.Params
		strReader := strings.NewReader(rasDataMetadata.Result.Metadatas[0].Value)
		decoder := json.NewDecoder(strReader)
		decoder.UseNumber()
		err := decoder.Decode(&data)
		if err != nil {
			SDKRes := exception.GetSDKRes(exception.SYSTEM_ERROR)
			resData.ErrorCode = SDKRes.ErrorCode
			resData.ErrorDesc = SDKRes.ErrorDesc
			return resData
		}
		balance, err := strconv.ParseInt(data.Balance, 10, 64)
		if err != nil {
			resData.ErrorCode = exception.INVALID_CONTRACTADDRESS_ERROR
			resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
			return resData
		}
		if balance <= 0 {
			resData.ErrorCode = exception.INVALID_CONTRACTADDRESS_ERROR
			resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
			return resData
		}
		if data.Decimals < 0 || data.Decimals > 8 {
			resData.ErrorCode = exception.INVALID_CONTRACTADDRESS_ERROR
			resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
			return resData
		}
		if len(data.Name) > 1024 {
			resData.ErrorCode = exception.INVALID_CONTRACTADDRESS_ERROR
			resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
			return resData
		}
		if len(data.Symbol) > 1024 {
			resData.ErrorCode = exception.INVALID_TOKEN_SIMBOL_ERROR
			resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
			return resData
		}
		totalSupply, err := strconv.ParseInt(data.TotalSupply, 10, 64)
		if err != nil {
			resData.ErrorCode = exception.INVALID_CONTRACTADDRESS_ERROR
			resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
			return resData
		}
		if totalSupply < 0 {
			resData.ErrorCode = exception.INVALID_CONTRACTADDRESS_ERROR
			resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
			return resData
		}
		if !keypair.CheckAddress(data.TokenOwner) {
			resData.ErrorCode = exception.INVALID_CONTRACTADDRESS_ERROR
			resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
			return resData
		}
		if data.Ctp != "1.0" {
			resData.ErrorCode = exception.INVALID_CONTRACTADDRESS_ERROR
			resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
			return resData
		}
		resData.Result.IsValid = true
		return resData
	} else if rasDataMetadata.ErrorCode == exception.NO_METADATA_ERROR {
		resData.ErrorCode = exception.INVALID_CONTRACTADDRESS_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	} else {
		resData.ErrorCode = rasDataMetadata.ErrorCode
		resData.ErrorDesc = rasDataMetadata.ErrorDesc
		return resData
	}
}

//获取Allowance
func (token *TokenOperation) Allowance(reqData model.TokenAllowanceRequest) model.TokenAllowanceResponse {
	var resData model.TokenAllowanceResponse
	var reqDataCheck model.TokenCheckValidRequest
	reqDataCheck.SetContractAddress(reqData.GetContractAddress())
	resDataCheck := token.CheckValid(reqDataCheck)
	if resDataCheck.ErrorCode != 0 {
		resData.ErrorCode = exception.INVALID_CONTRACTADDRESS_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	if !keypair.CheckAddress(reqData.GetTokenOwner()) {
		resData.ErrorCode = exception.INVALID_TOKENOWNER_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	if !keypair.CheckAddress(reqData.GetSpender()) {
		resData.ErrorCode = exception.INVALID_SPENDER_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	var Contract contract.ContractOperation
	Contract.Url = token.Url
	var reqDataCall model.ContractCallRequest
	reqDataCall.SetContractAddress("buQXoNR24p2pPqnXPyiDprmTWsU4SYLtBNCG")
	reqDataCall.SetOptType(2)
	var Input model.Input
	Input.Method = "allowance"
	Input.Params.Owner = reqData.GetTokenOwner()
	Input.Params.Spender = reqData.GetSpender()
	InputByte, err := json.Marshal(Input)
	if err != nil {
		resData.ErrorCode = exception.SYSTEM_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	reqDataCall.SetInput(string(InputByte))
	resDataCall := Contract.Call(reqDataCall)
	if resDataCall.ErrorCode != 0 {
		resData.ErrorCode = resDataCall.ErrorCode
		resData.ErrorDesc = resDataCall.ErrorDesc
		return resData
	} else {
		if resDataCall.Result.QueryRets[0].Error.Data.Exception != "" {
			resData.ErrorCode = exception.GET_ALLOWNANCE_ERROR
			resData.ErrorDesc = resDataCall.Result.QueryRets[0].Error.Data.Exception
			return resData
		}
		strReader := strings.NewReader(resDataCall.Result.QueryRets[0].Result.Value)
		decoder := json.NewDecoder(strReader)
		decoder.UseNumber()
		dataStr := make(map[string]string)
		err = decoder.Decode(&dataStr)
		if err != nil {
			resData.ErrorCode = exception.SYSTEM_ERROR
			resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
			return resData
		}
		resData.Result.Allowance, err = strconv.ParseInt(dataStr["allowance"], 10, 64)
		if err != nil {
			resData.ErrorCode = exception.SYSTEM_ERROR
			resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
			return resData
		}
		resData.ErrorCode = exception.SUCCESS
		return resData
	}
}

//获取合约token的信息
func (token *TokenOperation) GetInfo(reqData model.TokenGetInfoRequest) model.TokenGetInfoResponse {
	var resData model.TokenGetInfoResponse
	var reqDataCheck model.TokenCheckValidRequest
	reqDataCheck.SetContractAddress(reqData.GetContractAddress())
	resDataCheck := token.CheckValid(reqDataCheck)
	if resDataCheck.ErrorCode != 0 {
		resData.ErrorCode = exception.INVALID_CONTRACTADDRESS_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	var Contract contract.ContractOperation
	Contract.Url = token.Url
	var reqDataCall model.ContractCallRequest
	reqDataCall.SetContractAddress("buQXoNR24p2pPqnXPyiDprmTWsU4SYLtBNCG")
	reqDataCall.SetOptType(2)
	var Input model.Input
	Input.Method = "contractInfo"
	InputByte, err := json.Marshal(Input)
	if err != nil {
		resData.ErrorCode = exception.SYSTEM_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	reqDataCall.SetInput(string(InputByte))
	resDataCall := Contract.Call(reqDataCall)
	if resDataCall.ErrorCode != 0 {
		resData.ErrorCode = resDataCall.ErrorCode
		resData.ErrorDesc = resDataCall.ErrorDesc
		return resData
	} else {
		if resDataCall.Result.QueryRets[0].Error.Data.Exception != "" {
			resData.ErrorCode = exception.GET_TOKEN_INFO_ERROR
			resData.ErrorDesc = resDataCall.Result.QueryRets[0].Error.Data.Exception
			return resData
		}
		strReader := strings.NewReader(resDataCall.Result.QueryRets[0].Result.Value)
		decoder := json.NewDecoder(strReader)
		decoder.UseNumber()
		var valueData model.Value
		err = decoder.Decode(&valueData)
		if err != nil {
			resData.ErrorCode = exception.SYSTEM_ERROR
			resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
			return resData
		}
		resData.Result.Decimals = valueData.ContractInfo.Decimals
		resData.Result.Name = valueData.ContractInfo.Name
		resData.Result.Symbol = valueData.ContractInfo.Symbol
		resData.Result.TotalSupply, err = strconv.ParseInt(valueData.ContractInfo.TotalSupply, 10, 64)
		if err != nil {
			resData.ErrorCode = exception.SYSTEM_ERROR
			resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
			return resData
		}
		resData.ErrorCode = exception.SUCCESS
		return resData
	}
}

//获取合约token的名称
func (token *TokenOperation) GetName(reqData model.TokenGetNameRequest) model.TokenGetNameResponse {
	var resData model.TokenGetNameResponse
	var reqDataCheck model.TokenCheckValidRequest
	reqDataCheck.SetContractAddress(reqData.GetContractAddress())
	resDataCheck := token.CheckValid(reqDataCheck)
	if resDataCheck.ErrorCode != 0 {
		resData.ErrorCode = exception.INVALID_CONTRACTADDRESS_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	var Contract contract.ContractOperation
	Contract.Url = token.Url
	var reqDataCall model.ContractCallRequest
	reqDataCall.SetContractAddress("buQXoNR24p2pPqnXPyiDprmTWsU4SYLtBNCG")
	reqDataCall.SetOptType(2)
	var Input model.Input
	Input.Method = "name"
	InputByte, err := json.Marshal(Input)
	if err != nil {
		resData.ErrorCode = exception.SYSTEM_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	reqDataCall.SetInput(string(InputByte))
	resDataCall := Contract.Call(reqDataCall)
	if resDataCall.ErrorCode != 0 {
		resData.ErrorCode = resDataCall.ErrorCode
		resData.ErrorDesc = resDataCall.ErrorDesc
		return resData
	} else {
		if resDataCall.Result.QueryRets[0].Error.Data.Exception != "" {
			resData.ErrorCode = exception.GET_TOKEN_INFO_ERROR
			resData.ErrorDesc = resDataCall.Result.QueryRets[0].Error.Data.Exception
			return resData
		}
		strReader := strings.NewReader(resDataCall.Result.QueryRets[0].Result.Value)
		decoder := json.NewDecoder(strReader)
		decoder.UseNumber()
		valueData := make(map[string]string)
		err = decoder.Decode(&valueData)
		if err != nil {
			resData.ErrorCode = exception.SYSTEM_ERROR
			resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
			return resData
		}
		resData.Result.Name = valueData["name"]
		resData.ErrorCode = exception.SUCCESS
		return resData
	}
}

//获取合约token的符号
func (token *TokenOperation) GetSymbol(reqData model.TokenGetSymbolRequest) model.TokenGetSymbolResponse {

	var resData model.TokenGetSymbolResponse
	var reqDataCheck model.TokenCheckValidRequest
	reqDataCheck.SetContractAddress(reqData.GetContractAddress())
	resDataCheck := token.CheckValid(reqDataCheck)
	if resDataCheck.ErrorCode != 0 {
		resData.ErrorCode = exception.INVALID_CONTRACTADDRESS_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	var Contract contract.ContractOperation
	Contract.Url = token.Url
	var reqDataCall model.ContractCallRequest
	reqDataCall.SetContractAddress("buQXoNR24p2pPqnXPyiDprmTWsU4SYLtBNCG")
	reqDataCall.SetOptType(2)
	var Input model.Input
	Input.Method = "symbol"
	InputByte, err := json.Marshal(Input)
	if err != nil {
		resData.ErrorCode = exception.SYSTEM_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	reqDataCall.SetInput(string(InputByte))
	resDataCall := Contract.Call(reqDataCall)
	if resDataCall.ErrorCode != 0 {
		resData.ErrorCode = resDataCall.ErrorCode
		resData.ErrorDesc = resDataCall.ErrorDesc
		return resData
	} else {
		if resDataCall.Result.QueryRets[0].Error.Data.Exception != "" {
			resData.ErrorCode = exception.GET_TOKEN_INFO_ERROR
			resData.ErrorDesc = resDataCall.Result.QueryRets[0].Error.Data.Exception
			return resData
		}
		strReader := strings.NewReader(resDataCall.Result.QueryRets[0].Result.Value)
		decoder := json.NewDecoder(strReader)
		decoder.UseNumber()
		valueData := make(map[string]string)
		err = decoder.Decode(&valueData)
		if err != nil {
			resData.ErrorCode = exception.SYSTEM_ERROR
			resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
			return resData
		}
		resData.Result.Symbol = valueData["symbol"]
		resData.ErrorCode = exception.SUCCESS
		return resData
	}
}

//获取合约token的精度
func (token *TokenOperation) GetDecimals(reqData model.TokenGetDecimalsRequest) model.TokenGetDecimalsResponse {
	var resData model.TokenGetDecimalsResponse
	var reqDataCheck model.TokenCheckValidRequest
	reqDataCheck.SetContractAddress(reqData.GetContractAddress())
	resDataCheck := token.CheckValid(reqDataCheck)
	if resDataCheck.ErrorCode != 0 {
		resData.ErrorCode = exception.INVALID_CONTRACTADDRESS_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	var Contract contract.ContractOperation
	Contract.Url = token.Url
	var reqDataCall model.ContractCallRequest
	reqDataCall.SetContractAddress("buQXoNR24p2pPqnXPyiDprmTWsU4SYLtBNCG")
	reqDataCall.SetOptType(2)
	var Input model.Input
	Input.Method = "decimals"
	InputByte, err := json.Marshal(Input)
	if err != nil {
		resData.ErrorCode = exception.SYSTEM_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	reqDataCall.SetInput(string(InputByte))
	resDataCall := Contract.Call(reqDataCall)
	if resDataCall.ErrorCode != 0 {
		resData.ErrorCode = resDataCall.ErrorCode
		resData.ErrorDesc = resDataCall.ErrorDesc
		return resData
	} else {
		if resDataCall.Result.QueryRets[0].Error.Data.Exception != "" {
			resData.ErrorCode = exception.GET_TOKEN_INFO_ERROR
			resData.ErrorDesc = resDataCall.Result.QueryRets[0].Error.Data.Exception
			return resData
		}
		strReader := strings.NewReader(resDataCall.Result.QueryRets[0].Result.Value)
		decoder := json.NewDecoder(strReader)
		decoder.UseNumber()
		valueData := make(map[string]interface{})
		err = decoder.Decode(&valueData)
		if err != nil {
			resData.ErrorCode = exception.SYSTEM_ERROR
			resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
			return resData
		}
		str := valueData["decimals"].(json.Number)
		resData.Result.Decimals, err = strconv.ParseInt(string(str), 10, 64)
		if err != nil {
			resData.ErrorCode = exception.SYSTEM_ERROR
			resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
			return resData
		}
		resData.ErrorCode = exception.SUCCESS
		return resData
	}
}

//获取合约token的总供应量
func (token *TokenOperation) GetTotalSupply(reqData model.TokenGetTotalSupplyRequest) model.TokenGetTotalSupplyResponse {
	var resData model.TokenGetTotalSupplyResponse
	var reqDataCheck model.TokenCheckValidRequest
	reqDataCheck.SetContractAddress(reqData.GetContractAddress())
	resDataCheck := token.CheckValid(reqDataCheck)
	if resDataCheck.ErrorCode != 0 {
		resData.ErrorCode = exception.INVALID_CONTRACTADDRESS_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	var Contract contract.ContractOperation
	Contract.Url = token.Url
	var reqDataCall model.ContractCallRequest
	reqDataCall.SetContractAddress("buQXoNR24p2pPqnXPyiDprmTWsU4SYLtBNCG")
	reqDataCall.SetOptType(2)
	var Input model.Input
	Input.Method = "totalSupply"
	InputByte, err := json.Marshal(Input)
	if err != nil {
		resData.ErrorCode = exception.SYSTEM_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	reqDataCall.SetInput(string(InputByte))
	resDataCall := Contract.Call(reqDataCall)
	if resDataCall.ErrorCode != 0 {
		resData.ErrorCode = resDataCall.ErrorCode
		resData.ErrorDesc = resDataCall.ErrorDesc
		return resData
	} else {
		if resDataCall.Result.QueryRets[0].Error.Data.Exception != "" {
			resData.ErrorCode = exception.GET_TOKEN_INFO_ERROR
			resData.ErrorDesc = resDataCall.Result.QueryRets[0].Error.Data.Exception
			return resData
		}
		strReader := strings.NewReader(resDataCall.Result.QueryRets[0].Result.Value)
		decoder := json.NewDecoder(strReader)
		decoder.UseNumber()
		valueData := make(map[string]interface{})
		err = decoder.Decode(&valueData)
		if err != nil {
			resData.ErrorCode = exception.SYSTEM_ERROR
			resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
			return resData
		}
		str := valueData["totalSupply"].(string)
		resData.Result.TotalSupply, err = strconv.ParseInt(str, 10, 64)
		if err != nil {
			resData.ErrorCode = exception.SYSTEM_ERROR
			resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
			return resData
		}
		resData.ErrorCode = exception.SUCCESS
		return resData
	}
}

//获取合约token拥有者的账户余额
func (token *TokenOperation) GetBalance(reqData model.TokenGetBalanceRequest) model.TokenGetBalanceResponse {
	var resData model.TokenGetBalanceResponse
	var reqDataCheck model.TokenCheckValidRequest
	reqDataCheck.SetContractAddress(reqData.GetContractAddress())
	resDataCheck := token.CheckValid(reqDataCheck)
	if resDataCheck.ErrorCode != 0 {
		resData.ErrorCode = exception.INVALID_CONTRACTADDRESS_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	if !keypair.CheckAddress(reqData.GetTokenOwner()) {
		resData.ErrorCode = exception.INVALID_TOKENOWNER_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	var Contract contract.ContractOperation
	Contract.Url = token.Url
	var reqDataCall model.ContractCallRequest
	reqDataCall.SetContractAddress("buQXoNR24p2pPqnXPyiDprmTWsU4SYLtBNCG")
	reqDataCall.SetOptType(2)
	var Input model.Input
	Input.Method = "balanceOf"
	Input.Params.Address = reqData.GetTokenOwner()
	InputByte, err := json.Marshal(Input)
	if err != nil {
		resData.ErrorCode = exception.SYSTEM_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	reqDataCall.SetInput(string(InputByte))
	resDataCall := Contract.Call(reqDataCall)
	if resDataCall.ErrorCode != 0 {
		resData.ErrorCode = resDataCall.ErrorCode
		resData.ErrorDesc = resDataCall.ErrorDesc
		return resData
	} else {
		if resDataCall.Result.QueryRets[0].Error.Data.Exception != "" {
			resData.ErrorCode = exception.GET_TOKEN_INFO_ERROR
			resData.ErrorDesc = resDataCall.Result.QueryRets[0].Error.Data.Exception
			return resData
		}
		strReader := strings.NewReader(resDataCall.Result.QueryRets[0].Result.Value)
		decoder := json.NewDecoder(strReader)
		decoder.UseNumber()
		valueData := make(map[string]interface{})
		err = decoder.Decode(&valueData)
		if err != nil {
			resData.ErrorCode = exception.SYSTEM_ERROR
			resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
			return resData
		}
		str := valueData["balance"].(string)
		resData.Result.Balance, err = strconv.ParseInt(str, 10, 64)
		if err != nil {
			resData.ErrorCode = exception.SYSTEM_ERROR
			resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
			return resData
		}
		resData.ErrorCode = exception.SUCCESS
		return resData
	}
}
