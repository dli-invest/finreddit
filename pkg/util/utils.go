package util
import (
	"log"
	"fmt"
	"os"
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
