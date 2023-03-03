package maxi

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

var url string = "https://api.pcexpress.ca/product-facade/v3/products/search"

func GetAllProducts(keyword string) {
	requestBody := []byte(`{"pagination":{"from":0,"size":48},"banner":"maxi","cartId":"ec770625-d847-4646-8a75-334c906a7be5","lang":"fr","date":"02032023","storeId":"8662","pcId":null,"pickupType":"STORE","offerType":"ALL","term":"eggs","userData":{"domainUserId":"3d30badc-12a3-4fe7-a019-4a86eb1a9908","sessionId":"f1d2e424-416d-4996-a999-544affbb63ee"}}`)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))

	if err != nil {
		log.Println(err)
		return
	}

	req.Header.Add("Accept", "application/json, text/plain, */*")
	req.Header.Add("Accept-Language", "fr")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Referer", "https://www.maxi.ca/")
	req.Header.Add("Sec-Fetch-Dest", "empty")
	req.Header.Add("Sec-Fetch-Mode", "cors")
	req.Header.Add("Sec-Fetch-Site", "cross-site")
	req.Header.Add("Site-Banner", "maxi")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "Windows")
	// api key is the same for every client so no point of hiding it
	req.Header.Add("x-apikey", "1im1hL52q9xvta16GlSdYDsTsG0dmyhF")

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		log.Println(err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Println(err)
	}

	log.Println(string(body))
}
