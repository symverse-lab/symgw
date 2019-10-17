package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gocheat/symgw/config"
	"github.com/gocheat/symgw/config/db"
	"log"
	"net/http"
	"time"
)

//RPC API 요청
func NodeRpcRequest(url string, request *JsonRpcRequest, httpRequest *http.Request) (*http.Response, *DynamicParameters, error) {
	// Build the request
	body, _ := json.Marshal(request)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, nil, err
	}

	req.Header = httpRequest.Header

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{
		Timeout: time.Second * 5,
	}
	res, err := client.Do(req)
	log.Printf("Proxy ApI Call: %v, Error: $v", req, err)
	if err != nil {
		return nil, nil, err
	}
	defer res.Body.Close()

	fmt.Print("??", res)

	var data DynamicParameters
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return nil, nil, err
	}
	return res, &data, nil
}

//cached 조건에 따른 proxy 호출 처리
func CachedRpcProxy(proxyHost string, r *http.Request, rpcBody *JsonRpcRequest, cache bool) (*DynamicParameters, error) {
	var response *DynamicParameters
	if cache {
		temp := &DynamicParameters{}
		cachedResponse, _ := db.GetCache().Get(Hash(rpcBody))
		err := DecodeBytes(cachedResponse, temp)
		if err != nil {
			panic(err)
		}
		response = temp
	} else {
		res, body, err := NodeRpcRequest(proxyHost, rpcBody, r)
		if err != nil {
			return nil, err
		}
		//save database
		value, _ := GetBytes(body)
		if res.StatusCode == http.StatusOK {
			db.GetCache().Write(Hash(rpcBody), value, time.Minute*time.Duration(config.GetEnv().Cache.Interval))
		}
		response = body
	}
	return response, nil
}
