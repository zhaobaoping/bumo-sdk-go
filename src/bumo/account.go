// account
package bumo

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/bumoproject/bumo-sdk-go/src/3rd/proto"
	"github.com/bumoproject/bumo-sdk-go/src/keypair"
	"github.com/bumoproject/bumo-sdk-go/src/protocol"
)

type AccountOperation struct {
	url   string
	Asset AssetOperation
}

//生成地址
func (account *AccountOperation) CreateInactive() (publicKey string, privateKey string, address string, Err Error) {
	publicKey, privateKey, address, err := keypair.Create()
	if err != nil {
		Err.Code = KEYPAIR_CREATE_ERROR
		Err.Err = err
		return "", "", "", Err
	}
	Err.Code = SUCCESS
	Err.Err = nil
	return publicKey, privateKey, address, Err
}

//创建普通账户
func (account *AccountOperation) CreateActive(sourceAddress string, destAddress string, initBalance int64) ([]byte, Error) {
	_, baseReserve, Err := getFees(account.url)
	if sourceAddress != "" {
		if !keypair.CheckAddress(sourceAddress) {
			return nil, sdkErr(INVALID_SOURCEADDRESS)
		}
	}
	if initBalance < baseReserve || initBalance <= 0 {
		return nil, sdkErr(INVALID_INITBALANCE)
	}
	if !keypair.CheckAddress(destAddress) {
		return nil, sdkErr(INVALID_DESTADDRESS)
	}
	Operations := []*protocol.Operation{
		{
			SourceAddress: sourceAddress,
			Type:          protocol.Operation_CREATE_ACCOUNT,
			CreateAccount: &protocol.OperationCreateAccount{
				DestAddress: destAddress,
				Priv: &protocol.AccountPrivilege{
					MasterWeight: 1,
					Thresholds: &protocol.AccountThreshold{
						TxThreshold: 1,
					},
				},
				InitBalance: initBalance,
			},
		},
	}
	data, err := proto.Marshal(Operations[0])
	if err != nil {
		Err.Code = PROTO_MARSHAL_ERROR
		Err.Err = err
		return nil, Err
	}
	Err.Code = SUCCESS
	Err.Err = nil
	return data, Err

}

//设置metadata
func (account *AccountOperation) SetMetadata(sourceAddress string, key string, value string, version int64) ([]byte, Error) {
	if sourceAddress != "" {
		if !keypair.CheckAddress(sourceAddress) {
			return nil, sdkErr(INVALID_SOURCEADDRESS)
		}
	}
	if len([]rune(key)) <= 0 || len([]rune(key)) > 1024 {
		return nil, sdkErr(INVALID_KEY)
	}
	if len([]rune(value)) < 0 || len([]rune(value)) > (1024*256) {
		return nil, sdkErr(INVALID_VALUE)
	}
	if version < 0 {
		return nil, sdkErr(INVALID_VERSION)
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
		Err.Code = PROTO_MARSHAL_ERROR
		Err.Err = err
		return nil, Err
	}
	Err.Code = SUCCESS
	Err.Err = nil
	return data, Err

}

//设置权限
func (account *AccountOperation) SetPrivilege(sourceAddress string, signerAddress string, masterWeight int32, weight int64, txThreshold int64, thresholdsType int32, thresholds int64) ([]byte, Error) {
	if sourceAddress != "" {
		if !keypair.CheckAddress(sourceAddress) {
			return nil, sdkErr(INVALID_SOURCEADDRESS)
		}
	}
	if !keypair.CheckAddress(signerAddress) {
		return nil, sdkErr(INVALID_SIGNERADDRESS)
	}
	if masterWeight < 0 {
		return nil, sdkErr(INVALID_MASTERWEIGHT)
	}
	if weight < 0 {
		return nil, sdkErr(INVALID_WEIGHT)
	}
	if txThreshold < 0 {
		return nil, sdkErr(INVALID_TXTHRESHOLD)
	}
	if thresholdsType <= 0 || thresholdsType > 100 {
		return nil, sdkErr(INVALID_THRESHOLDSTYPE)
	}
	if thresholds < 0 {
		return nil, sdkErr(INVALID_THRESHOLDS)
	}
	masterWeightStr := strconv.Itoa(int(masterWeight))
	txThresholdStr := strconv.FormatInt(txThreshold, 10)
	Operations := []*protocol.Operation{
		{
			SourceAddress: sourceAddress,
			Type:          protocol.Operation_CREATE_ACCOUNT,
			SetPrivilege: &protocol.OperationSetPrivilege{
				MasterWeight: masterWeightStr,
				Signers: []*protocol.Signer{
					{
						Address: signerAddress,
						Weight:  weight,
					},
				},
				TxThreshold: txThresholdStr,
				TypeThresholds: []*protocol.OperationTypeThreshold{
					{
						Type:      protocol.Operation_Type(thresholdsType),
						Threshold: thresholds,
					},
				},
			},
		},
	}
	data, err := proto.Marshal(Operations[0])
	if err != nil {
		Err.Code = PROTO_MARSHAL_ERROR
		Err.Err = err
		return nil, Err
	}
	Err.Code = SUCCESS
	Err.Err = nil
	return data, Err
}

