package caesargo_test

import (
	"testing"

	"github.com/aZ4ziL/caesargo"
)

func TestEncryptText(t *testing.T) {
	result := caesargo.EncryptText("Hello World", 23)
	expected := "Ebiil Tloia"

	if result != expected {
		t.Errorf("Function EncryptText returned incorrect result, got: %s, want: %s", result, expected)
	}
}

func TestDecryptText(t *testing.T) {
	result := caesargo.DecryptText("Ebiil Tloia", 23)
	expected := "Hello World"

	if result != expected {
		t.Errorf("Function DecryptText returned incorrect result, got: %s, want: %s", result, expected)
	}
}
