package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func (c Connector) GetConnector() ([]byte, error) {
	url := fmt.Sprintf("%s://%s:%d/connectors", c.Protocol, c.Host, c.Port)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, err
	}
	//req.Header.Add("User-Agent", "Apifox/1.0.0 (https://www.apifox.cn)")
	//req.Header.Add("Host", "192.168.78.141:8083")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Connection", "keep-alive")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
