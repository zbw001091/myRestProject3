package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type TRCANDIDATEUSER struct {
	Id              int           `orm:"column(CU_ID);auto" description:"序列号"`
	QUALITYSCORE    int           `orm:"column(QUALITY_SCORE);null" description:"交付物质量得分"`
	EVALUATESCORE   int           `orm:"column(EVALUATE_SCORE);null" description:"BA能力评估得分"`
	CAPABILITYSCORE int           `orm:"column(CAPABILITY_SCORE);null" description:"行为能力得分"`
	BASKILLSCORE    int           `orm:"column(BA_SKILL_SCORE);null" description:"BA技能得分"`
	BATASKSSCORE    int           `orm:"column(BATASKS_SCORE);null" description:"BATasks得分"`
	MANAGESCORE     int           `orm:"column(MANAGE_SCORE);null" description:"PM管理技术能力得分"`
	PROCESSSCORE    int           `orm:"column(PROCESS_SCORE);null" description:"过程管理与监控能力得分"`
	SCORESTATE      int           `orm:"column(SCORE_STATE);null" description:"分数有效状态"`
	USID            *TMSYSTEMUSER `orm:"column(US_ID);rel(fk)" description:"用户-评委序列号"`
	CAID            *TMCANDIDATE  `orm:"column(CA_ID);rel(fk)" description:"参评人序列号"`
}

func (t *TRCANDIDATEUSER) TableName() string {
	return "TR_CANDIDATE_USER"
}

func init() {
	orm.RegisterModel(new(TRCANDIDATEUSER))
}

// AddTRCANDIDATEUSER insert a new TRCANDIDATEUSER into database and returns
// last inserted Id on success.
func AddTRCANDIDATEUSER(m *TRCANDIDATEUSER) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetTRCANDIDATEUSERById retrieves TRCANDIDATEUSER by Id. Returns error if
// Id doesn't exist
func GetTRCANDIDATEUSERById(id int) (v *TRCANDIDATEUSER, err error) {
	o := orm.NewOrm()
	v = &TRCANDIDATEUSER{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllTRCANDIDATEUSER retrieves all TRCANDIDATEUSER matches certain condition. Returns empty list if
// no records exist
func GetAllTRCANDIDATEUSER(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(TRCANDIDATEUSER))
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

	var l []TRCANDIDATEUSER
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

// UpdateTRCANDIDATEUSER updates TRCANDIDATEUSER by Id and returns error if
// the record to be updated doesn't exist
func UpdateTRCANDIDATEUSERById(m *TRCANDIDATEUSER) (err error) {
	o := orm.NewOrm()
	v := TRCANDIDATEUSER{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteTRCANDIDATEUSER deletes TRCANDIDATEUSER by Id and returns error if
// the record to be deleted doesn't exist
func DeleteTRCANDIDATEUSER(id int) (err error) {
	o := orm.NewOrm()
	v := TRCANDIDATEUSER{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&TRCANDIDATEUSER{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
