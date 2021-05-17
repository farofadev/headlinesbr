package config

import "os"

var Database = os.Getenv("MONGO_DATABASE")
