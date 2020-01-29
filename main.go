package main

import (
  "goNetworkBasic/fetchBasic"
  "fmt"
)

func main() {
  _, err := fetchBasic.Fetch(`https://google.com`, `google`)
  if err != nil {
    fmt.Println(`Error`, err)
  }
}
