package gobootstrap_test

import (
	"os"
	"testing"

	"github.com/arthureichelberger/gobootstrap"
	"github.com/stretchr/testify/assert"
)

func initEnv() {
	os.Setenv("DB_HOST", "127.0.0.1")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_USER", "arthur")
	os.Setenv("DB_PASSWORD", "arthurpwd")
	os.Setenv("DB_DATABASE", "arthurdb")
}

func TestConfig(t *testing.T) {
	t.Run("It can parse environment when every variable are defined.", func(t *testing.T) {
		initEnv()
		cfg, err := gobootstrap.ParseEnv()
		assert.NoError(t, err)
		assert.Equal(t, cfg.DbHost, "127.0.0.1")
		assert.Equal(t, cfg.DbPort, "5432")
		assert.Equal(t, cfg.DbUser, "arthur")
		assert.Equal(t, cfg.DbPassword, "arthurpwd")
		assert.Equal(t, cfg.DbDatabase, "arthurdb")
	})

	t.Run("It should throw an error when DB_HOST is not defined.", func(t *testing.T) {
		initEnv()
		os.Unsetenv("DB_HOST")

		_, err := gobootstrap.ParseEnv()
		assert.Error(t, err)
	})

	t.Run("It should throw an error when DB_PORT is not defined.", func(t *testing.T) {
		initEnv()
		os.Unsetenv("DB_PORT")

		_, err := gobootstrap.ParseEnv()
		assert.Error(t, err)
	})

	t.Run("It should throw an error when DB_USER is not defined.", func(t *testing.T) {
		initEnv()
		os.Unsetenv("DB_USER")

		_, err := gobootstrap.ParseEnv()
		assert.Error(t, err)
	})

	t.Run("It should throw an error when DB_PASSWORD is not defined.", func(t *testing.T) {
		initEnv()
		os.Unsetenv("DB_PASSWORD")

		_, err := gobootstrap.ParseEnv()
		assert.Error(t, err)
	})

	t.Run("It should throw an error when DB_DATABASE is not defined.", func(t *testing.T) {
		initEnv()
		os.Unsetenv("DB_DATABASE")

		_, err := gobootstrap.ParseEnv()
		assert.Error(t, err)
	})
}
