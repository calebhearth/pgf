package pgp

import (
	"bytes"
	"text/template"
)

type signed map[string][]string

func (s signed) html() []byte {
	out := bytes.Buffer{}
	err := template.Must(template.ParseFiles("signers.html")).ExecuteTemplate(&out, "T", s)
	if err != nil {
		panic(err.Error())
	}
	return out.Bytes()
}

func newSigned(connections map[string]set, keyIDToName map[string]string) signed {
	signed := signed{}
	for name, sigs := range connections {
		signers := sigs.Reduce(func(keyID string) string {
			return keyIDToName[keyID]
		})
		signed[name] = signers
	}
	return signed
}
