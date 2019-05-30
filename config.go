package main

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
	ReceivingAddressLabel      string
	SendingAddressLabel        string
	RegistrationAddressLabel   string
	DeregistrationAddressLabel string
	BugReportAddressBitmessage string
	BugReportAddressEmail      string
}
