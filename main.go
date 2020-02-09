package main

import (
	"bytes"
	"fmt"
	"goNetworkBasic/fetchBasic"
	"goNetworkBasic/logWriter"
	"goNetworkBasic/statusCheck"
	"io"
	"io/ioutil"
)

func main() {
	_, err := fetchBasic.Fetch(`https://google.com`, `google`)
	if err != nil {
		fmt.Println(`Error`, err)
	}
	ctx, err := ioutil.ReadFile(`google`)
	if err != nil {
		fmt.Println(`Error`, err)
	}
	fmt.Println("total byte", len(ctx), "byte")
	lw := logWriter.LogWriter{}
	io.Copy(lw, bytes.NewReader(ctx))

	links := []string{
		`https://google.com`,
		`https://naver.com`,
		`https://amazon.com`,
	}
	statusCheck.Check(links)
}
