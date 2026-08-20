package unifi

// jfb fork: DevicePortOverrides.NATiveNetworkID must serialize an explicit
// empty string (no native VLAN) and omit only when nil — the whole reason
// this fork exists (console rejects trunk PUTs otherwise with
// api.err.NativeVlanCannotBeExcluded).

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestNativeNetworkIDEmptyStringSerialized(t *testing.T) {
	empty := ""
	one := int64(1)
	b, err := json.Marshal(DevicePortOverrides{PortIDX: &one, NATiveNetworkID: &empty})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(b), `"native_networkconf_id":""`) {
		t.Fatalf("empty native id must be serialized explicitly, got: %s", b)
	}
}

func TestNativeNetworkIDNilOmitted(t *testing.T) {
	one := int64(1)
	b, err := json.Marshal(DevicePortOverrides{PortIDX: &one})
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(b), "native_networkconf_id") {
		t.Fatalf("nil native id must be omitted, got: %s", b)
	}
}

func TestNewerPortFieldsPointerSemantics(t *testing.T) {
	one := int64(1)
	yes := true
	edge := "disabled"
	b, err := json.Marshal(DevicePortOverrides{PortIDX: &one, StpBpduGuardEnabled: &yes, StpEdgeState: &edge})
	if err != nil {
		t.Fatal(err)
	}
	s := string(b)
	if !strings.Contains(s, `"stp_bpdu_guard_enabled":true`) || !strings.Contains(s, `"stp_edge_state":"disabled"`) {
		t.Fatalf("configured newer fields must serialize, got: %s", s)
	}
	b2, _ := json.Marshal(DevicePortOverrides{PortIDX: &one})
	for _, k := range []string{"eee_enabled", "link_debounce_auto", "multicast_router_mode", "sd_wan_underlay_port", "stp_bpdu_guard_enabled", "stp_edge_state", "stp_uplink"} {
		if strings.Contains(string(b2), k) {
			t.Fatalf("unset newer field %s must be omitted, got: %s", k, b2)
		}
	}
}
