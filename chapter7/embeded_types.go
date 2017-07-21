package main

import "fmt"

func main() {

  person := Person { "Edwin"}

  var android = new(Android)
  android.Model = "Samsung S3"
  android.Name = "Edwin Android"

  defer android.Talk()

  person.Talk()
}

type Person struct {

  Name string
}

func (p *Person) Talk() {

  fmt.Println(p.Name, "is talking...")
}

type Android struct {

  Person
  Model string
}
