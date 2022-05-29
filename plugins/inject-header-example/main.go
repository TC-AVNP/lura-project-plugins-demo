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
var ModifierRegisterer = registerer("inject-header-example")

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
	f(string(r)+"-top-secret", r.topSecret, true, false)

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

func (r registerer) topSecret(
	cfg map[string]interface{},
) func(interface{}) (interface{}, error) {
	// get required authorization from config
	authorization := cfg[string(r)+"-top-secret"].(map[string]interface{})["authorization"].(string)

	fmt.Println("[HANDLER INJECTION]: top-secret")
	return func(input interface{}) (interface{}, error) {
		req := input.(RequestWrapper)
		headersRef := req.Headers()
		headersRef["Authorization"] = []string{authorization}

		fmt.Println("Top secret injection!")
		return input, nil
	}
}
