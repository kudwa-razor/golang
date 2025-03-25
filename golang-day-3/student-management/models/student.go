package models

// student represents  a student record.
type Student struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	DOB      string `json:"dob"`
	Address  string `json:"address"`
	Subject  string `json:"subject"`
	Marks    int    `json:"marks"`
}
