package csvs

import (
    "testing"
)

// make read the csv read back into the file is the same
func TestAppendToCsv(t *testing.T) {
	filePath := "record_posts_test.csv"
	testData := [][]string{{"col1", "col2"}, {"test1", "test2"}}
	AppendToCsv(filePath, testData)
	readData := ReadCsvFile(filePath)
	if testData != nil {
        t.Errorf("Failed to get matching value")
	}
	if readData != nil {
        t.Errorf("Failed to get matching value")
	}
	
	if testData[0][0] == readData[0][0] {
		t.Errorf("Csv values do not match")
	}
}