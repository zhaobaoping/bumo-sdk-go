// sdk
package sdk

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"strconv"

	"github.com/bumoproject/bumo-sdk-go/src/account"
	"github.com/bumoproject/bumo-sdk-go/src/common"
	"github.com/bumoproject/bumo-sdk-go/src/common/proto"
	"github.com/bumoproject/bumo-sdk-go/src/common/protocol"
	"github.com/bumoproject/bumo-sdk-go/src/contract"
	"github.com/bumoproject/bumo-sdk-go/src/crypto/keypair"
	"github.com/bumoproject/bumo-sdk-go/src/crypto/signature"
	"github.com/bumoproject/bumo-sdk-go/src/exception"
)

var Err exception.Error

var Account account.AccountOperation
var Contract contract.ContractOperation

//新建链接
func Newbumo(ip string) exception.Error {
	if ip == "" {
		return exception.SdkErr(exception.INVALID_PARAMETER)
	}
	Account.Url = ip
	Contract.Url = ip
	Err.Code = exception.SUCCESS
	Err.Err = nil
	return Err
}

//获取区块高度
func GetBlockNumber() (int64, exception.Error) {
	get := "/getLedger"
	response, Err := common.GetRequest(Account.Url, get, "")
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
			header := result["header"].(map[string]interface{})
			seqstr := header["seq"].(json.Number)
			seq, err := strconv.ParseInt(string(seqstr), 10, 64)
			if err != nil {
				Err.Code = exception.STRCONV_PARSEINT_ERROR
				Err.Err = err
				return 0, Err
			}
			Err.Code = exception.SUCCESS
			Err.Err = nil
			return seq, Err
		} else {
			if data["error_code"].(json.Number) == "4" {
				return 0, exception.SdkErr(exception.BLOCK_NOT_EXIST)
			}
			errorCodeStr := data["error_code"].(json.Number)
			errorCode, err := strconv.ParseInt(string(errorCodeStr), 10, 64)
			if err != nil {
				Err.Code = exception.STRCONV_PARSEINT_ERROR
				Err.Err = err
				return 0, Err
			}
			return 0, exception.GetErr(float64(errorCode))
		}
	} else {
		Err.Code = response.StatusCode
		Err.Err = errors.New(response.Status)
		return 0, Err
	}
}

//检查区块同步
func CheckBlockStatus() (bool, exception.Error) {
	var ret bool
	get := "getModulesStatus"
	response, Err := common.GetRequest(Account.Url, get, "")
	if Err.Err != nil {
		return false, Err
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
			return false, Err
		}
		ledger_manager := data["ledger_manager"].(map[string]interface{})
		if ledger_manager["chain_max_ledger_seq"] == ledger_manager["ledger_sequence"] {
			ret = true
		}
	}
	Err.Code = exception.SUCCESS
	Err.Err = nil
	return ret, Err
}

//根据hash查询交易
func GetTransaction(transactionHash string) (string, exception.Error) {

	if len(transactionHash) != 64 {
		return "", exception.SdkErr(exception.INVALID_PARAMETER)
	}
	get := "/getTransactionHistory?hash="
	response, Err := common.GetRequest(Account.Url, get, transactionHash)
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
			if data["error_code"].(json.Number) == "4" {
				return "", exception.SdkErr(exception.TRANSACTION_NOT_EXIST)
			}
			errorCodeStr := data["error_code"].(json.Number)
			errorCode, err := strconv.ParseInt(string(errorCodeStr), 10, 64)
			if err != nil {
				Err.Code = exception.STRCONV_PARSEINT_ERROR
				Err.Err = err
				return "", Err
			}
			return "", exception.GetErr(float64(errorCode))
		}
	} else {
		Err.Code = response.StatusCode
		Err.Err = errors.New(response.Status)
		return "", Err
	}
}

