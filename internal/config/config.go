package config

import (
    "flag"
    "fmt"
    "os"
    "time"
    "gopkg.in/yaml.v2"
)

type FileConfig struct {
    Port     string `yaml:"port"`
    Greeting string `yaml:"greeting"`
    Timeout  string `yaml:"timeout"`
}

type Config struct{
	Port   string
	Greeting  string
	Timeout time.Duration
}
func Load() (Config, error) {
    fileConfig := loadConfigFile("config.yaml")
    
    defaultPort := getEnvOrDefault("PORT", fileConfig.Port)
    defaultGreeting := getEnvOrDefault("GREETING", fileConfig.Greeting)
    defaultTimeout := getEnvOrDefault("TIMEOUT", fileConfig.Timeout)
    
    timeoutDuration, err := time.ParseDuration(defaultTimeout)
    if err != nil {
        fmt.Printf("Invalid timeout format: %v, using 30s\n", err)
        timeoutDuration = 30 * time.Second
    }
    
    port := flag.String("port", defaultPort, "Server port")
    greeting := flag.String("greeting", defaultGreeting, "Default greeting for unnamed users")
    timeout := flag.Duration("timeout", timeoutDuration, "Server timeout")
    
    flag.Parse()
    
    return Config{
        Port:     *port,
        Greeting: *greeting,
        Timeout:  *timeout,
    }, nil
}

func loadConfigFile(filename string) FileConfig {
    defaults := FileConfig{
        Port:     ":8080",
        Greeting: "world",
        Timeout:  "30s",
    }
    data, err := os.ReadFile(filename) 
    if err != nil {
        fmt.Printf("Config file not found (%s), using defaults\n", filename)
        return defaults
    }
    
    var config FileConfig
    err = yaml.Unmarshal(data, &config)
    if err != nil {
        fmt.Printf("Error parsing config file: %v, using defaults\n", err)
        return defaults
    }
    
    fmt.Printf("Loaded config from %s\n", filename)
    return config
}

func getEnvOrDefault(envVar, defaultValue string) string {
    if value := os.Getenv(envVar); value != "" {
        return value
    }
    return defaultValue
}
