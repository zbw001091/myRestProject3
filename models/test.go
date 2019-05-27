package models

type Test struct {
	Idtest  int    `orm:"column(idtest);null"`
	Testcol string `orm:"column(testcol);size(45);null"`
}
