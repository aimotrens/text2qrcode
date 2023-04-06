package test_utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aimotrens/text2qrcode/app"
	"github.com/aimotrens/text2qrcode/app/text2qrcode"
	"github.com/stretchr/testify/assert"
)

func PostJson200(qrReq text2qrcode.QRCodeRequest, t *testing.T) {
	postJsonRequest("/api/text2qrcode/encode", qrReq, t, http.StatusOK)
}

func PostJson400(qrReq text2qrcode.QRCodeRequest, t *testing.T) {
	postJsonRequest("/api/text2qrcode/encode", qrReq, t, http.StatusBadRequest)
}

func PostJsonRaw400(data []byte, t *testing.T) {
	postJsonRawRequest("/api/text2qrcode/encode", data, t, http.StatusBadRequest)
}

func postJsonRequest(url string, qrReq text2qrcode.QRCodeRequest, t *testing.T, expectedStatus int) {
	data, err := json.Marshal(qrReq)
	if err != nil {
		t.Error("Json Encode failed", err)
	}

	postJsonRawRequest(url, data, t, expectedStatus)
}

func postJsonRawRequest(url string, data []byte, t *testing.T, expectedStatus int) {
	r := app.Setup()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, url, nil)
	req.Body = io.NopCloser(bytes.NewReader(data))
	r.ServeHTTP(w, req)

	assert.Equal(t, expectedStatus, w.Code)
	f := w.Body.String()

	fmt.Print(f)
}

func Get200(url string, t *testing.T) {
	getRequest(url, t, http.StatusOK)
}

func Get400(url string, t *testing.T) {
	getRequest(url, t, http.StatusBadRequest)
}

func getRequest(url string, t *testing.T, expectedStatus int) {
	r := app.Setup()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, expectedStatus, w.Code)
}
