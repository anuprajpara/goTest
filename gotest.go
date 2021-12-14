package gotest 

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)


func GoTestFunction() {
	fmt.Println("Hello from go package!!")	
}

func FromAPI(myURL string) (string, error) {
	u, err := url.Parse(myURL)
	if err != nil {
		log.Fatal(err)
	}
	u.Path += "/abc"
	params := url.Values{}
	params.Add("options", "version|destination|file|series|package|linear|network|episode|premiere|mediamanifest|zeusmetadata|status|mduversion|title|encom")
	u.RawQuery = params.Encode()
	modURL := u.String()

	client := &http.Client{}

	req, err := http.NewRequest("GET", modURL, nil)
	if err != nil {
		fmt.Println("http.NewRequest", err)
	}
	req.Header.Set("content-type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("X-Consumer-ID", "Lockstep-Go")

	fmt.Println("req", req)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("client.Do", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err, "ioutil.ReadAll failed", err)
	}
	if resp.StatusCode != http.StatusOK {
		err = errors.New("Got status code :" + resp.Status)
		return string{}, err
	}

	var payload string
	err = json.Unmarshal([]byte(body), &payload)
	if err != nil {
		fmt.Println("json.Unmarshal", err)
	}
	return payload, err
}
