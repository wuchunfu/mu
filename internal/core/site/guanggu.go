package site

import (
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const SITE_GUANGGU = "guanggu"

var GuangGuTabs = []map[string]string{
	{
		"tag":  "default",
		"name": "默认",
		"url":  "https://www.guozaoke.com/",
	},
	{
		"tag":  "latest",
		"name": "最新",
		"url":  "https://www.guozaoke.com/?tab=latest",
	},
}

type Guanggu struct {
	Site
}

var _ Spider = &Guanggu{}

func init() {
	RegistSite(SITE_GUANGGU, &Guanggu{
		Site{
			Name:     "光谷",
			Key:      SITE_GUANGGU,
			Root:     "https://www.guozaoke.com",
			Desc:     "武汉光谷社区",
			CrawType: CrawHtml,
			Tabs:     GuangGuTabs,
		},
	})
}

func (g *Guanggu) GetSite() *Site {
	return &g.Site
}

func (g *Guanggu) BuildUrl() ([]Link, error) {
	var list []Link
	for _, tab := range GuangGuTabs {
		link := Link{
			Key: tab["url"],
			Url: tab["url"],
			Tag: tab["tag"],
		}
		list = append(list, link)
	}

	return list, nil
}

func (g *Guanggu) CrawPage(link Link, headers map[string]string) (Page, error) {
	page, err := g.FetchData(link, nil, nil)
	if err != nil {
		return Page{}, err
	}
	var data []Hot
	doc := page.Doc
	doc.Find(".topic-item").Each(func(i int, s *goquery.Selection) {
		url, _ := s.Find(".main .title").Find("a").Attr("href")
		text := s.Find(".main .title").Find("a").Text()
		comment := s.Find(".count").Find("a").Text()
		if text == "" || url == "" {
			return
		}
		if comment == "" {
			comment = "0"
		}
		h := Hot{
			Title:     text,
			OriginUrl: fmt.Sprintf("%s%s", g.Root, url),
			Rank: (func() float64 {
				val, _ := strconv.ParseFloat(comment, 32)
				return float64(val)
			})(),
		}
		h.Key = g.FetchKey(h.OriginUrl)
		if h.Key == "" {
			return
		}
		data = append(data, h)
	})

	page.T = time.Now()
	page.List = data

	return page, nil
}

func (g *Guanggu) FetchKey(link string) string {
	reg := regexp.MustCompile(`.*/t/(\d+).*`)
	id := reg.ReplaceAllString(link, "$1")
	return id
}
