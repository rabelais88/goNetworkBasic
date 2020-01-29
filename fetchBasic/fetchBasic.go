package fetchBasic

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Fetch(url string, filename string) ([]byte, error) {
	// fetch body
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	// write file
	err = ioutil.WriteFile(filename, body, 0644)
	if err != nil {
		return nil, err
	}
	return body, nil
}
