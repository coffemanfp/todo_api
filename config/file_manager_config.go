package config

import (
	"fmt"
	"io/ioutil"
	"path"

	"gopkg.in/yaml.v3"
)

// FileManagerConfig is the Config implementation for the config files.
type FileManagerConfig struct {
	config ConfigInfo
}

func (f FileManagerConfig) Get() ConfigInfo {
	return f.config
}

// NewEnvManagerConfig initializes a new ConfigInfo instance by the config file provided.
//  @param env string: is the environment that must be used for get the config information.
//	 For example: "local" is for a local configuration.
//  @param configDir string: specifies the config dir path to locate the config files.
// 	@return conf ConfigInfo: new ConfigInfo instance with the env vars information.
// 	@return err error: error getting env vars values.
func NewFileManagerConfig(env, configDir string) (conf Config, err error) {
	var fc FileManagerConfig

	raw, err := readConfigFile(genConfigFileFullname(env, configDir))
	if err != nil {
		return
	}

	fc.config, err = newConfigWithBytes(raw)
	return
}

func readConfigFile(path string) (raw []byte, err error) {
	raw, err = ioutil.ReadFile(path)
	if err != nil {
		err = fmt.Errorf("not found: config filepath %s not found", path)
	}
	return
}

func genConfigFileFullname(env, configDir string) string {
	return path.Join(configDir, fmt.Sprintf("%s.yaml", env))
}

func newConfigWithBytes(b []byte) (c ConfigInfo, err error) {
	err = yaml.Unmarshal(b, &c)
	if err != nil {
		err = fmt.Errorf("invalid config: failed to get config info. Bad structure?")
	}
	return
}
