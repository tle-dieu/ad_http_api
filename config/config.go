package config

import (
	"context"
	"fmt"

	"github.com/etf1/go-config"
	"github.com/etf1/go-config/env"
	"github.com/heetch/confita/backend/flags"
)

// configuration for extract-ads app
type ExtractAds struct {
	HTTPServerPort int    `config:"http_server_port"`
	MySQLHost      string `config:"mysql_host"`
	MySQLPort      int    `config:"mysql_port"`
	MySQLUser      string `config:"mysql_user"`
	MySQLPassword  string `config:"mysql_password"`
	MySQLDbName    string `config:"mysql_db_name"`
	MySQLTableName string `config:"mysql_table_name"`
}

// NewExtractAds creates a new ExtractAds configuration from env vars
func New() *ExtractAds {
	// create default config
	cfg := &ExtractAds{
		HTTPServerPort: 8080,
	}

	// load from .env and flags
	loader := config.NewConfigLoader(
		env.NewBackend(),
		flags.NewBackend(),
	)

	loader.LoadOrFatal(context.Background(), cfg)
	fmt.Println(config.TableString(cfg))

	return cfg
}
