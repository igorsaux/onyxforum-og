package page

import (
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type User struct {
	Name      string
	AvatarUrl string
}

type Post struct {
	User
	Title  string
	Rating int
}

func parseAuthor(post *goquery.Selection) User {
	author := User{}

	author.Name = post.Find(".author-name").Text()
	author.AvatarUrl, _ = post.Find(".author-avatar img").Attr("src")

	return author
}

func parseRating(post *goquery.Selection) int {
	rating, _ := strconv.Atoi(post.Find(".post-rating-counter ").Text())

	return rating
}

func parsePost(doc *goquery.Document) Post {
	firstPost := doc.Find(".panel-body")

	author := parseAuthor(firstPost)
	rating := parseRating(firstPost)

	title := doc.Find(".forum-title").Text()

	return Post{
		User:   author,
		Rating: rating,
		Title:  title,
	}
}

func ParsePost(url string) Post {
	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		panic(err)
	}

	return parsePost(doc)
}
