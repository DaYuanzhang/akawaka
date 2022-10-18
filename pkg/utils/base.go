package utils

import (
	"akawaka/pkg/config"
	"errors"
	"strings"
)

func Transform(options *config.Options) error {
	if len(options.Keyword) != 0 {
		if strings.Contains(options.Keyword, ",") {
			options.SetKeywords()
		} else {
			options.Keywords = append(options.Keywords, options.Keyword)
		}
		if strings.Contains(options.Extens, ",") {
			options.SetExtensions()
			return nil
		} else {
			options.Extensions = append(options.Extensions, options.Extens)
			return nil
		}
	} else if len(options.Keywords_File) != 0 {
		var err error
		options.Keywords, err = ReadArrFromTxt(options.Keywords_File)
		if err != nil {
			return err
		}
		if strings.Contains(options.Extens, ",") {
			options.SetExtensions()
			return nil
		} else {
			options.Extensions = append(options.Extensions, strings.TrimSpace(options.Extens))
			return nil
		}
	} else {
		return errors.New("keyword is required...")
	}

}
