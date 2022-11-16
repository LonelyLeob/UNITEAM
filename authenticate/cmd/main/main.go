package main

import (
	"github.com/L0nelyleob/UNITEAM/oauth-service/internal"
	"github.com/ilyakaznacheev/cleanenv"
)

func main() {
	var cr internal.ConfigRemote
	var cd internal.ConfigDatabase
	cleanenv.ReadEnv(&cr)
	cleanenv.ReadEnv(&cd)

	a := internal.Initialize(cr.GetSKey())
	a.Start(cr.GetFullWebAddr(), cd.GetFullDBAddr(), cr.CSRFKey)
}
