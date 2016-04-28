package main

import (
  "math/rand"
  "time"
  "os"
  "fmt"
)

func main(){
  a := 123
  b := 532

  if a<b {
    print("A was less than B\n")
  } else if b>a {
    print("B was less than A\n")
  } else {
    print("I guess they were equal\n")
  }

  switch tmpNum := random(); tmpNum {
    case 0,2,4,6,8: print("Even number ", tmpNum, "\n")
    case 1,3,5,7,9: print("Odd number ", tmpNum, "\n")
    default: print("GODZILLAAA\n")
  }

  _, err := os.Open("c:\\randomfilethatdoesnotexists.txt")
  if err != nil {
    fmt.Print("Error found: ", err)
  }
}

func random() int {
  rand.Seed(time.Now().Unix())
  return rand.Intn(10)
}
