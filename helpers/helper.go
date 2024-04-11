package helpers

import "fmt"

func GenerateStaticSite(qrUrl, userId string) string{
	htmlString := `<!DOCTYPE html>
	<body>
		<h3>This is a memorize test static site for user %s</h3>
		<img src="%s" alt="Qr Image">
	</body>`

	hydratedString := fmt.Sprintf(htmlString, userId, qrUrl)

	return hydratedString
}