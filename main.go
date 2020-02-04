package main

import (
	"bytes"
	"fmt"
	"goNetworkBasic/fetchBasic"
	"goNetworkBasic/logWriter"
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
}
