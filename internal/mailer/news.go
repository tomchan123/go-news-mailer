package mailer

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/tomchan123/go-news-mailer/internal/db"
)

type NewsServer struct {
	ApiKey      string
	ApiEndpoint string
}

func CreateNewsServer() *NewsServer {
	return &NewsServer{
		"pub_30538d32bff31cca76612ea8c3dca77f9c323",
		"https://newsdata.io/api/1",
	}
}

func CreateNewsArticle() *db.NewsAritcle {
	return &db.NewsAritcle{
		Authors:     make([]string, 0),
		PublishedAt: time.Now(),
	}
}

// fetch top n headlines
func (ns *NewsServer) FetchRecentArticles(n int, t int) ([]*db.NewsAritcle, error) {
	url := fmt.Sprintf("%s/news?language=en&timeframe=%d&image=1&full_content=1", ns.ApiEndpoint, t)

	nas, err := ns.getArticlesAPIRequest(url)
	if err != nil {
		return nil, fmt.Errorf("NewsServer.FetchRecentArticles: %v", err)
	}

	if len(nas) > n {
		nas = nas[:n]
	}

	gpts := CreateGPTServer()
	for _, na := range nas {
		s, err := gpts.ChatSummarise(na.Content)
		if err == nil {
			na.AISummary = s
		}
	}

	return nas, nil
}

func (ns *NewsServer) getArticlesAPIRequest(url string) ([]*db.NewsAritcle, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("NewsServer.getArticlesAPIRequest: %v", err)
	}
	req.Header.Add("X-ACCESS-KEY", ns.ApiKey)
	req.Header.Add("accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("NewsServer.getArticlesAPIRequest: %v", err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("NewsServer.getArticlesAPIRequest: %v", err)
	}

	// fmt.Printf("%s", body)

	arts, err := ns.marshalArticles(body)
	if err != nil {
		return nil, fmt.Errorf("NewsServer.getArticlesAPIRequest: %v", err)
	}

	return arts, nil
}

func (ns *NewsServer) marshalArticles(bs []byte) ([]*db.NewsAritcle, error) {
	var d map[string]interface{}
	err := json.Unmarshal(bs, &d)
	if err != nil {
		return nil, fmt.Errorf("NewsServer.marshalArticles: %v", err)
	}

	var nas []*db.NewsAritcle
	r := d["results"].([]interface{})
	for _, v := range r {
		nr := v.(map[string]interface{})
		na := CreateNewsArticle()

		// if x := nr["creator"]; x != nil {
		// 	na.Authors = x.([]string)
		// }
		if x := nr["image_url"]; x != nil {
			na.ImgUrl = x.(string)
		}
		if x := nr["pubDate"]; x != nil {
			lay := "2006-01-02 03:04:05"
			y, err := time.Parse(lay, x.(string))
			if err == nil {
				na.PublishedAt = y
			}
		}
		if x := nr["source_id"]; x != nil {
			na.Source = x.(string)
		}
		if x := nr["content"]; x != nil {
			na.Content = x.(string)
		}
		if x := nr["title"]; x != nil {
			na.Title = x.(string)
		}
		if x := nr["link"]; x != nil {
			na.Url = x.(string)
		}

		nas = append(nas, na)
	}

	return nas, nil
}
