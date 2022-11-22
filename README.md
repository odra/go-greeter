# Greeter Plugin

This is a sample go plugin project to demonstrate how  a plugin based project is structured.

Upstream golang plugin documentation: https://pkg.go.dev/plugin

## Structure

* `pkg`: contains the plugin library that loads and runs a loded pluign;
* `contrib/plugins`: contains plugin code, this folder should contain N plugins to be loaded an used
* `cmd`: the cli that loads and runs plugins based on cli args.

## How it works

The plugin cli (greet) receives an agument of which language to greet the user, the language name
is used to load a plugin for that language and run the `Greet` method from that plugin.

Plugins are built as `.so` files and loaded with the `plugin` package.

All plugins should implement the `Greeter` interface and have a `Build` function that creates a new instance
of a struct that implements that interface.

## Building and Running

You can run `make build` and everything will be built within the `build` folder.

That will build all plugins, plugin library and cli.

Run it by executing `./build/greet $LANG` available values for `$LANG` are:

* en
* ptbr

The program will panic in case other values are used since a plugin file won't be found for other langs/values.

## LICENSE

[MIT](./LICENSE)
