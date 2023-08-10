package browse

import (
	"strings"

	"github.com/pkg/browser"
)

func OpenURL(url string) error {
	// Prepend "http://" to the URL if it doesn't have a scheme
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "http://" + url
	}

	err := browser.OpenURL(url)
	if err != nil {
		return err
	}

	return nil
}