//检查地址合法性
func (account *AccountOperation) CheckAddress(address string) bool {
	return keypair.CheckAddress(address)
}

//查询账户
func (account *AccountOperation) GetInfo(address string) (string, Error) {
	if !keypair.CheckAddress(address) {
		return "", sdkErr(INVALID_PARAMETER)
	}
	get := "/getAccount?address="
	response, Err := getRequest(account.url, get, address)
	if Err.Err != nil {
		return "", Err
	}
	if response.StatusCode == 200 {
		data := make(map[string]interface{})
		decoder := json.NewDecoder(response.Body)
		decoder.UseNumber()
		err := decoder.Decode(&data)
		if err != nil {
			Err.Code = DECODER_DECODE_ERROR
			Err.Err = err
			return "", Err
		}
		if data["error_code"].(json.Number) == "0" {
			result := data["result"]
			Mdata, err := json.Marshal(&result)
			if err != nil {
				Err.Code = JSON_MARSHAL_ERROR
				Err.Err = err
				return "", Err
			}
			Err.Code = SUCCESS
			Err.Err = nil
			return string(Mdata), Err
		} else {
			result := data["result"]
			Mdata, err := json.Marshal(&result)
			if err != nil {
				Err.Code = JSON_MARSHAL_ERROR
				Err.Err = err
				return "", Err
			}
			if data["error_code"].(json.Number) == "4" {
				return string(Mdata), sdkErr(ACCOUNT_NOT_EXIST)
			}
			strErr := data["error_code"].(json.Number)
			errInt, err := strconv.ParseInt(string(strErr), 10, 64)
			if err != nil {
				Err.Code = STRCONV_PARSEINT_ERROR
				Err.Err = err
				return "", Err
			}
			return string(Mdata), getErr(float64(errInt))
		}
	} else {
		Err.Code = response.StatusCode
		Err.Err = errors.New("response.Status")
		return "", Err
	}
}

//查询账户余额
func (account *AccountOperation) GetBalance(address string) (int64, Error) {
	if !keypair.CheckAddress(address) {
		return 0, sdkErr(INVALID_PARAMETER)
	}
	get := "/getAccount?address="
	response, Err := getRequest(account.url, get, address)
	if Err.Err != nil {
		return 0, Err
	}
	if response.StatusCode == 200 {
		data := make(map[string]interface{})
		decoder := json.NewDecoder(response.Body)
		decoder.UseNumber()
		err := decoder.Decode(&data)
		if err != nil {
			Err.Code = DECODER_DECODE_ERROR
			Err.Err = err
			return 0, Err
		}
		if data["error_code"].(json.Number) == "0" {
			result := data["result"].(map[string]interface{})
			balance := result["balance"].(json.Number)
			Err.Code = SUCCESS
			Err.Err = nil
			balanceint64, err := strconv.ParseInt(string(balance), 10, 64)
			if err != nil {
				Err.Code = STRCONV_PARSEINT_ERROR
				Err.Err = err
				return 0, Err
			}
			return balanceint64, Err
		} else {
			if err != nil {
				Err.Code = JSON_MARSHAL_ERROR
				Err.Err = err
				return 0, Err
			}
			if data["error_code"].(json.Number) == "4" {
				return 0, sdkErr(ACCOUNT_NOT_EXIST)
			}
			strErr := data["error_code"].(json.Number)
			errInt, err := strconv.ParseInt(string(strErr), 10, 64)
			if err != nil {
				Err.Code = STRCONV_PARSEINT_ERROR
				Err.Err = err
				return 0, Err
			}
			return 0, getErr(float64(errInt))
		}
	} else {
		Err.Code = response.StatusCode
		Err.Err = errors.New(response.Status)
		return 0, Err
	}
}
