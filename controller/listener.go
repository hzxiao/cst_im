package controller

import (
	"cst_im/model"
	pb "cst_im/model"
	"cst_im/util"
	"fmt"

	"github.com/golang/protobuf/proto"
)

const (
	FAIL    = "fail"
	SUCCESS = "success"
)

type iListener interface {
	onProcess(msg pb.Msg)
}

type SearchListener struct {
	Conn *MyConnType
}

func (this *SearchListener) onProcess(msg pb.Msg) {
	opt := msg.UserOpt
	if opt != util.SEEK_NAME && opt != util.SEEK_AROUD && opt != util.SEEK_CONDITION && opt != util.SEEK_GROUP {
		return
	}

	newMsg := &pb.Msg{UserOpt: *proto.Int32(opt)}
	switch opt {
	case util.SEEK_GROUP:
		group, err := model.SearchGroup(msg.GetSrchInfo())
		if err != nil || group == nil {
			newMsg.OptResult = *proto.String(FAIL) //fail to search for
			this.Conn.Write(newMsg)
			return
		}
		newMsg.OptResult = *proto.String(SUCCESS)
		newMsg.Groups = make([]*model.Group, 1)
		newMsg.Groups[0] = group
		fmt.Println(*group)
		this.Conn.Write(newMsg)
	case util.SEEK_NAME:
		user, err := model.SearchUserWithUserNamne(msg.SrchInfo)
		if err != nil || user == nil {
			newMsg.OptResult = *proto.String(FAIL) //fail to search for
			this.Conn.Write(newMsg)
			return
		}
		newMsg.OptResult = *proto.String(SUCCESS)
		newMsg.Friends = make([]*model.User, 1)
		newMsg.Friends[0] = user
		this.Conn.Write(newMsg)
	case util.SEEK_CONDITION:
		users, err := model.SearchUsers(msg.GetSrchInfo())
		if err != nil || users == nil || len(users) == 0 {
			newMsg.OptResult = *proto.String(FAIL) //fail to search for
			fmt.Println(*newMsg)
			this.Conn.Write(newMsg)
			return
		}
		newMsg.OptResult = *proto.String(SUCCESS)
		newMsg.Friends = make([]*model.User, len(users))
		copy(newMsg.Friends, users)
		this.Conn.Write(newMsg)
	case util.ADD_FRIEND:
		user := msg.User
		info := msg.SrchInfo
		if user == nil || info == nil {
			newMsg.OptResult = *proto.String(FAIL) //fail to search for
			this.Conn.Write(newMsg)
			return
		}
		_, err := model.AddFriend(user.UserID, info.GroupNO)
		if err != nil {
			newMsg.OptResult = *proto.String(FAIL) //fail to search for
			this.Conn.Write(newMsg)
			return
		}
		newMsg.OptResult = *proto.String(SUCCESS)
		this.Conn.Write(newMsg)
	}
}

// }
//	newMsg := &pb.Msg{UserOpt: *proto.Int32(opt), OptResult: *proto.String("success"), ReceiveResult: *proto.String("fail")}

//	user := pb.User{}
//	user.UesrName = "肖泽壕"
//	user.NickName = "hz"
//	//user.IconName = "hz.jpg"
//	file, err := os.Open("hz.jpg")
//	if err != nil {
//		fmt.Println(err)
//	}
//	buf := make([]byte, 10240)
//	n, _ := file.Read(buf)
//	user.Icon = buf[:n]
//	detail := pb.UserDetail{}
//	detail.Sex = "男"
//	user.UserDetail = &detail
//	newMsg.Friends = append(newMsg.GetFriends(), &user)

//	user = pb.User{}
//	user.UesrName = "肖泽壕"
//	user.NickName = "hz"
//	//user.IconName = "hz.jpg"
//	file, err = os.Open("hz.jpg")
//	if err != nil {
//		fmt.Println(err)
//	}
//	buf = make([]byte, 10240)
//	n, _ = file.Read(buf)
//	user.Icon = buf[:n]
//	detail = pb.UserDetail{}
//	detail.Sex = "男"
//	user.UserDetail = &detail
//	newMsg.Friends = append(newMsg.GetFriends(), &user)

