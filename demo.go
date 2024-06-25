package main

import (
	"fmt"
	"github.com/Baidu-AIP/golang-sdk/aip/censor"
)

func main() {

	//url := "https://aip.baidubce.com/oauth/2.0/token?client_id=xC4cPdJf8WFj6CLm4CqAnqWY&client_secret=6jgBJl6RkiGA7RxpDckKnaeoTKBw8BCY&grant_type=client_credentials"
	//payload := strings.NewReader(``)
	//client := &http.Client{}
	//req, err := http.NewRequest("POST", url, payload)
	//
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//req.Header.Add("Content-Type", "application/json")
	//req.Header.Add("Accept", "application/json")
	//
	//res, err := client.Do(req)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//defer res.Body.Close()
	//
	//body, err := ioutil.ReadAll(res.Body)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(string(body))

	client := censor.NewClient("ozciKOkPa5PijfyMpTwN9Tme", "5UJ2qGIqXfwcVe1x9UYmcnwSTq1dZXKr")
	//如果是百度云ak sk,使用下面的客户端
	//client := censor.NewCloudClient("ozciKOkPa5PijfyMpTwN9Tme", "5UJ2qGIqXfwcVe1x9UYmcnwSTq1dZXKr")
	res := client.TextCensor("cnm")
	fmt.Println(res)
}
