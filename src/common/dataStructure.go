// dataStructure
package common

const Conversion float64 = 100000000

type deal struct {
	Items []Items `json:"items"`
}
type Items struct {
	Transaction_blob string       `json:"transaction_blob"`
	Signatures       []Signatures `json:"signatures"`
}
type Signatures struct {
	Sign_data  string `json:"sign_data"`
	Public_key string `json:"public_key"`
}
type Signers struct {
	SignerAddress string
	Weight        int64
}
type TypeThresholds struct {
	Type      int64
	Threshold int64
}
type ActivateAccountReqData struct {
	SourceAddress string
	DestAddress   string
	InitBalance   int64
}
