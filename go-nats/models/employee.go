package models

type Employee struct {
	EmpId    int     `json:"empid"`
	EmpFName string  `json:"empfname"`
	EmpLName string  `json:"emplname"`
	Salary   float32 `json:"salary"`
}
