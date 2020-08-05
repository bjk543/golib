package utils

import (
	"fmt"
	"testing"
)

func TestGetArticles2(t *testing.T) {
	u := HttpGetProxy("https://myip.com.tw/", "http://119.81.71.27:8123")
	// u := HttpGetProxy("https://myip.com.tw/", "")
	fmt.Println(u)
}
