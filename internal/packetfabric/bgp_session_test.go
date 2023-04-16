package packetfabric

import (
	"encoding/json"
	"fmt"
	"testing"
)

const _bgpSettingsUUID = "da53d96c-7783-4a6f-a171-a289bbf6763b"
const _bgpRemoteAddress = "10.0.0.1"
const _cID = "PF-L3-CUST-1730653"
const _connID = "PF-CC-LAX-LAX-1730655-PF"
const _bgpMd5 = "daffed346e29c5654f54133d1fc65ccb"
const _bgpRemoteAsn = 4555
const _bgpL3Prefix = "10.0.0.1/30"
const _bgpPublicIP = "185.161.1.152/31"

var _bgpPrefixOut = "10.0.0.2/32"
var _bgpPrefixIn = "10.0.0.1/32"

var _bgpSessionSettings = make([]BgpSessionAssociatedResp, 0)

var _createdTime = "2019-08-24T14:15:22Z"
var _updatedTime = "2019-08-24T14:15:22Z"

func init() {
	_bgpSessionSettings = append(_bgpSessionSettings, BgpSessionAssociatedResp{
		BgpSettingsUUID: _bgpSettingsUUID,
		AddressFamily:   "v4",
		RemoteAddress:   "10.0.0.1",
		RemoteAsn:       4556,
		MultihopTTL:     1,
		LocalPreference: 1,
		Community:       "1",
		AsPrepend:       1,
		Med:             1,
		Orlonger:        true,
		BfdInterval:     300,
		BfdMultiplier:   3,
		Disabled:        false,
		TimeCreated:     _createdTime,
		TimeUpdated:     _updatedTime,
		Prefixes: []BgpPrefix{
			{
				Prefix: _bgpPrefixOut,
				Type:   "out",
				Med:    2,
			},
			{
				Prefix:          _bgpPrefixIn,
				Type:            "in",
				LocalPreference: 100,
			},
		},
	})
}

func Test_CreateBgpSession(t *testing.T) {
	expectedPayload := BgpSession{}
	expectedResp := BgpSessionCreateResp{}
	if err := json.Unmarshal(_buildBgpSessionPayload(), &expectedPayload); err != nil {
		t.Fatalf("Failed to unmarshal BgpSession: %s", err)
	}
	if err := json.Unmarshal(_buildBgpSessionCreateResp(), &expectedResp); err != nil {
		t.Fatalf("Failed to unmarshal BgpSessionCreateResp: %s", err)
	}
	cTest.runFakeHttpServer(_callCreateBgpSession, expectedPayload, expectedResp, _buildBgpSessionCreateResp(), "bgp-session-create", t)
}

func Test_ListBgpSessionSettings(t *testing.T) {
	cTest.runFakeHttpServer(_callListBgpSessionSettings, nil, _bgpSessionSettings, _buildBgpSessionSettings(_createdTime, _updatedTime, _bgpSettingsUUID), "test-list-bgp-settings", t)
}

func _callCreateBgpSession(payload interface{}) (interface{}, error) {
	return cTest.CreateBgpSession(payload.(BgpSession), _cID, _connID)
}

func _callListBgpSessionSettings(payload interface{}) (interface{}, error) {
	return cTest.ListBgpSessions(_cID, _connID)
}

func _buildBgpSessionPayload() []byte {
	return []byte(fmt.Sprintf(`{
		"md5": "%s",
		"l3_address": "%s",
		"address_family": "v4",
		"remote_address": "%s",
		"remote_asn": %v,
		"prefixes": [
			{
				"prefix": "%s",
				"type": "out",
				"med": 2
			},
			{
				"prefix": "%s",
				"type": "in",
				"local_preference": 100
			}
		]
	}`, _bgpMd5, _bgpL3Prefix, _bgpRemoteAddress, _bgpRemoteAsn, _bgpPrefixOut, _bgpPrefixIn))
}

func _buildBgpSessionCreateResp() []byte {
	return []byte(fmt.Sprintf(`{
		"bgp_settings_uuid": "%s",
		"address_family": "v4",
		"remote_address": "%s",
		"remote_asn": %v,
		"local_preference": null,
		"md5": "%s",
		"med": null,
		"as_prepend": null,
		"orlonger": null,
		"bfd_interval": null,
		"bfd_multiplier": null,
		"disabled": false,
		"time_updated": "%s",
		"time_created": "%s",
		"bgp_state": "Configuring",
		"prefixes": [
			{
				"prefix": "%s",
				"type": "out",
				"med": 2
			},
			{
				"prefix": "%s",
				"type": "in",
				"local_preference": 100
			}
		],
		"subnet": null,
		"public_ip": "%s",
		"nat": null
	}`, _bgpSettingsUUID, _bgpRemoteAddress, _bgpRemoteAsn, _bgpMd5, _createdTime, _updatedTime, _bgpPrefixOut, _bgpPrefixIn, _bgpPublicIP))
}

func _buildBgpSessionSettings(timeCreated, timeUpdated, bgpSettingsUUID string) []byte {
	return []byte(fmt.Sprintf(`[
		{
		  "bgp_settings_uuid": "%s",
		  "address_family": "v4",
		  "remote_address": "%s",
		  "remote_asn": 4556,
		  "multihop_ttl": 1,
		  "local_preference": 1,
		  "as_prepend": 1,
		  "med": 1,
		  "orlonger": true,
		  "bfd_interval": 300,
		  "bfd_multiplier": 3,
		  "disabled": false,
		  "time_created": "%s",
		  "time_updated": "%s",
		  "prefixes": [
			{
				"prefix": "%s",
				"type": "out",
				"med": 2
			},
			{
				"prefix": "%s",
				"type": "in",
				"local_preference": 100
			}
		  ]
		}
	  ]`, bgpSettingsUUID, _bgpRemoteAddress, timeCreated, timeUpdated, _bgpPrefixOut, _bgpPrefixIn))
}
