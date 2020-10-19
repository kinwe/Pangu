package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode == 200 {
		return string(body)
	}
	return ""
}
