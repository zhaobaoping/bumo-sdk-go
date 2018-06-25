// asset
package asset

import (
	"github.com/bumoproject/bumo-sdk-go/src/common/proto"
	"github.com/bumoproject/bumo-sdk-go/src/common/protocol"
	"github.com/bumoproject/bumo-sdk-go/src/contract"
	"github.com/bumoproject/bumo-sdk-go/src/crypto/keypair"
	"github.com/bumoproject/bumo-sdk-go/src/exception"
)

var Err exception.Error

type AssetOperation struct {
	url string
}

//发行资产
func (Asset *AssetOperation) Issue(sourceAddress string, code string, amount int64) ([]byte, exception.Error) {
	if amount <= 0 {
		return nil, exception.SdkErr(exception.INVALID_AMOUNT)
	}
	if len([]rune(code)) > 64 || len([]rune(code)) == 0 {
		return nil, exception.SdkErr(exception.INVALID_CODE)
	}
	if sourceAddress != "" {
		if !keypair.CheckAddress(sourceAddress) {
			return nil, exception.SdkErr(exception.INVALID_SOURCEADDRESS)
		}
	}
	Operations := []*protocol.Operation{
		{
			SourceAddress: sourceAddress,
			Type:          protocol.Operation_ISSUE_ASSET,
			IssueAsset: &protocol.OperationIssueAsset{
				Code:   code,
				Amount: amount,
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

//转移资产
func (Asset *AssetOperation) Pay(sourceAddress string, destAddress string, issueAddress string, amount int64, code string) ([]byte, exception.Error) {
	if !keypair.CheckAddress(issueAddress) {
		return nil, exception.SdkErr(exception.INVALID_ISSUEADDRESS)
	}
	if amount <= 0 {
		return nil, exception.SdkErr(exception.INVALID_AMOUNT)
	}
	if len([]rune(code)) > 64 || len([]rune(code)) == 0 {
		return nil, exception.SdkErr(exception.INVALID_CODE)
	}
	var contractOperation = &contract.ContractOperation{Url: Asset.url}
	return contractOperation.InvokeContractByAsset(sourceAddress, destAddress, issueAddress, amount, code, "")
}

//交易BU
func (Asset *AssetOperation) SendBU(sourceAddress string, destAddress string, amount int64) ([]byte, exception.Error) {
	if amount <= 0 {
		return nil, exception.SdkErr(exception.INVALID_AMOUNT)
	}
	var contractOperation = &contract.ContractOperation{Url: Asset.url}
	return contractOperation.InvokeContractByBU(sourceAddress, destAddress, amount, "")

}
