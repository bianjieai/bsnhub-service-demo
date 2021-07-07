package redconfig

/*
重写读取配置文件
*/
import (
	"github.com/pkg/errors"

	"github.com/BSNDA/fabric-sdk-go-gm/pkg/common/logging"
	"github.com/BSNDA/fabric-sdk-go-gm/pkg/common/providers/core"
	"github.com/spf13/viper"

	"strings"
)

var logModules = [...]string{"fabsdk", "fabsdk/client", "fabsdk/core", "fabsdk/fab", "fabsdk/common",
	"fabsdk/msp", "fabsdk/util", "fabsdk/context"}

type options struct {
	envPrefix    string
	templatePath string
}

const (
	cmdRoot = "FABRIC_SDK"
)

// Option configures the package.
type Option func(opts *options) error

type SetOption func(def *defConfigBackend) error

func FromFile(name string, setopt []SetOption, opts ...Option) core.ConfigProvider {
	return func() ([]core.ConfigBackend, error) {
		backend, err := newBackend(opts...)
		if err != nil {
			return nil, err
		}

		if name == "" {
			return nil, errors.New("filename is required")
		}

		// create new viper
		backend.configViper.SetConfigFile(name)

		for _, set := range setopt {
			set(backend)
		}

		// If a config file is found, read it in.
		err = backend.configViper.MergeInConfig()
		if err != nil {
			return nil, errors.Wrapf(err, "loading config file failed: %s", name)
		}

		setLogLevel(backend)

		return []core.ConfigBackend{backend}, nil
	}
}

func newBackend(opts ...Option) (*defConfigBackend, error) {
	o := options{
		envPrefix: cmdRoot,
	}

	for _, option := range opts {
		err := option(&o)
		if err != nil {
			return nil, errors.WithMessage(err, "Error in options passed to create new config backend")
		}
	}

	v := newViper(o.envPrefix)

	//default backend for config
	backend := &defConfigBackend{
		configViper: v,
		opts:        o,
	}

	err := backend.loadTemplateConfig()
	if err != nil {
		return nil, err
	}

	return backend, nil
}

func newViper(cmdRootPrefix string) *viper.Viper {
	myViper := viper.New()
	myViper.SetEnvPrefix(cmdRootPrefix)
	myViper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	myViper.SetEnvKeyReplacer(replacer)
	return myViper
}

// setLogLevel will set the log level of the client
func setLogLevel(backend core.ConfigBackend) {
	loggingLevelString, _ := backend.Lookup("client.logging.level")
	logLevel := logging.INFO
	if loggingLevelString != nil {
		var err error
		logLevel, err = logging.LogLevel(loggingLevelString.(string))
		if err != nil {
			panic(err)
		}
	}

	// TODO: allow separate settings for each
	for _, logModule := range logModules {
		logging.SetLevel(logModule, logLevel)
	}
}
