package model

import (
    "testing"
    "cst_im/util"
    "fmt"
)

func TestCache(t *testing.T) {
    //get a user from db
    s := new(SearchInfo)
    s.SearchType = util.SEEK_NAME
    s.SrchName = "zhangsan"
    user,err := SearchUserWithUserNamne(s)
    if err != nil {
        t.Error(err)
    }
    //add the user to cache
    if saved,err := AddUserIntoCache(user); saved && err != nil {
        t.Error(err)
    }
    
    //get a user from cache
    user,err = GetUserFromCache("zhangsan")
    if err != nil {
        t.Error(err)
    }
    fmt.Println(user)
}