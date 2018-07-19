package secret

import "log"

func Fetch() string {
	var signing_key string = "2hZ2cpjxFSz8sR2MbKqo7XLz4HS6Nx4tuBWlLpvIrXQPR5O36syvcefGZAdbZisog9LWPvDCYEJajl9X"

	if signing_key == "" {
		log.Fatal("Please change the signing key directly in the file platform-test/src/secret/secret.go method Fetch, updating line 4")
	}

	return signing_key
}
