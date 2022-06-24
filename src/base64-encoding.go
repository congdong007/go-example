package main

import (
	b64 "encoding/base64"
	"fmt"
)

func main() {
	pl := fmt.Println
	data := "this is base64 encoding example"

	sEnc := b64.StdEncoding.EncodeToString([]byte(data))

	pl(sEnc)

	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	pl(string(sDec))
	pl()

	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	pl(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	pl(string(uDec))

}
