package main

import (
	"bytes"
	"strings"
	"text/template"

	"golang.org/x/crypto/openpgp"
)

type signed map[string][]string

func (s signed) HTML() []byte {
	t := template.Must(template.New("signed").Parse(`
{{define "T"}}
<ul>
{{range $name, $signatures := .}}
	<li>
		{{$name}}
		<ul>
		{{range $idx, $sig := $signatures}}
		<li>{{$sig}}</li>
		{{end}}
		</ul>
	</li>
{{end}}
</ul>
{{end}}
`))
	out := bytes.Buffer{}
	err := t.ExecuteTemplate(&out, "T", s)
	if err != nil {
		return []byte(err.Error())
	}
	return out.Bytes()
}

func signersFromKeyring(keyring string) ([]byte, error) {
	var (
		connectionSet = make(map[string]map[string]struct{})
		keyIdToName   = make(map[string]string)
	)

	ring, err := openpgp.ReadArmoredKeyRing(strings.NewReader(keyring))
	if err != nil {
		return nil, err
	}

	for _, entity := range ring {
		for _, ident := range entity.Identities {
			keyID := entity.PrimaryKey.KeyIdString()
			name := ident.UserId.Name
			keyIdToName[keyID] = ident.Name
			connectionSet[name] = make(map[string]struct{})
			for _, sig := range ident.Signatures {
				if keys := ring.KeysById(*sig.IssuerKeyId); len(keys) > 0 {
					for _, key := range keys {
						connectionSet[name][key.PublicKey.KeyIdString()] = struct{}{}
					}
				}
			}
		}
	}

	signed := signed{}
	for name, sigs := range connectionSet {
		signers := []string{}
		for keyId, _ := range sigs {
			signers = append(signers, keyIdToName[keyId])
		}
		signed[name] = signers
	}
	return signed.HTML(), nil
}
