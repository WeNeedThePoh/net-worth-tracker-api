package config_test

import (
	"testing"

	"weneedthepoh/net-worth-tracker-api/internal/config"

	"github.com/stretchr/testify/require"
	"gotest.tools/assert"
)

func TestReadFile(t *testing.T) {
	buffer := []byte(`
serve:
  public:
    port: 23025
log:
  level: debug
db:
  host: database-host-name
  port: 3306
  name: schema-name
  user: username
  password: user-password
  ssl: disable
`)

	var expected config.Conf

	expected = *new(config.Conf)
	expected.Serve.Public.Port = 23025
	expected.Log.Level = "debug"
	expected.Db.Host = "database-host-name"
	expected.Db.Port = 3306
	expected.Db.Name = "schema-name"
	expected.Db.User = "username"
	expected.Db.Password = "user-password"
	expected.Db.Ssl = "disable"

	// when
	configs, err := config.ParseFile(buffer)

	//then
	require.Nil(t, err)
	require.NotEmpty(t, configs)
	assert.Equal(t, expected, *configs)
}

func TestReadFileOnYamlFormatError(t *testing.T) {
	buffer := []byte(`
	serve:
  		public:
port: 23025
log:
level: debug
`)

	// when
	configs, err := config.ParseFile(buffer)

	//then
	require.NotNil(t, err)
	require.Nil(t, configs)
}

func TestReadFileOnMissingPropertiesError(t *testing.T) {
	buffer := []byte(`
serve:
  public:
   port: 23025
log:
level: debug
`)

	// when
	configs, err := config.ParseFile(buffer)

	//then
	require.NotNil(t, err)
	require.Nil(t, configs)
}
