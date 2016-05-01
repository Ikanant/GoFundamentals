package main

import(
  "fmt"
)

func main() {

    type courseMeta struct {
      Author string
      Level string
      Rating float64
    }

    var DockerDeepDive courseMeta

    // Using the new Keyword gives us a pointer
    DockerDeepDive2 := new(courseMeta)

    DockerDeepDive3 := courseMeta{
      Author: "Nigel Poulton",
      Level: "Intermediate",
      Rating: 5,
    }

    fmt.Println(DockerDeepDive)
    fmt.Println(DockerDeepDive2)
    fmt.Println(DockerDeepDive3.Author)
}
