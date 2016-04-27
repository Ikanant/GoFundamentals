package main

import (
  "fmt"
)

func main(){
  firstName := "Jonathan"
  lastName := "Hernandez"
  age := 23
  module := &age

  fmt.Println("My name is: ", firstName, " ", lastName)
  fmt.Println("I am: ", age, " years old\n\n")

  fmt.Println("The AGE variable memory address is: ", module, " and its value is: ", *module)
}
