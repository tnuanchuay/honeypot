package pot

import (
	"errors"
	"github.com/tnuanchuay/honeypot/mysql"
)

var (
	ErrRequiredFieldIsMissing = errors.New("required fields is missing")
)

type Pot struct {
	Path     string `json:"path" form:"path" query:"path"`
	Redirect string `json:"redirect" form:"redirect" query:"redirect"`
	User     string `json:"user" form:"user" query:"user"`
}

func (p Pot) Validate() error {
	switch {
	case p.User == "":
		fallthrough
	case p.Redirect == "":
		fallthrough
	case p.Path == "":
		return ErrRequiredFieldIsMissing
	}

	return nil
}

func CreateTable() {
	err := mysql.Execute(`
CREATE TABLE IF NOT EXISTS POT (
		id 					INT 			PRIMARY KEY NOT NULL 	AUTO_INCREMENT,
		path 				VARCHAR(200)				NOT NULL,
		redirect_to 		VARCHAR(1000) 				NOT NULL,
		user	 			VARCHAR(100) 				NOT NULL,
		create_date 		TIMESTAMP 					NOT NULL
	);
`)

	if err != nil {
		panic(err)
	}
}
