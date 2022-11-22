# Greeter Plugin

This is a sample go plugin project to demonstrate how  a plugin based project is structured.

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

## LICENSE

[MIT](./LICENSE)
