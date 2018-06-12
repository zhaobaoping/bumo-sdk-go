// account
package bumo

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
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
func (account *AccountOperation) CreateActive(sourceAddress string, destaddress string, initBalance int64) ([]byte, Error) {
	if initBalance < 0 {
		return nil, sdkErr(INVALID_INITBALANCE)
	}
	if !keypair.CheckAddress(destaddress) {
		return nil, sdkErr(INVALID_DESTADDRESS)
	}
	if !keypair.CheckAddress(sourceAddress) {
		return nil, sdkErr(INVALID_SOURCEADDRESS)
	}
	Operations := []*protocol.Operation{
		{
			SourceAddress: sourceAddress,
			Type:          protocol.Operation_CREATE_ACCOUNT,
			CreateAccount: &protocol.OperationCreateAccount{
				DestAddress: destaddress,
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

//检查地址合法性
func (account *AccountOperation) CheckAddress(address string) bool {
	return keypair.CheckAddress(address)
}

//查询账户
func (account *AccountOperation) GetInfo(address string) (string, Error) {
	if !keypair.CheckAddress(address) {
		return "", sdkErr(INVALID_PARAMETER)
	}
	str := "/getAccount?address="
	var buf bytes.Buffer
	buf.WriteString(account.url)
	buf.WriteString(str)
	buf.WriteString(address)
	url := buf.String()
	client := &http.Client{}
	reqest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		Err.Code = HTTP_NEWREQUEST_ERROR
		Err.Err = err
		return "", Err
	}
	response, err := client.Do(reqest)
	if err != nil {
		Err.Code = CLIENT_DO_ERROR
		Err.Err = err
		return "", Err
	}
	if response.StatusCode == 200 {
		data := make(map[string]interface{})
		decoder := json.NewDecoder(response.Body)
		decoder.UseNumber()
		err = decoder.Decode(&data)
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
			return string(Mdata), getErr(data["error_code"].(float64))
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
	str := "/getAccount?address="
	var buf bytes.Buffer
	buf.WriteString(account.url)
	buf.WriteString(str)
	buf.WriteString(address)
	url := buf.String()
	client := &http.Client{}
	reqest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		Err.Code = HTTP_NEWREQUEST_ERROR
		Err.Err = err
		return 0, Err
	}
	response, err := client.Do(reqest)
	if err != nil {
		Err.Code = CLIENT_DO_ERROR
		Err.Err = err
		return 0, Err
	}
	if response.StatusCode == 200 {
		data := make(map[string]interface{})
		decoder := json.NewDecoder(response.Body)
		decoder.UseNumber()
		err = decoder.Decode(&data)
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
			return 0, getErr(data["error_code"].(float64))
		}
	} else {
		Err.Code = response.StatusCode
		Err.Err = errors.New(response.Status)
		return 0, Err
	}
}
