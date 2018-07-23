// common
package common

import (
	"container/list"
	"encoding/hex"
	"encoding/json"
	"strconv"

	"github.com/bumoproject/bumo-sdk-go/src/crypto/keypair"
	"github.com/bumoproject/bumo-sdk-go/src/crypto/protocol"
	"github.com/bumoproject/bumo-sdk-go/src/exception"
	"github.com/bumoproject/bumo-sdk-go/src/model"
)

//GetOperations
func GetOperations(operationsList list.List, url string) ([]*protocol.Operation, exception.SDKResponse) {
	operations := make([]*protocol.Operation, operationsList.Len())
	var i int = 0
	for e := operationsList.Front(); e != nil; e = e.Next() {
		operations[i] = new(protocol.Operation)
		operationsData, ok := e.Value.(model.OperationBase)
		if !ok {
			return operations, exception.GetSDKRes(exception.INVALID_OPERATION_ERROR)
		}
		switch operationsData.Get() {
		case 1:
			operationsReqData, ok := operationsData.(model.AccountActivateOperation)
			if !ok {
				return operations, exception.GetSDKRes(exception.INVALID_OPERATION_ERROR)
			}
			operationsResData := Activate(operationsReqData, url)
			if operationsResData.ErrorCode != 0 {
				return operations, exception.GetSDKRes(exception.INVALID_OPERATION_ERROR)
			}
			operations[i] = &operationsResData.Result.Operation
		case 2:
			operationsReqData, ok := operationsData.(model.AccountSetMetadataOperation)
			if !ok {
				return operations, exception.GetSDKRes(exception.INVALID_OPERATION_ERROR)
			}
			operationsResData := SetMetadata(operationsReqData)
			if operationsResData.ErrorCode != 0 {
				return operations, exception.GetSDKRes(exception.INVALID_OPERATION_ERROR)
			}
			operations[i] = &operationsResData.Result.Operation
		case 3:
			operationsReqData, ok := operationsData.(model.AccountSetPrivilegeOperation)
			if !ok {
				return operations, exception.GetSDKRes(exception.INVALID_OPERATION_ERROR)
			}
			operationsResData := SetPrivilege(operationsReqData)
			if operationsResData.ErrorCode != 0 {
				return operations, exception.GetSDKRes(exception.INVALID_OPERATION_ERROR)
			}
			operations[i] = &operationsResData.Result.Operation
		case 4:
			operationsReqData, ok := operationsData.(model.AssetIssueOperation)
			if !ok {
				return operations, exception.GetSDKRes(exception.INVALID_OPERATION_ERROR)
			}
			operationsResData := AssetIssue(operationsReqData)
			if operationsResData.ErrorCode != 0 {
				return operations, exception.GetSDKRes(exception.INVALID_OPERATION_ERROR)
			}
			operations[i] = &operationsResData.Result.Operation
		case 5:
			operationsReqData, ok := operationsData.(model.AssetSendOperation)
			if !ok {
				return operations, exception.GetSDKRes(exception.INVALID_OPERATION_ERROR)
			}
			operationsResData := AssetSend(operationsReqData)
			if operationsResData.ErrorCode != 0 {
				return operations, exception.GetSDKRes(exception.INVALID_OPERATION_ERROR)
			}
			operations[i] = &operationsResData.Result.Operation
		case 6:
			operationsReqData, ok := operationsData.(model.BUSendOperation)
			if !ok {
				return operations, exception.GetSDKRes(exception.INVALID_OPERATION_ERROR)
			}

			operationsResData := BUSend(operationsReqData)
			if operationsResData.ErrorCode != 0 {
				return operations, exception.GetSDKRes(exception.INVALID_OPERATION_ERROR)
			}
			operations[i] = &operationsResData.Result.Operation
		case 7:
			operationsReqData, ok := operationsData.(model.TokenIssueOperation)
			if !ok {
				return operations, exception.GetSDKRes(exception.INVALID_OPERATION_ERROR)
			}
			operationsResData := TokenIssue(operationsReqData)
			if operationsResData.ErrorCode != 0 {
				return operations, exception.GetSDKRes(exception.INVALID_OPERATION_ERROR)
			}
			operations[i] = &operationsResData.Result.Operation
		case 8:
			operationsReqData, ok := operationsData.(model.TokenTransferOperation)
			if !ok {
				return operations, exception.GetSDKRes(exception.INVALID_OPERATION_ERROR)
			}
			operationsResData := Transfer(operationsReqData)
			if operationsResData.ErrorCode != 0 {
				return operations, exception.GetSDKRes(exception.INVALID_OPERATION_ERROR)
			}
			operations[i] = &operationsResData.Result.Operation
		case 9:
			operationsReqData, ok := operationsData.(model.TokenTransferFromOperation)
			if !ok {
				return operations, exception.GetSDKRes(exception.INVALID_OPERATION_ERROR)
			}
			operationsResData := TransferFrom(operationsReqData)
			if operationsResData.ErrorCode != 0 {
				return operations, exception.GetSDKRes(exception.INVALID_OPERATION_ERROR)
			}
			operations[i] = &operationsResData.Result.Operation
		case 10:
			operationsReqData, ok := operationsData.(model.TokenApproveOperation)
			if !ok {
				return operations, exception.GetSDKRes(exception.INVALID_OPERATION_ERROR)
			}
			operationsResData := Approve(operationsReqData)
			if operationsResData.ErrorCode != 0 {
				return operations, exception.GetSDKRes(exception.INVALID_OPERATION_ERROR)
			}
			operations[i] = &operationsResData.Result.Operation
		case 11:
			operationsReqData, ok := operationsData.(model.TokenAssignOperation)
			if !ok {
				return operations, exception.GetSDKRes(exception.INVALID_OPERATION_ERROR)
			}
			operationsResData := Assign(operationsReqData)
			if operationsResData.ErrorCode != 0 {
				return operations, exception.GetSDKRes(exception.INVALID_OPERATION_ERROR)
			}
			operations[i] = &operationsResData.Result.Operation
		case 12:
			operationsReqData, ok := operationsData.(model.TokenChangeOwnerOperation)
			if !ok {
				return operations, exception.GetSDKRes(exception.INVALID_OPERATION_ERROR)
			}
			operationsResData := ChangeOwner(operationsReqData)
			if operationsResData.ErrorCode != 0 {
				return operations, exception.GetSDKRes(exception.INVALID_OPERATION_ERROR)
			}
			operations[i] = &operationsResData.Result.Operation
		case 13:
			operationsReqData, ok := operationsData.(model.ContractCreateOperation)
			if !ok {
				return operations, exception.GetSDKRes(exception.INVALID_OPERATION_ERROR)
			}
			operationsResData := ContractCreate(operationsReqData)
			if operationsResData.ErrorCode != 0 {
				return operations, exception.GetSDKRes(exception.INVALID_OPERATION_ERROR)
			}
			operations[i] = &operationsResData.Result.Operation
		case 14:
			operationsReqData, ok := operationsData.(model.ContractInvokeByAssetOperation)
			if !ok {
				return operations, exception.GetSDKRes(exception.INVALID_OPERATION_ERROR)
			}
			operationsResData := InvokeByAsset(operationsReqData)
			if operationsResData.ErrorCode != 0 {
				return operations, exception.GetSDKRes(exception.INVALID_OPERATION_ERROR)
			}
			operations[i] = &operationsResData.Result.Operation
		case 15:
			operationsReqData, ok := operationsData.(model.ContractInvokeByBUOperation)
			if !ok {
				return operations, exception.GetSDKRes(exception.INVALID_OPERATION_ERROR)
			}
			operationsResData := InvokeByBU(operationsReqData)
			if operationsResData.ErrorCode != 0 {
				return operations, exception.GetSDKRes(exception.INVALID_OPERATION_ERROR)
			}
			operations[i] = &operationsResData.Result.Operation
		case 16:
			operationsReqData, ok := operationsData.(model.LogCreateOperation)
			if !ok {
				return operations, exception.GetSDKRes(exception.INVALID_OPERATION_ERROR)
			}
			operationsResData := LogCreate(operationsReqData)
			if operationsResData.ErrorCode != 0 {
				return operations, exception.GetSDKRes(exception.INVALID_OPERATION_ERROR)
			}
			operations[i] = &operationsResData.Result.Operation
		default:
			return operations, exception.GetSDKRes(exception.INVALID_OPERATION_ERROR)
		}
		i++
	}
	return operations, exception.GetSDKRes(exception.SUCCESS)
}

