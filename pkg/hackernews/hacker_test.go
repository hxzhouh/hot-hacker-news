package hackernews

import (
	"fmt"
	"testing"
)

// test ParseDailyPage
func TestParseDailyPage(t *testing.T) {
	url := "https://www.daemonology.net/hn-daily/2025-03-11.html"
	pages, err := ParseDailyPage(url)
	if err != nil {
		t.Errorf("ParseDailyPage() error = %v", err)
		return
	}
	for _, v := range pages {
		fmt.Println(v)
	}
}
