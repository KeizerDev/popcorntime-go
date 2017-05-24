package main

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

const endpoint string = "https://tv-v2.api-fetch.website"

type Core struct {
}

func getURI(action string) string {
	return endpoint + action
}

func (c Core) Get(action string, d interface{}) error {
	return c.request("GET", action, d, nil)
}

func (c Core) Post(action string, d interface{}, postData interface{}) error {
	postStr, err := json.Marshal(postData)
	if err != nil {
		return err
	}

	reader := strings.NewReader(string(postStr))
	return c.request("POST", action, d, reader)
}

func (c Core) request(method, action string, d interface{}, reader io.Reader) error {
	req, err := http.NewRequest(method, getURI(action), reader)
	if err != nil {
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err := resp.Body.Close(); err != nil {
		return err
	}

	if resp.StatusCode >= 400 {
		return errors.New(resp.Status)
	}

	err = json.Unmarshal(data, &d)
	if err != nil {
		return err
	}
	return nil
}