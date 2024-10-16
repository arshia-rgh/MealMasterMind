package main

var serviceQueues = map[string][]string{
	"auth": {
		"auth_login",
		"auth_register",
		"auth_reset_password",
		"auth_change_password",
		"auth_delete",
		"auth_update",
	},
	// Other services
}
