// request
package bumo

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func getRequest(url string, get string, data string) (*http.Response, Error) {
	var buf bytes.Buffer
	buf.WriteString(url)
	buf.WriteString(get)
	buf.WriteString(data)
	url = buf.String()
	client := &http.Client{}
	newRequest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		Err.Code = HTTP_NEWREQUEST_ERROR
		Err.Err = err
		return nil, Err
	}
	response, err := client.Do(newRequest)
	if err != nil {
		Err.Code = CLIENT_DO_ERROR
		Err.Err = err
		return nil, Err
	}
	Err.Code = SUCCESS
	Err.Err = nil
	return response, Err
}
func postRequest(url string, post string, data []byte) (*http.Response, Error) {
	var buf bytes.Buffer
	buf.WriteString(url)
	buf.WriteString(post)
	url = buf.String()
	client := &http.Client{}
	newRequest, err := http.NewRequest("POST", url, bytes.NewReader(data))
	if err != nil {
		Err.Code = HTTP_NEWREQUEST_ERROR
		Err.Err = err
		return nil, Err
	}
	response, err := client.Do(newRequest)
	if err != nil {
		Err.Code = CLIENT_DO_ERROR
		Err.Err = err
		return nil, Err
	}
	Err.Code = SUCCESS
	Err.Err = nil
	return response, Err
}
func getRequestJson(transactionBlob string, signatures []Signatures) ([]byte, Error) {
	request := make(map[string]interface{})
	items := make([]map[string]interface{}, 1)
	items[0] = make(map[string]interface{})
	items[0]["transaction_blob"] = transactionBlob
	items[0]["signatures"] = signatures
	request["items"] = items
	requestJson, err := json.Marshal(request)
	if err != nil {
		Err.Code = JSON_MARSHAL_ERROR
		Err.Err = err
		return nil, Err
	}
	Err.Code = SUCCESS
	Err.Err = nil
	return requestJson, Err
}
