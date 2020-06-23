package shared

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Word struct {
	Word                 string `json:"w"`
	WordPronunciation    string `json:"wp"`
	Pron                 string `json:"p"`
	Def                  string `json:"d"`
	Example              string `json:"e"`
	ExamplePronunciation string `json:"ep"`
}

// 学习使用，请勿商业用途
func NewWord(word string) (*Word, error) {
	format := func(str string) string {
		str = strings.TrimSpace(str)
		return strings.Replace(str, "\n", "", -1)
	}

	res, err := http.Get("https://www.collinsdictionary.com/zh/dictionary/english/" + word)
	if err != nil {

		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	s := doc.Find(".hom").First()
	ep, _ := s.Find(".audio_play_button").First().Attr("data-src-mp3")
	wp, _ := doc.Find(".pron").First().Find(".audio_play_button").First().Attr("data-src-mp3")
	wordObj := &Word{
		Word:                 word,
		WordPronunciation:    format(wp),
		Pron:                 format(doc.Find(".pron").First().Text()),
		Def:                  format(s.Find(".def").First().Text()),
		Example:              format(s.Find(".type-example .quote").First().Text()),
		ExamplePronunciation: format(ep),
	}

	if wordObj.Def == "" {
		return nil, fmt.Errorf("get data fail")
	}

	return wordObj, nil
}
