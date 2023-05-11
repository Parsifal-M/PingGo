package monitor

import (
	"fmt"
	"net/http"
)

func CheckWebsite(website string) {
	resp, err := http.Get(website)
	if err != nil {
		fmt.Printf("%s is down!\n", website)
		return
	}

	defer resp.Body.Close()
	fmt.Printf("%s is up!\n", website)
}
