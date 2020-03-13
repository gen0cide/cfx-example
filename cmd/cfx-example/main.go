package main

import (
	"fmt"

	"github.com/gen0cide/cfx"
	"github.com/gen0cide/cfx-example/barpkg"
	"github.com/gen0cide/cfx-example/foopkg"
	"github.com/k0kubun/pp"
	"go.uber.org/fx"
)

// Perform is used as the actual FX register.
func Perform(b *barpkg.Bar, f *foopkg.Foo) {
	b.PrintInfo()
	f.PrintInfo()

	fmt.Printf("** ENVIRONMENT **\n")

	pp.Println(b.Env)
}

func main() {
	app := fx.New(
		barpkg.Module,
		foopkg.Module,
		cfx.Module,
		fx.Invoke(Perform),
	)

	_ = app
}
