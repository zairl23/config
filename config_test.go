package config

import (
	"fmt"
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

func TestViperManager(t *testing.T) {
	err := Manager.Add("default", "")

	if err != nil {
		panic(err)
	}

	err = Manager.Add("foo", "config.yaml")

	if err != nil {
		panic(err)
	}

	v1 := Manager.Get("default")

	if v1 != nil {
		fmt.Println(v1.Get("foo"))
	}

	v2 := Manager.Get("foo")

	if v2 != nil {
		fmt.Println(v2.Get("name"))
	}
}
