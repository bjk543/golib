package utils

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

func HttpGetProxy(urlStr, proxy string) string {

	up, err := url.Parse(proxy)
	if err != nil {
		log.Println(err)
	}

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	if len(proxy) > 0 {
		transport.Proxy = http.ProxyURL(up)
	}

	client := &http.Client{
		Timeout:   5 * time.Second,
		Transport: transport,
	}

	url, err := url.Parse(urlStr)
	if err != nil {
		log.Println(err)
		return ""
	}
	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		log.Println(err)
		return ""
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/46.0.2490.86 Safari/537.36")

	retries := 5
	res, err := client.Do(req)
	for {
		if err == nil {
			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				fmt.Println(err)
				return ""
			}
			defer res.Body.Close()
			fmt.Println(string(body))
			return string(body)
		}
		if retries < 0 {
			return ""
		}
		retries--
		log.Println("retries", retries)

		res, err = client.Do(req)
	}

	return ""
}

func HttpGet(url string) string {
	return HttpGetProxy(url, "")
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
	defer resp.Body.Close()
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
