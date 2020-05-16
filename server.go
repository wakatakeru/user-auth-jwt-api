package main

import "github.com/wakatakeru/user-auth-jwt-api/infrastructure"

func main() {
	infrastructure.Router.Run()
}
