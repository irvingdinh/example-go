package config

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/spf13/viper"
)

var _ = Describe("GetAppConfig", func() {
	viper.Set("app.env", "local")

	It("should returns", func() {
		Expect(GetAppConfig()).To(Equal(AppConfig{
			Env: "local",
		}))
	})
})

var _ = Describe("GetHTTPConfig", func() {
	viper.Set("http.port", "8080")

	It("should returns", func() {
		Expect(GetHTTPConfig()).To(Equal(HTTPConfig{
			Port: 8080,
		}))
	})
})

var _ = Describe("GetLoggingConfig", func() {
	Context("local environment", func() {
		viper.Set("app.env", "local")

		It("should returns", func() {
			Expect(GetLoggingConfig()).To(Equal(LoggingConfig{
				Formatter: LoggingTextFormatter,
				Level:     LoggingDebugLevel,
			}))
		})
	})

	Context("staging environment", func() {
		viper.Set("app.env", "staging")

		It("should returns", func() {
			Expect(GetLoggingConfig()).To(Equal(LoggingConfig{
				Formatter: LoggingJSONFormatter,
				Level:     LoggingDebugLevel,
			}))
		})
	})

	Context("production environment", func() {
		viper.Set("app.env", "production")

		It("should returns", func() {
			Expect(GetLoggingConfig()).To(Equal(LoggingConfig{
				Formatter: LoggingJSONFormatter,
				Level:     LoggingInfoLevel,
			}))
		})
	})
})
