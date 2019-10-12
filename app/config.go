package app

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/kirsle/configdir"
)

//TODO(pirmd): If command is called using --config FILE flag, configuration will be read
//from FILE and Config.Files will be ignored.

//ConfigFile represents a configuration file description (location and usage)
type ConfigFile struct {
	//Name is the path to the configuration file
	Name string
	//Usage is a short description of what this configuration file is about
	Usage string
}

//DefaultConfigFiles returns a commonly used ConfigFile that is to say an rc
//files from user's config dir (if any).
func DefaultConfigFiles() []*ConfigFile {
	return []*ConfigFile{
		DefaultSystemConfigFile(),
		DefaultUserConfigFile(),
	}
}

//DefaultUserConfigFile returns a commonly used ConfigFile that points to
//per-user config file
func DefaultUserConfigFile() *ConfigFile {
	appName := filepath.Base(os.Args[0])

	//TODO(pirmd): switch to go1.13 usrCfgDir, err := os.UserConfigDir()
	cfgDir := configdir.LocalConfig()

	return &ConfigFile{
		Name:  filepath.Join(cfgDir, appName),
		Usage: "Per-user configuration file for " + appName,
	}
}

//DefaultSystemConfigFile returns a commonly used ConfigFile that points to a
//system-wide config file
func DefaultSystemConfigFile() *ConfigFile {
	appName := filepath.Base(os.Args[0])
	if cfgDir := configdir.SystemConfig(); len(cfgDir) > 0 {
		return &ConfigFile{
			Name:  filepath.Join(cfgDir[0], "."+appName),
			Usage: "System-wide configuration file for " + appName,
		}
	}

	return &ConfigFile{}
}

//Config represents a set of Command's configuration information
type Config struct {
	//Var contains the unmarshalled configuration of Command
	Var interface{}

	//Unmarshaller is the function to use to read configuration from
	//configuration file.
	//It defaults to json.Unmarshal.
	Unmarshaller func([]byte, interface{}) error

	//Files contains the list of configuration file(s) to read from.  Config
	//will be read from this files set in the same order than the slice (giving
	//preference to latest path).
	//If a file within Files does not exist, config loading will ignore it and
	//move to the next one if any.  As a result, having loaded the
	//configuration from a config file is not mandatory (no valid config file
	//from Path will not triger any feedback/error), it is expected either to
	//test for nil/improper config from the main cmd.Execute routine or
	//provided a config with reasonable defaults.
	Files []*ConfigFile
}

//Load loads config look in order for each file in configuration's files set. Any
//non-existing file in configuration's Files is silently ignored.
func (cfg *Config) Load() error {
	if cfg.Unmarshaller == nil {
		cfg.Unmarshaller = json.Unmarshal
	}

	for _, rc := range cfg.Files {
		if err := cfg.loadFromFile(rc.Name); err != nil {
			return err
		}
	}

	return nil
}

func (cfg *Config) loadFromFile(path string) error {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	return cfg.Unmarshaller(b, &cfg.Var)
}
