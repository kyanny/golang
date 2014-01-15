package main

import (
	"os"
	"fmt"
	"encoding/csv"
)

func main() {
	file, err := os.Open("./smap.csv")
	if err != nil {
		panic(err)
	}
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}
	for i, record := range records {
		fmt.Println(i, record)
	}
}

// 0 [first_name last_name retired]
// 1 [masahiro nakai false]
// 2 [takuya kimura false]
// 3 [goro inagaki false]
// 4 [tsuyoshi kusanagi false]
// 5 [shingo katori false]
// 6 [katsuyuki mori true]
