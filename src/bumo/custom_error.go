// custom_error
package bumo

import (
	"errors"
)

type Error struct {
	Code int
	Err  error
}

const (
	SUCCESS               = 0
	ACCOUNT_NOT_EXIST int = iota + 10200
	TRANSACTION_NOT_EXIST
	BLOCK_NOT_EXIST
	INVALID_PARAMETER
	KEYPAIR_CREATE_ERROR
	PROTO_MARSHAL_ERROR
	PROTO_UNMARSHAL_ERROR
	HTTP_NEWREQUEST_ERROR
	CLIENT_DO_ERROR
	JSON_UNMARSHAL_ERROR
	JSON_MARSHAL_ERROR
	IOUTIL_READALL_ERROR
	TRANSACTION_INVALID
	KEYPAIR_GETENCPUBLICKEY_ERROR
	KEYPAIR_CHECKADDRESS_ERROR
	HEX_DECODESTRING_ERROR
	SIGNATURE_SIGN_ERROR
	DECODER_DECODE_ERROR
	STRCONV_PARSEINT_ERROR
	INVALID_AMOUNT
	INVALID_CODE
	INVALID_ISSUEADDRESS
	INVALID_SOURCEADDRESS
	INVALID_DESTADDRESS
	INVALID_INITBALANCE
	INVALID_PAYLOAD
	INVALID_NONCE
	INVALID_OPERATION
	INVALID_GASPRICE
	INVALID_FEELIMIT
	INVALID_SIGNATURENUMBER
	INVALID_TRANSACTIONBLOB
	INVALID_PRIVATEKEY
	INVALID_PUBLICKEY
	INVALID_SIGNDATA
	INVALID_SIGNATURES
	INVALID_KEY
	INVALID_VALUE
	INVALID_VERSION
	INVALID_SIGNERADDRESS
	INVALID_MASTERWEIGHT
	INVALID_WEIGHT
	INVALID_TXTHRESHOLD
	INVALID_THRESHOLDSTYPE
	INVALID_THRESHOLDS
	DESTADDRESS_EQUAL_SOURCEADDRESS
	NO_TRANSACTIONS_FOUND
)

var Err Error

