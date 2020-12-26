package csvs

// csv utilties to prevent duplicate entries
// probably manually clear out every now and then
import (
    "encoding/csv"
    "fmt"
    "os"
)

// read csv from file path
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

// append rows to csvs
func AppendToCsv(fileName string, data [][]string) {
    f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
    w := csv.NewWriter(f)
    for _, row:= range data {
        err = w.Write(row)
        if err != nil {
            fmt.Println("Append Error")
            fmt.Println(err)
        }
	}
	w.Flush()
}

// check if value exists in csv
func FindInCsv(filePath string, searchValue string, searchColumn int) (bool) {
    records := ReadCsvFile(filePath)
    foundValue := false
    for _, row:= range records {
        valueInRow := row[searchColumn]
        if searchValue == valueInRow {
            foundValue = true
            break
        }
	}
    return foundValue
}