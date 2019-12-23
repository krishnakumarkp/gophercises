package apputil

type Configuration struct {
	File string
}

var AppConfig Configuration

func init() {
	AppConfig = Configuration{}
}
