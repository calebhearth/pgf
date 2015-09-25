package pgp

import (
	"strings"

	"golang.org/x/crypto/openpgp"
)

func SignersFromKeyring(keyring string) ([]byte, error) {
	var (
		connectionSet = make(map[string]set)
		keyIDToName   = make(map[string]string)
	)

	ring, err := openpgp.ReadArmoredKeyRing(strings.NewReader(keyring))
	if err != nil {
		return nil, err
	}

	for _, entity := range ring {
		for _, ident := range entity.Identities {
			keyID := entity.PrimaryKey.KeyIdString()
			name := ident.UserId.Name
			keyIDToName[keyID] = ident.Name
			connectionSet[name] = set{}
			for _, sig := range ident.Signatures {
				if keys := ring.KeysById(*sig.IssuerKeyId); len(keys) > 0 {
					for _, key := range keys {
						connectionSet[name].Add(key.PublicKey.KeyIdString())
					}
				}
			}
		}
	}

	return newSigned(connectionSet, keyIDToName).html(), nil
}
