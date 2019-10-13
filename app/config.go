package app

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pirmd/cli/app/configdir"
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

//Load loads config from configuration's files set, looking at each file
//location one by one. Any non-existing file in configuration's Files is
//silently ignored.
//If no Unmarshaller is defined, the function will panic.
func (cfg *Config) Load() error {
	if cfg.Unmarshaller == nil {
		panic("no Unmarshaller is defined")
	}

	for _, rc := range cfg.Files {
		b, err := ioutil.ReadFile(rc.Name)
		if err != nil && !os.IsNotExist(err) {
			return err
		}

		if err := cfg.load(b); err != nil {
			return err
		}
	}

	return nil
}

func (cfg *Config) load(b []byte) error {
	return cfg.Unmarshaller(b, cfg.Var)
}

//DefaultConfigFiles returns a commonly used ConfigFile that is to say an rc
//files from user's config dir (if any).
func DefaultConfigFiles(rc string) []*ConfigFile {
	return []*ConfigFile{
		DefaultSystemConfigFile(rc),
		DefaultUserConfigFile(rc),
	}
}

//DefaultUserConfigFile returns a commonly used ConfigFile that points to
//per-user config file
func DefaultUserConfigFile(rc string) *ConfigFile {
	return &ConfigFile{
		Name:  filepath.Join(configdir.PerUser, filepath.Base(os.Args[0]), rc),
		Usage: "Per-user configuration location",
	}
}

//DefaultSystemConfigFile returns a commonly used ConfigFile that points to a
//system-wide config file
func DefaultSystemConfigFile(rc string) *ConfigFile {
	return &ConfigFile{
		Name:  filepath.Join(configdir.SystemWide, filepath.Base(os.Args[0]), rc),
		Usage: "System-wide configuration location",
	}
}
