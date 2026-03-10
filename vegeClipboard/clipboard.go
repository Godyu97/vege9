//go:build !wl

package vegeClipboard

import (
	"golang.design/x/clipboard"
)

func Copy(text string) error {
	err := clipboard.Init()
	if err != nil {
		return err
	}
	clipboard.Write(clipboard.FmtText, []byte(text))
	return nil
}

func Paste() (string, error) {
	err := clipboard.Init()
	if err != nil {
		return "", err
	}
	return string(clipboard.Read(clipboard.FmtText)), nil
}