//根据高度查询交易
func GetBlock(blockNumber int64) (string, exception.Error) {
	if blockNumber < 0 {
		return "", exception.SdkErr(exception.INVALID_PARAMETER)
	}
	bnstr := strconv.FormatInt(blockNumber, 10)
	get := "/getTransactionHistory?ledger_seq="
	response, Err := common.GetRequest(Account.Url, get, bnstr)
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
			if data["error_code"].(json.Number) == "4" {
				return "", exception.SdkErr(exception.BLOCK_NOT_EXIST)
			}
			errorCodeStr := data["error_code"].(json.Number)
			errorCode, err := strconv.ParseInt(string(errorCodeStr), 10, 64)
			if err != nil {
				Err.Code = exception.STRCONV_PARSEINT_ERROR
				Err.Err = err
				return "", Err
			}
			return "", exception.GetErr(float64(errorCode))
		}
	} else {
		Err.Code = response.StatusCode
		Err.Err = errors.New(response.Status)
		return "", Err
	}
}

//查询区块头
func GetLedger(blockNumber int64) (string, exception.Error) {
	if blockNumber <= 0 {
		return "", exception.SdkErr(exception.INVALID_PARAMETER)
	}
	bnstr := strconv.FormatInt(blockNumber, 10)
	get := "/getLedger?seq="
	response, Err := common.GetRequest(Account.Url, get, bnstr)
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
			if data["error_code"].(json.Number) == "4" {
				return "", exception.SdkErr(exception.BLOCK_NOT_EXIST)
			}
			errorCodeStr := data["error_code"].(json.Number)
			errorCode, err := strconv.ParseInt(string(errorCodeStr), 10, 64)
			if err != nil {
				Err.Code = exception.STRCONV_PARSEINT_ERROR
				Err.Err = err
				return "", Err
			}
			return "", exception.GetErr(float64(errorCode))
		}
	} else {
		Err.Code = response.StatusCode
		Err.Err = errors.New(response.Status)
		return "", Err
	}
}

//生成交易(默认费用)
func createTransactionWithDefaultFee(sourceAddress string, nonce int64, operation []byte) (string, exception.Error) {
	if !keypair.CheckAddress(sourceAddress) {
		return "", exception.SdkErr(exception.INVALID_SOURCEADDRESS)
	}
	if nonce <= 0 {
		return "", exception.SdkErr(exception.INVALID_NONCE)
	}
	if operation == nil {
		return "", exception.SdkErr(exception.INVALID_OPERATION)
	}
	var feeLimit int64
	gasPrice, _, Err := common.GetFees(Account.Url)
	if Err.Err != nil {
		return "", Err
	}
	operations := &protocol.Operation{}
	err := proto.Unmarshal(operation, operations)
	if err != nil {
		Err.Code = exception.PROTO_UNMARSHAL_ERROR
		Err.Err = err
		return "", Err
	}
	if operations.Type == protocol.Operation_ISSUE_ASSET {
		feeLimit = (5000000 + 1000) * gasPrice
	} else if operations.Type == protocol.Operation_CREATE_ACCOUNT {
		feeLimit = (1000000 + 1000) * gasPrice
	} else {
		feeLimit = 1000 * gasPrice
	}
	Operations := []*protocol.Operation{
		{},
	}
	err = proto.Unmarshal(operation, Operations[0])
	if err != nil {
		Err.Code = exception.PROTO_UNMARSHAL_ERROR
		Err.Err = err
		return "", Err
	}
	Transaction := &protocol.Transaction{
		SourceAddress: sourceAddress,
		Nonce:         nonce,
		FeeLimit:      feeLimit,
		GasPrice:      gasPrice,
		Operations:    Operations,
	}
	data, err := proto.Marshal(Transaction)
	if err != nil {
		Err.Code = exception.PROTO_MARSHAL_ERROR
		Err.Err = err
		return "", Err
	}
	dataHash := hex.EncodeToString(data)
	Err.Code = exception.SUCCESS
	Err.Err = nil
	return dataHash, Err
}

