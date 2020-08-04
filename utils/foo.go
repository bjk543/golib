package utils

import (
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func HttpGet(url string) string {

	req, e := http.NewRequest("GET", url, nil)
	if e != nil {
		return "nil, e"
	}

	client := http.Client{Timeout: 2 * time.Second}

	retries := 15
	res, err := client.Do(req)
	for {
		if err == nil || retries < 0 {
			break
		}
		retries--
		log.Println("retries", retries)

		res, err = client.Do(req)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		print(err)
	}
	// WriteFile("body.html", string(body))
	return string(body)
}

func GetKeyWord(ss, l, r string) string {

	a := strings.Index(ss, l)
	if a < 0 {
		return ""
	}
	ss = ss[a:]
	b := strings.Index(ss, r)
	if b < 0 {
		return ""
	}
	return ss[len(l):b]
}

func HttpGet2(url string, querys map[string]string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	q := req.URL.Query()

	for k, v := range querys {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()

	// fmt.Println(req.URL.String())
	// Output:
	// http://api.themoviedb.org/3/tv/popular?another_thing=foo+%26+bar&api_key=key_from_environment_or_flag
	var resp *http.Response
	retires := 10
	for {
		resp, err = http.DefaultClient.Do(req)
		if err == nil || retires == 0 {
			break
		}
		retires--
	}

	if err != nil {
		log.Fatal(err)
	}

	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	print(err)
	// }
	// WriteFile("body.html", string(body))

	return resp, err
}

func HttpPost(resp *http.Response, url string) string {

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		print(err)
	}

	for _, i := range resp.Cookies() {
		// fmt.Println(i.Name)
		// fmt.Println(i.Value)
		req.AddCookie(&http.Cookie{Name: i.Name, Value: i.Value})
	}
	resp2, err := client.Do(req)
	if err != nil {
		print(err)
		return ""
	} else {
		defer resp2.Body.Close()
	}

	body, err := ioutil.ReadAll(resp2.Body)
	if err != nil {
		print(err)
	}
	// WriteFile("body.html", string(body))
	return string(body)
}

func HttpGetUrlInfo(url string) string {
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/46.0.2490.86 Safari/537.36")
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	timeout := time.Duration(10 * time.Second) //超时时间50ms
	client := &http.Client{
		Timeout: timeout,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	var resp *http.Response
	retires := 10
	for {
		resp, err = client.Do(req)
		if err == nil || retires == 0 {
			break
		}
		retires--
	}

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	if err != nil {
		print(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}
	return string(body)

}

func GetMapKeys(mymap map[string]bool) []string {
	keys := make([]string, 0, len(mymap))
	for k := range mymap {
		keys = append(keys, k)
	}
	return keys
}

func HasKey(s string, ss []string) bool {
	for _, v := range ss {
		if strings.Contains(s, v) {
			return true
		}
	}
	return false
}
