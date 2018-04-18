package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	// var g1, g2 Girl
	// g1 = Girl{Name: "g1", Age: 17, Like: "sing", Email: "g1@gmail.com"}
	// g2 = Girl{"g2", 23, "tour", "g2@gmail.com"}
	// var girls = []Girl{g1, g2}
	// fmt.Printf("%q", marshal(girls))
	// fmt.Println(unmarshal(marshal(girls), girls[:0]))
	// leftData(marshal(girls))
	github()
}
func marshal(girls []Girl) []byte {
	data, err := json.Marshal(girls)
	// data, err := json.MarshalIndent(girls, "", "   ")
	if err != nil {
		fmt.Printf("JSON marshal filed:\t%v\n", err)
	}
	return data
}
func unmarshal(bytes []byte, girls []Girl) []Girl {
	if err := json.Unmarshal(bytes, &girls); err != nil {
		log.Fatalf("error:\t%v\n", err)
	}
	return girls
}
func leftData(bytes []byte) {
	var gs []SimpleGirl
	if err := json.Unmarshal(bytes, &gs); err != nil {
		log.Fatalf("error:\t%v\n", err)
	}
	fmt.Println(gs)
}
func github() {
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}

}
