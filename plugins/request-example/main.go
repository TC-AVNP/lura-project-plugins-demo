package main

import (
	"fmt"
	"io"
	"net/url"
)

func main() {}

func init() {
	fmt.Println(string(ModifierRegisterer), "loaded!!!")
}

// ModifierRegisterer is the symbol the plugin loader will be looking for. It must
// implement the plugin.Registerer interface
// https://github.com/luraproject/lura/blob/master/proxy/plugin/modifier.go#L71
var ModifierRegisterer = registerer("request-example")

type registerer string

// RegisterModifiers is the function the plugin loader will call to register the
// modifier(s) contained in the plugin using the function passed as argument.
// f will register the factoryFunc under the name and mark it as a request
// and/or response modifier.
func (r registerer) RegisterModifiers(f func(
	name string,
	factoryFunc func(map[string]interface{}) func(interface{}) (interface{}, error),
	appliesToRequest bool,
	appliesToResponse bool,
)) {
	f(string(r)+"-hello", r.hello, true, false)
	f(string(r)+"-dynamic", r.dynamic, true, false)
	f(string(r)+"-goodbye", r.goodbye, true, false)

	fmt.Printf("All modifiers from %s are registered\n", string(r))
}

// RequestWrapper is an interface for passing proxy request between the lura pipe and the loaded plugins
type RequestWrapper interface {
	Params() map[string]string
	Headers() map[string][]string
	Body() io.ReadCloser
	Method() string
	URL() *url.URL
	Query() url.Values
	Path() string
}

func (r registerer) hello(
	cfg map[string]interface{},
) func(interface{}) (interface{}, error) {
	fmt.Println("[HANDLER INJECTION]: hello")
	return func(input interface{}) (interface{}, error) {
		fmt.Println("Hello injection!")
		return input, nil
	}
}

func (r registerer) goodbye(
	cfg map[string]interface{},
) func(interface{}) (interface{}, error) {
	fmt.Println("[HANDLER INJECTION]: goodbye")
	return func(input interface{}) (interface{}, error) {
		fmt.Println("Goodbye injection!")
		return input, nil
	}
}
func (r registerer) dynamic(
	cfg map[string]interface{},
) func(interface{}) (interface{}, error) {
	// this proves we can get values dynamically from the configuration
	fooValue := cfg[string(r)+"-dynamic"].(map[string]interface{})["foo"].(string)

	fmt.Println("[HANDLER INJECTION]: dynamic")
	return func(input interface{}) (interface{}, error) {
		fmt.Printf("Dynamic injection comes with a: %s!\n", fooValue)
		return input, nil
	}
}
