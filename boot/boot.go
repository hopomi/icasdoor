package boot

import (
	"icasdoor/services/rsa"
	"os"
)

func genDefaultRsa() {
	var err error
	pri, pub := rsa.GenRsa()
	err = os.MkdirAll("files/rsa", 0777)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("files/rsa/rsa", pri, 0777)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("files/rsa/rsa.pub", pub, 0777)
	if err != nil {
		panic(err)
	}
}

func Boot() {
	// genDefaultRsa()
}
