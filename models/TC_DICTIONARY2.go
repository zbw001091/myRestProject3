package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type TCDICTIONARY2 struct {
	Id       int    `orm:"column(DICTIONARY2_ID);auto" description:"DICTIONARY2_ID"`
	CATEGORY string `orm:"column(CATEGORY);size(30);null" description:"种类"`
	RANK     string `orm:"column(RANK);size(30);null" description:"参评等级"`
	ITEM     string `orm:"column(ITEM);size(30);null" description:"评分项"`
	VALUE    string `orm:"column(VALUE);size(30);null" description:"VALUE"`
	TYPE     string `orm:"column(TYPE);size(30);null" description:"岗位类型"`
	SCORE    int    `orm:"column(SCORE);null" description:"得分"`
}

func (t *TCDICTIONARY2) TableName() string {
	return "TC_DICTIONARY2"
}

func init() {
	orm.RegisterModel(new(TCDICTIONARY2))
}

// AddTCDICTIONARY2 insert a new TCDICTIONARY2 into database and returns
// last inserted Id on success.
func AddTCDICTIONARY2(m *TCDICTIONARY2) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetTCDICTIONARY2ById retrieves TCDICTIONARY2 by Id. Returns error if
// Id doesn't exist
func GetTCDICTIONARY2ById(id int) (v *TCDICTIONARY2, err error) {
	o := orm.NewOrm()
	v = &TCDICTIONARY2{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllTCDICTIONARY2 retrieves all TCDICTIONARY2 matches certain condition. Returns empty list if
// no records exist
func GetAllTCDICTIONARY2(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(TCDICTIONARY2))
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

	var l []TCDICTIONARY2
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

// UpdateTCDICTIONARY2 updates TCDICTIONARY2 by Id and returns error if
// the record to be updated doesn't exist
func UpdateTCDICTIONARY2ById(m *TCDICTIONARY2) (err error) {
	o := orm.NewOrm()
	v := TCDICTIONARY2{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteTCDICTIONARY2 deletes TCDICTIONARY2 by Id and returns error if
// the record to be deleted doesn't exist
func DeleteTCDICTIONARY2(id int) (err error) {
	o := orm.NewOrm()
	v := TCDICTIONARY2{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&TCDICTIONARY2{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