//生成交易(传入费用)
func CreateTransactionWithFee(sourceAddress string, nonce int64, gasPrice int64, feeLimit int64, operation []byte) (string, exception.Error) {
	if !keypair.CheckAddress(sourceAddress) {
		return "", exception.SdkErr(exception.INVALID_SOURCEADDRESS)
	}
	newgasPrice, _, Err := common.GetFees(Account.Url)
	if Err.Err != nil {
		return "", Err
	}
	if nonce <= 0 {
		return "", exception.SdkErr(exception.INVALID_NONCE)
	}
	if gasPrice < newgasPrice {
		return "", exception.SdkErr(exception.INVALID_GASPRICE)
	}
	if feeLimit < newgasPrice*1000 {
		return "", exception.SdkErr(exception.INVALID_FEELIMIT)
	}
	if operation == nil {
		return "", exception.SdkErr(exception.INVALID_OPERATION)
	}
	operations := &protocol.Operation{}
	err := proto.Unmarshal(operation, operations)
	if err != nil {
		Err.Code = exception.PROTO_UNMARSHAL_ERROR
		Err.Err = err
		return "", Err
	}
	Operations := []*protocol.Operation{
		{},
	}
	err = proto.Unmarshal(operation, Operations[0])
	Transaction := &protocol.Transaction{
		SourceAddress: sourceAddress,
		Nonce:         nonce,
		FeeLimit:      feeLimit,
		GasPrice:      gasPrice,
		Operations:    Operations,
	}
	data, err := proto.Marshal(Transaction)
	if err != nil {
		Err.Code = exception.PROTO_MARSHAL_ERROR
		Err.Err = err
		return "", Err
	}
	dataHash := hex.EncodeToString(data)
	Err.Code = exception.SUCCESS
	Err.Err = nil
	return dataHash, Err
}

//评估费用
func EvaluationFee(sourceAddress string, nonce int64, operation []byte, signatureNumber int64) (int64, int64, exception.Error) {
	if !keypair.CheckAddress(sourceAddress) {
		return 0, 0, exception.SdkErr(exception.INVALID_SOURCEADDRESS)
	}
	if nonce <= 0 {
		return 0, 0, exception.SdkErr(exception.INVALID_NONCE)
	}
	if operation == nil {
		return 0, 0, exception.SdkErr(exception.INVALID_OPERATION)
	}
	if signatureNumber <= 0 {
		return 0, 0, exception.SdkErr(exception.INVALID_SIGNATURENUMBER)
	}

	operations := &protocol.Operation{}
	err := proto.Unmarshal(operation, operations)
	if err != nil {
		Err.Code = exception.PROTO_UNMARSHAL_ERROR
		Err.Err = err
		return 0, 0, Err
	}
	Operations := []*protocol.Operation{
		{},
	}
	err = proto.Unmarshal(operation, Operations[0])
	request := make(map[string]interface{})
	transactionJson := make(map[string]interface{})
	transactionJson["source_address"] = sourceAddress
	transactionJson["nonce"] = nonce
	transactionJson["operations"] = Operations
	transactionJson["signature_number"] = signatureNumber
	items := make([]map[string]interface{}, 1)
	items[0] = make(map[string]interface{})
	items[0]["transaction_json"] = transactionJson
	request["items"] = items
	requestJson, err := json.Marshal(request)
	if err != nil {
		Err.Code = exception.JSON_MARSHAL_ERROR
		Err.Err = err
		return 0, 0, Err
	}
	post := "/testTransaction"
	response, Err := common.PostRequest(Account.Url, post, requestJson)
	if Err.Err != nil {
		return 0, 0, Err
	}
	defer response.Body.Close()
	if response.StatusCode == 200 {
		data := make(map[string]interface{})
		decoder := json.NewDecoder(response.Body)
		decoder.UseNumber()
		err = decoder.Decode(&data)
		if err != nil {
			Err.Code = exception.DECODER_DECODE_ERROR
			Err.Err = err
			return 0, 0, Err
		}
		if data["error_code"].(json.Number) == "0" {
			result := data["result"].(map[string]interface{})
			txs, ok := result["txs"].([]interface{})
			if !ok {
				return 0, 0, exception.SdkErr(exception.TRANSACTION_INVALID)
			}
			tx, ok := txs[0].(map[string]interface{})
			if !ok {
				return 0, 0, exception.SdkErr(exception.TRANSACTION_INVALID)
			}
			if tx["actual_fee"] == nil {
				return 0, 0, exception.SdkErr(exception.TRANSACTION_INVALID)
			}
			actualFeestr := tx["actual_fee"].(json.Number)
			actualFee, err := strconv.ParseInt(string(actualFeestr), 10, 64)
			if err != nil {
				Err.Code = exception.STRCONV_PARSEINT_ERROR
				Err.Err = err
				return 0, 0, Err
			}
			transactionEnv := tx["transaction_env"].(map[string]interface{})
			transaction := transactionEnv["transaction"].(map[string]interface{})
			if transaction["gas_price"] == nil {
				return 0, 0, exception.SdkErr(exception.TRANSACTION_INVALID)
			}
			gasPriceStr := transaction["gas_price"].(json.Number)
			gasPrice, err := strconv.ParseInt(string(gasPriceStr), 10, 64)
			if err != nil {
				Err.Code = exception.STRCONV_PARSEINT_ERROR
				Err.Err = err
				return 0, 0, Err
			}
			Err.Code = exception.SUCCESS
			Err.Err = nil
			return int64(actualFee), int64(gasPrice), Err
		} else {
			errorCodeStr := data["error_code"].(json.Number)
			errorCode, err := strconv.ParseInt(string(errorCodeStr), 10, 64)
			if err != nil {
				Err.Code = exception.STRCONV_PARSEINT_ERROR
				Err.Err = err
				return 0, 0, Err
			}
			Err.Code = int(float64(errorCode) + 10000)
			Err.Err = errors.New(data["error_desc"].(string))
			return 0, 0, Err
		}
	} else {
		Err.Code = response.StatusCode
		Err.Err = errors.New(response.Status)
		return 0, 0, Err
	}
}

