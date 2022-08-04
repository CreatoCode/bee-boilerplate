package httpClient

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	GET = "GET"
)

func Get(url string, data []byte) (body []byte, err error) {

	url = fmt.Sprint(url, "?", string(data))
	data = nil

	client := http.Client{}
	req, err := http.NewRequest(GET, url, bytes.NewBuffer(data))
	if err != nil {
		return body, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err = ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return body, err
	}

	return body, err
}
