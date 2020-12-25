package util

import (
    "testing"
    "os"
    "fmt"
)


func TestGetEnvVar(t *testing.T) {
    test_value := "test"
    os.Setenv("TEST_VAR", test_value)
    test_result :=GetEnvVar("TEST_VAR")
    if test_value != test_result {
        t.Errorf("Failed to get matching value")
    }
}

func TestNewConfig(t *testing.T) {
    cfg_path := MkPathFromStr("test.yml")
    cfg, err := NewConfig(cfg_path)
    fmt.Println(cfg)
    if err != nil {
        t.Errorf("Config import error")
    }
    subreddits := cfg.Data.SubReddits

    num_subreddits := len(subreddits)
    AssertEqual(t, num_subreddits, 3)
    sr_minscore := subreddits[1].MinScore
    AssertEqual(t, sr_minscore, 50)
    sr_mincomments := subreddits[1].MinComments
    AssertEqual(t, sr_mincomments, 20)
    
    sr_name := subreddits[0].Name
    AssertEqual(t, sr_name, "investing")
}
