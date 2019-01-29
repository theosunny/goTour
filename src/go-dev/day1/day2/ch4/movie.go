package ch4

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title: "Casablance", Year: 2012, Color: false, Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: false, Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Bullitt", Year: 1968, Color: true, Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
}

func MarshalDemo() {
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("Json marshal faild %s", err)

	}
	fmt.Printf("%s\n", data)
	fmt.Println("整齐格式的输出")
	data, err = json.MarshalIndent(movies, "", "    ")
	if err != nil {
		log.Fatalf("Json marshal faild %s", err)

	}
	fmt.Printf("%s\n", data)
	fmt.Println("解码")
	var titles []struct{ Title string }
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("json unmarshaling failed:%s", err)
	}
	fmt.Println(titles)

}
