package config

import (
	"calorie_deficit/utils"
)

// Set the SERVER_PORT either from an environment variable or default to ":3000"
// SERVER_PORT is the port on which the server will listen for incoming requests.
var SERVER_PORT = utils.GetEnv("SERVER_PORT", ":3000")
