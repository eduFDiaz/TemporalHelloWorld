package app

import (
	"fmt"
)

func ComposeGreeting(name string) (string, error) {
	greeting := fmt.Sprintf("Hello %s!", name)
	return greeting, nil
}

func ComposeGreeting2(name string) (string, error) {
	greeting := fmt.Sprintf("Hello2 %s!", name)
	return greeting, nil
}
