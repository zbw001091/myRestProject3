package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type TRCANDIDATEPROJECT struct {
	Id                 int          `orm:"column(CP_ID);auto" description:"参评-项目序列号"`
	CAID               *TMCANDIDATE `orm:"column(CA_ID);rel(fk)" description:"参评人序列号"`
	PROJECTID          *TMPROJECT   `orm:"column(PROJECT_ID);rel(fk)" description:"项目编号"`
	PDTollgate         int          `orm:"column(PD_Tollgate);null" description:"P&D得分"`
	CTTollgate         int          `orm:"column(CT_Tollgate);null" description:"C&T得分"`
	PROJECTPERFORMANCE int          `orm:"column(PROJECT_PERFORMANCE);null" description:"项目绩效"`
	MANAGEPERFORMANCE  int          `orm:"column(MANAGE_PERFORMANCE);null" description:"项目经理绩效"`
	PRODUCTSCORE       int          `orm:"column(PRODUCT_SCORE);null" description:"产品、服务、成果评价得分"`
	BAPROJECTSCORE     int          `orm:"column(BA_PROJECT_SCORE);null" description:"BA项目得分"`
	PROJECTFLAG        int          `orm:"column(PROJECT_FLAG);null" description:"项目使用状态"`
}

func (t *TRCANDIDATEPROJECT) TableName() string {
	return "TR_CANDIDATE_PROJECT"
}

func init() {
	orm.RegisterModel(new(TRCANDIDATEPROJECT))
}

// AddTRCANDIDATEPROJECT insert a new TRCANDIDATEPROJECT into database and returns
// last inserted Id on success.
func AddTRCANDIDATEPROJECT(m *TRCANDIDATEPROJECT) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetTRCANDIDATEPROJECTById retrieves TRCANDIDATEPROJECT by Id. Returns error if
// Id doesn't exist
func GetTRCANDIDATEPROJECTById(id int) (v *TRCANDIDATEPROJECT, err error) {
	o := orm.NewOrm()
	v = &TRCANDIDATEPROJECT{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllTRCANDIDATEPROJECT retrieves all TRCANDIDATEPROJECT matches certain condition. Returns empty list if
// no records exist
func GetAllTRCANDIDATEPROJECT(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(TRCANDIDATEPROJECT))
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

	var l []TRCANDIDATEPROJECT
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

// UpdateTRCANDIDATEPROJECT updates TRCANDIDATEPROJECT by Id and returns error if
// the record to be updated doesn't exist
func UpdateTRCANDIDATEPROJECTById(m *TRCANDIDATEPROJECT) (err error) {
	o := orm.NewOrm()
	v := TRCANDIDATEPROJECT{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteTRCANDIDATEPROJECT deletes TRCANDIDATEPROJECT by Id and returns error if
// the record to be deleted doesn't exist
func DeleteTRCANDIDATEPROJECT(id int) (err error) {
	o := orm.NewOrm()
	v := TRCANDIDATEPROJECT{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&TRCANDIDATEPROJECT{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
