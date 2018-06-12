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
)

var Err Error

//自定义错误
func sdkErr(code int) Error {
	errm := map[int]string{
		SUCCESS:                       "Success",
		ACCOUNT_NOT_EXIST:             "Account does not exist",
		TRANSACTION_NOT_EXIST:         "Transaction does not exist",
		BLOCK_NOT_EXIST:               "Block does not exist",
		INVALID_PARAMETER:             "The parameter is wrong",
		KEYPAIR_CREATE_ERROR:          "Keypair_create call failed",
		PROTO_MARSHAL_ERROR:           "Proto_marshal call failed",
		PROTO_UNMARSHAL_ERROR:         "Proto_unmarshal call failed",
		HTTP_NEWREQUEST_ERROR:         "Http_newrequest call failed",
		CLIENT_DO_ERROR:               "Client_do call failed",
		JSON_UNMARSHAL_ERROR:          "Json_unmarshal call failed",
		JSON_MARSHAL_ERROR:            "Json_marshal call failed",
		IOUTIL_READALL_ERROR:          "Ioutil_readall call failed",
		TRANSACTION_INVALID:           "Transaction is invalid",
		KEYPAIR_GETENCPUBLICKEY_ERROR: "Keypair_getencpublickey call failed",
		KEYPAIR_CHECKADDRESS_ERROR:    "Keypair_checkaddress call failed",
		HEX_DECODESTRING_ERROR:        "Hex_decodestring call failed",
		SIGNATURE_SIGN_ERROR:          "Signature_sign call failed",
		DECODER_DECODE_ERROR:          "Decoder_decode call failed",
		STRCONV_PARSEINT_ERROR:        "strconv_ParseInt call failed",
		INVALID_AMOUNT:                "Invalid_Amount",
		INVALID_CODE:                  "Invalid_Code",
		INVALID_ISSUEADDRESS:          "Invalid_Issueaddress",
		INVALID_SOURCEADDRESS:         "Invalid_Sourceaddress",
		INVALID_DESTADDRESS:           "Invalid_Destaddress",
		INVALID_INITBALANCE:           "Invalid_Initbalance",
		INVALID_PAYLOAD:               "Invalid_Payload",
		INVALID_NONCE:                 "Invalid_Nonce",
		INVALID_OPERATION:             "Invalid_Operation",
		INVALID_GASPRICE:              "Invalid_GasPrice",
		INVALID_FEELIMIT:              "Invalid_FeeLimit",
		INVALID_SIGNATURENUMBER:       "Invalid_SignatureNumber",
		INVALID_TRANSACTIONBLOB:       "Invalid_TransactionBlob",
		INVALID_PRIVATEKEY:            "Invalid_PrivateKey",
		INVALID_PUBLICKEY:             "Invalid_PublicKey",
		INVALID_SIGNDATA:              "Invalid_SignData",
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
		93:  "Not enough weight",
		99:  "Nonce incorrect",
		100: "BU is not enough",
		101: "Sourc eaddress equal to dest address",
		102: "Dest account already exists",
		103: "Account not exist",
		111: "Fee not enough",
		160: "Discard transaction, because of lower fee in queue.",
	}

	if v, ok := errm[code]; ok {
		Err.Code = int(code) + 10000
		Err.Err = errors.New(v)
		return Err
	} else {
		Err.Code = int(code)
		Err.Err = errors.New("transactionfail")
		return Err
	}
}
