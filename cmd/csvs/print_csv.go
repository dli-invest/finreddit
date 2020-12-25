package main

import (
    "github.com/dli-invest/finreddit/pkg/scrape"
    "fmt"
)
// test script to illustrate csvs operations
func main() {
    records := scrape.ReadCsvFile("cmd/tasks.csv")
    fmt.Println(records)
    for i, r:= range records {
        // header column igonre
        if i == 0 {
            fmt.Println("header line")
        } else {
            fmt.Println(i, r[0], r[1])
        }
	}
}