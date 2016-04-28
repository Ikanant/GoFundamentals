package main

import (

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

  switch a {
    case 0: print("A is ZERO\n")
    case 1: print("A is ONE\n")
    default: print("A was not ONE or ZERO\n")
  }
}
