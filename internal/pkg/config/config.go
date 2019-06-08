package config

var c *C

// New creates a new singleton config
func New() *C {
	if c != nil {
		panic("config already initialised")
	}
	c = &C{
		Addresses: &AddressesConfig{
			ReceivingAddress:           "BM-2cVrrbMLZx6WoH1k8egoCLKk1R2ZEhATLJ",
			SendingAddress:             "BM-2cWdn2iLJmVLwYwvwWZ96kGyeWDwq97z9A",
			RegistrationAddress:        "BM-2cV6JbTkxiA7EDuAKvAE6oXxEEWUEWFE8v",
			DeregistrationAddress:      "BM-2cTQgr1iEnMnwQUsETpoVuCXwkCngsnbgj",
			BugReportAddressBitmessage: "BM-2cTYHGfV4HY5kfpD4M1TGCtnpxy2we4oTE",
			BugReportAddressEmail:      "bugs@bmail.dev",
		},
	}
	return c
}

// Get returns an exiting singleton config
func Get() *C {
	if c == nil {
		panic("config not initialised")
	}
	return c
}

// C is the config
type C struct {
	General    *GeneralConfig    `toml:"General"`
	Bitmessage *BitmessageConfig `toml:"Bitmessage"`
	Storage    *StorageConfig    `toml:"Storage"`
	Server     *ServerConfig     `toml:"Server"`
	Addresses  *AddressesConfig  `toml:"Addresses"`
}

// GeneralConfig is one of the subconfigs
type GeneralConfig struct {
	Debug            bool
	RespondToInvalid bool
	ProcessInterval  int
}

// BitmessageConfig is one of the subconfigs
type BitmessageConfig struct {
	Conn     string
	Username string
	Password string
	Host     string
	Port     string
}

// StorageConfig is one of the subconfigs
type StorageConfig struct {
	MailFolder  string
	LogFilename string
}

// ServerConfig is one of the subconfigs
type ServerConfig struct {
	DomainName string
}

// AddressesConfig is one of the subconfigs
type AddressesConfig struct {
	ReceivingAddress           string
	SendingAddress             string
	RegistrationAddress        string
	DeregistrationAddress      string
	BugReportAddressBitmessage string
	BugReportAddressEmail      string
}