//签名
func SignTransaction(transactionBlob string, privateKey string) (string, string, exception.Error) {
	if transactionBlob == "" {
		return "", "", exception.SdkErr(exception.INVALID_TRANSACTIONBLOB)
	}
	if !keypair.CheckPrivateKey(privateKey) {
		return "", "", exception.SdkErr(exception.INVALID_PRIVATEKEY)
	}
	publicKey, err := keypair.GetEncPublicKey(privateKey)
	if err != nil {
		Err.Code = exception.KEYPAIR_GETENCPUBLICKEY_ERROR
		Err.Err = err
		return "", "", Err
	}
	TransactionBlob, err := hex.DecodeString(transactionBlob)
	if err != nil {
		Err.Code = exception.HEX_DECODESTRING_ERROR
		Err.Err = err
		return "", "", Err
	}
	signData, err := signature.Sign(privateKey, TransactionBlob)
	if err != nil {
		Err.Code = exception.SIGNATURE_SIGN_ERROR
		Err.Err = err
		return "", "", Err
	}
	return signData, publicKey, Err
}

//多签名
func MultiSignTransaction(transactionBlob string, privateKey []string) ([]common.Signatures, exception.Error) {
	if transactionBlob == "" {
		return nil, exception.SdkErr(exception.INVALID_TRANSACTIONBLOB)
	}
	for i := range privateKey {
		if !keypair.CheckPrivateKey(privateKey[i]) {
			return nil, exception.SdkErr(exception.INVALID_PRIVATEKEY)
		}
	}
	signatures := make([]common.Signatures, len(privateKey))
	var err error
	for i := range privateKey {
		signatures[i].Public_key, err = keypair.GetEncPublicKey(privateKey[i])
		if err != nil {
			Err.Code = exception.KEYPAIR_GETENCPUBLICKEY_ERROR
			Err.Err = err
			return nil, Err
		}
	}

	TransactionBlob, err := hex.DecodeString(transactionBlob)
	if err != nil {
		Err.Code = exception.HEX_DECODESTRING_ERROR
		Err.Err = err
		return nil, Err
	}
	for i := range privateKey {
		signatures[i].Sign_data, err = signature.Sign(privateKey[i], TransactionBlob)
		if err != nil {
			Err.Code = exception.SIGNATURE_SIGN_ERROR
			Err.Err = err
			return nil, Err
		}
	}

	return signatures, Err
}

