package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type TCDICTIONARY1 struct {
	Id       int    `orm:"column(DICTIONARY1_ID);auto" description:"DICTIONARY1_ID"`
	CATEGORY string `orm:"column(CATEGORY);size(30);null" description:"种类"`
	VALUE    string `orm:"column(VALUE);size(30);null" description:"VALUE"`
}

func (t *TCDICTIONARY1) TableName() string {
	return "TC_DICTIONARY1"
}

func init() {
	orm.RegisterModel(new(TCDICTIONARY1))
}

// AddTCDICTIONARY1 insert a new TCDICTIONARY1 into database and returns
// last inserted Id on success.
func AddTCDICTIONARY1(m *TCDICTIONARY1) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetTCDICTIONARY1ById retrieves TCDICTIONARY1 by Id. Returns error if
// Id doesn't exist
func GetTCDICTIONARY1ById(id int) (v *TCDICTIONARY1, err error) {
	o := orm.NewOrm()
	v = &TCDICTIONARY1{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllTCDICTIONARY1 retrieves all TCDICTIONARY1 matches certain condition. Returns empty list if
// no records exist
func GetAllTCDICTIONARY1(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(TCDICTIONARY1))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []TCDICTIONARY1
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateTCDICTIONARY1 updates TCDICTIONARY1 by Id and returns error if
// the record to be updated doesn't exist
func UpdateTCDICTIONARY1ById(m *TCDICTIONARY1) (err error) {
	o := orm.NewOrm()
	v := TCDICTIONARY1{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteTCDICTIONARY1 deletes TCDICTIONARY1 by Id and returns error if
// the record to be deleted doesn't exist
func DeleteTCDICTIONARY1(id int) (err error) {
	o := orm.NewOrm()
	v := TCDICTIONARY1{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&TCDICTIONARY1{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
