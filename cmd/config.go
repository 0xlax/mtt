package main

import (
	"io/ioutil"

	"github.com/BurntSushi/toml"

	"github.com/cosmos/cosmos-sdk/types/bech32"
)

// A chain we sign txs on
type Chain struct {
	Name   string // chain name
	Binary string // binary to use for signing
	Prefix string // bech32 address prefix
	ID     string // chain id for signing
	Node   string // node to broadcast signed txs to
	Denom  string // denom used for fees
}

// A key we sign txs with
type Key struct {
	Name      string
	Address   string
	LocalName string
}

type Gcp struct {
	ClientEmail  string `json:"client_email" structs:"client_email" mapstructure:"client_email"`
	ClientId     string `json:"client_id" structs:"client_id" mapstructure:"client_id"`
	PrivateKeyId string `json:"private_key_id" structs:"private_key_id" mapstructure:"private_key_id"`
	PrivateKey   string `json:"private_key" structs:"private_key" mapstructure:"private_key"`
	ProjectId    string `json:"project_id" structs:"project_id" mapstructure:"project_id"`
}

// Config file
type Config struct {
	User           string
	KeyringBackend string

	// AWS    AWS
	GCP    Gcp
	Keys   []Key
	Chains []Chain
}

func (c *Config) GetChain(name string) (Chain, bool) {
	for _, chain := range c.Chains {
		if chain.Name == name {
			return chain, true
		}
	}
	return Chain{}, false
}

func (c *Config) GetKey(name string) (Key, bool) {
	for _, key := range c.Keys {
		if key.Name == name {
			return key, true
		}
	}
	return Key{}, false
}

// load toml config
func loadConfig(filename string) (*Config, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	c := &Config{}
	err = toml.Unmarshal(b, c)
	if err != nil {
		return nil, err
	}

	if c.AWS.BucketRegion == "" {
		c.AWS.BucketRegion = defaultBucketRegion
	}

	return c, nil
}

func bech32ify(addrBech, prefix string) (string, error) {
	hrp, addrBytes, err := bech32.DecodeAndConvert(addrBech)
	if err != nil {
		return "", err
	}
	_ = hrp

	newAddrBech, err := bech32.ConvertAndEncode(prefix, addrBytes)
	if err != nil {
		return "", err
	}
	return newAddrBech, nil
}
