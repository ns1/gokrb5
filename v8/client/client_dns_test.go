package client

import (
	"encoding/hex"
	"testing"

	"github.com/ns1/gokrb5/v8/config"
	"github.com/ns1/gokrb5/v8/keytab"
	"github.com/ns1/gokrb5/v8/test"
	"github.com/ns1/gokrb5/v8/test/testdata"
)

func TestClient_Login_DNSKDCs(t *testing.T) {
	test.Privileged(t)

	//ns := os.Getenv("DNSUTILS_OVERRIDE_NS")
	//if ns == "" {
	//	os.Setenv("DNSUTILS_OVERRIDE_NS", testdata.TEST_NS)
	//}
	c, _ := config.NewFromString(testdata.TEST_KRB5CONF)
	// Set to lookup KDCs in DNS
	c.LibDefaults.DNSLookupKDC = true
	//Blank out the KDCs to ensure they are not being used
	c.Realms = []config.Realm{}

	b, _ := hex.DecodeString(testdata.TESTUSER1_KEYTAB)
	kt := keytab.New()
	kt.Unmarshal(b)
	cl := NewWithKeytab("testuser1", "TEST.GOKRB5", kt, c)

	err := cl.Login()
	if err != nil {
		t.Errorf("error on logging in using DNS lookup of KDCs: %v\n", err)
	}
}
