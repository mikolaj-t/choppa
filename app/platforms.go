package main

import (
	"errors"
	"github.com/mileusna/useragent"
)

const (
	Android = iota + 1
	iOS
	Windows
	Linux
	MacOS
)

func platformFromUserAgent(userAgent string) (int, error) {
	parsedUserAgent := useragent.Parse(userAgent)
	return platformFromString(parsedUserAgent.OS)
}

func platformFromString(s string) (int, error) {
	switch s {
	case useragent.Android:
		return Android, nil
	case useragent.IOS:
		return iOS, nil
	case useragent.Windows:
		return Windows, nil
	case useragent.Linux:
		return Linux, nil
	case useragent.MacOS:
		return MacOS, nil
	}
	return 0, errors.New("platform not found!")
}