//	user = pb.User{}
//	user.UesrName = "肖泽壕"
//	user.NickName = "hz"
//	user.IconName = "/hz.jpg"
//	file, err = os.Open("e:/hz.jpg")
//	if err != nil {
//		fmt.Println(err)
//	}
//	buf = make([]byte, 102400)
//	n, _ = file.Read(buf)
//	user.Icon = buf[:n]
//	detail = pb.UserDetail{}
//	detail.Sex = "男"
//	user.UserDetail = &detail
//	newMsg.Friends = append(newMsg.GetFriends(), &user)

//	user = pb.User{}
//	user.UesrName = "肖泽壕"
//	user.NickName = "hz"
//	user.IconName = "/hz.jpg"
//	file, err = os.Open("e:/hz.jpg")
//	if err != nil {
//		fmt.Println(err)
//	}
//	buf = make([]byte, 102400)
//	n, _ = file.Read(buf)
//	user.Icon = buf[:n]
//	detail = pb.UserDetail{}
//	detail.Sex = "男"
//	user.UserDetail = &detail
//	newMsg.Friends = append(newMsg.GetFriends(), &user)

//	user = pb.User{}
//	user.UesrName = "肖泽壕"
//	user.NickName = "hz"
//	user.IconName = "/hz.jpg"
//	file, err = os.Open("e:/hz.jpg")
//	if err != nil {
//		fmt.Println(err)
//	}
//	buf = make([]byte, 102400)
//	n, _ = file.Read(buf)
//	user.Icon = buf[:n]
//	detail = pb.UserDetail{}
//	detail.Sex = "男"
//	user.UserDetail = &detail
//	newMsg.Friends = append(newMsg.GetFriends(), &user)

//	user = pb.User{}
//	user.UesrName = "肖泽壕"
//	user.NickName = "hz"
//	user.IconName = "/hz.jpg"
//	file, err = os.Open("e:/hz.jpg")
//	if err != nil {
//		fmt.Println(err)
//	}
//	buf = make([]byte, 102400)
//	n, _ = file.Read(buf)
//	user.Icon = buf[:n]
//	detail = pb.UserDetail{}
//	detail.Sex = "男"
//	user.UserDetail = &detail
//	newMsg.Friends = append(newMsg.GetFriends(), &user)

//	user = pb.User{}
//	user.UesrName = "肖泽壕"
//	user.NickName = "hz"
//	user.IconName = "/hz.jpg"
//	file, err = os.Open("e:/hz.jpg")
//	if err != nil {
//		fmt.Println(err)
//	}
//	buf = make([]byte, 102400)
//	n, _ = file.Read(buf)
//	user.Icon = buf[:n]
//	detail = pb.UserDetail{}
//	detail.Sex = "男"
//	user.UserDetail = &detail
//	newMsg.Friends = append(newMsg.GetFriends(), &user)

//	user = pb.User{}
//	user.UesrName = "肖泽壕"
//	user.NickName = "hz"
//	user.IconName = "/hz.jpg"
//	file, err = os.Open("e:/hz.jpg")
//	if err != nil {
//		fmt.Println(err)
//	}
//	buf = make([]byte, 102400)
//	n, _ = file.Read(buf)
//	user.Icon = buf[:n]
//	detail = pb.UserDetail{}
//	detail.Sex = "男"
//	user.UserDetail = &detail
//	newMsg.Friends = append(newMsg.GetFriends(), &user)

//	user = pb.User{}
//	user.UesrName = "肖泽壕"
//	user.NickName = "hz"
//	user.IconName = "/hz.jpg"
//	file, err = os.Open("e:/hz.jpg")
//	if err != nil {
//		fmt.Println(err)
//	}
//	buf = make([]byte, 102400)
//	n, _ = file.Read(buf)
//	user.Icon = buf[:n]
//	detail = pb.UserDetail{}
//	detail.Sex = "男"
//	user.UserDetail = &detail
//	newMsg.Friends = append(newMsg.GetFriends(), &user)

//	user = pb.User{}
//	user.UesrName = "肖泽壕"
//	user.NickName = "hz"
//	user.IconName = "/hz.jpg"
//	file, err = os.Open("e:/hz.jpg")
//	if err != nil {
//		fmt.Println(err)
//	}
//	buf = make([]byte, 102400)
//	n, _ = file.Read(buf)
//	user.Icon = buf[:n]
//	detail = pb.UserDetail{}
//	detail.Sex = "男"
//	user.UserDetail = &detail
//	newMsg.Friends = append(newMsg.GetFriends(), &user)

