package text2qrcode

type QRCodeRequest struct {
	// Der zu kodierende Text
	Text string

	// Der Fehlerkorrekturlevel (0-3)
	// 0 = L
	// 1 = M
	// 2 = Q
	// 3 = H
	ErrorCorrection int

	// Die Größe des QR-Codes in Pixel
	Size int

	// Gibt an, ob ein Rand um den QR-Code gezeichnet werden soll
	WhiteBorder bool
}
