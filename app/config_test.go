package app

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/pirmd/verify"
)

type TestStruct struct {
	Success bool   `json:"success"`
	Gobin   string `json:"gobin"`
}

func TestConfigUnmarshalling(t *testing.T) {
	testCases := []struct {
		in   string
		want TestStruct
	}{
		{in: `{ "success": true }`, want: TestStruct{Success: true}},
		{in: `{ "gobin": "$GOBIN" }`, want: TestStruct{Gobin: os.Getenv("GOBIN")}},
	}

	for _, tc := range testCases {
		cfg := TestStruct{}

		cmdCfg := Config{
			Unmarshaller: json.Unmarshal,
			Var:          &cfg,
		}

		if err := cmdCfg.load([]byte(tc.in)); err != nil {
			t.Errorf("cannot read config '%s': %s", tc.in, err)
		}

		verify.Equal(t, cfg, tc.want, "reading config for %s failed", tc.in)
	}
}
