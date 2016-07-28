package model

//查找群
func SearchGroup(info *SearchInfo) (*Group, error) {
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
	group := new(Group)
	if res.Next() {
		res.Scan(&(group.GroupID), &(group.GroupName), &(group.GroupIntro), &(group.CreateTime), &(group.Rank))
		return group, nil
	}
	return nil, nil
}
