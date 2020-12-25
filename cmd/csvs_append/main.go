package main

import (
	// "encoding/csv"
	// "fmt"
	// "os"
	"github.com/dli-invest/finreddit/pkg/csvs"
)

func main() {
	// original sample at https://play.golang.org/p/CzFNG1eDec
	// f, err := os.OpenFile("test.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// w := csv.NewWriter(f)
	// for i := 0; i < 10; i++ {
	// 	w.Write([]string{"a", "b", "c"})
	// }
	// w.Flush()
	testData := [][]string{{"col1", "col2"}, {"test1", "test2"}}
	csvs.AppendToCsv("tasks.csv", testData)
	// append rows to column
}
