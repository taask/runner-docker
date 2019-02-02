package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/user"
	"path"
	"path/filepath"

	"github.com/cohix/simplcrypto"
	"github.com/pkg/errors"
	"github.com/taask/taask-server/auth"
	sconfig "github.com/taask/taask-server/config"
	yaml "gopkg.in/yaml.v2"
)

// ConfigClientBadeDir is the path in $HOME where configs are stored/
const (
	ConfigClientBaseDir         = ".taask/client/config/"
	ConfigClientDefaultFilename = "admin-auth.yaml"
)

// LocalAuthConfig includes everything needed to auth with a member group
type LocalAuthConfig struct {
	sconfig.ClientAuthConfig
	Passphrase    string        `yaml:"passphrase,omitempty"`
	ActiveSession ActiveSession `yaml:"-"`
}

// ActiveSession represents an active session with the server
type ActiveSession struct {
	*auth.Session      `yaml:"-"`
	Keypair            *simplcrypto.KeyPair `yaml:"-"`
	MasterRunnerPubKey *simplcrypto.KeyPair `yaml:"-"`
}

// LocalAuthConfigFromFile reads a LocalAuthConfig from a file
func LocalAuthConfigFromFile(filepath string) (*LocalAuthConfig, error) {
	raw, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, errors.Wrap(err, "failed to ReadFile")
	}

	config := &LocalAuthConfig{}
	if err := yaml.Unmarshal(raw, config); err != nil {
		if jsonErr := json.Unmarshal(raw, config); jsonErr != nil {
			return nil, errors.Wrap(jsonErr, errors.Wrap(err, "failed to yaml and json Unmarshal").Error()) // stupid, but whatever
		}
	}

	return config, nil
}

// GroupKey returns the key for a group
func (la *LocalAuthConfig) GroupKey() (*simplcrypto.SymKey, error) {
	return auth.GroupDerivedKey(la.Passphrase)
}

// WriteServerConfig writes the admin groups's auth file to disk
func (la *LocalAuthConfig) WriteServerConfig(filename string) error {
	serverConfigPath := filepath.Join(sconfig.DefaultServerConfigDir(), filename)

	return la.ClientAuthConfig.WriteYAML(serverConfigPath)
}

// WriteYAML writes the YAML marshalled config to disk
func (la *LocalAuthConfig) WriteYAML(filepath string) error {
	rawYAML, err := yaml.Marshal(la)
	if err != nil {
		return errors.Wrap(err, "failed to yaml.Marshal")
	}

	if err := ioutil.WriteFile(filepath, rawYAML, 0666); err != nil {
		return errors.Wrap(err, "failed to WriteFile")
	}

	return nil
}

// DefaultClientConfigDir returns ~/.taask/client/config unless XDG_CONFIG_HOME is set
func DefaultClientConfigDir() string {
	u, err := user.Current()
	if err != nil {
		return ""
	}

	root := u.HomeDir
	xdgConfig, useXDG := os.LookupEnv("XDG_CONFIG_HOME")
	if useXDG {
		root = xdgConfig
	}

	return path.Join(root, ConfigClientBaseDir)
}
