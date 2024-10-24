package endpoints

import (
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"os"
	"strings"
)

func NewResolver() endpoints.Resolver {
	flag := os.Getenv("KHULNASOFT_DISABLE")

	if flag == "1" || strings.ToLower(flag) == "true" {
		return endpoints.DefaultResolver()
	} else {
		return NewKhulnasoftResolver()
	}
}