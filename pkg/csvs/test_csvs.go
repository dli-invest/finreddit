package csvs

import (
	"testing"
	"github.com/dli-invest/finreddit/pkg/util"
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

func TestFindInCsv(t *testing.T) {
	filePath := "view_values_test.csv"
	testData := [][]string{{"col1", "col2"}, {"test11", "test12"}, {"test21", "test22"}}
	AppendToCsv(filePath, testData)
	valueFound := FindInCsv(filePath, "test22", 1)
	util.AssertEqual(t, valueFound, true)
		
}