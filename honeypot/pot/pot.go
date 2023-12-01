package pot

import (
	"errors"
	"github.com/tnuanchuay/honeypot/mysql"
	"time"
)

var (
	ErrRequiredFieldIsMissing = errors.New("required fields is missing")
)

type Pot struct {
	Id         int       `json:"id" form:"id" query:"id"`
	Path       string    `json:"path" form:"path" query:"path"`
	Redirect   string    `json:"redirect" form:"redirect" query:"redirect"`
	User       string    `json:"user" form:"user" query:"user"`
	CreateDate time.Time `json:"create_at" form:"create_date" query:"create_date"`
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
	_, err := mysql.Execute(`
CREATE TABLE IF NOT EXISTS POT (
		id 					INT 			PRIMARY KEY NOT NULL 	AUTO_INCREMENT,
		path 				VARCHAR(200)				NOT NULL,
		redirect_to 		VARCHAR(1000) 				NOT NULL,
		user	 			VARCHAR(100) 				NOT NULL,
		create_at 			TIMESTAMP 					NOT NULL
	);
`)

	if err != nil {
		panic(err)
	}
}
