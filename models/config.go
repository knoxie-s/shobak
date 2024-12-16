package models

type Config struct {
	App App `json:"app"`
}

type App struct {
	Port int `json:"port"`
}
