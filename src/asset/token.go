// token
package asset

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/bumoproject/bumo-sdk-go/src/common"
	"github.com/bumoproject/bumo-sdk-go/src/crypto/keypair"
	"github.com/bumoproject/bumo-sdk-go/src/exception"
	"github.com/bumoproject/bumo-sdk-go/src/model"
)

type TokenOperation struct {
	Url string
}

//获取Allowance
func (token *TokenOperation) Allowance(reqData model.TokenAllowanceRequest) model.TokenAllowanceResponse {
	var resDataC model.TokenCallResponse
	var resData model.TokenAllowanceResponse
	if !keypair.CheckAddress(reqData.GetContractAddress()) {
		resData.ErrorCode = exception.INVALID_CONTRACTADDRESS_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	if !keypair.CheckAddress(reqData.GetTokenOwner()) {
		resData.ErrorCode = exception.INVALID_TOKENOWNER_ERRPR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	if !keypair.CheckAddress(reqData.GetSpender()) {
		resData.ErrorCode = exception.INVALID_SPENDER_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	var Input model.Input
	Input.Method = "allowance"
	Input.Params.Owner = reqData.GetTokenOwner()
	Input.Params.Spender = reqData.GetSpender()
	InputStr, err := json.Marshal(Input)
	if err != nil {
		resData.ErrorCode = exception.SYSTEM_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	callData := model.CallContractRequest{
		ContractAddress: reqData.GetContractAddress(),
		Code:            model.Payload,
		Input:           string(InputStr),
		OptType:         2,
	}
	callDataStr, err := json.Marshal(callData)
	if err != nil {
		resData.ErrorCode = exception.SYSTEM_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	post := "/callContract"
	response, SDKRes := common.PostRequest(token.Url, post, callDataStr)
	if SDKRes.ErrorCode != 0 {
		resData.ErrorCode = exception.SYSTEM_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	defer response.Body.Close()
	if response.StatusCode == 200 {
		decoder := json.NewDecoder(response.Body)
		decoder.UseNumber()
		err = decoder.Decode(&resDataC)
		if err != nil {
			resData.ErrorCode = exception.SYSTEM_ERROR
			resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
			return resData
		}
		if resDataC.ErrorCode == 0 {
			if resDataC.Result.QueryRets == nil {
				resData.ErrorCode = exception.THE_QUERY_FAILED
				resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
				return resData
			}
			if resDataC.Result.QueryRets[0].Error.Data.Exception != "" {
				resData.ErrorCode = exception.THE_QUERY_FAILED
				resData.ErrorDesc = resDataC.Result.QueryRets[0].Error.Data.Exception
				return resData
			}
			strReader := strings.NewReader(resDataC.Result.QueryRets[0].Result.Value)
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
		} else {
			resData.ErrorCode = resData.ErrorCode
			resData.ErrorDesc = resData.ErrorDesc
			return resData
		}
	} else {
		resData.ErrorCode = exception.CONNECTNETWORK_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
}

//获取合约token的信息
func (token *TokenOperation) GetInfo(reqData model.TokenGetInfoRequest) model.TokenGetInfoResponse {
	var resDataC model.TokenCallResponse
	var resData model.TokenGetInfoResponse
	callDataStr, SDKRes := common.GetCallDataStr("contractInfo", reqData.GetContractAddress(), "")
	if SDKRes.ErrorCode != 0 {
		resData.ErrorCode = SDKRes.ErrorCode
		resData.ErrorDesc = SDKRes.ErrorDesc
		return resData
	}
	post := "/callContract"
	response, SDKRes := common.PostRequest(token.Url, post, []byte(callDataStr))
	if SDKRes.ErrorCode != 0 {
		resData.ErrorCode = exception.SYSTEM_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	defer response.Body.Close()
	if response.StatusCode == 200 {
		decoder := json.NewDecoder(response.Body)
		decoder.UseNumber()
		err := decoder.Decode(&resDataC)
		if err != nil {
			resData.ErrorCode = exception.SYSTEM_ERROR
			resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
			return resData
		}
		if resDataC.ErrorCode == 0 {
			if resDataC.Result.QueryRets == nil {
				resData.ErrorCode = exception.THE_QUERY_FAILED
				resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
				return resData
			}
			if resDataC.Result.QueryRets[0].Error.Data.Exception != "" {
				resData.ErrorCode = exception.THE_QUERY_FAILED
				resData.ErrorDesc = resDataC.Result.QueryRets[0].Error.Data.Exception
				return resData
			}
			strReader := strings.NewReader(resDataC.Result.QueryRets[0].Result.Value)
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
		} else {
			resData.ErrorCode = resDataC.ErrorCode
			resData.ErrorDesc = resDataC.ErrorDesc
			return resData
		}
	} else {
		resData.ErrorCode = exception.CONNECTNETWORK_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
}

//获取合约token的名称
func (token *TokenOperation) GetName(reqData model.TokenGetNameRequest) model.TokenGetNameResponse {
	var resDataC model.TokenCallResponse
	var resData model.TokenGetNameResponse
	callDataStr, SDKRes := common.GetCallDataStr("name", reqData.GetContractAddress(), "")
	if SDKRes.ErrorCode != 0 {
		resData.ErrorCode = SDKRes.ErrorCode
		resData.ErrorDesc = SDKRes.ErrorDesc
		return resData
	}
	post := "/callContract"
	response, SDKRes := common.PostRequest(token.Url, post, []byte(callDataStr))
	if SDKRes.ErrorCode != 0 {
		resData.ErrorCode = exception.SYSTEM_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	defer response.Body.Close()
	if response.StatusCode == 200 {
		decoder := json.NewDecoder(response.Body)
		decoder.UseNumber()
		err := decoder.Decode(&resDataC)
		if err != nil {
			resData.ErrorCode = exception.SYSTEM_ERROR
			resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
			return resData
		}
		if resDataC.ErrorCode == 0 {
			if resDataC.Result.QueryRets == nil {
				resData.ErrorCode = exception.THE_QUERY_FAILED
				resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
				return resData
			}
			if resDataC.Result.QueryRets[0].Error.Data.Exception != "" {
				resData.ErrorCode = exception.THE_QUERY_FAILED
				resData.ErrorDesc = resDataC.Result.QueryRets[0].Error.Data.Exception
				return resData
			}
			strReader := strings.NewReader(resDataC.Result.QueryRets[0].Result.Value)
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
		} else {
			resData.ErrorCode = resDataC.ErrorCode
			resData.ErrorDesc = resDataC.ErrorDesc
			return resData
		}
	} else {
		resData.ErrorCode = exception.CONNECTNETWORK_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
}

//获取合约token的符号
func (token *TokenOperation) GetSymbol(reqData model.TokenGetSymbolRequest) model.TokenGetSymbolResponse {
	var resDataC model.TokenCallResponse
	var resData model.TokenGetSymbolResponse
	callDataStr, SDKRes := common.GetCallDataStr("symbol", reqData.GetContractAddress(), "")
	if SDKRes.ErrorCode != 0 {
		resData.ErrorCode = SDKRes.ErrorCode
		resData.ErrorDesc = SDKRes.ErrorDesc
		return resData
	}
	post := "/callContract"
	response, SDKRes := common.PostRequest(token.Url, post, []byte(callDataStr))
	if SDKRes.ErrorCode != 0 {
		resData.ErrorCode = exception.SYSTEM_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	defer response.Body.Close()
	if response.StatusCode == 200 {
		decoder := json.NewDecoder(response.Body)
		decoder.UseNumber()
		err := decoder.Decode(&resDataC)
		if err != nil {
			resData.ErrorCode = exception.SYSTEM_ERROR
			resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
			return resData
		}
		if resDataC.ErrorCode == 0 {
			if resDataC.Result.QueryRets == nil {
				resData.ErrorCode = exception.THE_QUERY_FAILED
				resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
				return resData
			}
			if resDataC.Result.QueryRets[0].Error.Data.Exception != "" {
				resData.ErrorCode = exception.THE_QUERY_FAILED
				resData.ErrorDesc = resDataC.Result.QueryRets[0].Error.Data.Exception
				return resData
			}
			strReader := strings.NewReader(resDataC.Result.QueryRets[0].Result.Value)
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
		} else {
			resData.ErrorCode = resDataC.ErrorCode
			resData.ErrorDesc = resDataC.ErrorDesc
			return resData
		}
	} else {
		resData.ErrorCode = exception.CONNECTNETWORK_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
}

//获取合约token的精度
func (token *TokenOperation) GetDecimals(reqData model.TokenGetDecimalsRequest) model.TokenGetDecimalsResponse {
	var resDataC model.TokenCallResponse
	var resData model.TokenGetDecimalsResponse
	callDataStr, SDKRes := common.GetCallDataStr("decimals", reqData.GetContractAddress(), "")
	if SDKRes.ErrorCode != 0 {
		resData.ErrorCode = SDKRes.ErrorCode
		resData.ErrorDesc = SDKRes.ErrorDesc
		return resData
	}
	post := "/callContract"
	response, SDKRes := common.PostRequest(token.Url, post, []byte(callDataStr))
	if SDKRes.ErrorCode != 0 {
		resData.ErrorCode = exception.SYSTEM_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	defer response.Body.Close()
	if response.StatusCode == 200 {
		decoder := json.NewDecoder(response.Body)
		decoder.UseNumber()
		err := decoder.Decode(&resDataC)
		if err != nil {
			resData.ErrorCode = exception.SYSTEM_ERROR
			resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
			return resData
		}
		if resDataC.ErrorCode == 0 {
			if resDataC.Result.QueryRets == nil {
				resData.ErrorCode = exception.THE_QUERY_FAILED
				resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
				return resData
			}
			if resDataC.Result.QueryRets[0].Error.Data.Exception != "" {
				resData.ErrorCode = exception.THE_QUERY_FAILED
				resData.ErrorDesc = resDataC.Result.QueryRets[0].Error.Data.Exception
				return resData
			}
			strReader := strings.NewReader(resDataC.Result.QueryRets[0].Result.Value)
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
		} else {
			resData.ErrorCode = resDataC.ErrorCode
			resData.ErrorDesc = resDataC.ErrorDesc
			return resData
		}
	} else {
		resData.ErrorCode = exception.CONNECTNETWORK_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
}

//获取合约token的总供应量
func (token *TokenOperation) GetTotalSupply(reqData model.TokenGetTotalSupplyRequest) model.TokenGetTotalSupplyResponse {
	var resDataC model.TokenCallResponse
	var resData model.TokenGetTotalSupplyResponse
	callDataStr, SDKRes := common.GetCallDataStr("totalSupply", reqData.GetContractAddress(), "")
	if SDKRes.ErrorCode != 0 {
		resData.ErrorCode = SDKRes.ErrorCode
		resData.ErrorDesc = SDKRes.ErrorDesc
		return resData
	}
	post := "/callContract"
	response, SDKRes := common.PostRequest(token.Url, post, []byte(callDataStr))
	if SDKRes.ErrorCode != 0 {
		resData.ErrorCode = exception.SYSTEM_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	defer response.Body.Close()
	if response.StatusCode == 200 {
		decoder := json.NewDecoder(response.Body)
		decoder.UseNumber()
		err := decoder.Decode(&resDataC)
		if err != nil {
			resData.ErrorCode = exception.SYSTEM_ERROR
			resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
			return resData
		}
		if resDataC.ErrorCode == 0 {
			if resDataC.Result.QueryRets == nil {
				resData.ErrorCode = exception.THE_QUERY_FAILED
				resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
				return resData
			}
			if resDataC.Result.QueryRets[0].Error.Data.Exception != "" {
				resData.ErrorCode = exception.THE_QUERY_FAILED
				resData.ErrorDesc = resDataC.Result.QueryRets[0].Error.Data.Exception
				return resData
			}
			strReader := strings.NewReader(resDataC.Result.QueryRets[0].Result.Value)
			decoder := json.NewDecoder(strReader)
			decoder.UseNumber()
			valueData := make(map[string]string)
			err = decoder.Decode(&valueData)
			if err != nil {
				resData.ErrorCode = exception.SYSTEM_ERROR
				resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
				return resData
			}
			resData.Result.TotalSupply, err = strconv.ParseInt(valueData["totalSupply"], 10, 64)
			if err != nil {
				resData.ErrorCode = exception.SYSTEM_ERROR
				resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
				return resData
			}
			resData.ErrorCode = exception.SUCCESS
			return resData
		} else {
			resData.ErrorCode = resDataC.ErrorCode
			resData.ErrorDesc = resDataC.ErrorDesc
			return resData
		}
	} else {
		resData.ErrorCode = exception.CONNECTNETWORK_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
}

//获取合约token拥有者的账户余额
func (token *TokenOperation) GetBalance(reqData model.TokenGetBalanceRequest) model.TokenGetBalanceResponse {
	var resDataC model.TokenCallResponse
	var resData model.TokenGetBalanceResponse
	callDataStr, SDKRes := common.GetCallDataStr("balanceOf", reqData.GetContractAddress(), reqData.GetTokenOwner())
	if SDKRes.ErrorCode != 0 {
		resData.ErrorCode = SDKRes.ErrorCode
		resData.ErrorDesc = SDKRes.ErrorDesc
		return resData
	}
	post := "/callContract"
	response, SDKRes := common.PostRequest(token.Url, post, []byte(callDataStr))
	if SDKRes.ErrorCode != 0 {
		resData.ErrorCode = exception.SYSTEM_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	defer response.Body.Close()
	if response.StatusCode == 200 {
		decoder := json.NewDecoder(response.Body)
		decoder.UseNumber()
		err := decoder.Decode(&resDataC)
		if err != nil {
			resData.ErrorCode = exception.SYSTEM_ERROR
			resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
			return resData
		}
		if resDataC.ErrorCode == 0 {
			if resDataC.Result.QueryRets == nil {
				resData.ErrorCode = exception.THE_QUERY_FAILED
				resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
				return resData
			}
			if resDataC.Result.QueryRets[0].Error.Data.Exception != "" {
				resData.ErrorCode = exception.THE_QUERY_FAILED
				resData.ErrorDesc = resDataC.Result.QueryRets[0].Error.Data.Exception
				return resData
			}
			strReader := strings.NewReader(resDataC.Result.QueryRets[0].Result.Value)
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
		} else {
			resData.ErrorCode = resDataC.ErrorCode
			resData.ErrorDesc = resDataC.ErrorDesc
			return resData
		}
	} else {
		resData.ErrorCode = exception.CONNECTNETWORK_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
}
