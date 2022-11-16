package main

import (
	"github.com/L0nelyleob/UNITEAM/golang-forms/internal/forms"
	"github.com/ilyakaznacheev/cleanenv"
)

func main() {
	var cr forms.ConfigRemote
	var cd forms.ConfigDatabase
	cleanenv.ReadEnv(&cr)
	cleanenv.ReadEnv(&cd)

	srv := forms.Spawn([]byte(cr.GetSKey()))
	srv.Start(cr.GetFullWebAddr(), cd.GetFullDBAddr())
}
