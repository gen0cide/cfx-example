package foopkg

import "fmt"

// NewFoo is a constructor that requires a populated config and returns a Foo.
func NewFoo(cfg *Config) *Foo {
	return &Foo{
		Config: cfg,
	}
}

// Foo is an example container that holds a configuration.
type Foo struct {
	Config *Config
}

// PrintInfo is used to show the contents of the populated configuration.
func (f *Foo) PrintInfo() {
	fmt.Printf("[FOO] Name = %v\n", f.Config.Name)
	fmt.Printf("[FOO] URL = %v\n", f.Config.URL)
	fmt.Printf("[FOO] Age = %v\n", f.Config.Age)
	fmt.Printf("[FOO] Tags = %v\n", f.Config.Tags)
}
