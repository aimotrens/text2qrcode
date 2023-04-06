package text2qrcode_test

import (
	"testing"

	"github.com/aimotrens/text2qrcode/app/text2qrcode"
	"github.com/aimotrens/text2qrcode/testing/test_utils"
)

func TestEncodeWithQueryString_OK(t *testing.T) {
	test_utils.Get200("/api/text2qrcode/encode?text=HelloWorld", t)
}

func TestEncodeWithQueryString_SizeTooLow(t *testing.T) {
	test_utils.Get400("/api/text2qrcode/encode?text=HelloWorld&size=99", t)
}

func TestEncodeWithQueryString_SizeTooHigh(t *testing.T) {
	test_utils.Get400("/api/text2qrcode/encode?text=HelloWorld&size=1001", t)
}

func TestEncodeWithQueryString_EclTooLow(t *testing.T) {
	test_utils.Get400("/api/text2qrcode/encode?text=HelloWorld&errorCorrection=-1", t)
}

func TestEncodeWithQueryString_EclTooHigh(t *testing.T) {
	test_utils.Get400("/api/text2qrcode/encode?text=HelloWorld&errorCorrection=5", t)
}

func TestEncodeWithQueryString_SizeNotInteger(t *testing.T) {
	test_utils.Get400("/api/text2qrcode/encode?text=HelloWorld&size=abc", t)
}

func TestEncodeWithQueryString_EclNotInteger(t *testing.T) {
	test_utils.Get400("/api/text2qrcode/encode?text=HelloWorld&errorCorrection=abc", t)
}

func TestEncodeWithQueryString_EmptyText(t *testing.T) {
	test_utils.Get400("/api/text2qrcode/encode?text=", t)
}

func TestEncodeWithQueryString_NoText(t *testing.T) {
	test_utils.Get400("/api/text2qrcode/encode", t)
}

func TestEncodeWithQueryString_WhiteBorderNotBool(t *testing.T) {
	test_utils.Get400("/api/text2qrcode/encode?text=HelloWorld&whiteBorder=abc", t)
}

func TestEncode_OK(t *testing.T) {
	qrReq := text2qrcode.QRCodeRequest{
		Text:            "HelloWorld",
		ErrorCorrection: 1,
		Size:            250,
		WhiteBorder:     true,
	}
	test_utils.PostJson200(qrReq, t)
}

func TestEncode_SizeTooLow(t *testing.T) {
	qrReq := text2qrcode.QRCodeRequest{
		Text:            "HelloWorld",
		ErrorCorrection: 1,
		Size:            99,
		WhiteBorder:     true,
	}
	test_utils.PostJson400(qrReq, t)
}

func TestEncode_SizeTooHigh(t *testing.T) {
	qrReq := text2qrcode.QRCodeRequest{
		Text:            "HelloWorld",
		ErrorCorrection: 1,
		Size:            1001,
		WhiteBorder:     true,
	}
	test_utils.PostJson400(qrReq, t)
}

func TestEncode_EclTooLow(t *testing.T) {
	qrReq := text2qrcode.QRCodeRequest{
		Text:            "HelloWorld",
		ErrorCorrection: -1,
		Size:            250,
		WhiteBorder:     true,
	}
	test_utils.PostJson400(qrReq, t)
}

func TestEncode_EclTooHigh(t *testing.T) {
	qrReq := text2qrcode.QRCodeRequest{
		Text:            "HelloWorld",
		ErrorCorrection: 4,
		Size:            250,
		WhiteBorder:     true,
	}
	test_utils.PostJson400(qrReq, t)
}

func TestEncode_EmptyText(t *testing.T) {
	qrReq := text2qrcode.QRCodeRequest{
		Text:            "",
		ErrorCorrection: 1,
		Size:            250,
		WhiteBorder:     true,
	}
	test_utils.PostJson400(qrReq, t)
}

func TestEncode_InvalidJson(t *testing.T) {
	test_utils.PostJsonRaw400([]byte("invalid json"), t)
}
