package catch

import "github.com/tnuanchuay/honeypot/mysql"

type Catch struct{}

func CreateTable() {
	err := mysql.Execute(`
CREATE TABLE IF NOT EXISTS CATCH (
    id 			INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    pot_id 		INT 			NOT NULL,
    remote_ip	VARCHAR(12)		NOT NULL,
    create_date	TIMESTAMP		NOT NULL
);
`)

	if err != nil {
		panic(err)
	}
}
