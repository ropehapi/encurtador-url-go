package model

import (
	"fmt"

	"github.com/ropehapi/encurtador-url-go/config"
)

type Relation struct {
	Url, Code string
}

func (r *Relation) Store() (err error) {
	db := config.GetConexao()
	defer db.Close()

	_, err = db.Exec(fmt.Sprintf("INSERT INTO relation (url, code) VALUES ('%s','%s')", r.Url, r.Code))
	if err != nil {
		return err
	}
	return
}