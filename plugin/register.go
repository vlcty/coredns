package plugin

import "github.com/caddyserver/caddy"

// Register registers your plugin with CoreDNS and allows it to be called when the server is running.
func Register(name string, action SetupFunc) {
	caddy.RegisterPlugin(name, caddy.Plugin{
		ServerType: "dns",
		Action:     func(c *caddy.Controller) error { return action(c) },
	})
}

// Controller is given to the setup function of directives which gives them
// access to be able to read tokens with which to configure themselves. It also
// stores state for the setup functions, can get the current context, and can
// be used to identify a particular server block using the Key field.
type Controller struct {
	*caddy.Controller
}

// SetupFunc is used to set up a plugin, or in other words, execute a
// directive. It will be called once per key for each server block it appears
// in.
type SetupFunc func(c *Controller) error
