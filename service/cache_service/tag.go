package cacheservice

import (
	"strconv"
	"strings"

	"github.com/aifuxi/jianyu-blog-api/pkg/e"
)

type Tag struct {
	ID    int
	Name  string
	State int

	PageNum  int
	PageSize int
}

func (tag *Tag) GetTagsKey() string {
	keys := []string{
		e.CACHE_ARTICLE,
		"LIST",
	}

	if tag.Name != "" {
		keys = append(keys, tag.Name)
	}
	if tag.State > 0 {
		keys = append(keys, strconv.Itoa(tag.State))
	}
	if tag.PageNum > 0 {
		keys = append(keys, strconv.Itoa(tag.PageNum))
	}
	if tag.PageSize > 0 {
		keys = append(keys, strconv.Itoa(tag.PageSize))
	}

	return strings.Join(keys, "_")

}
