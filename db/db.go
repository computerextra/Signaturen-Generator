package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"strings"
	"time"
)

type Database struct {
	db *sql.DB
}

type Employee struct {
	Id        string
	Name      string
	Mail      string
	Short     string
	Tags      []string
	Focus     []string
	Abteilung string
}

func NewDatabase(connectionString string) (*Database, error) {
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return &Database{
		db: db,
	}, nil
}

func (db Database) GetEmployees() ([]Employee, error) {
	var employees []Employee
	rows, err := db.db.Query("SELECT Mitarbeiter.id, Mitarbeiter.name, Mitarbeiter.short, Mitarbeiter.tags, Mitarbeiter.focus, Abteilung.name FROM Mitarbeiter RIGHT JOIN Abteilung ON Mitarbeiter.abteilungId = Abteilung.id;")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var employee Employee
		var Id sql.NullString
		var Name sql.NullString
		var Short sql.NullString
		var Tags sql.NullString
		var Focus sql.NullString
		var Abteilung sql.NullString
		err = rows.Scan(&Id, &Name, &Short, &Tags, &Focus, &Abteilung)
		if err != nil {
			return nil, err
		}
		if Id.Valid {
			employee.Id = Id.String
		}
		if Name.Valid {
			employee.Name = Name.String
			s := strings.Split(Name.String, " ")
			employee.Mail = fmt.Sprintf("%s.%s@%s", s[0], s[1], os.Getenv("COMPANY_WEBSITE"))
		}
		if Short.Valid {
			employee.Short = Short.String
		}
		if Tags.Valid {
			employee.Tags = strings.Split(Tags.String, ",")
		}
		if Focus.Valid {
			employee.Focus = strings.Split(Focus.String, ",")
		}
		if Abteilung.Valid {
			employee.Abteilung = Abteilung.String
		}
		employees = append(employees, employee)
	}

	err = rows.Close()
	if err != nil {
		return nil, err
	}
	err = db.db.Close()
	if err != nil {
		return nil, err
	}
	return employees, nil
}
