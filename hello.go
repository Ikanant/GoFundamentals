package main

import (
  "fmt"
)

func main(){
  fmt.Println( average(1,2,3,4,5) )

}

func average(countList ...int) int {
  sum := 0
  for _, val := range countList {
    sum += val
  }

  return sum
}
