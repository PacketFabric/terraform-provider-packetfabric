package testutil

import (
	"os"
	"testing"
)

func GetAccountUUID() string {
	return os.Getenv("PF_ACCOUNT_UUID")
}

func PreCheck(t *testing.T, additionalEnvVars *[]string) {
	requiredEnvVars := []string{
		"PF_HOST",
		"PF_TOKEN",
		"PF_ACCOUNT_UUID",
	}
	if additionalEnvVars != nil {
		requiredEnvVars = append(requiredEnvVars, *additionalEnvVars...)
	}
	missing := false
	for _, variable := range requiredEnvVars {
		if _, ok := os.LookupEnv(variable); !ok {
			missing = true
			t.Errorf("`%s` must be set for this acceptance test!", variable)
		}
	}
	if missing {
		t.Fatalf("Some environment variables missing.")
	}
}
