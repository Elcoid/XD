package config

import (
	"xd/lib/configparser"
)

type Config struct {
	I2P        I2PConfig
	Storage    StorageConfig
	RPC        RPCConfig
	Log        LogConfig
	Bittorrent BittorrentConfig
}

// Configurable interface for entity serializable to/from config parser section
type Configurable interface {
	Load(s *configparser.Section) error
	Save(c *configparser.Section) error
}

// Load loads a config from file by filename
func (cfg *Config) Load(fname string) (err error) {
	sects := map[string]Configurable{
		"i2p":        &cfg.I2P,
		"storage":    &cfg.Storage,
		"rpc":        &cfg.RPC,
		"log":        &cfg.Log,
		"bittorrent": &cfg.Bittorrent,
	}
	var c *configparser.Configuration
	c, err = configparser.Read(fname)
	for sect, conf := range sects {
		if c == nil {
			err = conf.Load(nil)
		} else {
			s, _ := c.Section(sect)
			err = conf.Load(s)
		}
		if err != nil {
			return
		}
	}
	return
}

// Save saves a loaded config to file by filename
func (cfg *Config) Save(fname string) (err error) {
	sects := map[string]Configurable{
		"i2p":        &cfg.I2P,
		"storage":    &cfg.Storage,
		"rpc":        &cfg.RPC,
		"log":        &cfg.Log,
		"bittorrent": &cfg.Bittorrent,
	}
	c := configparser.NewConfiguration()
	for sect, conf := range sects {
		s := c.NewSection(sect)
		err = conf.Save(s)
		if err != nil {
			return
		}
	}
	err = configparser.Save(c, fname)
	return
}
