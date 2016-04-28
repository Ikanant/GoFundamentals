package main

import (
  "fmt"
  "os"
)

func main(){
  loggedName := os.Getenv("USERNAME")

  fmt.Println(loggedName)
}
