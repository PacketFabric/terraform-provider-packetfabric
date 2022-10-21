package packetfabric

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

var host = "https://packetfabric.fakeurl.com"
var token = "2512d9bf-fd4c-46ae-a340-1d663f4fb01b"

var cTest, _ = NewPFClient(&host, &token)

var regions = make([]Region, 0)

func init() {
	regions = append(regions, Region{Name: "Continental U.S.", Code: "US"})
	regions = append(regions, Region{Name: "Europe", Code: "UK"})
}

var clrExpectedResp = CloudRouterResponse{
	CircuitID: "PF-L3-CUST-2",
	Scope:     "private",
	Asn:       4556,
	Name:      "Super Cool Cloud router",
	Capacity:  "10Gbps",
	Regions: []Region{
		{Name: "Continental U.S.", Code: "US"},
	},
	TimeCreated: "2020-08-20T22:08:37.000000+0000",
	TimeUpdated: "2020-08-20T22:08:37.000000+0000",
}

func Test_CreateCloudRouter(t *testing.T) {

	router := CloudRouter{
		Scope:       "private",
		Asn:         4556,
		Name:        "New Cloud Router",
		AccountUUID: "3482182c-b483-45e0-b8f7-5562bba57e6b",
		Regions:     []string{"UK", "US"},
		Capacity:    "10Gbps",
	}

	cTest.runFakeHttpServer(_callCreateRouter, router, clrExpectedResp, _buildMockCloudRouterResp(), "cloud-router-test", t)
}

func Test_UpdateCloudRouter(t *testing.T) {
	routerUpdt := CloudRouterUpdate{
		Name:     "New Cloud Router",
		Regions:  []string{"UK"},
		Capacity: "1Gbps",
	}

	cTest.runFakeHttpServer(_callUpdateRouter, routerUpdt, clrExpectedResp, _buildMockCloudRouterUpdateResp(), "cloud-router-update-test", t)

}

func Test_ReadCloudRouter(t *testing.T) {
	router := CloudRouter{
		AccountUUID: "3482182c-b483-45e0-b8f7-5562bba57e6b",
	}

	cTest.runFakeHttpServer(_callReadRouter, router, clrExpectedResp, _buildMockCloudRouterUpdateResp(), "cloud-router-read-test", t)
}

func Test_ListCloudRouters(t *testing.T) {
	var expectedResp []CloudRouterResponse
	_ = json.Unmarshal(_buildMockCloudRouterResps(), &expectedResp)
	cTest.runFakeHttpServer(_callListCloudRouters, nil, expectedResp, _buildMockCloudRouterResps(), "list-cloud-routers", t)
}

func Test_DeleteCoudRouter(t *testing.T) {
	router := CloudRouter{
		AccountUUID: "3482182c-b483-45e0-b8f7-5562bba57e6b",
	}

	delResp := CloudRouterDelResp{
		Message: "Cloud router deleted.",
	}

	cTest.runFakeHttpServer(_callDeleteRouter, router, delResp, _buildMockCloudRouterDeleteResp(), "cloud-router-delete-test", t)
}

func _callCreateRouter(payload interface{}) (interface{}, error) {
	return cTest.CreateCloudRouter(payload.(CloudRouter))
}

func _callUpdateRouter(payload interface{}) (interface{}, error) {
	return cTest.UpdateCloudRouter(payload.(CloudRouterUpdate), clrExpectedResp.CircuitID)
}

func _callReadRouter(payload interface{}) (interface{}, error) {
	return cTest.ReadCloudRouter(payload.(CloudRouter).AccountUUID)
}

func _callListCloudRouters(payload interface{}) (interface{}, error) {
	return cTest.ListCloudRouters()
}

func _callDeleteRouter(payload interface{}) (interface{}, error) {
	return cTest.DeleteCloudRouter(payload.(CloudRouter).AccountUUID)
}

func _buildMockCloudRouterResp() []byte {
	return []byte(`{
		"circuit_id": "PF-L3-CUST-2",
		"scope": "private",
		"asn": 4556,
		"name": "Super Cool Cloud router",
		"capacity": "10Gbps",
		"regions": [
		  {
			"name": "Continental U.S.",
			"code": "US"
		  }
		],
		"time_created": "2020-08-20T22:08:37.000000+0000",
		"time_updated": "2020-08-20T22:08:37.000000+0000"
	  }`)
}

func _buildMockCloudRouterResps() []byte {
	return []byte(`[{
		"circuit_id": "PF-L3-CUST-2",
		"scope": "private",
		"asn": 4556,
		"name": "Super Cool Cloud router",
		"capacity": "10Gbps",
		"regions": [
		  {
			"name": "Continental U.S.",
			"code": "US"
		  }
		],
		"time_created": "2020-08-20T22:08:37.000000+0000",
		"time_updated": "2020-08-20T22:08:37.000000+0000"
	  }]`)
}

func _buildMockCloudRouterUpdateResp() []byte {
	return []byte(`{
		"circuit_id": "PF-L3-CUST-2",
		"scope": "private",
		"asn": 4556,
		"name": "Super Cool Cloud router",
		"capacity": "10Gbps",
		"regions": [
		  {
			"name": "Continental U.S.",
			"code": "US"
		  }
		],
		"time_created": "2020-08-20T22:08:37.000000+0000",
		"time_updated": "2020-08-20T22:08:37.000000+0000"
	  }`)
}

func _buildMockCloudRouterDeleteResp() []byte {
	return []byte(`{
		"message": "Cloud router deleted."
	  }`)
}

func (cTest *PFClient) runFakeHttpServer(fn func(pl interface{}) (interface{}, error), payload interface{}, expectedResp interface{}, mockRespStr []byte, testName string, t *testing.T) {
	testServers := []struct {
		name         string
		server       *httptest.Server
		expectedResp interface{}
		expectedErr  error
	}{
		{
			name: testName,
			server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				_, err := w.Write(mockRespStr)
				if err != nil {
					t.Errorf("Error: %v", err)
				}
			})),
			expectedResp: expectedResp,
			expectedErr:  nil,
		},
	}
	for _, tes := range testServers {
		t.Run(tes.name, func(t *testing.T) {
			defer tes.server.Close()
			cTest.HostURL = tes.server.URL
			resp, err := fn(payload)
			respJSON, _ := json.Marshal(resp)
			expectJSON, _ := json.Marshal(tes.expectedResp)
			if !reflect.DeepEqual(respJSON, expectJSON) {
				t.Errorf("Expected: (%v), but got (%v)", string(expectJSON), string(respJSON))
			}
			if !errors.Is(err, tes.expectedErr) {
				t.Errorf("Expected err: (%v), but got: (%v)", tes.expectedErr, err)
			}
		})
	}
}
