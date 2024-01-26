package initializers

import (
	"os"
	"testing"
)

func TestLoadEnvVar(t *testing.T) {
	var vars = []string{
		"PORT",
		"SECRET",
		"DATABASE_NAME",
		"DATABASE_COLLECTION",
		"MONGODB_URI",
	}

	for _, j := range vars {
		if _, ok := os.LookupEnv(j); ok == false {
			t.Errorf("'%s' not found in the env", j)
		}
	}
}
