package model

import (
	//	"fmt"
	"os"
	"testing"
)

func TestGetSQL(t *testing.T) {
	s := &SearchInfo{}
	s.AgeHigh = 19
	//	//////s.SearchType = 39
	//	s.SrchName = "attr"
	s.SelectMale = true
	s.Address = "广州"
	//	s.SrchAttrb = "篮球"
	sql, _ := getSQL(s)
	t.Log(sql)
}

func Te5stInserUser(t *testing.T) {
	//--------------------------1
	file, err := os.Open("photos/t1.jpg")
	if err != nil {
		t.Error(err)
	}
	fileBuf := make([]byte, 0)
	buf := make([]byte, 1024*4)
	for {
		n, _ := file.Read(buf)
		if n <= 0 {
			break
		}
		fileBuf = append(fileBuf, buf[:n]...)
	}
	t.Log(fileBuf)
	file.Close()
	user := NewUser("zhangsan", "password", "zs", "t1.jpg", fileBuf)
	id, err := InsertUser(user)
	if err != nil {
		t.Error(err)
	}
	detail := NewDetail(int32(id), 19, "广东广州番禺区", "男")
	_, err = AddUserDetail(detail)
	if err != nil {
		t.Error(err)
	}
	//----------------2
	file, err = os.Open("photos/t2.jpg")
	if err != nil {
		t.Error(err)
	}
	fileBuf = make([]byte, 0)
	for {
		n, _ := file.Read(buf)
		if n <= 0 {
			break
		}
		fileBuf = append(fileBuf, buf[:n]...)
	}
	file.Close()
	user = NewUser("lisi", "password", "ls", "t2.jpg", fileBuf)
	id, err = InsertUser(user)
	if err != nil {
		t.Error(err)
	}
	detail = NewDetail(int32(id), 17, "广东广州天河区", "男")
	_, err = AddUserDetail(detail)
	if err != nil {
		t.Error(err)
	}
	//---------------------------3
	file, err = os.Open("photos/timg.jpg")
	if err != nil {
		t.Error(err)
	}
	fileBuf = make([]byte, 0)
	for {
		n, _ := file.Read(buf)
		if n <= 0 {
			break
		}
		fileBuf = append(fileBuf, buf[:n]...)
	}
	file.Close()
	user = NewUser("wanwu", "password", "wu", "timg.jpg", fileBuf)
	id, err = InsertUser(user)
	if err != nil {
		t.Error(err)
	}
	detail = NewDetail(int32(id), 23, "广东广州天河区", "男")
	_, err = AddUserDetail(detail)
	if err != nil {
		t.Error(err)
	}
	//---------------------------------4
	file, err = os.Open("photos/hz.jpg")
	if err != nil {
		t.Error(err)
	}
	fileBuf = make([]byte, 0)
	for {
		n, _ := file.Read(buf)
		if n <= 0 {
			break
		}
		fileBuf = append(fileBuf, buf[:n]...)
	}
	file.Close()
	user = NewUser("xiaozehao", "password", "hz", "hz.jpg", fileBuf)
	id, err = InsertUser(user)
	if err != nil {
		t.Error(err)
	}
	detail = NewDetail(int32(id), 20, "广东广州番禺区", "男")
	_, err = AddUserDetail(detail)
	if err != nil {
		t.Error(err)
	}

}
func Te7stSelectUser(t *testing.T) {
	info := &SearchInfo{
		SearchType: 39,
		SrchName:   "zhangsan",
	}
	user, err := SearchUserWithUserNamne(info)
	if err != nil {
		t.Error(err)
	}
	t.Log(*user)
	t.Log(user.UserDetail.Sex)
}
func NewUser(name, pwd, nick, iconName string, icon []byte) *User {
	user := &User{
		UesrName: name,
		UserPwd:  pwd,
		NickName: nick,
		IconName: iconName,
	}
	user.Icon = make([]byte, len(icon))
	copy(user.Icon, icon)
	return user
}

func NewDetail(id, age int32, addr, sex string) *UserDetail {
	return &UserDetail{
		UID:  id,
		Addr: addr,
		Sex:  sex,
		Age:  age,
	}
}
