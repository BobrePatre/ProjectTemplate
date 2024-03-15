package keycloak_redis

import (
	"github.com/lestrrat-go/jwx/jwk"
	"testing"
)

func TestSerialize(t *testing.T) {

	provider := Provider{}

	t.Run("Success serialize key set", func(t *testing.T) {

		keyRsa := jwk.NewRSAPublicKey()
		err := keyRsa.Set("alg", "RS256")
		if err != nil {
			t.Fatal("Failed to set key", err)
		}

		keySet := jwk.NewSet()
		keySet.Add(keyRsa)

		serializedKey, err := provider.SerializeJwkSet(keySet)
		t.Log(serializedKey)
		if err != nil {
			t.Error(err)
		}
		if serializedKey == "" {
			t.Error("serializedKey is empty")
		}
	})

}

func TestDeserialize(t *testing.T) {
	provider := Provider{}

	t.Run("Success deserialize key set", func(t *testing.T) {
		serializedKeySet := "{\"keys\":[{\"kid\":\"0b82GdNVQLQsINMkSC0CbhEJj9NQj62RmHY05-rS2ec\",\"kty\":\"RSA\",\"alg\":\"RS256\",\"use\":\"sig\",\"n\":\"qE1YRFh6WC9FuJu5s7_4vvZmoQfJ4Yd6NOX7TsT3il2SFbTHw5mQ0ePetDovim14sz6Ebox3BpuP1QposcqQAgAH3izjlzj0QIvAWisYscNOgPDTM9uDfoBbLPLn-ZZRiHdUq5sAp01Tz-xSIsAReXVhLjY62fjNXk52U1Rwe7K4ynQFFjlTZtgW-tdy-eEhsfRvajvenR4uuuMrRkHoe6Q0D-OvvTnQ6oOIIjf3svDezEujjuwfKwDLbQ1m_1p0uVGM5Y2ZfyOybVWXx8raPf9TGCAyD_VbbYJfa7tBYjMx81geCYE91hOJySopL3ZHPWjdUecT8HrZaEpBas9_uw\",\"e\":\"AQAB\",\"x5c\":[\"MIICqzCCAZMCBgGONgX92zANBgkqhkiG9w0BAQsFADAZMRcwFQYDVQQDDA5jaGFuZ2VpdC1yZWFsbTAeFw0yNDAzMTMwNDE0MDZaFw0zNDAzMTMwNDE1NDZaMBkxFzAVBgNVBAMMDmNoYW5nZWl0LXJlYWxtMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAqE1YRFh6WC9FuJu5s7/4vvZmoQfJ4Yd6NOX7TsT3il2SFbTHw5mQ0ePetDovim14sz6Ebox3BpuP1QposcqQAgAH3izjlzj0QIvAWisYscNOgPDTM9uDfoBbLPLn+ZZRiHdUq5sAp01Tz+xSIsAReXVhLjY62fjNXk52U1Rwe7K4ynQFFjlTZtgW+tdy+eEhsfRvajvenR4uuuMrRkHoe6Q0D+OvvTnQ6oOIIjf3svDezEujjuwfKwDLbQ1m/1p0uVGM5Y2ZfyOybVWXx8raPf9TGCAyD/VbbYJfa7tBYjMx81geCYE91hOJySopL3ZHPWjdUecT8HrZaEpBas9/uwIDAQABMA0GCSqGSIb3DQEBCwUAA4IBAQAs9ZV00HqLlzz6sMd079DwVEuaS3l5tXbGrmOEy/aCWxcgTtYP1VXVTL1mJHAFtNaIbT3NRvS/rc6+MoO/1ZY7GTRjl5I4bckQ/Tf6nLIRL5+pszodgzdlf6kxEchzfuLigWxVr18H6nZhZxxefg3nQp5vvv1h/MDsPSk6j7Wpgr/6yq77+Rs/WoNccTXaD0DM+lwkHCPAv3deYg26Vwn7k+cKdE5cLQXILgfWSgwjsf1lhhHMRftzoHDixupErHA7yEZ2zmfnE0p/lNoi/PaOXTr68wLF3xX4CmSyvY/17QAYSfcee9Ljg52hYc3+rKbV3dyIkEFCmoDJ9ppBIhIe\"],\"x5t\":\"6QEpXh6psa8ZYZAc5wq8Q7P3fQM\",\"x5t#S256\":\"FDTVU70idOrLeAFpRdypqYipWGloNbXv3ZL7ynr77pI\"},{\"kid\":\"cAF3eYJ7crTbj600b9lNIUFE36RwrhorzKfvQpRu6Qg\",\"kty\":\"RSA\",\"alg\":\"RSA-OAEP\",\"use\":\"enc\",\"n\":\"p0XKmzywLgOYGBQJQmvDCKmAu3zP1OqhJvAGdFXR5J8JT0g7GsUxSJtVVYYqTzCEtOTOfdRMEUNIbPy5etSLIqQ2__CFIn_F4zus22SiY_BFhlKdTko7miMZQurIxrApO2Nz2z-fRV1iEdak0BrYl1zmEZgjouvao9iSlmyn-qWNTCkX9StUBOB2D_2heT93UyXl1a5ojHC6eXQqiX5jqS9ENimyjS75LBKlERSMi3ifx80AK_ZZOKBhrUA6Ty0xv4Xqc6ciNSNlZPFdwE2Oy8-EN5EDSAN1Ae6jTAYiAjWQSYgKm2_Po_p0E7YlU3N2RRKdhsutWdDQ4cRWfwKg-Q\",\"e\":\"AQAB\",\"x5c\":[\"MIICqzCCAZMCBgGONgX+nDANBgkqhkiG9w0BAQsFADAZMRcwFQYDVQQDDA5jaGFuZ2VpdC1yZWFsbTAeFw0yNDAzMTMwNDE0MDZaFw0zNDAzMTMwNDE1NDZaMBkxFzAVBgNVBAMMDmNoYW5nZWl0LXJlYWxtMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAp0XKmzywLgOYGBQJQmvDCKmAu3zP1OqhJvAGdFXR5J8JT0g7GsUxSJtVVYYqTzCEtOTOfdRMEUNIbPy5etSLIqQ2//CFIn/F4zus22SiY/BFhlKdTko7miMZQurIxrApO2Nz2z+fRV1iEdak0BrYl1zmEZgjouvao9iSlmyn+qWNTCkX9StUBOB2D/2heT93UyXl1a5ojHC6eXQqiX5jqS9ENimyjS75LBKlERSMi3ifx80AK/ZZOKBhrUA6Ty0xv4Xqc6ciNSNlZPFdwE2Oy8+EN5EDSAN1Ae6jTAYiAjWQSYgKm2/Po/p0E7YlU3N2RRKdhsutWdDQ4cRWfwKg+QIDAQABMA0GCSqGSIb3DQEBCwUAA4IBAQBRskV6jt2YQN7BLCXU/Mpz9+tSKyLarFnkKkHuFyawbWzMnXhrDFDWg8Y50Ah6P7wMvGp1G20hgsCjqIGggB0k+LrDeqnwMbfp0u/rVEVc6um6FNJ9mV4J96So1LJUvEI+yLhllRlIIga4iY/1bZzu/3aoF0IU0G4j0xXPbNpNbR8E7mch1HJ0pm4fW42iYYDFLOWbbYbt+5+T6YakBq280ND4Vq/0eQ3/pTVndtrsoNJWfANWfX9xhzRkQsz4e26gdko2ZX/IWsnc2yRb707BgzDqdqelBrtNhw6o/WvQ81rYRIooA2yv07JDJSZgqw4nnJZ+YrMQOFQb223kPNGP\"],\"x5t\":\"8rF-_NDqdytz7gm9zMCNvCAkH0E\",\"x5t#S256\":\"4u4lWxn1tNxPgh3sUmOYa2Q_wCS1xRAw-Rr__GhnoWw\"}]}"

		keySet, err := provider.DeserializeJwkSet(serializedKeySet)
		if err != nil {
			t.Log(err)
			t.Fatal(err)
		}
		key, _ := keySet.Get(0)
		t.Log(key.KeyID())
		t.Log(key.Algorithm())
	})
}
