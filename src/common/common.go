// request
package common

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/bumoproject/bumo-sdk-go/src/exception"
)

var Err exception.Error

func GetRequest(url string, get string, data string) (*http.Response, exception.Error) {
	var buf bytes.Buffer
	buf.WriteString(url)
	buf.WriteString(get)
	buf.WriteString(data)
	url = buf.String()
	client := &http.Client{}
	newRequest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		Err.Code = exception.HTTP_NEWREQUEST_ERROR
		Err.Err = err
		return nil, Err
	}
	response, err := client.Do(newRequest)
	if err != nil {
		Err.Code = exception.CLIENT_DO_ERROR
		Err.Err = err
		return nil, Err
	}
	Err.Code = exception.SUCCESS
	Err.Err = nil
	return response, Err
}
func PostRequest(url string, post string, data []byte) (*http.Response, exception.Error) {
	var buf bytes.Buffer
	buf.WriteString(url)
	buf.WriteString(post)
	url = buf.String()
	client := &http.Client{}
	newRequest, err := http.NewRequest("POST", url, bytes.NewReader(data))
	if err != nil {
		Err.Code = exception.HTTP_NEWREQUEST_ERROR
		Err.Err = err
		return nil, Err
	}
	response, err := client.Do(newRequest)
	if err != nil {
		Err.Code = exception.CLIENT_DO_ERROR
		Err.Err = err
		return nil, Err
	}
	Err.Code = exception.SUCCESS
	Err.Err = nil
	return response, Err
}
func GetRequestJson(transactionBlob string, signatures []Signatures) ([]byte, exception.Error) {
	request := make(map[string]interface{})
	items := make([]map[string]interface{}, 1)
	items[0] = make(map[string]interface{})
	items[0]["transaction_blob"] = transactionBlob
	items[0]["signatures"] = signatures
	request["items"] = items
	requestJson, err := json.Marshal(request)
	if err != nil {
		Err.Code = exception.JSON_MARSHAL_ERROR
		Err.Err = err
		return nil, Err
	}
	Err.Code = exception.SUCCESS
	Err.Err = nil
	return requestJson, Err
}

//获取最新fees
func GetFees(url string) (int64, int64, exception.Error) {
	get := "/getLedger?with_fee=true"
	response, Err := GetRequest(url, get, "")
	if Err.Err != nil {
		return 0, 0, Err
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
			return 0, 0, Err
		}
		if data["error_code"].(json.Number) == "0" {
			result := data["result"].(map[string]interface{})
			fees := result["fees"].(map[string]interface{})
			gasPriceStr, ok := fees["gas_price"].(json.Number)
			if ok != true {
				Err.Code = exception.SUCCESS
				Err.Err = nil
				return 0, 0, Err
			}
			baseReserveStr, ok := fees["base_reserve"].(json.Number)
			if ok != true {
				Err.Code = exception.SUCCESS
				Err.Err = nil
				return 0, 0, Err
			}
			gasPrice, err := strconv.ParseInt(string(gasPriceStr), 10, 64)
			if err != nil {
				Err.Code = exception.STRCONV_PARSEINT_ERROR
				Err.Err = err
				return 0, 0, Err
			}
			baseReserve, err := strconv.ParseInt(string(baseReserveStr), 10, 64)
			if err != nil {
				Err.Code = exception.STRCONV_PARSEINT_ERROR
				Err.Err = err
				return 0, 0, Err
			}
			Err.Code = exception.SUCCESS
			Err.Err = nil
			return gasPrice, baseReserve, Err
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
