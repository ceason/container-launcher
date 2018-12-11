package container_launcher

import (
	"io"
	"strings"
)

func init() {
	RegisterResolver(&awsSecretsmanagerResolver{})
}

type awsSecretsmanagerResolver struct{}

func (*awsSecretsmanagerResolver) IsDefinedAt(str string) bool {
	return strings.HasPrefix(str, "arn:aws:secretsmanager:")
}

func (*awsSecretsmanagerResolver) UsageText() string {
	panic("implement me")
}

func (*awsSecretsmanagerResolver) Resolve(str string, w io.WriterAt) error {
	panic("implement me")
}
