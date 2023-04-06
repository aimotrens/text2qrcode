package healthcheck_test

import (
	"testing"

	"github.com/aimotrens/text2qrcode/testing/test_utils"
)

func TestHealthcheck_OK(t *testing.T) {
	test_utils.Get200("/api/healthcheck/", t)
}
