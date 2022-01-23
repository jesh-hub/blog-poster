package app

import (
	"io/ioutil"
	"log"
	"net/http"
)

type Tistory struct {
	BlogName    string
	AccessToken string
}

func NewTistory(blogName string, token string) *Tistory {
	return &Tistory{
		BlogName:    blogName,
		AccessToken: token,
	}
}

func (t *Tistory) List() {
	req, err := http.NewRequest("GET", "https://www.tistory.com/apis/post/list", nil)
	if err != nil {
		log.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("access_token", t.AccessToken)
	q.Add("blogName", t.BlogName)
	q.Add("output", "json")
	q.Add("page", "1")

	req.URL.RawQuery = q.Encode()

	resp, err := http.Get(req.URL.String())
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		str := string(respBody)
		println(str)
	}
}
