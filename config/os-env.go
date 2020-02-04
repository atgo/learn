package pkg

import (
	"bytes"
	"io/ioutil"
	"os"
	"regexp"

	"gopkg.in/yaml.v2"
)

func NewConfig(path string) (*Config, error) {
	cnf := &Config{}

	return cnf.read(path)
}

type (
	Config struct {
		Version    string           `yaml:"version"`
		Env        string           `yaml:"env"`
	}
)

func (this *Config) read(path string) (*Config, error) {
	raw, err := ioutil.ReadFile(path)
	if nil != err {
		return nil, err
	} else if content, err := this.replaceEnvVariables(raw); nil != err {
		return nil, err
	} else if err := yaml.Unmarshal(content, &this); nil != err {
		return nil, err
	}

	return this, nil
}

func (this Config) replaceEnvVariables(inBytes []byte) ([]byte, error) {
	if envRegex, err := regexp.Compile(`\${[0-9A-Za-z_]+(:((\${[^}]+})|[^}])+)?}`); err != nil {
		return nil, err
	} else if escapedEnvRegex, err := regexp.Compile(`\${({[0-9A-Za-z_]+(:((\${[^}]+})|[^}])+)?})}`); err != nil {
		return nil, err
	} else {
		replaced := envRegex.ReplaceAllFunc(inBytes, func(content []byte) []byte {
			var value string
			if len(content) > 3 {
				if colonIndex := bytes.IndexByte(content, ':'); colonIndex == -1 {
					value = os.Getenv(string(content[2 : len(content)-1]))
				} else {
					targetVar := content[2:colonIndex]
					defaultVal := content[colonIndex+1 : len(content)-1]

					value = os.Getenv(string(targetVar))
					if len(value) == 0 {
						value = string(defaultVal)
					}
				}
			}
			return []byte(value)
		})

		return escapedEnvRegex.ReplaceAll(replaced, []byte("$$$1")), nil
	}
}
