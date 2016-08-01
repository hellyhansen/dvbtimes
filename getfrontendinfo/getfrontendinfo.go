package main

import (
  "flag"
  "fmt"
  "log"
  "os"
  "github.com/hellyhansen/dvbtimes/dvb"
)

func main() {
  flag.Parse()
  for _, n := range flag.Args() {
    if f, err := os.Open(n); err != nil {
      log.Printf("Open: %s: %v\n", n, err)
    } else if fi, err := dvb.GetFrontendInfo(f); err != nil {
      log.Printf("GetFrontentInfo: %s: %v", n, err)
    } else {
      fmt.Printf("%s: %#v\n", n, *fi)
    }
  }
}