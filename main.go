package main

import (
	http_proxy "github.com/sadn1ck/http_proxy/http_proxy"
)

func main() {
	banned := http_proxy.LoadBanned("./assets/banned.csv")
	http_proxy.Start(banned)
}
