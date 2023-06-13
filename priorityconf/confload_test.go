package priorityconf

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Conf struct {
	MainDatabaseConf DBconf `toml:"main_database"`
}

type DBconf struct {
	DbUser             string `toml:"db_user"`
	DbHost             string `toml:"db_host"`
	DbPass             string `toml:"db_pass"`
	DbPort             int    `toml:"db_port""`
	DbName             string `toml:"db_name"`
	MaxConnectionCount int    `toml:"max_connection_count"`
}

func TestLoadConf(t *testing.T) {
	var conf Conf
	err := LoadConf("./conf", "test", "th", &conf)
	assert.NoError(t, err)
}
