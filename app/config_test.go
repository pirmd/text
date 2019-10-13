package app

import (
	"encoding/json"
	"testing"

	"github.com/pirmd/verify"
)

type TestStruct struct {
	TestSucceed bool `json:"succeed"`
}

func TestConfigUnmarshalling(t *testing.T) {
	testCases := []struct {
		in   string
		want TestStruct
	}{
		{in: `{ "succeed": true }`, want: TestStruct{TestSucceed: true}},
	}

	cfg := TestStruct{}

	cmdCfg := Config{
		Unmarshaller: json.Unmarshal,
		Var:          &cfg,
	}

	for _, tc := range testCases {
		if err := cmdCfg.load([]byte(tc.in)); err != nil {
			t.Errorf("cannot read config '%s': %s", tc.in, err)
		}

		verify.Equal(t, cfg, tc.want, "reading config for %s failed", tc.in)
	}
}
