package greet

import "fmt"

func Greet(name string) string {
    if name == "" {
        name = "world"
    }
    return fmt.Sprintf("Hello, %s!", name)
}
