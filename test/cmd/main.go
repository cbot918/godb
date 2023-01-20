package main

import "godb/test/httpService"

func main() {
	httpService.New().Run(":2266")
}
