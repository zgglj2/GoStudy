package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

func GetYAMLFromString() {
	viper.SetConfigType("yaml") // or viper.SetConfigType("YAML")

	// any approach to require this configuration into your program.
	var yamlExample = []byte(`
Hacker: true
name: steve
hobbies:
- skateboarding
- snowboarding
- go
clothing:
jacket: leather
trousers: denim
age: 35
eyes : brown
beard: true
`)

	err := viper.ReadConfig(bytes.NewBuffer(yamlExample))
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	fmt.Printf("name: %v\n", viper.Get("name"))         // this would be "steve"
	fmt.Printf("hobbies: %v\n", viper.Get("hobbies"))   // this would be []string{"skateboarding", "snowboarding", "go"}
	fmt.Printf("jacket: %v\n", viper.Get("jacket"))     // this would be "leather"
	fmt.Printf("trousers: %v\n", viper.Get("trousers")) // this would be "denim"
	fmt.Printf("age: %v\n", viper.Get("age"))           // this would be "35"
	fmt.Printf("eyes: %v\n", viper.Get("eyes"))         // this would be "brown"
	fmt.Printf("beard: %v\n", viper.Get("beard"))       // this would be "true"
	fmt.Printf("Hacker: %v\n", viper.Get("Hacker"))     // this would be "true"

}

func GetYAMLFromFile() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			fmt.Println("Config file not found")
		} else {
			// Config file was found but another error was produced
			fmt.Printf("Config file found but error: %s", err)
		}
	}
	fmt.Printf("name: %v\n", viper.Get("name"))         // this would be "steve"
	fmt.Printf("hobbies: %v\n", viper.Get("hobbies"))   // this would be []string{"skateboarding", "snowboarding", "go"}
	fmt.Printf("jacket: %v\n", viper.Get("jacket"))     // this would be "leather"
	fmt.Printf("trousers: %v\n", viper.Get("trousers")) // this would be "denim"
	fmt.Printf("age: %v\n", viper.Get("age"))           // this would be "35"
	fmt.Printf("eyes: %v\n", viper.Get("eyes"))         // this would be "brown"
	fmt.Printf("beard: %v\n", viper.Get("beard"))       // this would be "true"
	fmt.Printf("Hacker: %v\n", viper.Get("Hacker"))     // this would be "true"
}

func WriteYAMLToFile() {
	viper.SetConfigName("config2")
	viper.AddConfigPath(".")
	viper.Set("name", "steve")
	viper.Set("hobbies", []string{"skateboarding", "snowboarding", "go"})
	viper.Set("jacket", "leather")
	viper.Set("trousers", "denim")
	viper.Set("age", 35)
	viper.Set("eyes", "brown")
	viper.Set("beard", true)
	viper.Set("Hacker", true)

	viper.WriteConfig() // writes current config to predefined path set by 'viper.AddConfigPath()' and 'viper.SetConfigName'
	// viper.SafeWriteConfig()
	// viper.WriteConfigAs("/path/to/my/.config")
	// viper.SafeWriteConfigAs("/path/to/my/.config") // will error since it has already been written
	// viper.SafeWriteConfigAs("/path/to/my/.other_config")
}

func EnvExample() {
	viper.SetEnvPrefix("spf") // will be uppercased automatically
	viper.BindEnv("id")

	os.Setenv("SPF_ID", "13") // typically done outside of the app

	id := viper.Get("id") // 13
	fmt.Println(id)
}

func BindPflagsExample() {
	pflag.IntP("flagname", "f", 1234, "help message for flagname")

	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	i := viper.GetInt("flagname") // retrieve values from viper instead of pflag
	fmt.Println(i)
}

func BindPflagsExample2() {
	flag.String("username", "ggg", "help message for username")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	u := viper.GetString("username") // retrieve values from viper instead of pflag
	fmt.Println(u)
}

func WatchConfigFile() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		fmt.Println("Config file op:", e.Op)
	})
	viper.WatchConfig()
}

type myFlag struct{}

func (f myFlag) HasChanged() bool    { return false }
func (f myFlag) Name() string        { return "my-flag-name" }
func (f myFlag) ValueString() string { return "my-flag-value" }
func (f myFlag) ValueType() string   { return "string" }

func FlagInterfaceExample() {
	viper.BindFlagValue("my-flag-name", myFlag{})

	u := viper.GetString("my-flag-name")
	fmt.Println(u)
}

func GetNestedJsonKey() {
	viper.SetConfigName("config3")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			fmt.Println("Config file not found")
		} else {
			// Config file was found but another error was produced
			fmt.Printf("Config file found but error: %s", err)
		}
	}

	fmt.Printf("host: %v\n", viper.GetString("datastore.metric.host")) // "127.0.0.1"
	fmt.Printf("port: %v\n", viper.GetInt("datastore.metric.port"))    // 3099
	fmt.Printf("host port0: %v\n", viper.GetString("host.ports.0"))    // 5799
	fmt.Printf("host port1: %v\n", viper.GetString("host.ports.1"))    // 6029
}

