package internal_test

import (
	"testing"

	"github.com/aimotrens/text2qrcode/testing/test_utils"
)

func TestRoot_Redirect(t *testing.T) {
	test_utils.Get307("/", t)
}

func TestSwagger_OK(t *testing.T) {
	test_utils.Get200("/swagger/index.html", t)
}
