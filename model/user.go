package model

import (
	"bytes"
	"cst_im/util"
	opt "cst_im/util"
	"errors"
	"reflect"
	"strconv"
)

//InsertUser returns a new user id and error if it does
func InsertUser(user *User) (int64, error) {
	if user == nil {
		return 0, errors.New("the user is nil")
	}
	SQL := "insert into userBrief(name,pwd,nick,iconName,iconPath) values(?,?,?,?,?);"
	stmt, err := g_db.Prepare(SQL)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	iconPath := util.GetIconPath(user.Icon)
	res, err := stmt.Exec(user.UesrName, user.UserPwd, user.NickName, user.IconName, iconPath)
	if err != nil {
		return 0, err
	}
	if effected, err := res.RowsAffected(); int(effected) <= 0 || err != nil {
		return 0, err
	}
	lastId, err := res.LastInsertId()
	return lastId, err
}

//添加用户详情
func AddUserDetail(detail *UserDetail) (bool, error) {
	if detail == nil {
		return false, errors.New("the user detail info is nil")
	}
	SQL := `insert into userDetail(UID,address,sex,age) values(?,?,?,?);`
	stmt, err := g_db.Prepare(SQL)
	if err != nil {
		return false, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(detail.UID, detail.Addr, detail.Sex, detail.Age)
	if err != nil {
		return false, nil
	}
	if effected, err := res.RowsAffected(); int(effected) <= 0 || err != nil {
		return false, err
	}
	return true, nil
}
func AddFriend(ownerID, friendID int32) (bool, error) {
	if ownerID == 0 || friendID == 0 {
		return false, errors.New("the arg is invalid")
	}
	SQL := `insert into friendship(ownerID,friendID) values(?,?);`
	stmt, err := g_db.Prepare(SQL)
	if err != nil {
		return false, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(ownerID, friendID)
	if err != nil {
		return false, err
	}
	effected, _ := res.RowsAffected()
	if int(effected) > 0 {
		return true, nil
	}
	return false, nil
}

func SearchUserWithUserNamne(info *SearchInfo) (*User, error) {
	if info == nil {
		return nil, errors.New("the search info is nil")
	}
	if info.SearchType != opt.SEEK_NAME {
		return nil, errors.New("the option type is invalid")
	}
	SQL := `select UID,name,nick,iconName,iconPath,sex,age,addr from userBriefView where name = ?`
	stmt, err := g_db.Prepare(SQL)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	res, err := stmt.Query(info.SrchName)
	if err != nil {
		return nil, err
	}
	defer res.Close()
	var user *User
	iconPath := ""
	if res.Next() {
		user = new(User)
		detail := new(UserDetail)
		res.Scan(&(user.UserID), &(user.UesrName), &(user.NickName), &(user.IconName), &iconPath, &(detail.Sex), &(detail.Age), &(detail.Addr))
		user.UserDetail = detail
		if buf, err := util.GetImageBytes(iconPath); err == nil {
			user.Icon = make([]byte, len(buf))
			copy(user.Icon, buf)
		}
	}
	return user, nil
}

//SearchUsers will return the slice of *User, or nil when there is 
//no result for searching
func SearchUsers(info *SearchInfo) ([]*User, error) {
	if info == nil {
		return nil, errors.New("the search info is nil")
	}
	SQL, err := getSQL(info)
	if err != nil {
		return nil, err
	}
	stmt, err := g_db.Prepare(SQL)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	res, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer res.Close()
	users := make([]*User, 0)
	iconPath := ""
	for res.Next() {
		user := new(User)
		detail := new(UserDetail)
		res.Scan(&(user.UserID), &(user.UesrName), &(user.NickName), &(user.IconName), &iconPath, &(detail.Sex), &(detail.Age), &(detail.Addr))
		user.UserDetail = detail
		if buf, err := util.GetImageBytes(iconPath); err == nil {
			user.Icon = make([]byte, len(buf))
			copy(user.Icon, buf)
		}
		users = append(users, user)
	}
	return users, nil
}

//getSQL returns a SQL accoding to the search info
func getSQL(info *SearchInfo) (string, error) {
	if info == nil {
		return "", errors.New("the searchInfo is nil")
	}
	var SQL bytes.Buffer
	//if searching for group
	if info.SearchType == opt.SEEK_GROUP {
		SQL.WriteString("SELECT groupID,nick,groupIntro,createTime,rank FROM IMgroup where groupID= ")
		SQL.WriteString(strconv.Itoa(int(info.GroupNO)))
		return SQL.String(), nil
	}

	SQL.WriteString("select UID,name,nick,iconName,iconPath,sex,age,addr from userBriefView where ")
	if info.SearchType == opt.SEEK_NAME { // if search the user with te username
		SQL.WriteString("name=")
		SQL.WriteString(info.SrchName)
		SQL.WriteString(";")
		return SQL.String(), nil
	}
	
	v := reflect.ValueOf(info).Elem()
	typeOfInfo := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)                   //field
		fieldName := typeOfInfo.Field(i).Name //the name of field
		switch field.Type().Kind() {
		case reflect.Int32:
			if fieldName == "SrchType" {
				continue
			}
			value := field.Interface().(int32)
			if value == 0 {
				continue
			}
			sqlColName, exists := g_struName_sqlName[fieldName]
			if !exists {
				continue
			}
			SQL.WriteString(sqlColName)
			SQL.WriteString(strconv.Itoa(int(value)))
			SQL.WriteString(" and ")
		case reflect.Bool:
			value := field.Interface().(bool)
			if !value {
				continue
			}
			SQL.WriteString(g_struName_sqlName[fieldName])
			SQL.WriteString(" and ")
		case reflect.String:
			if fieldName != "SrchAttrb" && fieldName != "Address" {
				continue
			}
			value := field.Interface().(string)
			if len(value) <= 0 {
				continue
			}
			SQL.WriteString(g_struName_sqlName[fieldName])
			if fieldName == "SrchAttrb" {
				SQL.WriteString("'%" + value + "%') and ")
			} else {
				SQL.WriteString("'%" + value + "%' and ")
			}

		}
	}
	sql := (SQL.String())[:len(SQL.String())-len(" and ")]
	SQL.Reset()
	SQL.WriteString(sql)
	SQL.WriteString(" limit ")
	SQL.WriteString(strconv.Itoa(int(info.SinceId)))
	SQL.WriteString(",15;")
	return SQL.String(), nil
}
