# Text2QRCode
Eine einfache API zur Umwandlung von Text in einen QRCode.

## Start
Nach dem Start des Dienstes, ist er per HTTP unter Port 8080 erreichbar.

Über z.B. `http://localhost:8080/` lässt sich die Swagger-Seite aufrufen. Dort sind die verfügbaren Endpunkte beschrieben.

# Endpunkte
Es gibt je einen Endpunkt, der per GET oder per POST angesprochen werden kann und diverse Parameter annimmt.

## Parameter
|Name|Beschreibung|
|-|-|
|text|Der zu kodierende Text|
|errorCorrection|Der Fehlerkorrekturlevel (0-3)<br/>0 = L<br/>1 = M<br/>2 = Q<br/>3 = H|
|size|Die Größe des QR-Codes in Pixel|
|whiteBorder|Gibt an, ob ein Rand um den QR-Code gezeichnet werden soll|