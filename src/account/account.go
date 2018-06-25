// account
package account

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/bumoproject/bumo-sdk-go/src/asset"
	"github.com/bumoproject/bumo-sdk-go/src/common"
	"github.com/bumoproject/bumo-sdk-go/src/common/proto"
	"github.com/bumoproject/bumo-sdk-go/src/common/protocol"
	"github.com/bumoproject/bumo-sdk-go/src/crypto/keypair"
	"github.com/bumoproject/bumo-sdk-go/src/exception"
)

var Err exception.Error

type AccountOperation struct {
	Url   string
	Asset asset.AssetOperation
}

//创建账户公私玥对
func (account *AccountOperation) Create() (publicKey string, privateKey string, address string, Err exception.Error) {
	publicKey, privateKey, address, err := keypair.Create()
	if err != nil {
		Err.Code = exception.KEYPAIR_CREATE_ERROR
		Err.Err = err
		return "", "", "", Err
	}
	Err.Code = exception.SUCCESS
	Err.Err = nil
	return publicKey, privateKey, address, Err
}

//生成激活账户
func (account *AccountOperation) ActiveAccount(ActivateData common.ActivateAccountReqData) (string, exception.Error) {
	_, baseReserve, Err := common.GetFees(account.Url)
	if ActivateData.SourceAddress != "" {
		if !keypair.CheckAddress(ActivateData.SourceAddress) {
			return "", exception.SdkErr(exception.INVALID_SOURCEADDRESS)
		}
	}
	if ActivateData.SourceAddress == ActivateData.DestAddress {
		return "", exception.SdkErr(exception.DESTADDRESS_EQUAL_SOURCEADDRESS)
	}
	if ActivateData.InitBalance < baseReserve || ActivateData.InitBalance <= 0 {
		return "", exception.SdkErr(exception.INVALID_INITBALANCE)
	}
	if !keypair.CheckAddress(ActivateData.DestAddress) {
		return "", exception.SdkErr(exception.INVALID_DESTADDRESS)
	}
	Operations := []*protocol.Operation{
		{
			SourceAddress: ActivateData.SourceAddress,
			Type:          protocol.Operation_CREATE_ACCOUNT,
			CreateAccount: &protocol.OperationCreateAccount{
				DestAddress: ActivateData.DestAddress,
				Priv: &protocol.AccountPrivilege{
					MasterWeight: 1,
					Thresholds: &protocol.AccountThreshold{
						TxThreshold: 1,
					},
				},
				InitBalance: ActivateData.InitBalance,
			},
		},
	}
	data, err := proto.Marshal(Operations[0])
	if err != nil {
		Err.Code = exception.PROTO_MARSHAL_ERROR
		Err.Err = err
		return "", Err
	}
	str := string(data[:])
	Err.Code = exception.SUCCESS
	Err.Err = nil
	return str, Err

}

