package utils

import (
	"akawaka/pkg/config"
	"errors"
	"strings"
)

func SetArray(str string, strArray *[]string) bool {
	if len(str) > 0 {
		arr := strings.Split(str, ",")
		if len(arr) > 0 {
			for _, v := range arr {
				*strArray = append(*strArray, strings.TrimSpace(v))
			}
			println(*strArray)
			str = "a"
			return true
		}
	}
	return false
}

/*
将keyword转换为keywords数组
*/
func KeywordTransform(options *config.Options) error {
	// 判断关键词输入方式是文本还是多个字符串
	if len(options.Keyword) != 0 {
		if strings.Contains(options.Keyword, ",") { //输入形式为：-k "a,b,c"
			SetArray(options.Keyword, &options.Keywords)
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

/*
将extens转换成extensions数组
*/
func ExtensionTransform(options *config.Options) error {
	if len(options.Extension) != 0 {
		if strings.Contains(options.Extension, ",") {
			SetArray(options.Extension, &options.Extensions)
			return nil
		} else {
			options.Extensions = append(options.Extensions, options.Extension)
			return nil
		}
	} else if len(options.Extensions_File) != 0 {
		var err error
		options.Extensions, err = ReadArrFromTxt(options.Extensions_File)
		if err != nil {
			return err
		}
		return nil
	} else {
		return errors.New("extension is required...")
	}
}

func Transform(options *config.Options) error {

	/*
		优先判断用户受否有输入keyword，然后再判断extension后缀名
	*/
	// 转换关键字属性
	err1 := KeywordTransform(options)
	if err1 != nil {
		return err1
	}

	// 转换后缀名属性
	err2 := ExtensionTransform(options)
	if err2 != nil {
		return err2
	}

	return nil
}
