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
)

var birdNames = []string{
	"albatross",
	"blackbird",
	"canary",
	"dove",
	"eagle",
	"falcon",
	"goldfinch",
	"hawk",
	"ibis",
	"jay",
	"kite",
	"lark",
	"magpie",
	"nightingale",
	"owl",
	"parrot",
	"quail",
	"raven",
	"sparrow",
	"toucan",
	"vulture",
	"woodpecker",
	"xantus",
	"yellowhammer",
	"zebra finch",
}

func GenerateUniqueName() string {
	rand.Seed(time.Now().UnixNano())
	birdName := birdNames[rand.Intn(len(birdNames))]
	return fmt.Sprintf("terraform_testacc_%s", birdName)
}
func GenerateUniqueResourceName(resource string) (resourceName, hclName string) {
	uuid := uuid.NewString()
	shortUuid := uuid[0:8]
	randomNumber := rand.Intn(9000) + 1000
	hclName = fmt.Sprintf("terraform_testacc_%s_%d", shortUuid, randomNumber)
	resourceName = fmt.Sprintf("%s.%s", resource, hclName)
	return
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

func _contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func GetPopAndZoneWithAvailablePort(desiredSpeed string, skipDesiredMarket *string) (pop, zone, media, market string, availabilityErr error) {

	c, err := _createPFClient()
	if err != nil {
		log.Println("Error creating PF client: ", err)
		return "", "", "", "", err
	}

	locations, err := c.ListLocations()
	if err != nil {
		log.Println("Error getting locations list: ", err)
		return "", "", "", "", fmt.Errorf("error getting locations list: %w", err)
	}

	// We need to shuffle the list of locations. Otherwise, we may try to run
	// all tests on the same port.
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(locations), func(i, j int) { locations[i], locations[j] = locations[j], locations[i] })

	testingInLab := strings.Contains(os.Getenv("PF_HOST"), "api.dev")

	for _, l := range locations {
		// log.Printf("Checking PoP: %s\n", l.Pop)
		// Skip Colt locations
		if l.Vendor == "Colt" {
			continue
		}

		// Do not select a port in the same market as the one set in skipDesiredMarket
		if skipDesiredMarket != nil && l.Market == *skipDesiredMarket {
			continue
		}
		portAvailability, err := c.GetLocationPortAvailability(l.Pop)
		if err != nil {
			log.Println("Error getting location port availability for ", l.Pop, ": ", err)
			return "", "", "", "", fmt.Errorf("error getting location port availability: %w", err)
		}

		for _, p := range portAvailability {
			if p.Speed == desiredSpeed && p.Count > 0 && (!testingInLab || _contains(labPopsPort, l.Pop)) {
				pop = l.Pop
				zone = p.Zone
				media = p.Media
				market = l.Market
				log.Println("Found available port at ", pop, zone, media, market)
				if skipDesiredMarket == nil {
					log.Println("Not specified Market to avoid.")
				} else {
					log.Println("Specified Market to avoid: ", *skipDesiredMarket)
				}

				return
			}
		}
	}
	log.Println("No pops with available ports found.")
	return "", "", "", "", errors.New("no pops with available ports")
}

func (details PortDetails) FindAvailableCloudPopZone() (pop, zone, region string) {
	popsWithZones, _ := details.FetchCloudPopsAndZones()
	popsToSkip := make([]string, 0)

	log.Println("Starting to search for available Cloud PoP and zone...")
	log.Printf("Available PoPs with Zones: %v\n", popsWithZones)

	testingInLab := strings.Contains(os.Getenv("PF_HOST"), "api.dev")

	for popAvailable, zones := range popsWithZones {
		// log.Printf("Checking PoP: %s\n", popAvailable)
		if len(popsToSkip) == len(popsWithZones) {
			log.Fatal(errors.New("there's no port available on any pop"))
		}
		if _contains(popsToSkip, popAvailable) {
			log.Printf("PoP %s is in popsToSkip, skipping...\n", popAvailable)
			continue
		} else {
			if len(zones) > 1 && (!testingInLab || _contains(labPopsCloud, popAvailable)) {
				pop = popAvailable
				zone = zones[0]
				region = zones[len(zones)-1]
				log.Printf("Found available PoP: %s, Zone: %s, Region: %s\n", pop, zone, region)
				return
			} else {
				popsToSkip = append(popsToSkip, popAvailable)
			}
		}
	}

	log.Println("No available Cloud PoP, zone, and region found.")
	return
}

func (details PortDetails) FetchCloudPopsAndZones() (popsWithZones map[string][]string, err error) {
	if details.DesiredProvider == "" {
		err = errors.New("please provide a valid cloud provider to fetch pop")
	}
	if details.PFClient == nil {
		err = errors.New("please create PFClient to fetch cloud pop")
		return
	}
	popsWithZones = make(map[string][]string)
	if cloudLocations, locErr := details.PFClient.GetCloudLocations(
		details.DesiredProvider,
		details.DesiredConnectionType,
		details.IsNatCapable,
		details.HasCloudRouter,
		details.AnyType,
		details.DesiredPop,
		details.DesiredCity,
		details.DesiredState,
		details.DesiredMarket,
		details.DesiredRegion); locErr != nil {
		err = locErr
		return
	} else {
		for _, loc := range cloudLocations {
			popsWithZones[loc.Pop] = append(loc.Zones, loc.CloudConnectionDetails.Region)

		}
	}
	return
}

func CreateBasePortDetails() PortDetails {
	c, err := _createPFClient()
	if err != nil {
		log.Panic(err)
	}
	return PortDetails{
		PFClient:          c,
		DesiredSpeed:      portSpeed,
		skipDesiredMarket: nil,
	}
}

func setAzureLocations(host string) (string, string, string) {
	var AzureLocation string
	var AzurePeeringLocation string
	var AzureServiceProviderName string

	testingInLab := strings.Contains(host, "api.dev")

	if testingInLab {
		AzureLocation = AzureLocationDev
		AzurePeeringLocation = AzurePeeringLocationDev
		AzureServiceProviderName = AzureServiceProviderNameDev
	} else {
		AzureLocation = AzureLocationProd
		AzurePeeringLocation = AzurePeeringLocationProd
		AzureServiceProviderName = AzureServiceProviderNameProd
	}

	return AzureLocation, AzurePeeringLocation, AzureServiceProviderName
}
