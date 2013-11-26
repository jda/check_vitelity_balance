package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type VitelityResponse struct {
	XMLName  xml.Name `xml:"content"`
	Status   string   `xml:"status"`
	Response string   `xml:"response"`
}

func get_balance(user string, pass string) (float64, error) {
	// set up request
	c := &http.Client{}

	req_url := fmt.Sprintf("%s?login=%s&pass=%s&cmd=balance&xml=yes", api_url, user, pass)

	if verbose >= 3 {
		fmt.Println("Request URL:", req_url)
	}

	req, err := http.NewRequest("GET", req_url, nil)
	if err != nil {
		return 0, err
	}

	req.Header.Set("User-Agent", version)

	// make request
	resp, err := c.Do(req)
	if err != nil {
		return 0, err
	}

	// read response
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	// parse xml
	res := VitelityResponse{}
	err = xml.Unmarshal(body, &res)
	if err != nil {
		return 0, err
	}

	if verbose >= 3 {
		fmt.Printf("%+v\n", &res)
	}

	if res.Status != "ok" {
		return 0, fmt.Errorf("Vitelity API Response: %s", res.Status)
	}

	balance, err := strconv.ParseFloat(res.Response, 2)
	if err != nil {
		return 0, fmt.Errorf("Can't parse API responce for balance: %s", res.Response)
	}

	return balance, nil
}
