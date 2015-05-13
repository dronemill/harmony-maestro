package main

import (
	"reflect"
	"testing"
)

func TestInitConfigDefaultConfig(t *testing.T) {
	LogSetLevel("warning")
	want := Config{
		LogLevel: "info",
		Harmony: HarmonyConfig{
			API:       "http://harmony.dev:4774",
			VerifySSL: true,
		},
		Eventsocket: EventsocketConfig{
			Port: 4775,
		},
	}
	if err := initConfig(); err != nil {
		t.Errorf(err.Error())
	}
	if !reflect.DeepEqual(want, config) {
		t.Errorf("initConfig() = %v, want %v", config, want)
	}
}
