package main

import (
	"encoding/json"
	"fmt"
	"os"

	"golang.org/x/crypto/openpgp"
)

func getKeyByEmail(keyring openpgp.EntityList, email string) *openpgp.Entity {
	for _, entity := range keyring {
		for _, ident := range entity.Identities {
			if ident.UserId.Email == email {
				return entity
			}
		}
	}

	return nil
}

func main() {
	privringFile, _ := os.Open("/Users/caleb/.gnupg/secring.gpg")
	privring, err := openpgp.ReadKeyRing(privringFile)
	if err != nil {
		panic(err)
	}

	var connectionSet map[string]struct{}
	key := getKeyByEmail(privring, "caleb@calebthompson.io")

	for _, ident := range key.Identities {
		fmt.Println(ident.Name)
		fmt.Println(len(ident.Signatures))
		for _, sig := range ident.Signatures {
			keyID := sig.IssuerKeyId
			fmt.Println(keyID)
			connectionSet[fmt.Sprint(keyID)] = struct{}{}
		}
	}

	b, err := json.MarshalIndent(connectionSet, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Print(string(b))
}
