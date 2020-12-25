package util
import (
	"log"
	"fmt"
	"os"
    "gopkg.in/yaml.v2"
    "testing"
    "reflect"
    "path/filepath"
    "github.com/dli-invest/finreddit/pkg/types"
)

// Gets enviroment variable if available
// throws error if not available
func GetEnvVar(varName string) string {
	al, present := os.LookupEnv(varName)
	if present {
		return al 
	} else {
		var error_message = fmt.Sprintf("Environment Variable: %s missing", varName)
		log.Fatal(error_message)
		return ""
	}
}


func ValidateConfigPath(path string) error {
    s, err := os.Stat(path)
    if err != nil {
        return err
    }
    if s.IsDir() {
        return fmt.Errorf("'%s' is a directory, not a normal file", path)
    }
    return nil
}

// NewConfig returns a new decoded Config struct
func NewConfig(configPath string) (*types.SearchConfig, error) {
	// check if path is valid
	err := ValidateConfigPath(configPath)
    // Create config structure
    config := &types.SearchConfig{}

    // Open config file
    file, err := os.Open(configPath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    // Init new YAML decode
    d := yaml.NewDecoder(file)

    // Start YAML decoding from file
    if err := d.Decode(&config); err != nil {
        return nil, err
    }
    log.Println(config)
    return config, nil
}

func MkPathFromStr(pathStr string) (string) {
    p := filepath.FromSlash(pathStr)
    return p
}

// AssertEqual checks if values are equal
func AssertEqual(t *testing.T, a interface{}, b interface{}) {
	if a == b {
		return
	}
	// debug.PrintStack()
	t.Errorf("Received %v (type %v), expected %v (type %v)", a, reflect.TypeOf(a), b, reflect.TypeOf(b))
}