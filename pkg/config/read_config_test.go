// License: OpenFaaS Community Edition (CE) EULA
// Copyright (c) 2017,2019-2024 OpenFaaS Author(s)

// Copyright (c) Alex Ellis 2017. All rights reserved.
// Copyright 2020 OpenFaaS Author(s)

package config

import (
	"testing"
)

type EnvBucket struct {
	Items map[string]string
}

func NewEnvBucket() EnvBucket {
	return EnvBucket{
		Items: make(map[string]string),
	}
}

func (e EnvBucket) Getenv(key string) string {
	return e.Items[key]
}

func (e EnvBucket) Setenv(key string, value string) {
	e.Items[key] = value
}
func TestRead_EmptyProbeConfig(t *testing.T) {
	defaults := NewEnvBucket()
	readConfig := ReadConfig{}
	config, err := readConfig.Read(defaults)
	if err != nil {
		t.Fatalf("Unexpected error while reading env %s", err.Error())
	}
	want := false
	if config.HTTPProbe != want {
		t.Logf("EnableFunctionReadinressProbe incorrect, want: %t, got: %t", want, config.HTTPProbe)
		t.Fail()
	}
}

func TestRead_HTTPProbeConfig(t *testing.T) {
	defaults := NewEnvBucket()
	defaults.Setenv("http_probe", "false")

	readConfig := ReadConfig{}
	config, err := readConfig.Read(defaults)
	if err != nil {
		t.Fatalf("Unexpected error while reading env %s", err.Error())
	}

	if config.HTTPProbe {
		t.Logf("HTTPProbe incorrect, got: %v\n", config.HTTPProbe)
		t.Fail()
	}
}

func TestRead_HTTPProbeConfig_true(t *testing.T) {
	defaults := NewEnvBucket()
	defaults.Setenv("http_probe", "true")

	readConfig := ReadConfig{}
	config, err := readConfig.Read(defaults)
	if err != nil {
		t.Fatalf("Unexpected error while reading env %s", err.Error())
	}

	if !config.HTTPProbe {
		t.Logf("HTTPProbe incorrect, got: %v\n", config.HTTPProbe)
		t.Fail()
	}
}
