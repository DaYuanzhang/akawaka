package config

import "strings"

type Options struct {
	DirPath       string
	Keywords_File string
	Keyword       string
	Keywords      []string
	Extens        string
	Extensions    []string
}

func (o *Options) SetExtensions() bool {
	if len(o.Extens) > 0 {
		arr := strings.Split(o.Extens, ",")
		if len(arr) > 0 {
			for _, v := range arr {
				o.Extensions = append(o.Extensions, strings.TrimSpace(v))
			}
			return true
		}
	}
	return false
}

func (o *Options) SetKeywords() bool {
	if len(o.Keyword) > 0 {
		arr := strings.Split(o.Keyword, ",")
		if len(arr) > 0 {
			for _, v := range arr {
				o.Keywords = append(o.Keywords, strings.TrimSpace(v))
			}
			return true
		}
	}
	return false
}
