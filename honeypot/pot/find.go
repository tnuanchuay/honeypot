package pot

import (
	"github.com/tnuanchuay/honeypot/mysql"
	"time"
)

func FindPotByPath(path string) (*Pot, error) {
	rows, err := mysql.Query(`SELECT id, path, redirect_to, user, create_at FROM POT WHERE path = ?`, path)
	if err != nil {
		return nil, err
	}

	hasRow := rows.Next()
	if !hasRow {
		return nil, nil
	}

	p := Pot{}
	dt := ""

	err = rows.Scan(&p.Id, &p.Path, &p.Redirect, &p.User, &dt)
	if err != nil {
		return nil, err
	}

	p.CreateDate, err = time.Parse("2006-01-02 15:04:05", dt)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
