package testutil

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var ErrNoPops = errors.New("no pops with available ports")

func GenerateUniqueName(prefix string) string {
	return fmt.Sprintf("%s-%s", prefix, uuid.NewString())
}

func GenerateUniqueResourceName() string {
	return fmt.Sprintf("pf_%s", strings.ReplaceAll(uuid.NewString(), "-", "_"))
}

func GetAccountUUID() string {
	return os.Getenv("PF_ACCOUNT_UUID")
}

func GetPopAndZoneWithAvailablePort(speed string) (string, string, error) {
	host := os.Getenv("PF_HOST")
	token := os.Getenv("PF_TOKEN")
	c, err := packetfabric.NewPFClient(&host, &token)
	if err != nil {
		return "", "", fmt.Errorf("error creating PFClient: %w", err)
	}

	locations, err := c.ListLocations()
	if err != nil {
		return "", "", fmt.Errorf("error getting locations list: %w", err)
	}
	for _, l := range locations {
		if l.Vendor == "Colt" {
			continue
		}
		portAvailability, err := c.GetLocationPortAvailability(l.Pop)
		if err != nil {
			return "", "", fmt.Errorf("error getting location port availability: %w", err)
		}
		for _, p := range portAvailability {
			if p.Count > 0 && p.Speed == speed {
				return l.Pop, p.Zone, nil
			}
		}
	}
	return "", "", ErrNoPops
}

func PreCheck(t *testing.T, additionalEnvVars []string) {
	requiredEnvVars := []string{
		"PF_HOST",
		"PF_TOKEN",
		"PF_ACCOUNT_UUID",
	}
	if additionalEnvVars != nil {
		requiredEnvVars = append(requiredEnvVars, additionalEnvVars...)
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

func SkipIfEnvNotSet(t *testing.T) {
	if os.Getenv(resource.EnvTfAcc) == "" {
		t.Skip()
	}
}
