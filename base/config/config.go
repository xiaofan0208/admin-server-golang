package config

type Configer interface {
	String(section , key string)  string
	Int( section , key string ) int
}


