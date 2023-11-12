package types

import (
	"github.com/caarlos0/env/v6"
	"golang.org/x/xerrors"

	"github.com/khulnasoft/tunnel/pkg/fanal/types"
)

// DockerConfig holds the config of Docker
type DockerConfig struct {
	UserName      string `env:"TUNNEL_USERNAME"`
	Password      string `env:"TUNNEL_PASSWORD"`
	RegistryToken string `env:"TUNNEL_REGISTRY_TOKEN"`
	NonSSL        bool   `env:"TUNNEL_NON_SSL" envDefault:"false"`
}

// GetDockerOption returns the Docker scanning options using DockerConfig
func GetDockerOption(insecureTlsSkip bool, Platform string) (types.DockerOption, error) {
	cfg := DockerConfig{}
	if err := env.Parse(&cfg); err != nil {
		return types.DockerOption{}, xerrors.Errorf("unable to parse environment variables: %w", err)
	}

	return types.DockerOption{
		UserName:              cfg.UserName,
		Password:              cfg.Password,
		RegistryToken:         cfg.RegistryToken,
		InsecureSkipTLSVerify: insecureTlsSkip,
		NonSSL:                cfg.NonSSL,
		Platform:              Platform,
	}, nil
}
