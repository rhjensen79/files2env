package main

import "files2env/secrets"

func main() {
	secrets.Import("./bindings")
}
