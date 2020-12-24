package util

import (
    "testing"
    "os"
)


func TestGetEnvVar(t *testing.T) {
    test_value := "test"
    os.Setenv("TEST_VAR", test_value)
    test_result :=GetEnvVar("TEST_VAR")
    if test_value != test_result {
        t.Errorf("Failed to get matching value")
    }
}