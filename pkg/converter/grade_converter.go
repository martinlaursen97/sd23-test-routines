package converter

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type GradeType int

const (
	DK GradeType = iota
	US
)

type Grade struct {
	system GradeType
}

const (
	host     = "db"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

var validDenmarkGrades = []int{12, 10, 7, 4, 2, 0, -3}
var validUSGrades = []string{"A", "B", "C", "D", "F"}

func checkIsValidGrade(value any, system GradeType) bool {
	switch system {
	case DK:
		for _, grade := range validDenmarkGrades {
			if grade == value {
				return true
			}
		}
	case US:
		for _, grade := range validUSGrades {
			if grade == value {
				return true
			}
		}
	}
	return false
}

func (g *Grade) Convert(value any, targetSystem GradeType) (any, error) {
	if g.system == targetSystem {
		return nil, fmt.Errorf("cannot convert between the same systems")
	}

	if !checkIsValidGrade(value, g.system) {
		return nil, fmt.Errorf("invalid grade")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	log.Println("Connecting to the database...")
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		return nil, err
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("Successfully connected to the database!")

	var query string
	if g.system == DK {
		query = `SELECT usa FROM grades WHERE denmark = $1`
	} else {
		query = `SELECT denmark FROM grades WHERE usa = $1`
	}

	if err := db.QueryRow(query, value).Scan(&value); err != nil {
		return nil, err
	}

	return value, nil
}

func NewGrade(system GradeType) *Grade {
	return &Grade{
		system: system,
	}
}
