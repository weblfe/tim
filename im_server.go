/*
 *
 * im_server.go
 * server
 *
 * Created by lintao on 2020/4/16 2:20 下午
 * Copyright © 2020-2020 LINTAO. All rights reserved.
 *
 */

package tim

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	sign2 "github.com/tencentyun/tls-sig-api-v2-golang/tencentyun"
)

type IMServer struct {
	AppId      int
	Identifier string
	SecretKey  string
	Expire     int
	Sig        string
}

func NewIMServer(appId, expire int, identifier, secretKey string, opts ...ServerOption) (IMServer, error) {
	server := IMServer{
		AppId:      appId,
		Identifier: identifier,
		SecretKey:  secretKey,
		Expire:     expire,
	}
	var err error
	if server.Sig, err = server.userSig(); err != nil {
		return IMServer{}, err
	}
	for _, opt := range opts {
		if err := opt.SetOption(&server); err != nil {
			return IMServer{}, err
		}
	}

	return server, nil
}

func (s IMServer) ListenCallback() {
	Run(strconv.Itoa(s.AppId))
}

func (s IMServer) request(url string, requestJson []byte) ([]byte, error) {
	body := bytes.NewBuffer(requestJson)
	// Create client
	log.Println(string(requestJson))
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		log.Panic(err)
	}
	// Headers
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	// Fetch Request
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer func(){
		if err:=resp.Body.Close();err!=nil{
			log.Println("response body close error:",err)
		}
	}()

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http erro status code is %d", resp.StatusCode)
	}
	var respCheck struct {
		ErrorInfo string `json:"ErrorInfo"`
		ErrorCode int    `json:"ErrorCode"`
	}
	if err = json.Unmarshal(respBody, &respCheck); err != nil {
		return nil, err
	}

	if respCheck.ErrorCode != 0 {
		return nil, fmt.Errorf("操作失败，错误码 %d ,错误信息 %s \n 详情请查询 "+
			"https://cloud.tencent.com/document/product/269/1671", respCheck.ErrorCode, respCheck.ErrorInfo)
	}

	log.Println(string(respBody))
	return respBody, nil

}

func (s IMServer) requestWithPath(path ApiPath, v interface{}) (jsons []byte, err error) {
	var b []byte
	switch a := v.(type) {
	case []byte:
		b = a
	default:
		b, err = json.Marshal(&v)
		if err != nil {
			return nil, err
		}
	}

	return s.request(s.combineURL(path), b)
}

func (s IMServer) userSig() (string, error) {
	userSig, err := sign2.GenUserSig(s.AppId, s.SecretKey, s.Identifier, s.Expire)
	if err != nil {
		return "", err
	}
	return userSig, nil
}

func (s IMServer) combineURL(path ApiPath) string {
	rand.Seed(time.Now().Unix())

	return fmt.Sprintf("%s/%s%s?sdkappid=%d&identifier=%s&usersig=%s&random=%d&contenttype=json",
		BaseUrl, VERSION, path, s.AppId, s.Identifier, s.Sig, rand.Intn(4294967294))
}

func (s IMServer) Do(path ApiPath, v interface{}) (jsons []byte, err error) {
	return s.requestWithPath(path,v)
}
