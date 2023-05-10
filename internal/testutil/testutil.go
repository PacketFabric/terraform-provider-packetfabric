package testutil

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func GenerateUniqueName(prefix string) string {
	return fmt.Sprintf("%s-%s", prefix, uuid.NewString())
}

func GenerateUniqueResourceName() string {
	return fmt.Sprintf("pf_%s", strings.ReplaceAll(uuid.NewString(), "-", "_"))
}

func GetAccountUUID() string {
	return os.Getenv("PF_ACCOUNT_ID")
}

func GetPopAndZoneWithAvailablePort(desiredSpeed string) (pop, zone, media string, availabilityErr error) {
	c, err := _createPFClient()
	if err != nil {
		log.Println("Error creating PF client: ", err)
		return "", "", "", err
	}

	locations, err := c.ListLocations()
	if err != nil {
		log.Println("Error getting locations list: ", err)
		return "", "", "", fmt.Errorf("error getting locations list: %w", err)
	}

	// We need to shuffle the list of locations. Otherwise, we may try to run
	// all tests on the same port.
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(locations), func(i, j int) { locations[i], locations[j] = locations[j], locations[i] })

	testingInLab := strings.Contains(os.Getenv(PF_HOST_KEY), "api.dev")

	for _, l := range locations {
		if l.Vendor == "Colt" {
			continue
		}

		portAvailability, err := c.GetLocationPortAvailability(l.Pop)
		if err != nil {
			log.Println("Error getting location port availability for ", l.Pop, ": ", err)
			return "", "", "", fmt.Errorf("error getting location port availability: %w", err)
		}

		for _, p := range portAvailability {
			if p.Speed == desiredSpeed && p.Count > 0 && (!testingInLab || _contains(listPortsLabMetro, l.Pop)) {
				pop = l.Pop
				zone = p.Zone
				media = p.Media
				log.Println("Found available port at ", pop, zone, media)
				return
			}
		}
	}
	log.Println("No pops with available ports found.")
	return "", "", "", errors.New("no pops with available ports")
}

func PreCheck(t *testing.T, additionalEnvVars []string) {
	requiredEnvVars := []string{
		"PF_HOST",
		"PF_TOKEN",
		"PF_ACCOUNT_ID",
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

func _createPFClient() (*packetfabric.PFClient, error) {
	host := os.Getenv("PF_HOST")
	token := os.Getenv("PF_TOKEN")
	c, err := packetfabric.NewPFClient(&host, &token)
	if err != nil {
		return nil, fmt.Errorf("error creating PFClient: %w", err)
	}
	return c, nil
}
