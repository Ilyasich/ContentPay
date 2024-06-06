package repositories

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type SQLrepo struct {
	DB *sql.DB
}

func New(connstr string) (*SQLrepo, error) {
	conn, err := sql.Open("postgres", connstr)
	if err != nil {
		return &SQLrepo{DB: conn}, err
	}
	err = conn.Ping()
	if err != nil {
		return &SQLrepo{DB: conn}, err
	}
	return &SQLrepo{DB: conn}, nil
}

func (rep *SQLrepo) AddContent(copyrighterID int, contractAddr string, title string) error {
	row := rep.DB.QueryRow("INSERT INTO content (contractaddr,totaltondonations,totalusddonations,copyrighter,title) VALUES (" + contractAddr + ",0,0," + string(copyrighterID) + "," + title + ")")
	if row.Err() != nil {
		return row.Err()
	}
	return nil
}

func (rep *SQLrepo) DeleteContent(internalid int) error {
	row := rep.DB.QueryRow("DELETE FROM content WHERE internalid =" + fmt.Sprint(internalid))
	if row.Err() != nil {
		return row.Err()
	}
	return nil
}

func (rep *SQLrepo) AddCopyrighter(contractAddr string, email string, password string) error {
	row := rep.DB.QueryRow("INSERT INTO copyrighter (email,pass,contractaddr,totaltondonations,totalusddonations) VALUES (" + email + "," + password + "," + contractAddr + ",0,0)")
	if row.Err() != nil {
		return row.Err()
	}
	return nil
}

func (rep *SQLrepo) AddIMDB(internalid int, IMDBid string) error {
	row := rep.DB.QueryRow("INSERT INTO imdbid VALUES (" + string(internalid) + IMDBid + ")")
	if row.Err() != nil {
		return row.Err()
	}
	return nil
}

func (rep *SQLrepo) AddKinopoisk(internalid int, Kinopoiskid int) error {
	row := rep.DB.QueryRow("INSERT INTO kinopoiskid VALUES (" + string(Kinopoiskid) + "," + string(internalid) + ")")
	if row.Err() != nil {
		return row.Err()
	}
	return nil
}