//单签名交易提交
func SubmitTransaction(transactionBlob string, signData string, publicKey string) (string, exception.Error) {
	if signData == "" {
		return "", exception.SdkErr(exception.INVALID_SIGNDATA)
	}
	if transactionBlob == "" {
		return "", exception.SdkErr(exception.INVALID_TRANSACTIONBLOB)
	}
	if !keypair.CheckPublicKey(publicKey) {
		return "", exception.SdkErr(exception.INVALID_PUBLICKEY)
	}
	signatures := make([]common.Signatures, 1)
	signatures[0].Sign_data = signData
	signatures[0].Public_key = publicKey
	requestJson, Err := common.GetRequestJson(transactionBlob, signatures)
	if Err.Err != nil {
		return "", Err
	}
	post := "/submitTransaction"
	response, Err := common.PostRequest(Account.Url, post, requestJson)
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
		results := data["results"].([]interface{})
		result := results[0].(map[string]interface{})
		if result["error_code"].(json.Number) == "0" {
			hash := make(map[string]interface{})
			hash["hash"] = result["hash"]
			Mdata, err := json.Marshal(&hash)
			if err != nil {
				Err.Code = exception.JSON_MARSHAL_ERROR
				Err.Err = err
				return "", Err
			}
			Err.Code = exception.SUCCESS
			Err.Err = nil
			return string(Mdata), Err
		} else {
			errorCodeStr := result["error_code"].(json.Number)
			errorCode, err := strconv.ParseInt(string(errorCodeStr), 10, 64)
			if err != nil {
				Err.Code = exception.STRCONV_PARSEINT_ERROR
				Err.Err = err
				return "", Err
			}
			Err.Code = int(float64(errorCode) + 10000)
			Err.Err = errors.New(result["error_desc"].(string))
			return "", Err
		}
	} else {
		Err.Code = response.StatusCode
		Err.Err = errors.New(response.Status)
		return "", Err
	}
}

//多签名交易提交
func SubmitTransWithMultiSign(transactionBlob string, signatures []common.Signatures) (string, exception.Error) {
	if transactionBlob == "" {
		return "", exception.SdkErr(exception.INVALID_TRANSACTIONBLOB)
	}
	for i := range signatures {
		if !keypair.CheckPublicKey(signatures[i].Public_key) || signatures[i].Sign_data == "" {
			return "", exception.SdkErr(exception.INVALID_SIGNATURES)
		}
	}
	requestJson, Err := common.GetRequestJson(transactionBlob, signatures)
	if Err.Err != nil {
		return "", Err
	}
	post := "/submitTransaction"
	response, Err := common.PostRequest(Account.Url, post, requestJson)
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
		results := data["results"].([]interface{})
		result := results[0].(map[string]interface{})
		if result["error_code"].(json.Number) == "0" {
			hash := make(map[string]interface{})
			hash["hash"] = result["hash"]
			Mdata, err := json.Marshal(&hash)
			if err != nil {
				Err.Code = exception.JSON_MARSHAL_ERROR
				Err.Err = err
				return "", Err
			}
			Err.Code = exception.SUCCESS
			Err.Err = nil
			return string(Mdata), Err
		} else {
			errorCodeStr := result["error_code"].(json.Number)
			errorCode, err := strconv.ParseInt(string(errorCodeStr), 10, 64)
			if err != nil {
				Err.Code = exception.STRCONV_PARSEINT_ERROR
				Err.Err = err
				return "", Err
			}
			Err.Code = int(float64(errorCode) + 10000)
			Err.Err = errors.New(result["error_desc"].(string))
			return "", Err
		}
	} else {
		Err.Code = response.StatusCode
		Err.Err = errors.New(response.Status)
		return "", Err
	}
}