//自定义错误
func sdkErr(code int) Error {
	errm := map[int]string{
		SUCCESS:                         "Success",
		ACCOUNT_NOT_EXIST:               "Account does not exist",
		TRANSACTION_NOT_EXIST:           "Transaction does not exist",
		NO_TRANSACTIONS_FOUND:           "No transactions found",
		BLOCK_NOT_EXIST:                 "Block does not exist",
		INVALID_PARAMETER:               "The parameter is wrong",
		KEYPAIR_CREATE_ERROR:            "The function 'keypair.Create' failed",
		PROTO_MARSHAL_ERROR:             "The function 'proto.Marshal' failed",
		PROTO_UNMARSHAL_ERROR:           "The function 'proto.Unmarshal' failed",
		HTTP_NEWREQUEST_ERROR:           "The function 'http.NewRequest' failed",
		CLIENT_DO_ERROR:                 "The function 'client.Do' failed",
		JSON_UNMARSHAL_ERROR:            "The function 'json.Unarshal' failed",
		JSON_MARSHAL_ERROR:              "The function 'json.Marshal' failed",
		IOUTIL_READALL_ERROR:            "The function 'ioutil.ReadAll' failed",
		TRANSACTION_INVALID:             "The function 'Transaction is invalid' failed",
		KEYPAIR_GETENCPUBLICKEY_ERROR:   "The function 'keypair.GetEncPublicKey' failed",
		KEYPAIR_CHECKADDRESS_ERROR:      "The function 'keypair.CheckAddress' failed",
		HEX_DECODESTRING_ERROR:          "The function 'hex.DecodeString' failed",
		SIGNATURE_SIGN_ERROR:            "The function 'signature.Sign' failed",
		DECODER_DECODE_ERROR:            "The function 'decoder.Decode' failed",
		STRCONV_PARSEINT_ERROR:          "The function 'strconv.ParseInt' failed",
		INVALID_AMOUNT:                  "The parameter 'amount' is invalid",
		INVALID_CODE:                    "The parameter 'code' is invalid",
		INVALID_ISSUEADDRESS:            "The parameter 'issueAddress' is invalid",
		INVALID_SOURCEADDRESS:           "The parameter 'sourceAddress' is invalid",
		INVALID_DESTADDRESS:             "The parameter 'destAddress' is invalid",
		INVALID_INITBALANCE:             "The parameter 'initBalance' is invalid",
		INVALID_PAYLOAD:                 "The parameter 'payload' is invalid",
		INVALID_NONCE:                   "The parameter 'nonce' is invalid",
		INVALID_OPERATION:               "The parameter 'operation' is invalid",
		INVALID_GASPRICE:                "The parameter 'gasPrice' is invalid",
		INVALID_FEELIMIT:                "The parameter 'feeLimit' is invalid",
		INVALID_SIGNATURENUMBER:         "The parameter 'signatureNumber' is invalid",
		INVALID_TRANSACTIONBLOB:         "The parameter 'transactionBlob' is invalid",
		INVALID_PRIVATEKEY:              "The parameter 'privateKey' is invalid",
		INVALID_PUBLICKEY:               "The parameter 'publicKey' is invalid",
		INVALID_SIGNDATA:                "The parameter 'signData' is invalid",
		INVALID_SIGNATURES:              "The parameter 'signatures' is invalid",
		INVALID_KEY:                     "The parameter 'key' is invalid",
		INVALID_VALUE:                   "The parameter 'value' is invalid",
		INVALID_VERSION:                 "The parameter 'version' is invalid",
		INVALID_SIGNERADDRESS:           "The parameter 'signerAddress' is invalid",
		INVALID_MASTERWEIGHT:            "The parameter 'masterWeight' is invalid",
		INVALID_WEIGHT:                  "The parameter 'weight' is invalid",
		INVALID_TXTHRESHOLD:             "The parameter 'txThreshold' is invalid",
		INVALID_THRESHOLDSTYPE:          "The parameter 'thresholdsType' is invalid",
		INVALID_THRESHOLDS:              "The parameter 'thresholds' is invalid",
		DESTADDRESS_EQUAL_SOURCEADDRESS: "'sourceAddress' is equal to 'destAddress'",
	}
	v, _ := errm[code]
	Err.Code = int(code)
	Err.Err = errors.New(v)
	return Err
}

//查询错误判断
func getErr(code float64) Error {
	errm := map[float64]string{
		0:  "Success",
		1:  "Invalid private key",
		2:  "Invalid public key",
		3:  "Invalid adress",
		4:  "Objects do not exist, such as null account, transactions and blocks etc.",
		5:  "Transaction fail",
		6:  "Nonce too small",
		7:  "Not enough weight",
		8:  "Invalid number of arguments to the function.",
		9:  "Invalid type of argument to the function.",
		10: "Argument cannot be empty",
		11: "Internal Server Error",
		12: "Nonce incorrect",
		13: "BU is not enough",
		14: "Source address equal to dest address",
		15: "Dest account already exists",
		16: "Fee not enough",
		17: "Query result not exist",
		18: "Discard transaction because of lower fee  in queue.",
		19: "Include invalid arguments",
		20: "Fail",
	}

	if v, ok := errm[code]; ok {
		Err.Code = int(code) + 10000
		Err.Err = errors.New(v)
		return Err
	} else {
		Err.Code = 20
		Err.Err = errors.New("Fail")
		return Err
	}
}

//交易错误判断
func submitErr(code float64) Error {
	errm := map[float64]string{
		93:  "Not enough weight.",
		99:  "Nonce incorrect.",
		100: "BU is not enough.",
		101: "Source address equal to dest address.",
		102: "Dest account already exists.",
		103: "Account not exist.",
		111: "Fee not enough.",
		160: "Discard transaction, because of lower fee in queue.",
	}

	if v, ok := errm[code]; ok {
		Err.Code = int(code) + 10000
		Err.Err = errors.New(v)
		return Err
	} else {
		Err.Code = int(code)
		Err.Err = errors.New("transactionfail.")
		return Err
	}
}
