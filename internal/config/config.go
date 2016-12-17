package config

import (
	"fmt"
	"sort"
	"strings"
)

func Optional(defaultValue string) Option {
	return Option{
		required:     false,
		defaultValue: defaultValue,
	}
}

func Required() Option {
	return Option{
		required: true,
	}
}

type Option struct {
	required     bool
	defaultValue string
}

type Schema map[string]Option

type Source func(string) (string, bool)

func (s Schema) NewConfig(source Source) (Config, error) {
	config := Config{
		values: map[string]string{},
	}
	var missingKeys []string

	for key, option := range s {
		value, ok := source(key)
		if !ok && option.required {
			missingKeys = append(missingKeys, key)
			continue
		}

		if !ok {
			value = option.defaultValue
		}
		config.values[key] = value
	}

	if len(missingKeys) > 0 {
		sort.Strings(missingKeys)

		return Config{}, fmt.Errorf(
			"missing configuration keys: %s",
			strings.Join(missingKeys, ", "),
		)
	}

	return config, nil
}

type Config struct {
	values map[string]string
}

func (c Config) Get(key string) string {
	if v, ok := c.values[key]; ok {
		return v
	}
	panic(fmt.Sprintf("configuration key was not in schema %s", key))
}
