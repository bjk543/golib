package im

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	TG_URL string
	MM_URL string
)

func init() {
	TG_URL = os.Getenv("TG_URL")
	MM_URL = os.Getenv("MM_URL")
}

func getPayloadTG(id int, text string) []string {
	textLen := 4096
	// if len(string([]rune(text))) > textLen {
	// 	text = string([]rune(text)[:textLen/2])
	// }
	ss := []string{}
	if len(text) > textLen {
		i := 1
		for ; i*textLen < len(text); i++ {
			ss = append(ss, fmt.Sprintf(`{"chat_id": %d,"text": "%s"}`, id, text[(i-1)*textLen:i*textLen]))
		}
		ss = append(ss, fmt.Sprintf(`{"chat_id": %d,"text": "%s"}`, id, text[(i-1)*textLen:]))
	} else {
		ss = append(ss, fmt.Sprintf(`{"chat_id": %d,"text": "%s"}`, id, text))
	}

	return ss
}
func getPayloadTG2(id int, text string) string {
	return fmt.Sprintf(`{"chat_id": %d,"text": "%s"}`, id, text)
}
func getPayloadMM(id string, text string) string {
	textLen := 16380
	if len(text) > textLen {
		text = string([]rune(text)[:textLen])
	}
	return fmt.Sprintf(`{"channel": "%s","text": "%s"}`, id, text)
}

func post(url string, text string) error {
	// log.Println("Post text ", text)

	payload := strings.NewReader(text)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/json")
	client := http.Client{Timeout: 5 * time.Second}

	res, err := client.Do(req)

	if err != nil {
		textLen := 128
		if len(text) > textLen {
			text = string([]rune(text)[:textLen])
		}
		log.Println("Post err!= nil", err, text)
	} else {
		defer res.Body.Close()
	}

	return err
	// body, _ := ioutil.ReadAll(res.Body)
	// fmt.Println(res)
	// fmt.Println(string(body))

}

func PostTG(id int, text string) error {

	p := getPayloadTG(id, text)
	var err error
	for _, v := range p {
		err = post(TG_URL, v)
	}
	return err
}
func PostTG2(id int, text string) error {
	p := getPayloadTG2(id, text)
	return post(TG_URL, p)
}
func PostMM(id string, text string) error {
	p := getPayloadMM(id, text)
	return post(MM_URL, p)
}
