package text2qrcode_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/aimotrens/text2qrcode/text2qrcode"
	"github.com/gin-gonic/gin"
)

const (
	TestPort = "8080"
	TestHost = "localhost"
	TestURL  = "http://" + TestHost + ":" + TestPort
)

func init() {
	go func() {
		r := gin.Default()
		text2qrcode.SetRoutes(r)
		err := r.Run(":8080")
		if err != nil {
			panic(err)
		}
	}()
	<-time.After(100 * time.Millisecond)
}

func postJson200(qrReq text2qrcode.QRCodeRequest, t *testing.T) {
	data, err := json.Marshal(qrReq)
	if err != nil {
		t.Error("Json Encode failed", err)
	}

	httpClient := &http.Client{}
	req, err := httpClient.Post(TestURL+"/api/text2qrcode/encode", "application/json", bytes.NewReader(data))
	if err != nil {
		t.Error("HTTP Post failed", err)
	}

	if req.StatusCode != http.StatusOK {
		t.Error("Status Code is not 200, but", req.StatusCode)
	}
}

func postJson400(qrReq text2qrcode.QRCodeRequest, t *testing.T) {
	data, err := json.Marshal(qrReq)
	if err != nil {
		t.Error("Json Encode failed", err)
	}

	httpClient := &http.Client{}
	req, err := httpClient.Post(TestURL+"/api/text2qrcode/encode", "application/json", bytes.NewReader(data))
	if err != nil {
		t.Error("HTTP Post failed", err)
	}

	if req.StatusCode != http.StatusBadRequest {
		t.Error("Status Code is not 400, but", req.StatusCode)
	}
}

func TestHealthcheck_OK(t *testing.T) {
	httpClient := &http.Client{}
	_, err := httpClient.Get(TestURL + "/api/healthcheck")
	if err != nil {
		t.Error("Healthcheck failed")
	}
}

func TestEncodeWithQueryString_OK(t *testing.T) {
	httpClient := &http.Client{}
	_, err := httpClient.Get(TestURL + "/api/text2qrcode/encode?text=HelloWorld")
	if err != nil {
		t.Error("Request failed", err)
	}
}

func TestEncode_OK(t *testing.T) {
	qrReq := text2qrcode.QRCodeRequest{
		Text:            "HelloWorld",
		ErrorCorrection: 1,
		Size:            250,
		WhiteBorder:     true,
	}
	postJson200(qrReq, t)
}

func TestEncode_SizeTooLow(t *testing.T) {
	qrReq := text2qrcode.QRCodeRequest{
		Text:            "HelloWorld",
		ErrorCorrection: 1,
		Size:            99,
		WhiteBorder:     true,
	}
	postJson400(qrReq, t)
}

func TestEncode_SizeTooHigh(t *testing.T) {
	qrReq := text2qrcode.QRCodeRequest{
		Text:            "HelloWorld",
		ErrorCorrection: 1,
		Size:            1001,
		WhiteBorder:     true,
	}
	postJson400(qrReq, t)
}

func TestEncode_EclTooLow(t *testing.T) {
	qrReq := text2qrcode.QRCodeRequest{
		Text:            "HelloWorld",
		ErrorCorrection: -1,
		Size:            250,
		WhiteBorder:     true,
	}
	postJson400(qrReq, t)
}

func TestEncode_EclTooHigh(t *testing.T) {
	qrReq := text2qrcode.QRCodeRequest{
		Text:            "HelloWorld",
		ErrorCorrection: 4,
		Size:            250,
		WhiteBorder:     true,
	}
	postJson400(qrReq, t)
}
