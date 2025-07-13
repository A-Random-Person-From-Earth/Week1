package config

import (
    "flag"
    "fmt"
    "os"
    "time"
    "gopkg.in/yaml.v2"
    "github.com/A-Random-Person-From-Earth/go-camp/internal/server"  // Replace with your actual module name
)

// FileConfig represents the YAML file structure
type FileConfig struct {
    Port     string `yaml:"port"`
    Greeting string `yaml:"greeting"`
    Timeout  string `yaml:"timeout"`
}

// Load handles the complete configuration loading process
// Priority: CLI flags > Environment variables > YAML file > Hardcoded defaults
func Load() (server.Config, error) {
    // Step 1: Load YAML config file (base layer)
    fileConfig := loadConfigFile("config.yaml")
    
    // Step 2: Apply environment variables on top of file config
    defaultPort := getEnvOrDefault("PORT", fileConfig.Port)
    defaultGreeting := getEnvOrDefault("GREETING", fileConfig.Greeting)
    defaultTimeout := getEnvOrDefault("TIMEOUT", fileConfig.Timeout)
    
    // Step 3: Parse timeout string to duration
    timeoutDuration, err := time.ParseDuration(defaultTimeout)
    if err != nil {
        fmt.Printf("Invalid timeout format: %v, using 30s\n", err)
        timeoutDuration = 30 * time.Second
    }
    
    // Step 4: Define and parse command line flags (highest priority)
    port := flag.String("port", defaultPort, "Server port")
    greeting := flag.String("greeting", defaultGreeting, "Default greeting for unnamed users")
    timeout := flag.Duration("timeout", timeoutDuration, "Server timeout")
    
    flag.Parse()
    
    // Step 5: Return final configuration
    return server.Config{
        Port:     *port,
        Greeting: *greeting,
        Timeout:  *timeout,
    }, nil
}

// loadConfigFile loads configuration from YAML file
func loadConfigFile(filename string) FileConfig {
    // Set hardcoded defaults
    defaults := FileConfig{
        Port:     ":8080",
        Greeting: "world",
        Timeout:  "30s",
    }
    
    // Try to read file
    data, err := os.ReadFile(filename)
    if err != nil {
        fmt.Printf("Config file not found (%s), using defaults\n", filename)
        return defaults
    }
    
    // Try to parse YAML
    var config FileConfig
    err = yaml.Unmarshal(data, &config)
    if err != nil {
        fmt.Printf("Error parsing config file: %v, using defaults\n", err)
        return defaults
    }
    
    fmt.Printf("Loaded config from %s\n", filename)
    return config
}

// getEnvOrDefault gets environment variable or returns default value
func getEnvOrDefault(envVar, defaultValue string) string {
    if value := os.Getenv(envVar); value != "" {
        return value
    }
    return defaultValue
}
