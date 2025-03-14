package hackernews

import (
	"fmt"
	"hot-hacker-new/internal/models"
	"net/http"
	"regexp"
	"time"

	"golang.org/x/net/html"
)

func ParseDailyPage(url string) ([]*models.PostLink, error) {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	dataByUrl, err := parseDate(url)
	if err != nil {
		return nil, err
	}
	// 发送HTTP GET请求
	resp, err := client.Get(url)
	if err != nil {
		fmt.Printf("获取页面时出错: %v\n", err)
		return nil, fmt.Errorf("获取页面时出错: %v\n", err)
	}
	defer resp.Body.Close()

	// 检查HTTP响应状态
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("HTTP错误: %v\n", resp.Status)
		return nil, fmt.Errorf("HTTP错误: %v\n", resp.Status)
	}
	pList := make([]*models.PostLink, 0)

	// 使用goquery解析HTML
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	var f func(*html.Node)

	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "li" {
			p := &models.PostLink{}
			p.Date = dataByUrl
			p.CreatedAt = time.Now().Unix()
			p.UpdatedAt = time.Now().Unix()
			p.DeletedAt = 0
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				if c.Type == html.ElementNode && c.Data == "span" {
					for a := range c.Attr {
						if c.Attr[a].Key == "class" && c.Attr[a].Val == "storylink" {
							for l := c.FirstChild; l != nil; l = l.NextSibling {
								if l.Type == html.ElementNode && l.Data == "a" {
									for a2 := range l.Attr {
										if l.Attr[a2].Key == "href" {
											p.PostLink = l.Attr[a2].Val
										}
									}
									p.Title = l.FirstChild.Data
								}
							}
						} else if c.Attr[a].Key == "class" && c.Attr[a].Val == "postlink" {
							for l := c.FirstChild; l != nil; l = l.NextSibling {
								if l.Type == html.ElementNode && l.Data == "a" {
									p.CommentsLink = l.Attr[0].Val
								}
							}
						}
					}
				}
			}
			if p.PostLink != "" {
				pList = append(pList, p)
			}

		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return pList, nil
}

func parseDate(url string) (string, error) {
	// 使用正则表达式匹配URL中的日期部分
	re := regexp.MustCompile(`/(\d{4}-\d{2}-\d{2})\.html`)
	matches := re.FindStringSubmatch(url)

	if len(matches) < 2 {
		return "", fmt.Errorf("在URL中找不到日期")
	}

	date := matches[1]

	// 验证日期格式是否有效
	_, err := time.Parse("2006-01-02", date)
	if err != nil {
		return "", fmt.Errorf("无法解析日期: %v", err)
	}
	return date, nil
}
