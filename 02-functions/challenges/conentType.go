package challenges

import (
	"fmt"
	"net/http"
)

func GetContentTypeUsage(url string) {
	if contentType, err := getContentType(url); err != nil {
		fmt.Errorf("ERROR: %s\n", err)
	} else {
		fmt.Println(contentType)
	}
}

func getContentType(url string) (string, error) {
	resp, err := http.Get(url)

	if err != nil {
		return err.Error(), err
	}

	defer resp.Body.Close()

	contentType := resp.Header.Get("Content-Type")

	if contentType == "" {
		return "", fmt.Errorf("can't find Content-Type header")
	}

	return contentType, nil
}
