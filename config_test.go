package config

import (
	"testing"
)

func TestGet(t *testing.T) {
	// init config
	if err := Init("config.yaml"); err != nil {
		panic(err)
	}

	if Get("name") == "config" && Get("nokey") == "" {
		t.Log("test get pass")
		return
	}

	t.Error("test get fail")
	return
}