//设置metadata
func (account *AccountOperation) SetMetadata(sourceAddress string, key string, value string, version int64) ([]byte, exception.Error) {
	if sourceAddress != "" {
		if !keypair.CheckAddress(sourceAddress) {
			return nil, exception.SdkErr(exception.INVALID_SOURCEADDRESS)
		}
	}
	if len([]rune(key)) <= 0 || len([]rune(key)) > 1024 {
		return nil, exception.SdkErr(exception.INVALID_KEY)
	}
	if len([]rune(value)) < 0 || len([]rune(value)) > (1024*256) {
		return nil, exception.SdkErr(exception.INVALID_VALUE)
	}
	if version < 0 {
		return nil, exception.SdkErr(exception.INVALID_VERSION)
	}

	Operations := []*protocol.Operation{
		{
			SourceAddress: sourceAddress,
			Type:          protocol.Operation_SET_METADATA,
			SetMetadata: &protocol.OperationSetMetadata{
				Key:     key,
				Value:   value,
				Version: version,
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

//设置权限
func (account *AccountOperation) SetPrivilege(sourceAddress string, masterWeight string, signers []common.Signers, txThreshold string, typeThresholds []common.TypeThresholds) ([]byte, exception.Error) {
	if sourceAddress != "" {
		if !keypair.CheckAddress(sourceAddress) {
			return nil, exception.SdkErr(exception.INVALID_SOURCEADDRESS)
		}
	}
	masterWeightInt, err := strconv.ParseInt(masterWeight, 10, 32)
	if err != nil || masterWeightInt < 0 {
		return nil, exception.SdkErr(exception.INVALID_MASTERWEIGHT)
	}
	for i := range signers {
		if !keypair.CheckAddress(signers[i].SignerAddress) {
			return nil, exception.SdkErr(exception.INVALID_SIGNERADDRESS)
		}
		if signers[i].Weight > 4294967295 || signers[i].Weight < 0 {
			return nil, exception.SdkErr(exception.INVALID_WEIGHT)
		}
	}
	txThresholdInt, err := strconv.ParseInt(txThreshold, 10, 64)
	if err != nil || txThresholdInt < 0 {
		return nil, exception.SdkErr(exception.INVALID_TXTHRESHOLD)
	}
	for i := range typeThresholds {
		if typeThresholds[i].Type > 100 || typeThresholds[i].Type <= 0 {
			return nil, exception.SdkErr(exception.INVALID_THRESHOLDSTYPE)
		}
		if typeThresholds[i].Type < 0 {
			return nil, exception.SdkErr(exception.INVALID_THRESHOLDS)
		}
	}
	Signers := make([]*protocol.Signer, len(signers))
	for i := range signers {
		Signers[i] = new(protocol.Signer)
		Signers[i].Address = signers[i].SignerAddress
		Signers[i].Weight = signers[i].Weight
	}
	TypeThresholds := make([]*protocol.OperationTypeThreshold, len(typeThresholds))
	for i := range typeThresholds {
		TypeThresholds[i] = new(protocol.OperationTypeThreshold)
		TypeThresholds[i].Threshold = typeThresholds[i].Threshold
		TypeThresholds[i].Type = (protocol.Operation_Type)(typeThresholds[i].Type)
	}
	Operations := []*protocol.Operation{
		{
			SourceAddress: sourceAddress,
			Type:          protocol.Operation_SET_PRIVILEGE,
			SetPrivilege: &protocol.OperationSetPrivilege{
				MasterWeight:   masterWeight,
				Signers:        Signers,
				TxThreshold:    txThreshold,
				TypeThresholds: TypeThresholds,
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

//检查地址合法性
func (account *AccountOperation) CheckAddress(address string) bool {
	return keypair.CheckAddress(address)
}

//查询账户
func (account *AccountOperation) GetInfo(address string) (string, exception.Error) {
	if !keypair.CheckAddress(address) {
		return "", exception.SdkErr(exception.INVALID_PARAMETER)
	}
	get := "/getAccount?address="
	response, Err := common.GetRequest(account.Url, get, address)
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
			result := data["result"]
			Mdata, err := json.Marshal(&result)
			if err != nil {
				Err.Code = exception.JSON_MARSHAL_ERROR
				Err.Err = err
				return "", Err
			}
			Err.Code = exception.SUCCESS
			Err.Err = nil
			return string(Mdata), Err
		} else {
			result := data["result"]
			Mdata, err := json.Marshal(&result)
			if err != nil {
				Err.Code = exception.JSON_MARSHAL_ERROR
				Err.Err = err
				return "", Err
			}
			if data["error_code"].(json.Number) == "4" {
				return string(Mdata), exception.SdkErr(exception.ACCOUNT_NOT_EXIST)
			}
			strErr := data["error_code"].(json.Number)
			errInt, err := strconv.ParseInt(string(strErr), 10, 64)
			if err != nil {
				Err.Code = exception.STRCONV_PARSEINT_ERROR
				Err.Err = err
				return "", Err
			}
			return string(Mdata), exception.GetErr(float64(errInt))
		}
	} else {
		Err.Code = response.StatusCode
		Err.Err = errors.New("response.Status")
		return "", Err
	}
}

//查询账户余额
func (account *AccountOperation) GetBalance(address string) (int64, exception.Error) {
	if !keypair.CheckAddress(address) {
		return 0, exception.SdkErr(exception.INVALID_PARAMETER)
	}
	get := "/getAccount?address="
	response, Err := common.GetRequest(account.Url, get, address)
	if Err.Err != nil {
		return 0, Err
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
			return 0, Err
		}
		if data["error_code"].(json.Number) == "0" {
			result := data["result"].(map[string]interface{})
			balance := result["balance"].(json.Number)
			Err.Code = exception.SUCCESS
			Err.Err = nil
			balanceint64, err := strconv.ParseInt(string(balance), 10, 64)
			if err != nil {
				Err.Code = exception.STRCONV_PARSEINT_ERROR
				Err.Err = err
				return 0, Err
			}
			return balanceint64, Err
		} else {
			if err != nil {
				Err.Code = exception.JSON_MARSHAL_ERROR
				Err.Err = err
				return 0, Err
			}
			if data["error_code"].(json.Number) == "4" {
				return 0, exception.SdkErr(exception.ACCOUNT_NOT_EXIST)
			}
			strErr := data["error_code"].(json.Number)
			errInt, err := strconv.ParseInt(string(strErr), 10, 64)
			if err != nil {
				Err.Code = exception.STRCONV_PARSEINT_ERROR
				Err.Err = err
				return 0, Err
			}
			return 0, exception.GetErr(float64(errInt))
		}
	} else {
		Err.Code = response.StatusCode
		Err.Err = errors.New(response.Status)
		return 0, Err
	}
}
