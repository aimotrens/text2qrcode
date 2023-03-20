package text2qrcode

import "fmt"

func validateInput(qrReq QRCodeRequest) error {
	if len(qrReq.Text) == 0 {
		return fmt.Errorf("Der Text darf nicht leer sein")
	}

	if qrReq.ErrorCorrection < 0 || qrReq.ErrorCorrection > 3 {
		return fmt.Errorf("Der Fehlerkorrekturlevel muss zwischen 0 und 3 liegen")
	}

	if qrReq.Size < 100 || qrReq.Size > 1000 {
		return fmt.Errorf("Die Größe des QR-Codes muss zwischen 100 und 1000 liegen.")
	}

	return nil
}
