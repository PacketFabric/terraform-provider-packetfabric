package packetfabric

import (
	"encoding/json"
	"fmt"
	"testing"
)

const _bgpPrefixUUID = "3d78949f-1396-4163-b0ca-3eba3592efef"
const _bgpSettingsUUID = "da53d96c-7783-4a6f-a171-a289bbf6763b"
const _bgpRemoteAddress = "10.0.0.1"
const _cID = "PF-L3-CUST-1730653"
const _connID = "PF-CC-LAX-LAX-1730655-PF"
const _bgpSessionUUID = "32e0be91-7c92-4cad-95da-9001be89c917"
const _bgpMd5 = "daffed346e29c5654f54133d1fc65ccb"
const _bgpRemoteAsn = 4555
const _bgpPrefixInUUID = "434ab7c4-7e55-423c-a897-8a40b6ccf215"
const _bgpPrefixOutUUID = "ffc86e74-8c03-49ea-bf87-f3919d7c9d0a"
const _bgpL3Prefix = "10.0.0.1/30"
const _bgpPublicIP = "185.161.1.152/31"

var _bgpPrefixOut = "10.0.0.2/32"
var _bgpPrefixIn = "10.0.0.1/32"

var _bgpSessionPrefixes = make([]BgpSessionResponse, 0)
var _bgpSessionSettings = make([]BgpSessionAssociatedResp, 0)

var _createdTime = "2019-08-24T14:15:22Z"
var _updatedTime = "2019-08-24T14:15:22Z"

func init() {
	_bgpSessionPrefixes = append(_bgpSessionPrefixes, BgpSessionResponse{
		BgpPrefixUUID: _bgpPrefixOutUUID,
		Prefix:        _bgpPrefixOut,
		Type:          "out",
		Order:         1,
	})
	_bgpSessionPrefixes = append(_bgpSessionPrefixes, BgpSessionResponse{
		BgpPrefixUUID: _bgpPrefixInUUID,
		Prefix:        _bgpPrefixIn,
		Type:          "in",
		Order:         1,
	})
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
	})
}

func Test_CreateBgpSession(t *testing.T) {
	expectedPayload := BgpSession{}
	expectedResp := BgpSessionCreateResp{}
	_ = json.Unmarshal(_buildBgpSessionPayload(), &expectedPayload)
	_ = json.Unmarshal(_buildBgpSessionCreateResp(), &expectedResp)
	cTest.runFakeHttpServer(_callCreateBgpSession, expectedPayload, expectedResp, _buildBgpSessionCreateResp(), "bgp-session-create", t)
}

func Test_CreateBgpSessionPrefixes(t *testing.T) {
	var payload []BgpPrefix
	var expectedResp []BgpSessionResponse
	_ = json.Unmarshal(_buildBgpPrefixList(), &payload)
	_ = json.Unmarshal(_buildCreateBgpSessionPrefixes(), &expectedResp)
	cTest.runFakeHttpServer(_callCreateBgpSessionPrefixes, payload, expectedResp, _buildCreateBgpSessionPrefixes(), "test-create-bgp-session-prefixes", t)
}

func Test_ReadBgpSession(t *testing.T) {
	cTest.runFakeHttpServer(_callReadBgpSession, _bgpPrefixUUID, _bgpSessionPrefixes, _buildBgpPrefixList(), "test-read-bgp-sessions", t)
}

func Test_ReadBgpSessionPrefixes(t *testing.T) {
	var expectedResp []BgpSessionResponse
	_ = json.Unmarshal(_buildBgpPrefixList(), &expectedResp)
	cTest.runFakeHttpServer(_callReadBgpSessionPrefixes, _bgpSettingsUUID, expectedResp, _buildBgpPrefixList(), "test-read-bgp-session-prefixes", t)
}

func Test_ListBgpSessionSettings(t *testing.T) {
	cTest.runFakeHttpServer(_callListBgpSessionSettings, nil, _bgpSessionSettings, _buildBgpSessionSettings(_createdTime, _updatedTime, _bgpSettingsUUID), "test-list-bgp-settings", t)
}

func _callCreateBgpSession(payload interface{}) (interface{}, error) {
	return cTest.CreateBgpSession(payload.(BgpSession), _cID, _connID)
}

func _callCreateBgpSessionPrefixes(payload interface{}) (interface{}, error) {
	return cTest.CreateBgpSessionPrefixes(payload.([]BgpPrefix), _bgpSessionUUID)
}

func _callReadBgpSession(payload interface{}) (interface{}, error) {
	return cTest.ReadBgpSession(payload.(string))
}

func _callReadBgpSessionPrefixes(payload interface{}) (interface{}, error) {
	return cTest.ReadBgpSessionPrefixes(payload.(string))
}

func _callListBgpSessionSettings(payload interface{}) (interface{}, error) {
	return cTest.ListBgpSessions()
}

func _buildBgpSessionPayload() []byte {
	return []byte(fmt.Sprintf(`{
		"md5": "%s",
		"l3_address": "%s",
		"address_family": "v4",
		"remote_address": "%s",
		"remote_asn": %v,
		"multihop_ttl": 1,
		"orlonger": false
	}`, _bgpMd5, _bgpL3Prefix, _bgpRemoteAddress, _bgpRemoteAsn))
}

func _buildBgpSessionCreateResp() []byte {
	return []byte(fmt.Sprintf(`{
		"bgp_settings_uuid": "%s",
		"address_family": "v4",
		"remote_address": "%s",
		"remote_asn": %v,
		"multihop_ttl": 1,
		"local_preference": null,
		"md5": "%s",
		"med": null,
		"community": null,
		"as_prepend": null,
		"orlonger": false,
		"bfd_interval": null,
		"bfd_multiplier": null,
		"disabled": false,
		"time_updated": "%s",
		"time_created": "%s",
		"bgp_state": "Configuring",
		"subnet": null,
		"public_ip": "%s",
		"nat": null
	}`, _bgpSettingsUUID, _bgpRemoteAddress, _bgpRemoteAsn, _bgpMd5, _createdTime, _updatedTime, _bgpPublicIP))
}

func _buildBgpPrefixList() []byte {
	return []byte(fmt.Sprintf(`[
		{
			"bgp_prefix_uuid": "%s",
			"prefix": "%s",
			"type": "out",
			"order": 1
		},
		{
			"bgp_prefix_uuid": "%s",
			"prefix": "%s",
			"type": "in",
			"order": 1
		}
	]`, _bgpPrefixOutUUID, _bgpPrefixOut, _bgpPrefixInUUID, _bgpPrefixIn))
}

func _buildCreateBgpSessionPrefixes() []byte {
	return []byte(fmt.Sprintf(`[
		{
			"prefix": "%s",
			"type": "out",
			"order": 1
		},
		{
			"prefix": "%s",
			"type": "in",
			"order": 1
		}
	]`, _bgpPrefixOut, _bgpPrefixIn))
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
		  "community": "1",
		  "as_prepend": 1,
		  "med": 1,
		  "orlonger": true,
		  "bfd_interval": 300,
		  "bfd_multiplier": 3,
		  "disabled": false,
		  "time_created": "%s",
		  "time_updated": "%s"
		}
	  ]`, bgpSettingsUUID, _bgpRemoteAddress, timeCreated, timeUpdated))
}
