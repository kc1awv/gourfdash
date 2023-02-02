package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/BurntSushi/toml"
)

type Config struct {
	HostPort            string `toml:"hostport"`
	IPv4                string `toml:"ipv4"`
	IPv6                string `toml:"ipv6"`
	Refresh             int    `toml:"refresh"`
	LastHeard           int    `toml:"lastheard"`
	UrfFile             string `toml:"urffile"`
	UrfPidFile          string `toml:"urfpidfile"`
	Email               string `toml:"email"`
	Debug               bool   `toml:"debug"`
	SubPath             string `toml:"subpath"`
	CallsignCountryFile string `toml:"callsigncountryfile"`
}

func checkEnvOrErrString(k string) (string, error) {
	v, ok := os.LookupEnv(k)
	if !ok {
		return "", fmt.Errorf("%s not set", k)
	}
	return v, nil
}

func checkEnvOrErrInt(k string) (int, error) {
	v, ok := os.LookupEnv(k)
	if !ok {
		return 0, fmt.Errorf("%s not set", k)
	}
	i, err := strconv.ParseInt(v, 10, 32)
	return int(i), err
}

func NewConfigFromEnv() (*Config, error) {
	prefix := "GOMREFDASH"
	var config Config
	var err error

	config.HostPort, err = checkEnvOrErrString(prefix + "_HOSTPORT")
	if err != nil {
		return nil, err
	}

	config.Email, err = checkEnvOrErrString(prefix + "_EMAIL")
	if err != nil {
		return nil, err
	}

	config.IPv4, err = checkEnvOrErrString(prefix + "_IPV4")
	if err != nil {
		return nil, err
	}

	config.IPv6, err = checkEnvOrErrString(prefix + "_IPV6")
	if err != nil {
		return nil, err
	}

	config.LastHeard, err = checkEnvOrErrInt(prefix + "_LASTHEARD")
	if err != nil {
		return nil, err
	}

	config.UrfFile, err = checkEnvOrErrString(prefix + "_URFFILE")
	if err != nil {
		return nil, err
	}

	config.UrfPidFile, err = checkEnvOrErrString(prefix + "_URFPIDFILE")
	if err != nil {
		return nil, err
	}

	config.Refresh, err = checkEnvOrErrInt(prefix + "_REFRESH")
	if err != nil {
		return nil, err
	}

	config.SubPath, err = checkEnvOrErrString(prefix + "_SUBPATH")
	if err != nil {
		log.Println("warning: No SUBPATH defined, this is ok default to /")
	}

	config.CallsignCountryFile, err = checkEnvOrErrString(prefix + "_CALLSIGNCOUNTRYFILE")
	if err != nil {
		log.Println("warning: No CALLSIGNCOUNTRYFILE defined")
	}
	return &config, nil
}

func NewConfigFromFile(searchPaths []string) (*Config, error) {
	foundFile := ""
	for _, f := range searchPaths {
		_, err := os.Stat(f)
		if err != nil {
			log.Printf("%s not found", f)
			continue
		}
		log.Printf("%s found.", f)
		foundFile = f
		break
	}
	if foundFile == "" {
		log.Println("no config files found, let's try reading from environment")
		return NewConfigFromEnv()
	}
	b, err := os.ReadFile(foundFile)
	if err != nil {
		return nil, fmt.Errorf("unable to read configfile %s: %s", foundFile, err)
	}
	var config Config
	err = toml.Unmarshal(b, &config)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal %s into a Config: %s", foundFile, err)
	}
	return &config, nil
}
