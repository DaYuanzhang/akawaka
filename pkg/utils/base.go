package utils

import (
	"akawaka/pkg/config"
	"errors"
	"strings"
)

func KeywordTransform(options *config.Options) error {
	// 判断关键词输入方式是文本还是多个字符串
	if len(options.Keyword) != 0 {
		if strings.Contains(options.Keyword, ",") { //输入形式为：-k "a,b,c"
			options.SetKeywords()
			return nil
		} else { // 输入形式为 -k "a"
			options.Keywords = append(options.Keywords, options.Keyword)
			return nil
		}
	} else if len(options.Keywords_File) != 0 {
		var err error
		options.Keywords, err = ReadArrFromTxt(options.Keywords_File)
		if err != nil {
			return err
		}
		return nil
	} else {
		return errors.New("keyword is required...")
	}
}

func ExtensTransform(options *config.Options) error {
	if len(options.Extens) != 0 {
		if strings.Contains(options.Extens, ",") {
			options.SetExtensions()
			return nil
		} else {
			options.Extensions = append(options.Extensions, options.Extens)
			return nil
		}
	} else {
		return errors.New("extension is required...")
	}
}

func Transform(options *config.Options) error {

	// 转换关键字属性
	err1 := KeywordTransform(options)
	if err1 != nil {
		return err1
	}

	// 转换后缀名属性
	err2 := ExtensTransform(options)
	if err2 != nil {
		return err2
	}

	return nil
}