//	user = pb.User{}
//	user.UesrName = "肖泽壕"
//	user.NickName = "hz"
//	user.IconName = "/hz.jpg"
//	file, err = os.Open("e:/hz.jpg")
//	if err != nil {
//		fmt.Println(err)
//	}
//	buf = make([]byte, 102400)
//	n, _ = file.Read(buf)
//	user.Icon = buf[:n]
//	detail = pb.UserDetail{}
//	detail.Sex = "男"
//	user.UserDetail = &detail
//	newMsg.Friends = append(newMsg.GetFriends(), &user)

//	user = pb.User{}
//	user.UesrName = "肖泽壕"
//	user.NickName = "hz"
//	user.IconName = "/hz.jpg"
//	file, err = os.Open("e:/hz.jpg")
//	if err != nil {
//		fmt.Println(err)
//	}
//	buf = make([]byte, 102400)
//	n, _ = file.Read(buf)
//	user.Icon = buf[:n]
//	detail = pb.UserDetail{}
//	detail.Sex = "男"
//	user.UserDetail = &detail
//	newMsg.Friends = append(newMsg.GetFriends(), &user)

//	user = pb.User{}
//	user.UesrName = "肖泽壕"
//	user.NickName = "hz"
//	user.IconName = "/hz.jpg"
//	file, err = os.Open("e:/hz.jpg")
//	if err != nil {
//		fmt.Println(err)
//	}
//	buf = make([]byte, 102400)
//	n, _ = file.Read(buf)
//	user.Icon = buf[:n]
//	detail = pb.UserDetail{}
//	detail.Sex = "男"
//	user.UserDetail = &detail
//	newMsg.Friends = append(newMsg.GetFriends(), &user)

//	user = pb.User{}
//	user.UesrName = "肖泽壕"
//	user.NickName = "hz"
//	user.IconName = "/hz.jpg"
//	file, err = os.Open("e:/hz.jpg")
//	if err != nil {
//		fmt.Println(err)
//	}
//	buf = make([]byte, 102400)
//	n, _ = file.Read(buf)
//	user.Icon = buf[:n]
//	detail = pb.UserDetail{}
//	detail.Sex = "男"
//	user.UserDetail = &detail
//	newMsg.Friends = append(newMsg.GetFriends(), &user)

//	user = pb.User{}
//	user.UesrName = "肖泽壕"
//	user.NickName = "hz"
//	user.IconName = "/hz.jpg"
//	file, err = os.Open("e:/hz.jpg")
//	if err != nil {
//		fmt.Println(err)
//	}
//	buf = make([]byte, 102400)
//	n, _ = file.Read(buf)
//	user.Icon = buf[:n]
//	detail = pb.UserDetail{}
//	detail.Sex = "男"
//	user.UserDetail = &detail
//	newMsg.Friends = append(newMsg.GetFriends(), &user)

//	user = pb.User{}
//	user.UesrName = "肖泽壕"
//	user.NickName = "hz"
//	user.IconName = "/hz.jpg"
//	file, err = os.Open("e:/hz.jpg")
//	if err != nil {
//		fmt.Println(err)
//	}
//	buf = make([]byte, 102400)
//	n, _ = file.Read(buf)
//	user.Icon = buf[:n]
//	detail = pb.UserDetail{}
//	detail.Sex = "男"
//	user.UserDetail = &detail
//	newMsg.Friends = append(newMsg.GetFriends(), &user)

//	user = pb.User{}
//	user.UesrName = "肖泽壕"
//	user.NickName = "hz"
//	user.IconName = "/hz.jpg"
//	file, err = os.Open("e:/hz.jpg")
//	if err != nil {
//		fmt.Println(err)
//	}
//	buf = make([]byte, 102400)
//	n, _ = file.Read(buf)
//	user.Icon = buf[:n]
//	detail = pb.UserDetail{}
//	detail.Sex = "男"
//	user.UserDetail = &detail
//	newMsg.Friends = append(newMsg.GetFriends(), &user)

//	user = pb.User{}
//	user.UesrName = "肖泽壕"
//	user.NickName = "hz"
//	user.IconName = "/hz.jpg"
//	file, err = os.Open("e:/hz.jpg")
//	if err != nil {
//		fmt.Println(err)
//	}
//	buf = make([]byte, 102400)
//	n, _ = file.Read(buf)
//	user.Icon = buf[:n]
//	detail = pb.UserDetail{}
//	detail.Sex = "男"
//	user.UserDetail = &detail
//	newMsg.Friends = append(newMsg.GetFriends(), &user)
//	this.Conn.Write(newMsg)