//激活账户 activate the account 1
func Activate(reqData model.AccountActivateOperation, url string) model.AccountActivateResponse {
	var resData model.AccountActivateResponse
	_, baseReserve, SDKRes := GetLatestFees(url)
	if SDKRes.ErrorCode != 0 {
		resData.ErrorCode = SDKRes.ErrorCode
		resData.ErrorDesc = SDKRes.ErrorDesc
		return resData
	}
	if reqData.GetSourceAddress() != "" {
		if !keypair.CheckAddress(reqData.GetSourceAddress()) {
			resData.ErrorCode = exception.INVALID_SOURCEADDRESS_ERROR
			resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
			return resData
		}
	}
	if reqData.GetSourceAddress() == reqData.GetDestAddress() {
		resData.ErrorCode = exception.SOURCEADDRESS_EQUAL_DESTADDRESS_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	if reqData.GetInitBalance() < baseReserve || reqData.GetInitBalance() <= 0 {
		resData.ErrorCode = exception.INVALID_INITBALANCE_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	if !keypair.CheckAddress(reqData.GetDestAddress()) {
		resData.ErrorCode = exception.INVALID_DESTADDRESS_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	_, err := hex.DecodeString(reqData.GetMetadata())
	if err != nil {
		resData.ErrorCode = exception.METADATA_NOT_HEX_STRING_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	Operations := []*protocol.Operation{
		{
			SourceAddress: reqData.GetSourceAddress(),
			Metadata:      []byte(reqData.GetMetadata()),
			Type:          protocol.Operation_CREATE_ACCOUNT,
			CreateAccount: &protocol.OperationCreateAccount{
				DestAddress: reqData.GetDestAddress(),
				Priv: &protocol.AccountPrivilege{
					MasterWeight: 1,
					Thresholds: &protocol.AccountThreshold{
						TxThreshold: 1,
					},
				},
				InitBalance: reqData.GetInitBalance(),
			},
		},
	}
	resData.Result.Operation = *(Operations[0])
	return resData

}

//设置metadata SetMetadata 2
func SetMetadata(reqData model.AccountSetMetadataOperation) model.AccountSetMetadataResponse {
	var resData model.AccountSetMetadataResponse
	if reqData.GetSourceAddress() != "" {
		if !keypair.CheckAddress(reqData.GetSourceAddress()) {
			SDKRes := exception.GetSDKRes(exception.INVALID_SOURCEADDRESS_ERROR)
			resData.ErrorCode = SDKRes.ErrorCode
			resData.ErrorDesc = SDKRes.ErrorDesc
			return resData
		}
	}
	if len(reqData.GetKey()) <= 0 || len(reqData.GetKey()) > 1024 {
		SDKRes := exception.GetSDKRes(exception.INVALID_DATAKEY_ERROR)
		resData.ErrorCode = SDKRes.ErrorCode
		resData.ErrorDesc = SDKRes.ErrorDesc
		return resData
	}
	if len(reqData.GetValue()) < 0 || len(reqData.GetValue()) > (1024*256) {
		SDKRes := exception.GetSDKRes(exception.INVALID_DATAVALUE_ERROR)
		resData.ErrorCode = SDKRes.ErrorCode
		resData.ErrorDesc = SDKRes.ErrorDesc
		return resData
	}
	if reqData.GetVersion() < 0 {
		SDKRes := exception.GetSDKRes(exception.INVALID_DATAVERSION_ERROR)
		resData.ErrorCode = SDKRes.ErrorCode
		resData.ErrorDesc = SDKRes.ErrorDesc
		return resData
	}
	_, err := hex.DecodeString(reqData.GetMetadata())
	if err != nil {
		resData.ErrorCode = exception.METADATA_NOT_HEX_STRING_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	Operations := []*protocol.Operation{
		{
			SourceAddress: reqData.GetSourceAddress(),
			Metadata:      []byte(reqData.GetMetadata()),
			Type:          protocol.Operation_SET_METADATA,
			SetMetadata: &protocol.OperationSetMetadata{
				Key:        reqData.GetKey(),
				Value:      reqData.GetValue(),
				Version:    reqData.GetVersion(),
				DeleteFlag: reqData.GetDeleteFlag(),
			},
		},
	}
	resData.Result.Operation = *(Operations[0])
	return resData
}

//设置权限 SetPrivilege 3
func SetPrivilege(reqData model.AccountSetPrivilegeOperation) model.AccountSetPrivilegeResponse {
	var resData model.AccountSetPrivilegeResponse
	if reqData.GetSourceAddress() != "" {
		if !keypair.CheckAddress(reqData.GetSourceAddress()) {
			SDKRes := exception.GetSDKRes(exception.INVALID_SOURCEADDRESS_ERROR)
			resData.ErrorCode = SDKRes.ErrorCode
			resData.ErrorDesc = SDKRes.ErrorDesc
			return resData
		}
	}
	if reqData.GetMasterWeight() != "" {
		masterWeightInt, err := strconv.ParseInt(reqData.GetMasterWeight(), 10, 32)
		if err != nil || masterWeightInt < 0 {
			SDKRes := exception.GetSDKRes(exception.INVALID_MASTERWEIGHT_ERROR)
			resData.ErrorCode = SDKRes.ErrorCode
			resData.ErrorDesc = SDKRes.ErrorDesc
			return resData
		}
	}

	for i := range reqData.GetSigners() {
		if !keypair.CheckAddress(reqData.GetSigners()[i].Address) {
			SDKRes := exception.GetSDKRes(exception.INVALID_SIGNER_ADDRESS_ERROR)
			resData.ErrorCode = SDKRes.ErrorCode
			resData.ErrorDesc = SDKRes.ErrorDesc
			return resData
		}
		if reqData.GetSigners()[i].Weight > 4294967295 || reqData.GetSigners()[i].Weight < 0 {
			SDKRes := exception.GetSDKRes(exception.INVALID_SIGNER_WEIGHT_ERROR)
			resData.ErrorCode = SDKRes.ErrorCode
			resData.ErrorDesc = SDKRes.ErrorDesc
			return resData
		}
	}
	if reqData.GetTxThreshold() != "" {
		txThresholdInt, err := strconv.ParseInt(reqData.GetTxThreshold(), 10, 64)
		if err != nil || txThresholdInt < 0 {
			SDKRes := exception.GetSDKRes(exception.INVALID_TX_THRESHOLD_ERROR)
			resData.ErrorCode = SDKRes.ErrorCode
			resData.ErrorDesc = SDKRes.ErrorDesc
			return resData
		}
	}
	for i := range reqData.GetTypeThresholds() {
		if reqData.GetTypeThresholds()[i].Type > 100 || reqData.GetTypeThresholds()[i].Type <= 0 {
			SDKRes := exception.GetSDKRes(exception.INVALID_OPERATION_TYPE_ERROR)
			resData.ErrorCode = SDKRes.ErrorCode
			resData.ErrorDesc = SDKRes.ErrorDesc
			return resData
		}
		if reqData.GetTypeThresholds()[i].Type < 0 {
			SDKRes := exception.GetSDKRes(exception.INVALID_TYPE_THRESHOLD_ERROR)
			resData.ErrorCode = SDKRes.ErrorCode
			resData.ErrorDesc = SDKRes.ErrorDesc
			return resData
		}
	}
	Signers := make([]*protocol.Signer, len(reqData.GetSigners()))
	for i := range reqData.GetSigners() {
		Signers[i] = new(protocol.Signer)
		Signers[i].Address = reqData.GetSigners()[i].Address
		Signers[i].Weight = reqData.GetSigners()[i].Weight
	}
	TypeThresholds := make([]*protocol.OperationTypeThreshold, len(reqData.GetTypeThresholds()))
	for i := range reqData.GetTypeThresholds() {
		TypeThresholds[i] = new(protocol.OperationTypeThreshold)
		TypeThresholds[i].Threshold = reqData.GetTypeThresholds()[i].Threshold
		TypeThresholds[i].Type = (protocol.Operation_Type)(reqData.GetTypeThresholds()[i].Type)
	}
	_, err := hex.DecodeString(reqData.GetMetadata())
	if err != nil {
		resData.ErrorCode = exception.METADATA_NOT_HEX_STRING_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	Operations := []*protocol.Operation{
		{
			SourceAddress: reqData.GetSourceAddress(),
			Metadata:      []byte(reqData.GetMetadata()),
			Type:          protocol.Operation_SET_PRIVILEGE,
			SetPrivilege: &protocol.OperationSetPrivilege{
				MasterWeight:   reqData.GetMasterWeight(),
				Signers:        Signers,
				TxThreshold:    reqData.GetTxThreshold(),
				TypeThresholds: TypeThresholds,
			},
		},
	}
	resData.Result.Operation = *(Operations[0])
	return resData
}

//发行资产 AssetIssue 4
func AssetIssue(reqData model.AssetIssueOperation) model.AssetIssueResponse {
	var resData model.AssetIssueResponse
	if reqData.GetSourceAddress() != "" {
		if !keypair.CheckAddress(reqData.GetSourceAddress()) {
			resData.ErrorCode = exception.INVALID_SOURCEADDRESS_ERROR
			resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
			return resData
		}
	}
	if len(reqData.GetCode()) > 64 || len(reqData.GetCode()) == 0 {
		resData.ErrorCode = exception.INVALID_ASSET_CODE_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	if reqData.GetAmount() <= 0 {
		resData.ErrorCode = exception.INVALID_ASSET_AMOUNT_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	_, err := hex.DecodeString(reqData.GetMetadata())
	if err != nil {
		resData.ErrorCode = exception.METADATA_NOT_HEX_STRING_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	Operations := []*protocol.Operation{
		{
			SourceAddress: reqData.GetSourceAddress(),
			Metadata:      []byte(reqData.GetMetadata()),
			Type:          protocol.Operation_ISSUE_ASSET,
			IssueAsset: &protocol.OperationIssueAsset{
				Code:   reqData.GetCode(),
				Amount: reqData.GetAmount(),
			},
		},
	}
	resData.Result.Operation = *(Operations[0])
	return resData
}

//转移资产 AssetSend 5
func AssetSend(reqData model.AssetSendOperation) model.AssetSendResponse {
	var resData model.AssetSendResponse
	if !keypair.CheckAddress(reqData.GetIssuer()) {
		resData.ErrorCode = exception.INVALID_ISSUER_ADDRESS_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	if reqData.GetAmount() <= 0 {
		resData.ErrorCode = exception.INVALID_ASSET_AMOUNT_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	if len(reqData.GetCode()) > 64 || len(reqData.GetCode()) == 0 {
		resData.ErrorCode = exception.INVALID_ASSET_CODE_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	var data model.ContractInvokeByAssetOperation
	data.SetSourceAddress(reqData.GetSourceAddress())
	data.SetContractAddress(reqData.GetDestAddress())
	data.SetAmount(reqData.GetAmount())
	data.SetCode(reqData.GetCode())
	data.SetIssuer(reqData.GetIssuer())
	data.SetMetadata(reqData.GetMetadata())
	contractData := InvokeByAsset(data)
	resData.ErrorCode = exception.SUCCESS
	resData.Result.Operation = contractData.Result.Operation
	return resData
}

//交易BU BUSend 6
func BUSend(reqData model.BUSendOperation) model.BUSendResponse {
	var resData model.BUSendResponse
	if reqData.GetAmount() <= 0 {
		SDKRes := exception.GetSDKRes(exception.INVALID_ADDRESS_ERROR)
		resData.ErrorCode = SDKRes.ErrorCode
		resData.ErrorDesc = SDKRes.ErrorDesc
		return resData

	}
	var data model.ContractInvokeByBUOperation
	data.SetSourceAddress(reqData.GetSourceAddress())
	data.SetContractAddress(reqData.GetDestAddress())
	data.SetAmount(reqData.GetAmount())
	data.SetMetadata(reqData.GetMetadata())
	contractData := InvokeByBU(data)
	resData.ErrorCode = exception.SUCCESS
	resData.Result.Operation = contractData.Result.Operation
	return resData
}

//发行合约token 7
func TokenIssue(reqData model.TokenIssueOperation) model.TokenIssueResponse {
	var resData model.TokenIssueResponse
	if reqData.GetSourceAddress() != "" {
		if !keypair.CheckAddress(reqData.GetSourceAddress()) {
			resData.ErrorCode = exception.INVALID_SOURCEADDRESS_ERROR
			resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
			return resData
		}
	}
	if reqData.GetInitBalance() <= 0 {
		resData.ErrorCode = exception.INVALID_INITBALANCE_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	if reqData.GetDecimals() < 0 && reqData.GetDecimals() > 8 {
		resData.ErrorCode = exception.INVALID_TOKEN_DECIMALS_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	if len(reqData.GetName()) <= 0 && len(reqData.GetName()) > 1024 {
		resData.ErrorCode = exception.INVALID_TOKEN_NAME_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	if len(reqData.GetSymbol()) <= 0 && len(reqData.GetSymbol()) > 1024 {
		resData.ErrorCode = exception.INVALID_TOKEN_SIMBOL_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	if reqData.GetSupply() < 0 {
		resData.ErrorCode = exception.INVALID_TOKEN_TOTALSUPPLY_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	var Input model.Input
	Input.Params.Decimals = reqData.GetDecimals()
	Input.Params.Name = reqData.GetName()
	Input.Params.Symbol = reqData.GetSymbol()
	Input.Params.Supply = strconv.FormatInt(reqData.GetSupply(), 10)
	InitInput, err := json.Marshal(Input)
	if err != nil {
		resData.ErrorCode = exception.SYSTEM_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	_, err = hex.DecodeString(reqData.GetMetadata())
	if err != nil {
		resData.ErrorCode = exception.METADATA_NOT_HEX_STRING_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	Operations := []*protocol.Operation{
		{
			SourceAddress: reqData.GetSourceAddress(),
			Metadata:      []byte(reqData.GetMetadata()),
			Type:          protocol.Operation_CREATE_ACCOUNT,
			CreateAccount: &protocol.OperationCreateAccount{
				Contract: &protocol.Contract{
					Payload: model.Payload,
				},
				InitBalance: reqData.GetInitBalance(),
				InitInput:   string(InitInput),
				Priv: &protocol.AccountPrivilege{
					MasterWeight: 0,
					Thresholds: &protocol.AccountThreshold{
						TxThreshold: 1,
					},
				},
			},
		},
	}
	resData.Result.Operation = *(Operations[0])
	return resData
}

//转移合约token 8
func Transfer(reqData model.TokenTransferOperation) model.TokenTransferResponse {
	var resData model.TokenTransferResponse
	if !keypair.CheckAddress(reqData.GetDestAddress()) {
		resData.ErrorCode = exception.INVALID_DESTADDRESS_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	if reqData.GetDestAddress() == reqData.GetSourceAddress() {
		resData.ErrorCode = exception.SOURCEADDRESS_EQUAL_DESTADDRESS_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	if reqData.GetAmount() < 0 {
		resData.ErrorCode = exception.INVALID_TOKEN_AMOUNT_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	var Input model.Input
	Input.Method = "transfer"
	Input.Params.To = reqData.GetDestAddress()
	Input.Params.Value = strconv.FormatInt(reqData.GetAmount(), 10)
	InputStr, err := json.Marshal(Input)
	if err != nil {
		resData.ErrorCode = exception.SYSTEM_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	var data model.ContractInvokeByBUOperation
	data.SetSourceAddress(reqData.GetSourceAddress())
	data.SetContractAddress(reqData.GetContractAddress())
	data.SetAmount(0)
	data.SetInput(string(InputStr))
	data.SetMetadata(reqData.GetMetadata())
	contractData := InvokeByBU(data)
	resData.ErrorCode = exception.SUCCESS
	resData.Result.Operation = contractData.Result.Operation
	return resData
}

//转移合约token 9
func TransferFrom(reqData model.TokenTransferFromOperation) model.TokenTransferFromResponse {
	var resData model.TokenTransferFromResponse
	if !keypair.CheckAddress(reqData.GetDestAddress()) {
		resData.ErrorCode = exception.INVALID_DESTADDRESS_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	if !keypair.CheckAddress(reqData.GetFromAddress()) {
		resData.ErrorCode = exception.INVALID_FROMADDRESS_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	if reqData.GetDestAddress() == reqData.GetSourceAddress() {
		resData.ErrorCode = exception.SOURCEADDRESS_EQUAL_DESTADDRESS_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	if reqData.GetAmount() < 0 {
		resData.ErrorCode = exception.INVALID_TOKEN_AMOUNT_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	var Input model.Input
	Input.Method = "transferFrom"
	Input.Params.To = reqData.GetDestAddress()
	Input.Params.Value = strconv.FormatInt(reqData.GetAmount(), 10)
	Input.Params.From = reqData.GetFromAddress()
	InputStr, err := json.Marshal(Input)
	if err != nil {
		resData.ErrorCode = exception.SYSTEM_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	var data model.ContractInvokeByBUOperation
	data.SetSourceAddress(reqData.GetSourceAddress())
	data.SetContractAddress(reqData.GetContractAddress())
	data.SetAmount(0)
	data.SetInput(string(InputStr))
	data.SetMetadata(reqData.GetMetadata())
	contractData := InvokeByBU(data)
	resData.ErrorCode = exception.SUCCESS
	resData.Result.Operation = contractData.Result.Operation
	return resData
}

//授权从交易发送者账户转出合约token 10
func Approve(reqData model.TokenApproveOperation) model.TokenApproveResponse {
	var resData model.TokenApproveResponse
	if !keypair.CheckAddress(reqData.GetSpender()) {
		resData.ErrorCode = exception.INVALID_SPENDER_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	if reqData.GetAmount() < 0 {
		resData.ErrorCode = exception.INVALID_TOKEN_AMOUNT_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	var Input model.Input
	Input.Method = "approve"
	Input.Params.Spender = reqData.GetSpender()
	Input.Params.Value = strconv.FormatInt(reqData.GetAmount(), 10)
	InputStr, err := json.Marshal(Input)
	if err != nil {
		resData.ErrorCode = exception.SYSTEM_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	var data model.ContractInvokeByBUOperation
	data.SetSourceAddress(reqData.GetSourceAddress())
	data.SetContractAddress(reqData.GetContractAddress())
	data.SetAmount(0)
	data.SetInput(string(InputStr))
	data.SetMetadata(reqData.GetMetadata())
	contractData := InvokeByBU(data)
	resData.ErrorCode = exception.SUCCESS
	resData.Result.Operation = contractData.Result.Operation
	return resData
}

//分配合约token 11
func Assign(reqData model.TokenAssignOperation) model.TokenAssignResponse {
	var resData model.TokenAssignResponse
	if !keypair.CheckAddress(reqData.GetDestAddress()) {
		resData.ErrorCode = exception.INVALID_SPENDER_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	if reqData.GetAmount() < 0 {
		resData.ErrorCode = exception.INVALID_TOKEN_AMOUNT_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	var Input model.Input
	Input.Method = "assign"
	Input.Params.To = reqData.GetDestAddress()
	Input.Params.Value = strconv.FormatInt(reqData.GetAmount(), 10)
	InputStr, err := json.Marshal(Input)
	if err != nil {
		resData.ErrorCode = exception.SYSTEM_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	var data model.ContractInvokeByBUOperation
	data.SetSourceAddress(reqData.GetSourceAddress())
	data.SetContractAddress(reqData.GetContractAddress())
	data.SetAmount(0)
	data.SetInput(string(InputStr))
	data.SetMetadata(reqData.GetMetadata())
	contractData := InvokeByBU(data)
	resData.ErrorCode = exception.SUCCESS
	resData.Result.Operation = contractData.Result.Operation
	return resData
}

//转移合约token拥有权 12
func ChangeOwner(reqData model.TokenChangeOwnerOperation) model.TokenChangeOwnerResponse {
	var resData model.TokenChangeOwnerResponse
	if !keypair.CheckAddress(reqData.GetTokenOwner()) {
		resData.ErrorCode = exception.INVALID_TOKENOWNER_ERRPR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	var Input model.Input
	Input.Method = "changeOwner"
	Input.Params.Address = reqData.GetTokenOwner()
	InputStr, err := json.Marshal(Input)
	if err != nil {
		resData.ErrorCode = exception.SYSTEM_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	var data model.ContractInvokeByBUOperation
	data.SetSourceAddress(reqData.GetSourceAddress())
	data.SetContractAddress(reqData.GetContractAddress())
	data.SetAmount(0)
	data.SetInput(string(InputStr))
	data.SetMetadata(reqData.GetMetadata())
	contractData := InvokeByBU(data)
	resData.ErrorCode = exception.SUCCESS
	resData.Result.Operation = contractData.Result.Operation
	return resData
}

//创建合约账户 ContractCreate 13
func ContractCreate(reqData model.ContractCreateOperation) model.ContractCreateResponse {
	var resData model.ContractCreateResponse
	if reqData.GetSourceAddress() != "" {
		if !keypair.CheckAddress(reqData.GetSourceAddress()) {
			SDKRes := exception.GetSDKRes(exception.INVALID_SOURCEADDRESS_ERROR)
			resData.ErrorCode = SDKRes.ErrorCode
			resData.ErrorDesc = SDKRes.ErrorDesc
			return resData
		}
	}
	if reqData.GetInitBalance() < 0 {
		SDKRes := exception.GetSDKRes(exception.INVALID_INITBALANCE_ERROR)
		resData.ErrorCode = SDKRes.ErrorCode
		resData.ErrorDesc = SDKRes.ErrorDesc
		return resData
	}
	if reqData.GetPayload() == "" {
		SDKRes := exception.GetSDKRes(exception.INVALID_PAYLOAD_ERROR)
		resData.ErrorCode = SDKRes.ErrorCode
		resData.ErrorDesc = SDKRes.ErrorDesc
		return resData
	}
	Operations := []*protocol.Operation{
		{
			SourceAddress: reqData.GetSourceAddress(),
			Metadata:      []byte(reqData.GetMetadata()),
			Type:          protocol.Operation_CREATE_ACCOUNT,
			CreateAccount: &protocol.OperationCreateAccount{
				Contract: &protocol.Contract{
					Payload: reqData.GetPayload(),
				},
				InitBalance: reqData.GetInitBalance(),
				InitInput:   reqData.GetInitInput(),
				Priv: &protocol.AccountPrivilege{
					MasterWeight: 0,
					Thresholds: &protocol.AccountThreshold{
						TxThreshold: 1,
					},
				},
			},
		},
	}
	resData.Result.Operation = *(Operations[0])
	return resData
}

//转移资产并触发合约 InvokeByAsset 14
func InvokeByAsset(reqData model.ContractInvokeByAssetOperation) model.ContractInvokeByBUResponse {
	var resData model.ContractInvokeByBUResponse
	if reqData.GetSourceAddress() != "" {
		if !keypair.CheckAddress(reqData.GetSourceAddress()) {
			SDKRes := exception.GetSDKRes(exception.INVALID_SOURCEADDRESS_ERROR)
			resData.ErrorCode = SDKRes.ErrorCode
			resData.ErrorDesc = SDKRes.ErrorDesc
			return resData
		}
	}
	if !keypair.CheckAddress(reqData.GetContractAddress()) {
		SDKRes := exception.GetSDKRes(exception.INVALID_CONTRACTADDRESS_ERROR)
		resData.ErrorCode = SDKRes.ErrorCode
		resData.ErrorDesc = SDKRes.ErrorDesc
		return resData
	}
	if reqData.GetSourceAddress() == reqData.GetContractAddress() {
		SDKRes := exception.GetSDKRes(exception.SOURCEADDRESS_EQUAL_DESTADDRESS_ERROR)
		resData.ErrorCode = SDKRes.ErrorCode
		resData.ErrorDesc = SDKRes.ErrorDesc
		return resData
	}
	if len(reqData.GetCode()) > 64 {
		SDKRes := exception.GetSDKRes(exception.INVALID_ASSET_CODE_ERROR)
		resData.ErrorCode = SDKRes.ErrorCode
		resData.ErrorDesc = SDKRes.ErrorDesc
		return resData
	}
	if reqData.GetAmount() < 0 {
		SDKRes := exception.GetSDKRes(exception.INVALID_ASSET_AMOUNT_ERROR)
		resData.ErrorCode = SDKRes.ErrorCode
		resData.ErrorDesc = SDKRes.ErrorDesc
		return resData
	}
	if reqData.GetIssuer() != "" && !keypair.CheckAddress(reqData.GetIssuer()) {
		SDKRes := exception.GetSDKRes(exception.INVALID_ISSUER_ADDRESS_ERROR)
		resData.ErrorCode = SDKRes.ErrorCode
		resData.ErrorDesc = SDKRes.ErrorDesc
		return resData
	}
	_, err := hex.DecodeString(reqData.GetMetadata())
	if err != nil {
		resData.ErrorCode = exception.METADATA_NOT_HEX_STRING_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	var PayAsset protocol.OperationPayAsset
	if reqData.GetCode() != "" && reqData.GetIssuer() != "" && reqData.GetAmount() > 0 {
		PayAsset = protocol.OperationPayAsset{
			DestAddress: reqData.GetContractAddress(),
			Asset: &protocol.Asset{
				Key: &protocol.AssetKey{
					Issuer: reqData.GetIssuer(),
					Code:   reqData.GetCode(),
				},
				Amount: reqData.GetAmount(),
			},
			Input: reqData.GetInput(),
		}
	} else {
		PayAsset = protocol.OperationPayAsset{
			DestAddress: reqData.GetContractAddress(),
			Input:       reqData.GetInput(),
		}
	}

	Operations := []*protocol.Operation{
		{
			SourceAddress: reqData.GetSourceAddress(),
			Metadata:      []byte(reqData.GetMetadata()),
			Type:          protocol.Operation_PAY_ASSET,
			PayAsset:      &PayAsset,
		},
	}
	resData.Result.Operation = *(Operations[0])
	return resData
}

//发送BU并触发合约 InvokeByBU 15
func InvokeByBU(reqData model.ContractInvokeByBUOperation) model.ContractInvokeByBUResponse {
	var resData model.ContractInvokeByBUResponse
	if reqData.GetSourceAddress() != "" {
		if !keypair.CheckAddress(reqData.GetSourceAddress()) {
			SDKRes := exception.GetSDKRes(exception.INVALID_SOURCEADDRESS_ERROR)
			resData.ErrorCode = SDKRes.ErrorCode
			resData.ErrorDesc = SDKRes.ErrorDesc
			return resData
		}
	}
	if !keypair.CheckAddress(reqData.GetContractAddress()) {
		SDKRes := exception.GetSDKRes(exception.INVALID_CONTRACTADDRESS_ERROR)
		resData.ErrorCode = SDKRes.ErrorCode
		resData.ErrorDesc = SDKRes.ErrorDesc
		return resData
	}
	if reqData.GetSourceAddress() == reqData.GetContractAddress() {
		resData.ErrorCode = exception.SOURCEADDRESS_EQUAL_DESTADDRESS_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	if reqData.GetAmount() < 0 {
		SDKRes := exception.GetSDKRes(exception.INVALID_BU_AMOUNT_ERROR)
		resData.ErrorCode = SDKRes.ErrorCode
		resData.ErrorDesc = SDKRes.ErrorDesc
		return resData
	}

	Operations := []*protocol.Operation{
		{
			SourceAddress: reqData.GetSourceAddress(),
			Metadata:      []byte(reqData.GetMetadata()),
			Type:          protocol.Operation_PAY_COIN,
			PayCoin: &protocol.OperationPayCoin{
				DestAddress: reqData.GetContractAddress(),
				Amount:      reqData.GetAmount(),
				Input:       reqData.GetInput(),
			},
		},
	}
	resData.Result.Operation = *(Operations[0])
	return resData
}

//在区块链上写日志信息 LogCreate 16
func LogCreate(reqData model.LogCreateOperation) model.LogCreateResponse {
	var resData model.LogCreateResponse
	if reqData.GetSourceAddress() != "" {
		if !keypair.CheckAddress(reqData.GetSourceAddress()) {
			SDKRes := exception.GetSDKRes(exception.INVALID_SOURCEADDRESS_ERROR)
			resData.ErrorCode = SDKRes.ErrorCode
			resData.ErrorDesc = SDKRes.ErrorDesc
			return resData
		}
	}
	if len(reqData.GetTopic()) >= 128 && len(reqData.GetTopic()) < 0 {
		SDKRes := exception.GetSDKRes(exception.INVALID_LOG_TOPIC_ERROR)
		resData.ErrorCode = SDKRes.ErrorCode
		resData.ErrorDesc = SDKRes.ErrorDesc
		return resData
	}
	for i := range reqData.GetDatas() {
		if len(reqData.GetDatas()[i]) >= 1024 && len(reqData.GetTopic()) < 0 {
			SDKRes := exception.GetSDKRes(exception.INVALID_LOG_DATA_ERROR)
			resData.ErrorCode = SDKRes.ErrorCode
			resData.ErrorDesc = SDKRes.ErrorDesc
			return resData
		}
	}
	_, err := hex.DecodeString(reqData.GetMetadata())
	if err != nil {
		resData.ErrorCode = exception.METADATA_NOT_HEX_STRING_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	Operations := []*protocol.Operation{
		{
			SourceAddress: reqData.GetSourceAddress(),
			Metadata:      []byte(reqData.GetMetadata()),
			Type:          protocol.Operation_LOG,
			Log: &protocol.OperationLog{
				Topic: reqData.GetTopic(),
				Datas: reqData.GetDatas(),
			},
		},
	}
	resData.Result.Operation = *(Operations[0])
	return resData
}