type Cache struct {
	MaxItems int
	ItemSize int
}

func NewCache(v *viper.Viper) *Cache {
	return &Cache{
		MaxItems: v.GetInt("max-items"),
		ItemSize: v.GetInt("item-size"),
	}
}

func GetSubConfig() {
	viper.SetConfigName("config4")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			fmt.Println("Config file not found")
		} else {
			// Config file was found but another error was produced
			fmt.Printf("Config file found but error: %s", err)
		}
	}
	cache1Config := viper.Sub("cache.cache1")
	if cache1Config == nil { // Sub returns nil if the key cannot be found
		panic("cache configuration not found")
	}

	cache1 := NewCache(cache1Config)
	fmt.Printf("cache1: %v\n", cache1)
}

func UnmarshalExample() {
	v := viper.NewWithOptions(viper.KeyDelimiter("::"))

	v.SetDefault("chart::values", map[string]interface{}{
		"ingress": map[string]interface{}{
			"annotations": map[string]interface{}{
				"traefik.frontend.rule.type":                 "PathPrefix",
				"traefik.ingress.kubernetes.io/ssl-redirect": "true",
			},
		},
	})

	type config struct {
		Chart struct {
			Values map[string]interface{}
		}
	}

	var C config

	v.Unmarshal(&C)
	fmt.Printf("%+v\n", C)
}

func yamlStringSettings() string {
	c := viper.AllSettings()
	bs, err := yaml.Marshal(c)
	if err != nil {
		fmt.Printf("unable to marshal config to YAML: %v\n", err)
	}
	result := string(bs)
	fmt.Printf("%s\n", result)
	return result
}

func jsonStringSettings() string {
	c := viper.AllSettings()
	bs, err := json.Marshal(c)
	if err != nil {
		fmt.Printf("unable to marshal config to YAML: %v\n", err)
	}
	result := string(bs)
	fmt.Printf("%s\n", result)
	return result
}

type Config struct {
	Redis string
	MySQL MySQLConfig
}

type MySQLConfig struct {
	Port     int
	Host     string
	Username string
	Password string
}

func UnmarshalYamlExample() {
	var config Config
	v := viper.New()
	v.SetConfigFile("./config5.yaml")
	if err := v.ReadInConfig(); err != nil {
		fmt.Printf("unable to read config file: %v\n", err)
		return
	}
	err := v.Unmarshal(&config)
	if err != nil {
		fmt.Printf("unable to unmarshal config: %v\n", err)
		return
	}
	fmt.Printf("%+v\n", config)
}

type LogWriterConfig struct {
	LogPath     string `ini:"log_path" json:"log_path" yaml:"log_path" mapstructure:"log_path"`
	Level       string `ini:"level" json:"level" yaml:"level" mapstructure:"level"`
	MaxMegaSize int    `ini:"max_size" json:"max_size" yaml:"max_size" mapstructure:"max_size"`
	MaxBackups  int    `ini:"max_backups" json:"max_backups" yaml:"max_backups" mapstructure:"max_backups"`
	Compress    bool   `ini:"compress" json:"compress" yaml:"compress" mapstructure:"compress"`
}
type LogConfig struct {
	InfoLogWriterConfig  LogWriterConfig `ini:"infolog" json:"infolog" yaml:"infolog" mapstructure:"infolog"`
	ErrorLogWriterConfig LogWriterConfig `ini:"errorlog" json:"errorlog" yaml:"errorlog" mapstructure:"errorlog"`
}

func UnmarshalYamlExample2() {
	var config LogConfig
	v := viper.New()
	v.SetConfigFile("./config6.yaml")
	if err := v.ReadInConfig(); err != nil {
		fmt.Printf("unable to read config file: %v\n", err)
		return
	}
	err := v.Unmarshal(&config)
	if err != nil {
		fmt.Printf("unable to unmarshal config: %v\n", err)
		return
	}
	fmt.Printf("%+v\n", config)
}

func main() {
	GetYAMLFromString()
	GetYAMLFromFile()
	WriteYAMLToFile()
	EnvExample()
	BindPflagsExample()
	BindPflagsExample2()
	FlagInterfaceExample()
	GetNestedJsonKey()
	GetSubConfig()
	UnmarshalExample()
	yamlStringSettings()
	jsonStringSettings()
	WatchConfigFile()
	UnmarshalYamlExample()
	UnmarshalYamlExample2()
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT)
	<-ch
}
