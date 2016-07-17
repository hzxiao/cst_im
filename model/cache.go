package model

import (
	"bytes"
	"errors"
	//	pb "cst_im/model"
	"encoding/base64"
	"encoding/gob"
	//	"fmt"

	"github.com/garyburd/redigo/redis"
)

var c redis.Conn

func init() {
	var err error
	c, err = redis.Dial("tcp", "localhost:6379")
	if err != nil {
		panic(err)
	}
}

//在redis中，存储用户的信息时，key为用户的id，value为User

//存入一个用户进redis
func AddUserIntoCache(user *User) (bool, error) {
	if user == nil || user.UserID <= 0 {
		return false, errors.New("the user is invalid")
	}
	str, err := UserToGob64(user)
	if err != nil {
		return false, err
	}
	reply, err := c.Do("SET", user.UesrName, str)
	if err != nil {
		return false, err
	}
	if (reply.(string)) == "OK" {
		return true, nil
	}
	return false, nil
}

//根据用户名获得一个用户
func GetUserFromCache(uname string) (*User, error) {
	if len(uname) <= 0 {
		return nil, errors.New("the user name is invalid")
	}
	reply, err := redis.String(c.Do("GET", uname))
	if err != nil {
		return nil, err
	}
	user, err := FromGob64ToUser(reply)
	if err != nil {
		return nil, err
	}
	return user, nil
}

//使用gob编码，把结构转化为string存进redis的value中
func UserToGob64(user *User) (string, error) {
	buf := bytes.Buffer{}
	e := gob.NewEncoder(&buf)
	err := e.Encode(user)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(buf.Bytes()), nil
}

//将gob编码转化为特定的结构体
func FromGob64ToUser(gob64 string) (*User, error) {
	u := User{}
	bys, err := base64.StdEncoding.DecodeString(gob64)
	if err != nil {
		return nil, err
	}
	buf := bytes.Buffer{}
	buf.Write(bys)
	d := gob.NewDecoder(&buf)
	err = d.Decode(&u)
	return &u, err
}
