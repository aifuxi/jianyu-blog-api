package cacheservice

import (
	"strconv"
	"strings"

	"github.com/aifuxi/jianyu-blog-api/pkg/e"
)

type Article struct {
	ID    int
	TagID int
	State int

	PageNum  int
	PageSize int
}

func (article *Article) GetArticleKey() string {
	return e.CACHE_ARTICLE + "_" + strconv.Itoa(article.ID)
}

func (article *Article) GetArticlesKey() string {
	keys := []string{
		e.CACHE_ARTICLE,
		"LIST",
	}

	if article.ID > 0 {
		keys = append(keys, strconv.Itoa(article.ID))
	}
	if article.TagID > 0 {
		keys = append(keys, strconv.Itoa(article.TagID))
	}
	if article.State > 0 {
		keys = append(keys, strconv.Itoa(article.State))
	}
	if article.PageNum > 0 {
		keys = append(keys, strconv.Itoa(article.PageNum))
	}
	if article.PageSize > 0 {
		keys = append(keys, strconv.Itoa(article.PageSize))
	}

	return strings.Join(keys, "_")

}
