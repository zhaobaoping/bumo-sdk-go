// asset
package bumo

import (
	"github.com/bumoproject/bumo-sdk-go/src/3rd/proto"
	"github.com/bumoproject/bumo-sdk-go/src/keypair"
	"github.com/bumoproject/bumo-sdk-go/src/protocol"
)

type AssetOperation struct {
	url string
}

//发行资产
func (Asset *AssetOperation) Issue(sourceAddress string, code string, amount int64) ([]byte, Error) {
	if amount <= 0 {
		return nil, sdkErr(INVALID_AMOUNT)
	}
	if len([]rune(code)) > 64 || len([]rune(code)) == 0 {
		return nil, sdkErr(INVALID_CODE)
	}
	if sourceAddress != "" {
		if !keypair.CheckAddress(sourceAddress) {
			return nil, sdkErr(INVALID_SOURCEADDRESS)
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
		Err.Code = PROTO_MARSHAL_ERROR
		Err.Err = err
		return nil, Err
	}
	Err.Code = SUCCESS
	Err.Err = nil
	return data, Err
}

//转移资产
func (Asset *AssetOperation) Pay(sourceAddress string, destAddress string, issueAddress string, amount int64, code string) ([]byte, Error) {
	if !keypair.CheckAddress(issueAddress) {
		return nil, sdkErr(INVALID_ISSUEADDRESS)
	}
	if amount <= 0 {
		return nil, sdkErr(INVALID_AMOUNT)
	}
	if len([]rune(code)) > 64 || len([]rune(code)) == 0 {
		return nil, sdkErr(INVALID_CODE)
	}
	var contractOperation = &ContractOperation{url: Asset.url}
	return contractOperation.InvokeContractByAsset(sourceAddress, destAddress, issueAddress, amount, code, "")
}

//交易BU
func (Asset *AssetOperation) SendBU(sourceAddress string, destAddress string, amount int64) ([]byte, Error) {
	if amount <= 0 {
		return nil, sdkErr(INVALID_AMOUNT)
	}
	var contractOperation = &ContractOperation{url: Asset.url}
	return contractOperation.InvokeContractByBU(sourceAddress, destAddress, amount, "")

}
