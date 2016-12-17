package config_test

import (
	. "github.com/ooesili/dead-drop/internal/config"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("configuration validation", func() {
	Describe("an empty schema", func() {
		schema := Schema{}
		source := mapSource(nil)
		config, err := schema.NewConfig(source)

		It("can successfully create a config", func() {
			Expect(err).ToNot(HaveOccurred())
		})

		It("panics when trying to get an option not in the schema", func() {
			Expect(func() {
				config.Get("anything")
			}).To(Panic())
		})
	})

	Describe("a schema with a required value", func() {
		schema := Schema{
			"REQUIRED": Required(),
		}

		Context("when the value is present in the source", func() {
			source := mapSource(map[string]string{
				"REQUIRED": "wohoo",
			})

			It("can get the value from the config", func() {
				config, err := schema.NewConfig(source)
				Expect(err).ToNot(HaveOccurred())
				Expect(config.Get("REQUIRED")).To(Equal("wohoo"))
			})
		})

		Context("when the value is missing from the source", func() {
			source := mapSource(nil)

			It("returns an error when creating a new config", func() {
				_, err := schema.NewConfig(source)
				Expect(err).To(MatchError("missing configuration keys: REQUIRED"))
			})
		})
	})

	Describe("a schema with multiple required values", func() {
		schema := Schema{
			"REQUIRED":  Required(),
			"MANDITORY": Required(),
		}

		Context("when the source is missng all values", func() {
			source := mapSource(nil)

			It("return an error describing all missing values", func() {
				_, err := schema.NewConfig(source)
				Expect(err).To(MatchError("missing configuration keys: MANDITORY, REQUIRED"))
			})
		})
	})

	Describe("a schema with an optional value", func() {
		schema := Schema{
			"OPTIONAL": Optional("default value"),
		}

		Context("when the value is present in the source", func() {
			source := mapSource(map[string]string{
				"OPTIONAL": "source value",
			})

			It("uses the value from the source", func() {
				config, err := schema.NewConfig(source)
				Expect(err).ToNot(HaveOccurred())
				Expect(config.Get("OPTIONAL")).To(Equal("source value"))
			})
		})

		Context("when the value is missing from the source", func() {
			source := mapSource(nil)

			It("uses the default value from the schema", func() {
				config, err := schema.NewConfig(source)
				Expect(err).ToNot(HaveOccurred())
				Expect(config.Get("OPTIONAL")).To(Equal("default value"))
			})
		})
	})

	Describe("a schema with both a required and an optional value", func() {
		schema := Schema{
			"DB_HOST": Required(),
			"PORT":    Optional("3000"),
		}
		source := mapSource(map[string]string{
			"DB_HOST": "db.example.com",
		})
		config, err := schema.NewConfig(source)

		It("can successfully create a config", func() {
			Expect(err).ToNot(HaveOccurred())
		})

		It("can get the required value", func() {
			Expect(config.Get("DB_HOST")).To(Equal("db.example.com"))
		})

		It("can get the optional value", func() {
			Expect(config.Get("PORT")).To(Equal("3000"))
		})
	})
})

func mapSource(m map[string]string) Source {
	return func(key string) (string, bool) {
		value, ok := m[key]
		return value, ok
	}
}
