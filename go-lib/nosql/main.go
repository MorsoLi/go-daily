package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"time"
)

const flags = log.LstdFlags

type ServerProperties struct {
	Bind           string `cfg:"bind"`
	Port           int    `cfg:"port"`
	AppendOnly     bool   `cfg:"appendOnly"`
	AppendFilename string `cfg:"appendFilename"`
	MaxClients     int    `cfg:"maxclients"`
	RequirePass    string `cfg:"requirepass"`
	Databases      int    `cfg:"databases"`

	Peers []string `cfg:"peers"`
	Self  string   `cfg:"self"`
}

func (pro *ServerProperties) String() string {
	return fmt.Sprintf("Bind: %v, Port: %d", pro.Bind, pro.Port)
}

var Properties *ServerProperties

var (
	logFile            *os.File
	defaultPrefix      = ""
	defaultCallerDepth = 2
	logger             *log.Logger
	mu                 sync.Mutex
	logPrefix          = ""
	levelFlags         = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

type Settings struct {
	Path       string `yaml:"path"`
	Name       string `yaml:"name"`
	Ext        string `yaml:"ext"`
	TimeFormat string `yaml:"time-format"`
}

var defaultPro = &ServerProperties{
	Bind:           "0.0.0.0",
	Port:           6399,
	AppendOnly:     false,
	AppendFilename: "",
	MaxClients:     1000,
}

func Setup(settings *Settings) {
	var err error
	dir := settings.Path
	fileName := fmt.Sprintf("%s-%s.%s",
		settings.Name,
		time.Now().Format(settings.TimeFormat),
		settings.Ext)

	logFile, err := mustOpen(fileName, dir)
	if err != nil {
		log.Fatalf("logging.Setup err: %s", err)
	}

	mw := io.MultiWriter(os.Stdout, logFile)
	logger = log.New(mw, defaultPrefix, flags)
}

func fileExits(filename string) bool {
	info, err := os.Stat(filename)
	return err == nil && !info.IsDir()
}
func parse(src io.Reader) *ServerProperties {
	return defaultPro
}
func SetConfig(configName string) {
	file, err := os.Open(configName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	parse(file)
}
func init() {
	logger = log.New(os.Stdout, defaultPrefix, flags)
}
func main() {
	var banner = `
   ______          ___
  / ____/___  ____/ (_)____
 / / __/ __ \/ __  / / ___/
/ /_/ / /_/ / /_/ / (__  )
\____/\____/\__,_/_/____/
`
	print(banner)
	Setup(&Settings{
		Path:       "logs",
		Name:       "godis",
		Ext:        "log",
		TimeFormat: "2021-01-06",
	})
	configFile := os.Getenv("CONFIG")
	if configFile == "" {
		if fileExits("redis.conf") {
			SetConfig("redis.conf")
		} else {
			Properties = defaultPro
		}
	} else {
		SetConfig(configFile)
	}
	err := ListenAndServerWithSignal(&Config{
		Address: fmt.Sprintf("%s:%d", Properties.Bind, Properties.Port), 
	})
}
