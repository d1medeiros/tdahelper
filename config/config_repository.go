package config

import (
	"tdahelper/internal/repository"
)

var Path = "/Users/diegobraga/Documents/projetos/tdahelper/data"

var Mng Manager

func init() {
	var db = DBStore{}
	db.New(Path)
	er := repository.NewEventRepository(db.GetDB())
	Mng = NewManager(er)
}
