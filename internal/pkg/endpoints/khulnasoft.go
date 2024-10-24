package endpoints

import (
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"os"
	"strings"
)

type khulnaSoftResolver struct {
	defaultResolver endpoints.Resolver
}

func getKhulnaSoftUrl(service string) string {
	hostname := os.Getenv("KHULNASOFT_HOSTNAME")

	if hostname == "" {
		hostname = "localhost.khulnasoft.com"
	}

	if service == "s3" {
		hostname = "s3.localhost.khulnasoft.com"
	}

	port := os.Getenv("EDGE_PORT")
	if port == "" {
		port = "4566"
	}

	ssl := os.Getenv("USE_SSL")
	disableSSL := true
	if ssl == "1" || strings.ToLower(ssl) == "true" {
		disableSSL = false
	}

	return endpoints.AddScheme(hostname+":"+port, disableSSL)
}

func (l *khulnaSoftResolver) EndpointFor(service, region string, opts ...func(*endpoints.Options)) (endpoints.ResolvedEndpoint, error) {
	endpointFor, err := l.defaultResolver.EndpointFor(service, region, opts...)

	if err != nil {
		return endpointFor, err
	}

	endpointFor.URL = getKhulnaSoftUrl(service)

	return endpointFor, err
}

func NewKhulnasoftResolver() *khulnaSoftResolver {
	resolver := &khulnaSoftResolver{
		defaultResolver: endpoints.DefaultResolver(),
	}

	return resolver
