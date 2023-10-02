package data

import (
	"testing"
)

func TestUrlWithWrongUrl(t *testing.T) {
	input := ShortUrl{
		Url:  "badformatted.com",
		Code: "abc123",
	}

	ok, errors := input.IsValid()

	if ok {
		t.Errorf("Expected %t, but got %t", false, ok)
	}
	if len(errors) != 1 {
		t.Errorf("Expected %d, but got %d", 1, len(errors))
	}
	if errors[0].Tag != "url" {
		t.Errorf("Expected %s, but got %s", "url", errors[0].Tag)
	}
}

func TestUrlWithWrongCode(t *testing.T) {
	input := ShortUrl{
		Url:  "https://goodurl.com.br",
		Code: "a1i2",
	}

	ok, errors := input.IsValid()

	if ok {
		t.Errorf("Expected %t, but got %t", false, ok)
	}
	if len(errors) != 1 {
		t.Errorf("Expected %d, but got %d", 1, len(errors))
	}
	if errors[0].Tag != "len" {
		t.Errorf("Expected %s, but got %s", "len", errors[0].Tag)
	}
}

func TestUrlRequired(t *testing.T) {
	input := ShortUrl{
		Url:  "",
		Code: "123qwe",
	}
	ok, errors := input.IsValid()

	if ok {
		t.Errorf("Expected %t, but got %t", false, ok)
	}
	if len(errors) != 1 {
		t.Errorf("Expected %d, but got %d", 1, len(errors))
	}
	if errors[0].Tag != "required" {
		t.Errorf("Expected %s, but got %s", "required", errors[0].Tag)
	}
}

func TestCodeRequired(t *testing.T) {
	input := ShortUrl{
		Url:  "https://goodurl.com.br",
		Code: "",
	}
	ok, errors := input.IsValid()

	if ok {
		t.Errorf("Expected %t, but got %t", false, ok)
	}
	if len(errors) != 1 {
		t.Errorf("Expected %d, but got %d", 1, len(errors))
	}
	if errors[0].Tag != "required" {
		t.Errorf("Expected %s, but got %s", "required", errors[0].Tag)
	}
}

func TestWithNoError(t *testing.T) {
	input := ShortUrl{
		Url:  "https://goodurl.com.br",
		Code: "aBc123",
	}
	ok, errors := input.IsValid()

	if !ok {
		t.Errorf("Expected %t, but got %t", false, ok)
	}
	if len(errors) != 0 {
		t.Errorf("Expected %d, but got %d", 1, len(errors))
	}

}
