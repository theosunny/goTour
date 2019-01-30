package main

import "fmt"

type PersonInfo struct {
	ID      string
	Name    string
	Address string
}

func main() {
	var personDB map[string]PersonInfo
	personDB = make(map[string]PersonInfo)
	//insert data
	personDB["123"] = PersonInfo{"123", "tjheo", "nj"}
	personDB["1234"] = PersonInfo{"1234", "tjheo", "nj"}
	//query value by key
	person, ok := personDB["1234"]
	if ok {
		fmt.Println("Found person", person.Name, "with ID 1234.")
	} else {
		fmt.Println("Did not find person with ID 1234.")

	}
	delete(personDB, "123")
}
