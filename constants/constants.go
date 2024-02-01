package constants

import "os"

var ViacepBaseUrl string = os.Getenv("VIACEP_BASE_URL")
var CorreiosBaseUrl string = os.Getenv("CORREIOS_BASE_URL")
