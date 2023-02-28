package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-proxy/utils"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"
)

/*
https://cn.developers.tron.network/docs/exchangewallet-integrate-with-the-tron-network

curl -X POST  https://api.trongrid.io/wallet/getblockbynum -d '{"num": 45014224}'
*/

var client = &http.Client{
	Timeout: 20 * time.Second,
}

func main() {
	// uri, _ := url.Parse("")
	client = &http.Client{
		Timeout: 20 * time.Second,
		// Transport: &http.Transport{
		// 	Proxy: http.ProxyURL(uri),
		// },
	}
	/* start 11:38 */
	/* start 12:00 Wifi */
	/* start 13:43 Wifi pm */
	/* start 15:45 Wifi pm proxy  */
	/* start 15:56 Wifi pm proxy v2 109148 */
	/* start 17:18 Wifi pm */
	var allNumber = 38000000
	var limit = 1000000
	startTime := time.Now()
	wg := &sync.WaitGroup{}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(strconv.Itoa(allNumber-(limit*i)) + "### start")
			for num := allNumber - (limit * i); num > allNumber-(limit*(i+1)); num-- {
				fmt.Println("...")
				Ex3(num, "./v3testWifiPm"+strconv.Itoa(i)+".txt")
				fmt.Println(num)
			}
		}(i)
		time.Sleep(1 * time.Second)
	}
	wg.Wait()

	endTime := time.Now()
	fmt.Println(endTime.Sub(startTime))
}

func Ex3(blockNumber int, fileName string) {
	urlPath := "https://api.trongrid.io/wallet/getblockbynum"
	type query struct {
		Num int `json:"num"`
	}

	contentType := "application/json"
	jsonData, _ := json.Marshal(query{
		Num: blockNumber,
	})

	resp, err := client.Post(urlPath, contentType, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println(err)
		Ex3(blockNumber, fileName)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		Ex3(blockNumber, fileName)
		return
	}
	rs := string(body)
	if !strings.HasPrefix(rs, "{\"blockID") {
		fmt.Println("超出速率了，请调整" + rs)
		Ex3(blockNumber, fileName)
		return
	}

	utils.WriteFile(fileName, rs)
	_ = err
}

func Ex1() {
	urlValues := url.Values{}
	urlValues.Add("num", "45014224")
	resp, _ := http.PostForm("https://api.trongrid.io/wallet/getblockbynum", urlValues)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func Ex2() {
	urlValues := url.Values{
		"num": {"45014224"},
	}
	reqBody := urlValues.Encode()
	// resp, _ := http.Post("https://api.trongrid.io/wallet/getblockbynum", "text/html", strings.NewReader(reqBody))
	resp, _ := http.Post("https://api.trongrid.io/wallet/getblockbynum", "application/x-www-form-urlencoded", strings.NewReader(reqBody))
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
