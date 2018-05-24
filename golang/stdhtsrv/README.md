# Syncthing DHT POC

Based on https://github.com/syncthing/syncthing/issues/3388

## Tools

`openssl x509 -in cert.pem -text`
`openssl x509 -pubkey -noout -in cert.pem`
`openssl ec -in key.pem -text`

## Resources

* <https://kjur.github.io/jsrsasign/sample/sample-ecdsa.html>
* <https://raymii.org/s/tutorials/Sign_and_verify_text_files_to_public_keys_via_the_OpenSSL_Command_Line.html>
* <https://github.com/gtank/cryptopasta>
* <https://github.com/ethereum/go-ethereum/blob/ba975dc0931b9f2962b2f163675772458ed339fd/crypto/crypto.go>
* <https://github.com/warner/python-ecdsa>
* <https://bitcoin.stackexchange.com/questions/2376/ecdsa-r-s-encoding-as-a-signature>
* <https://tools.ietf.org/html/rfc7518#section-3.3>
