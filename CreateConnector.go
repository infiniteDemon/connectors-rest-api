package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (c Connector) CreateConnector(data interface{}) ([]byte, error) {

	url := fmt.Sprintf("%s://%s:%d/connectors", c.Protocol, c.Host, c.Port)
	method := "POST"

	bt, _ := json.Marshal(data)

	payload := bytes.NewReader(bt)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
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
