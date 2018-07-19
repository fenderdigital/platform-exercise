package secret

import "log"

func Fetch() string {
	var signing_key string = ""

	if signing_key == "" {
		log.Fatal("Please change the signing key directly in the file platform-test/src/secret/secret.go method Fetch, updating line 4")
	}

	return signing_key
}
