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
