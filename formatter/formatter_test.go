package formatter

import (
	"errors"
	"testing"
	"verify"
)

type testStruct struct {
	PropA string
	PropB string
}

func (t *testStruct) Type() string {
	return "test"
}

func TestJSONFormatter(t *testing.T) {
	expected := "{\n  \"PropA\": \"testa\",\n  \"PropB\": \"testb\"\n}"

	got, err := JSONFormatter(testStruct{"testa", "testb"})
	if err != nil {
		t.Fatalf("Fail to format interface: %v", err)
	}
	verify.EqualString(t, got, expected, "JSONFormatter does not work as expected.")
}

func TestTemplateFormatter(t *testing.T) {
	t.Run("Test simple", func(t *testing.T) {
		expected := "test<testA, testB>"
		formatter := TemplateNewFormatter("{{.Type}}<{{.PropA}}, {{.PropB}}>")
		got, err := formatter(&testStruct{"testA", "testB"})
		if err != nil {
			t.Fatalf("Fail to format interface: %v", err)
		}
		if got != expected {
			t.Errorf("TemplateFormatter does not work as expected.\nGot    : %s\n Expected: %s", got, expected)
		}
	})

	t.Run("Test panic on malformed template", func(t *testing.T) {
		verify.Panic(t, func() { _ = TemplateNewFormatter("{{.Type}") }, "bad formatter did not panic")
	})
}

func TestMustFormat(t *testing.T) {
	pprint := Formatters{
		DefaultFormatter: JSONFormatter,
	}

	t.Run("Test default", func(t *testing.T) {
		expected := "{\n  \"PropA\": \"testa\",\n  \"PropB\": \"testb\"\n}"
		got := pprint.MustFormat(&testStruct{"testa", "testb"})
		verify.EqualString(t, got, expected, "MustFormat does not work as expected.")
	})

	t.Run("Test with type", func(t *testing.T) {
		pprint["test"] = TemplateNewFormatter("{{.Type}}<{{.PropA}}, {{.PropB}}>")

		expected := "test<testA, testB>"
		got := pprint.MustFormat(&testStruct{"testA", "testB"})
		if got != expected {
			t.Errorf("MustFormat does not work as expected.\nGot    : %s\n Expected: %s", got, expected)
		}
	})

	t.Run("Fallback without default", func(t *testing.T) {
		pprint = Formatters{}

		expected := "&{PropA:testA PropB:testB}"
		got := pprint.MustFormat(&testStruct{"testA", "testB"})
		if got != expected {
			t.Errorf("MustFormat does not work as expected.\nGot    : %s\n Expected: %s", got, expected)
		}
	})

	t.Run("Fallback due to Formatter error", func(t *testing.T) {
		expected := "!Err(mock error)"
		pprint["test"] = func(v interface{}) (string, error) { return "", errors.New("mock error") }

		got := pprint.MustFormat(&testStruct{"testA", "testB"})
		if got != expected {
			t.Errorf("MustFormat does not work as expected.\nGot     : %s\nExpected: %s", got, expected)
		}
	})
}
