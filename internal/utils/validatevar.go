package utils

import (
	"fmt"
	"log"
	"net/url"
	"strings"
)

func ValidateListenAddr(listenAddr string) (string, error) {

	// The only valid protocol defined for a Kubernetes KMS plugin is "unix".
	const (
		proto = "unix"
	)

	url, err := url.Parse(listenAddr)
	if err != nil {
		return url.Path, fmt.Errorf("/!\\ invalid listen-addr %q, error: %v", listenAddr, err)
	}

	if len(listenAddr) == 0 {
		return url.Path, fmt.Errorf("/!\\ can not be an empty string")
	}

	if url.Scheme != proto {
		return url.Scheme, fmt.Errorf("/!\\ proto can be different than unix")
	}

	if strings.HasPrefix(url.Path, "/@") {
		return strings.TrimPrefix(url.Path, "/"), nil
	}

	log.Println("INFO: flag -listen set to", listenAddr)
	return url.Path, nil
}

func ValidateConfigfile(providerConfigFile string) (string, error) {

	if len(providerConfigFile) == 0 {
		return providerConfigFile, fmt.Errorf("/!\\ can not be an empty string")
	}

	log.Println("INFO: flag -configfile set to", providerConfigFile)
	return providerConfigFile, nil

}
