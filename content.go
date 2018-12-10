package container_launcher

import (
	"io"
	"strings"
)

func init() {
	RegisterResolver(&contentResolver{})
}

type contentResolver struct{}

func (contentResolver) IsDefinedAt(str string) bool {
	return strings.HasPrefix(str, "content:")
}

func (contentResolver) UsageText() string {
	return "content:<content-which-will-be-used-unmodified>"
}

func (contentResolver) Resolve(str string, w io.WriterAt) error {
	_, err := w.WriteAt([]byte(str), 0)
	return err
}
