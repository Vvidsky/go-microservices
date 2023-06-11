package main

import (
	"database/sql"

	"github.com/pingcap-inc/tidb-example-golang/util"
)

type Employee struct {
	EmpId    int
	EmpFName string
	EmpLName string
	Salary   float32
}

func getAllEmployee(db *sql.DB) ([]Employee, error) {
	var employees []Employee

	rows, err := db.Query(GetAllEmployeeSQL)
	if err != nil {
		return employees, err
	}
	defer rows.Close()

	for rows.Next() {
		employee := Employee{}
		err = rows.Scan(&employee.EmpId, &employee.EmpFName, &employee.EmpLName, &employee.Salary)
		if err == nil {
			employees = append(employees, employee)
		} else {
			return employees, err
		}
	}

	return employees, nil
}

func UpdateEmployee(tx *util.TiDBSqlTx, e Employee) error {
	empid := e.EmpId
	updateStmt, err := tx.Prepare(UpdateEmployeeSQL)
	if err != nil {
		return err
	}
	defer updateStmt.Close()

	if _, err := updateStmt.Exec(e.EmpFName, e.EmpLName, e.Salary, empid); err != nil {
		return err
	}

	return nil
}
