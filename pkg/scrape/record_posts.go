package scrape

// csv utilties to prevent duplicate entries
// probably manually clear out every now and then
import (
    "encoding/csv"
    "fmt"
    "os"
)

func ReadCsvFile(filePath string) [][]string {
    f, err := os.Open(filePath)
    if err != nil {
        fmt.Println("Unable to read input file " + filePath, err)
    }
    defer f.Close()

    csvReader := csv.NewReader(f)
    records, err := csvReader.ReadAll()
    if err != nil {
        fmt.Println("Unable to parse file as CSV for " + filePath, err)
    }

    return records
}

func AppendToCsv(fileName string, data [][]string) {
    f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
    w := csv.NewWriter(f)
    for _, row:= range data {
        w.Write(row)
	}
	w.Flush()
}