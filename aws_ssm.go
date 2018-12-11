package container_launcher

import (
	"io"
	"strings"
)

func init() {
	RegisterResolver(&awsSsmResolver{})
}

type awsSsmResolver struct{}

func (*awsSsmResolver) IsDefinedAt(str string) bool {
	return strings.HasPrefix(str, "arn:aws:ssm:")
}

func (*awsSsmResolver) UsageText() string {
	panic("implement me")
}

func (*awsSsmResolver) Resolve(str string, w io.WriterAt) error {
	panic("implement me")
}
