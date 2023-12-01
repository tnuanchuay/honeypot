package catch

import (
	"github.com/tnuanchuay/honeypot/mysql"
	"time"
)

type Catch struct {
	Id         int       `json:"id" form:"id" query:"id"`
	PotId      int       `json:"pot_id" form:"pot_id" query:"pot_id"`
	RemoteIp   string    `json:"remote_ip" form:"remote_ip" query:"remote_ip"`
	Referer    string    `json:"referrer" form:"referrer" query:"referrer"`
	CreateDate time.Time `json:"create_at" form:"create_date" query:"create_date"`
}

func CreateTable() {
	_, err := mysql.Execute(`
CREATE TABLE IF NOT EXISTS CATCH (
    id 			INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    pot_id 		INT 			NOT NULL,
    remote_ip	VARCHAR(12)		NOT NULL,
    referer 	VARCHAR(500),
    create_at	TIMESTAMP		NOT NULL
);
`)

	if err != nil {
		panic(err)
	}
}
