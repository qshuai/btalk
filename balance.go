package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/tidwall/gjson"
)

const (
	bitcoinCashAPI  = "https://bch-tchain.api.btc.com/v3"
	defaultPageSize = 50
)

// get balance for the specified address, and the address should be
// base58 encoded format
func getBalance(addr string) (int64, error) {
	url := bitcoinCashAPI + "/address/" + addr
	res, err := http.Get(url)
	if err != nil {
		return 0, err
	}

	if res.StatusCode != 200 {
		return 0, errors.New("request failed")
	}

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0, err
	}

	return gjson.Get(string(content), "data.balance").Int(), nil
}

// get raw string of unspent list for the specified address
func getUnspent(addr string, page int) (string, error) {
	url := bitcoinCashAPI + "/address/" + addr + "/unspent?pagesize=" +
		strconv.Itoa(defaultPageSize) + "&page=" + strconv.Itoa(page)

	fmt.Println(url)
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}

	if res.StatusCode != 200 {
		return "", errors.New("request failed")
	}

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(content), nil
}
