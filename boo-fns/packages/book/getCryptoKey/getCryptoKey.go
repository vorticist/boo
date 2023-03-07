package main

import "os"

func Main(args map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{
		"key": os.Getenv("CRYPTO_KEY"),
	}
}
