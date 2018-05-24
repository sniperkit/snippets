package transip

import (
	"encoding/hex"
	"testing"
)

func TestAuthSha512Asn1(t *testing.T) {
	dataInput := "0=xor-gate.org&__method=getInfo&__service=DomainService&__hostname=api.transip.nl&__timestamp=1485612088&__nonce=588ca438c0b565.23541991"
	expectHex := "3051300d06096086480165030402030500044035eabe93b520cdab65a735e9f1664da67d18f2297223c4ed03beb0a97d405547d136e3b9602f432207fe644626bd317cf695818159f1c12a95a78c7b483060a7"

	outData := authSha512Asn1([]byte(dataInput))
	outputHex := hex.EncodeToString(outData)
	if outputHex != expectHex {
		t.Errorf("expected \"%s\", actual \"%s\"", expectHex, outputHex)
	}
}

func TestAuthSign(t *testing.T) {

}
