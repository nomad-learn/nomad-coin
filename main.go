package main

import (
	"fmt"

	"github.com/nomad-learn/nomad-coin/person"
)

func main() {
	juno := person.Person{}
	juno.SetDetails("nico", 12)
	fmt.Println(juno)
}
