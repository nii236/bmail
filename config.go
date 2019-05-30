package main

var c *Config

func NewConfig() {
	if c != nil {
		panic("config already initialised")
	}
	c = &Config{
		Addresses: &AddressesConfig{
			ReceivingAddress:           "BM-2cVrrbMLZx6WoH1k8egoCLKk1R2ZEhATLJ",
			SendingAddress:             "BM-2cWdn2iLJmVLwYwvwWZ96kGyeWDwq97z9A",
			RegistrationAddress:        "BM-2cV6JbTkxiA7EDuAKvAE6oXxEEWUEWFE8v",
			DeregistrationAddress:      "BM-2cTQgr1iEnMnwQUsETpoVuCXwkCngsnbgj",
			BugReportAddressBitmessage: "BM-2cTYHGfV4HY5kfpD4M1TGCtnpxy2we4oTE",
			BugReportAddressEmail:      "bugs@bmail.dev",
		},
	}
}

func GetConfig() *Config {
	if c == nil {
		panic("config not initialised")
	}
	return c
}

type Config struct {
	General    *GeneralConfig    `toml:"General"`
	Bitmessage *BitmessageConfig `toml:"Bitmessage"`
	Storage    *StorageConfig    `toml:"Storage"`
	Server     *ServerConfig     `toml:"Server"`
	Addresses  *AddressesConfig  `toml:"Addresses"`
}

type GeneralConfig struct {
	Debug            bool
	RespondToInvalid bool
	ProcessInterval  int
}

type BitmessageConfig struct {
	Conn     string
	Username string
	Password string
	Host     string
	Port     string
}

type StorageConfig struct {
	MailFolder  string
	LogFilename string
}

type ServerConfig struct {
	DomainName string
}

type AddressesConfig struct {
	ReceivingAddress           string
	SendingAddress             string
	RegistrationAddress        string
	DeregistrationAddress      string
	BugReportAddressBitmessage string
	BugReportAddressEmail      string
}
