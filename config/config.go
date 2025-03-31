package config

import "settings/domain"

type Config struct {
	DP  domain.DataProvider
	Log domain.Logger
}
