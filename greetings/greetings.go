package greetings
import (
	"fmt"
	"rsc.io/quote"
	"errors"
	"math/rand"
	"time"
)
// init sets initial values for variables used in the function
func init() {
	rand.Seed(time.Now().UnixNano())
}
// In Go, a function whose name starts with a capital letter can be called by a function not in the same package. 
// This is known in Go as an exported name. For more about exported names, see Exported names in the Go tour.
func Greeting(name string) (string, error) {
	message := ""
	if name != "" {
		//message := fmt.Sprintf("Hi, %v. Welcome!", name)
		message = fmt.Sprintf("Hello %v! %s\n", name, quote.Go())
	} else {
		message = fmt.Sprintf("Hello! %s\n", quote.Go())
	}
	return message, nil
}
func Greetings(names []string) (map[string]string, error) {
	if names == nil {
		return nil, errors.New("Invalid empty names!")
	}
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Greeting(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}
func randomFormat() string  {
	formats := []string {
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	return formats[rand.Intn(len(formats))]
}