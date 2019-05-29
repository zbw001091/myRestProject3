package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type TMCANDIDATE struct {
	Id             int            `orm:"column(CA_ID);auto" description:"序列号"`
	CANDIDATEID    int            `orm:"column(CANDIDATE_ID);null" description:"参评人工号"`
	CANDIDATENAME  string         `orm:"column(CANDIDATE_NAME);size(30);null" description:"参评人姓名"`
	CANDIDATEDATE  time.Time      `orm:"column(CANDIDATE_DATE);type(date);null" description:"参评年份"`
	BAABILITY      int            `orm:"column(BA_ABILITY);null" description:"BA能力评估"`
	CANDIDATEGRADE int            `orm:"column(CANDIDATE_GRADE);null" description:"交付物质量排名"`
	PMPSCORE       int            `orm:"column(PMP_SCORE);null" description:"PMP认证加分"`
	STATES         int            `orm:"column(STATES);null" description:"参评状态"`
	CANDIDATETYPE  *TCDICTIONARY1 `orm:"column(CANDIDATE_TYPE);rel(fk)" description:"岗位类型"`
	CANDIDATEDEPT  *TCDICTIONARY1 `orm:"column(CANDIDATE_DEPT);rel(fk)" description:"所属科室"`
	CANDIDATESITE  *TCDICTIONARY1 `orm:"column(CANDIDATE_SITE);rel(fk)" description:"所属基地"`
	CANDIDATERANK  *TCDICTIONARY1 `orm:"column(CANDIDATE_RANK);rel(fk)" description:"参评等级"`
}

type candidateinfo struct {
	CANDIDATENAME  string         `orm:"column(CANDIDATE_NAME);size(30);null" description:"参评人姓名"`
	CANDIDATETYPE  string		  `orm:"column(value);size(30);null" description:"岗位类型"`
}

func (t *TMCANDIDATE) TableName() string {
	return "TM_CANDIDATE"
}

func init() {
	orm.RegisterModel(new(TMCANDIDATE))
}

// AddTMCANDIDATE insert a new TMCANDIDATE into database and returns
// last inserted Id on success.
func AddTMCANDIDATE(m *TMCANDIDATE) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// Raw SQL方式select // by zbw
func GetTMCANDIDATEWithType() (list []candidateinfo, err error) {
	var num int64
	o := orm.NewOrm()
	candidates := make([]candidateinfo, 20)
	sqlselect := "select a.CANDIDATE_NAME,b.value from isbprs.TM_CANDIDATE a,isbprs.TC_DICTIONARY1 b where a.candidate_type=b.dictionary1_id;"
	num, err = o.Raw(sqlselect).QueryRows(&candidates)
	if err != nil {
		fmt.Println("查询CANDIDATE信息出错")
	} else {
		fmt.Printf("查询CANDIDATE信息成功,共:%d条\n", num)
		return candidates, err
	}
	return nil, err
}

// Raw SQL方式insert // by zbw
func AddTMCANDIDATEWithRawSql() (id int64, err error) {
	o := orm.NewOrm()
	sqlinsert := "insert into isbprs.TM_CANDIDATE(CA_ID,CANDIDATE_TYPE,CANDIDATE_DEPT,CANDIDATE_SITE,CANDIDATE_RANK) values (24,35,13,3,7);" //需要修改第一个变量24，唯一主键
	_, err = o.Raw(sqlinsert).Exec()
	if err != nil {
		fmt.Println("新增CANDIDATE信息出错")
	} else {
		fmt.Println("新增CANDIDATE信息成功")
	}
	return
}

// GetTMCANDIDATEById retrieves TMCANDIDATE by Id. Returns error if
// Id doesn't exist
func GetTMCANDIDATEById(id int) (v *TMCANDIDATE, err error) {
	o := orm.NewOrm()
	v = &TMCANDIDATE{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllTMCANDIDATE retrieves all TMCANDIDATE matches certain condition. Returns empty list if
// no records exist
func GetAllTMCANDIDATE(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(TMCANDIDATE))
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

	var l []TMCANDIDATE
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

// UpdateTMCANDIDATE updates TMCANDIDATE by Id and returns error if
// the record to be updated doesn't exist
func UpdateTMCANDIDATEById(m *TMCANDIDATE) (err error) {
	o := orm.NewOrm()
	v := TMCANDIDATE{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteTMCANDIDATE deletes TMCANDIDATE by Id and returns error if
// the record to be deleted doesn't exist
func DeleteTMCANDIDATE(id int) (err error) {
	o := orm.NewOrm()
	v := TMCANDIDATE{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&TMCANDIDATE{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
