package goviper

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestViper(t *testing.T) {
	var config *viper.Viper = viper.New()
	assert.NotNil(t, config)
}

func TestJSON(t *testing.T) {
	config := viper.New()
	config.SetConfigName("config")
	config.SetConfigType("json")
	config.AddConfigPath(".")

	// read a config
	err := config.ReadInConfig()
	assert.Nil(t, err)

}

func TestGetDataJSON(t *testing.T) {
	config := viper.New()
	config.SetConfigName("config")
	config.SetConfigType("json")
	config.AddConfigPath(".")

	// read a config
	err := config.ReadInConfig()
	assert.Equal(t, "golang-viper", config.GetString("app.name"))
	assert.Equal(t, "Thomas Ardiansah", config.GetString("app.author"))
	assert.Equal(t, 3306, config.GetInt("database.port"))
	assert.True(t, config.GetBool("database.show_sql"))
	assert.Nil(t, err)
}

func TestYAML(t *testing.T) {
	config := viper.New()
	config.SetConfigFile("config.yaml")
	config.AddConfigPath(".")

	err := config.ReadInConfig()
	assert.Nil(t, err)

	assert.Equal(t, "golang-viper", config.GetString("app.name"))
	assert.Equal(t, "Thomas Ardiansah", config.GetString("app.author"))
	assert.Equal(t, 3306, config.GetInt("database.port"))
	assert.True(t, config.GetBool("database.show_sql"))
}

func TestENV(t *testing.T) {
	config := viper.New()
	config.SetConfigFile("config.env")
	config.AddConfigPath(".")

	err := config.ReadInConfig()
	assert.Nil(t, err)

	assert.Equal(t, "golang_viper", config.GetString("APP_NAME"))
	assert.Equal(t, "Thomas Ardiansah", config.GetString("APP_AUTHOR"))
	assert.Equal(t, 3306, config.GetInt("DATABASE_PORT"))
	assert.True(t, config.GetBool("DATABASE_SHOW_SQL"))

}

func TestENVExport(t *testing.T) {
	config := viper.New()
	config.SetConfigFile("config.env")
	config.AddConfigPath(".")
	config.AutomaticEnv()

	err := config.ReadInConfig()
	assert.Nil(t, err)

	assert.Equal(t, "golang_viper", config.GetString("APP_NAME"))
	assert.Equal(t, "Thomas Ardiansah", config.GetString("APP_AUTHOR"))
	assert.Equal(t, 3306, config.GetInt("DATABASE_PORT"))
	assert.True(t, config.GetBool("DATABASE_SHOW_SQL"))

	assert.Equal(t, "Hello", config.GetString("FROM_ENV"))
}
