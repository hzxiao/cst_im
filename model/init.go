package model

import (
	"database/sql"
	"fmt"

	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
)

var (
	g_db               *sql.DB
	cache              redis.Conn
	g_struName_sqlName map[string]string
)

func init() {
	initDB()
	initCache()

	g_struName_sqlName = make(map[string]string)
	g_struName_sqlName["SrchName"] = "name="
	g_struName_sqlName["SrchAttrb"] = "userBrief.UID in (select distinct userID from userCustomAttr where attrContent like "
	g_struName_sqlName["AgeLow"] = "age <= "
	g_struName_sqlName["AgeHigh"] = "age >= "
	g_struName_sqlName["SelectMale"] = "sex='男'"
	g_struName_sqlName["SelectFemale"] = "sex='女'"
	g_struName_sqlName["Address"] = "addr like "
}

func initCache() {
	var err error
	cache, err = redis.Dial("tcp", "localhost:6379")
	if err != nil {
		panic(err)
	}
}

func initDB() {
	rootDBPwd := "abc123"
	connStr := "root:" + rootDBPwd + "@/mysql?charset=utf8&loc=Local&parseTime=true"
	var err error
	g_db, err = sql.Open("mysql", connStr)
	if err != nil {
		panic(err)
	}
	err = g_db.Ping()
	if err != nil {
		panic(err)
	}
	cr_db := "CREATE DATABASE IF NOT EXISTS cstdb DEFAULT CHARSET utf8;"

	stmt, err := g_db.Prepare(cr_db)
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec()
	if err != nil {
		fmt.Println(err.Error())
	}
	stmt.Close()

	grantSQL := "grant all on cstdb.* to cstAdmin identified by 'cstDb4ever';"
	stmt, err = g_db.Prepare(grantSQL)
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec()
	if err != nil {
		fmt.Println(err.Error())
	}
	stmt.Close()
	grantSQL = "grant all on cstdb.* to cstAdmin@'' identified by 'cstDb4ever';"
	stmt, err = g_db.Prepare(grantSQL)
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec()
	if err != nil {
		fmt.Println(err.Error())
	}
	stmt.Close()
	grantSQL = "grant all on cstdb.* to cstAdmin@'localhost' identified by 'cstDb4ever';"
	stmt, err = g_db.Prepare(grantSQL)
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec()
	if err != nil {
		fmt.Println(err.Error())
	}
	stmt.Close()
	grantSQL = "grant all on cstdb.* to cstAdmin@'127.0.0.1' identified by 'cstDb4ever';"
	stmt, err = g_db.Prepare(grantSQL)
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec()
	if err != nil {
		fmt.Println(err.Error())
	}
	stmt.Close()
	g_db.Close()
	dbPwd := "cstDb4ever"
	connStr = "cstAdmin:" + dbPwd + "@/cstdb?charset=utf8&loc=Local&parseTime=true"
	g_db, err = sql.Open("mysql", connStr)
	if err != nil {
		panic(err.Error())
	}
	err = g_db.Ping()
	if err != nil {
		panic(err.Error())
	}
	//创建角色类型表
	SQL := `create table IF NOT EXISTS role
		(
		rID int(2) primary key auto_increment,
		name varchar(10) not null
		);`
	stmt, err = g_db.Prepare(SQL)
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec()
	if err != nil {
		panic(err.Error())
	}
	stmt.Close()

	SQL = `create table if not exists userBrief
		(
		UID int(11) primary key auto_increment,
		name varchar(18) not null unique,
		pwd varchar(16) not null,
		nick varchar(20),
		iconName varchar(40),
		iconPath varchar(30),
		role int(2),
		constraint FK_ROLE_ID foreign key(role) references role(rID)
		)ENGINE=InnoDB DEFAULT CHARSET=utf8;`
	stmt, err = g_db.Prepare(SQL)
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec()
	if err != nil {
		panic(err.Error())
	}
	stmt.Close()

	SQL = `create table if not exists userDetail
		(
		UID int(11) not null unique,
		phone varchar(11),
		address varchar(60),
		sex varchar(2),
		age int(3),
		brithday Date,
		mail varchar(20),
		qq varchar(11),
		wechat varchar(40),
		idCard varchar(18),
		creCard varchar(16),
		debtCard varchar(19),
		stuNo varchar(15),
		realPhoto varchar(30),
		isAuth int(1),
		intro varchar(100),
		constraint FK_USER_ID foreign key(UID) references userBrief(UID)
		)ENGINE=InnoDB DEFAULT CHARSET=utf8;`
	stmt, err = g_db.Prepare(SQL)
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec()
	if err != nil {
		panic(err.Error())
	}
	stmt.Close()

	SQL = `create table if not exists groupRank
		(
		rankID int(1) primary key auto_increment,
		name varchar(10) not null
		)ENGINE=InnoDB DEFAULT CHARSET=utf8;`
	stmt, err = g_db.Prepare(SQL)
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec()
	if err != nil {
		panic(err.Error())
	}
	stmt.Close()

	SQL = `create table if not exists IMgroup
		(
		groupID int(10) primary key auto_increment,
		nick varchar(20) not null,
		createrID int(11) not null,
		createTime datetime,
		groupIntro varchar(100),
		rank int(1),
		constraint FK_RANK_ID foreign key(rank) references groupRank(rankID),
		constraint FK_CREATER_ID foreign key(createrID) references userBrief(UID)
		)ENGINE=InnoDB DEFAULT CHARSET=utf8;`
	stmt, err = g_db.Prepare(SQL)
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec()
	if err != nil {
		panic(err.Error())
	}
	stmt.Close()

	//创建用户自定义属性表
	SQL = `create table if not exists userCustomAttr 
	(
	userID int(11),
	attrName varchar(30),
	attrContent varchar(60),
	constraint FK_USER1_ID foreign key(userID) references userBrief(UID)
	)ENGINE=InnoDB DEFAULT CHARSET=utf8;`
	stmt, err = g_db.Prepare(SQL)
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec()
	if err != nil {
		panic(err.Error())
	}
	stmt.Close()
	//建立用户的视图
	SQL = `create or replace View userBriefView
		(UID,name,nick,iconName,iconPath,role,age,addr,sex,intro)
		AS
		select userBrief.UID,name,nick,iconName,iconPath,role,age,address,sex,intro
		from userBrief,userDetail
		where userBrief.UID=userDetail.UID;
		`
	stmt, err = g_db.Prepare(SQL)
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec()
	if err != nil {
		panic(err.Error())
	}
	stmt.Close()

	//建立分组表
	SQL = `create table if not exists friendLists
	(
		userID int(11),
		listNO int(2),
		listName varchar(20),
		constraint FK_USER3_ID foreign key(userID) references userBrief(UID)
		)ENGINE=InnoDB DEFAULT CHARSET=utf8;
		`
	stmt, err = g_db.Prepare(SQL)
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec()
	if err != nil {
		panic(err.Error())
	}
	stmt.Close()

	//建立好友关系表
	SQL = `create table if not exists friendship
	(
		ownerID int(11),
		friendID int(11),
		addTime timestamp,
		remark varchar(20),
		listNO int(2),
		constraint PK_FRISHIP primary key(ownerID,friendID),
		constraint FK_owner_ID foreign key(ownerID) references userBrief(UID),
		constraint FK_FRI_ID foreign key(friendID) references userBrief(UID)
		)ENGINE=InnoDB DEFAULT CHARSET=utf8;
		`
	stmt, err = g_db.Prepare(SQL)
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec()
	if err != nil {
		panic(err.Error())
	}
	stmt.Close()

}
