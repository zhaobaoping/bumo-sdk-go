// contract
package contract

import (
	"encoding/json"
	"errors"

	"github.com/bumoproject/bumo-sdk-go/src/common"
	"github.com/bumoproject/bumo-sdk-go/src/common/proto"
	"github.com/bumoproject/bumo-sdk-go/src/common/protocol"
	"github.com/bumoproject/bumo-sdk-go/src/crypto/keypair"
	"github.com/bumoproject/bumo-sdk-go/src/exception"
)

type ContractOperation struct {
	Url string
}

var Err exception.Error

//创建合约账户
func (contract *ContractOperation) Create(sourceAddress string, destAddress string, initBalance int64, payload string, input string) ([]byte, exception.Error) {
	if initBalance < 0 {
		return nil, exception.SdkErr(exception.INVALID_INITBALANCE)
	}
	if !keypair.CheckAddress(destAddress) {
		return nil, exception.SdkErr(exception.INVALID_DESTADDRESS)
	}
	if sourceAddress != "" {
		if !keypair.CheckAddress(sourceAddress) {
			return nil, exception.SdkErr(exception.INVALID_SOURCEADDRESS)
		}
	}
	if sourceAddress == destAddress {
		return nil, exception.SdkErr(exception.DESTADDRESS_EQUAL_SOURCEADDRESS)
	}
	if payload == "" {
		return nil, exception.SdkErr(exception.INVALID_PAYLOAD)
	}
	Operations := []*protocol.Operation{
		{
			SourceAddress: sourceAddress,
			Type:          protocol.Operation_CREATE_ACCOUNT,
			CreateAccount: &protocol.OperationCreateAccount{
				DestAddress: destAddress,
				Contract: &protocol.Contract{
					Payload: payload,
				},
				InitBalance: initBalance,
				InitInput:   input,
				Priv: &protocol.AccountPrivilege{
					MasterWeight: 0,
					Thresholds: &protocol.AccountThreshold{
						TxThreshold: 1,
					},
				},
			},
		},
	}
	data, err := proto.Marshal(Operations[0])
	if err != nil {
		Err.Code = exception.PROTO_MARSHAL_ERROR
		Err.Err = err
		return nil, Err
	}
	Err.Code = exception.SUCCESS
	Err.Err = nil
	return data, Err

}

//获取合约
func (contract *ContractOperation) GetContract(address string) (string, exception.Error) {
	if !keypair.CheckAddress(address) {
		return "", exception.SdkErr(exception.INVALID_PARAMETER)
	}
	get := "/getAccount?address="
	response, Err := common.GetRequest(contract.Url, get, address)
	if Err.Err != nil {
		return "", Err
	}
	defer response.Body.Close()
	if response.StatusCode == 200 {
		data := make(map[string]interface{})
		decoder := json.NewDecoder(response.Body)
		decoder.UseNumber()
		err := decoder.Decode(&data)
		if err != nil {
			Err.Code = exception.DECODER_DECODE_ERROR
			Err.Err = err
			return "", Err
		}
		if data["error_code"].(json.Number) == "0" {
			result := data["result"].(map[string]interface{})
			contract := make(map[string]interface{})
			contract["contract"] = result["contract"]
			Mdata, err := json.Marshal(&contract)
			if err != nil {
				Err.Code = exception.JSON_MARSHAL_ERROR
				Err.Err = err
				return "", Err
			}
			Err.Code = exception.SUCCESS
			Err.Err = nil
			return string(Mdata), Err
		} else {
			if data["error_code"].(json.Number) == "4" {
				return "", exception.SdkErr(exception.ACCOUNT_NOT_EXIST)
			}
			return "", exception.GetErr(data["error_code"].(float64))
		}
	} else {
		Err.Code = response.StatusCode
		Err.Err = errors.New(response.Status)
		return "", Err

	}
}

//转移资产并触发合约
func (contract *ContractOperation) InvokeContractByAsset(sourceAddress string, destAddress string, issueAddress string, amount int64, code string, input string) ([]byte, exception.Error) {
	if sourceAddress != "" {
		if !keypair.CheckAddress(sourceAddress) {
			return nil, exception.SdkErr(exception.INVALID_SOURCEADDRESS)
		}
	}
	if !keypair.CheckAddress(destAddress) {
		return nil, exception.SdkErr(exception.INVALID_DESTADDRESS)
	}
	if sourceAddress == destAddress {
		return nil, exception.SdkErr(exception.DESTADDRESS_EQUAL_SOURCEADDRESS)
	}
	if amount < 0 {
		return nil, exception.SdkErr(exception.INVALID_AMOUNT)
	}
	if issueAddress != "" && !keypair.CheckAddress(issueAddress) {
		return nil, exception.SdkErr(exception.INVALID_ISSUEADDRESS)
	}
	if len([]rune(code)) > 64 {
		return nil, exception.SdkErr(exception.INVALID_CODE)
	}
	var PayAsset protocol.OperationPayAsset
	if code != "" && issueAddress != "" && amount > 0 {
		PayAsset = protocol.OperationPayAsset{
			DestAddress: destAddress,
			Asset: &protocol.Asset{
				Key: &protocol.AssetKey{
					Issuer: issueAddress,
					Code:   code,
				},
				Amount: amount,
			},
			Input: input,
		}
	} else {
		PayAsset = protocol.OperationPayAsset{
			DestAddress: destAddress,
			Input:       input,
		}
	}

	Operations := []*protocol.Operation{
		{
			SourceAddress: sourceAddress,
			Type:          protocol.Operation_PAY_ASSET,
			PayAsset:      &PayAsset,
		},
	}
	data, err := proto.Marshal(Operations[0])
	if err != nil {
		Err.Code = exception.PROTO_MARSHAL_ERROR
		Err.Err = err
		return nil, Err
	}
	Err.Code = exception.SUCCESS
	Err.Err = nil
	return data, Err
}

//发送BU并触发合约
func (contract *ContractOperation) InvokeContractByBU(sourceAddress string, destAddress string, amount int64, input string) ([]byte, exception.Error) {
	if sourceAddress != "" {
		if !keypair.CheckAddress(sourceAddress) {
			return nil, exception.SdkErr(exception.INVALID_SOURCEADDRESS)
		}
	}
	if !keypair.CheckAddress(destAddress) {
		return nil, exception.SdkErr(exception.INVALID_DESTADDRESS)
	}
	if sourceAddress == destAddress {
		return nil, exception.SdkErr(exception.DESTADDRESS_EQUAL_SOURCEADDRESS)
	}
	if amount < 0 {
		return nil, exception.SdkErr(exception.INVALID_AMOUNT)
	}

	Operations := []*protocol.Operation{
		{
			SourceAddress: sourceAddress,
			Type:          protocol.Operation_PAY_COIN,
			PayCoin: &protocol.OperationPayCoin{
				DestAddress: destAddress,
				Amount:      amount,
				Input:       input,
			},
		},
	}
	data, err := proto.Marshal(Operations[0])
	if err != nil {
		Err.Code = exception.PROTO_MARSHAL_ERROR
		Err.Err = err
		return nil, Err
	}
	Err.Code = exception.SUCCESS
	Err.Err = nil
	return data, Err
}
