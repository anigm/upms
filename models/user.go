package models

import "github.com/astaxie/beego"
import "strings"

type User struct {
	Model
	Account    string      `gorm:"not null;unique" json:",omitempty"`
	Password   string      `gorm:"not null" json:",omitempty"`
	Name       string      `gorm:"not null;unique" json:",omitempty" form:"Name"`
	Tag        string      `json:",omitempty"`
	QQ         string      `json:",omitempty"`
	EMail      string      `json:",omitempty"`
	Roles      []Role      `gorm:"many2many:user_roles;" json:",omitempty"`
	UserGroups []UserGroup `gorm:"many2many:group_users;" json:",omitempty"`
}

type UserGroup struct {
	Model
	Name       string      `gorm:"unique_index:idx_name_parent"`
	ParentID   int64       `gorm:"unique_index:idx_name_parent"`
	UserGroups []UserGroup `gorm:"ForeignKey:ParentID;AssociationForeignKey:ID" json:",omitempty"`
	Users      []User      `gorm:"many2many:group_users;" json:",omitempty"`
}

func (u *User) Validate() (int, bool) {
	u.Account = strings.TrimSpace(u.Account)
	u.Name = strings.TrimSpace(u.Name)
	u.Password = strings.TrimSpace(u.Password)
	if u.Account == "" || u.Name == "" || u.Password == "" {
		return 1, false
	}
	if strings.Contains(u.Account, " ") || strings.Contains(u.Name, " ") || strings.Contains(u.Password, " ") {
		return 2, false
	}
	return 0, true
}

func SaveUser(user *User) error {
	if db, err := DB(); err != nil {
		beego.Error(err)
		return err
	} else {
		defer db.Close()
		if err := db.Save(user).Error; err != nil {
			return err
		}
	}
	return nil
}

func DeleteUser(user *User) error {
	if db, err := DB(); err != nil {
		beego.Error(err)
		return err
	} else {
		defer db.Close()
		db.Delete(user)
	}
	return nil
}

func DeleteUsers(ids []string) error {
	if db, err := DB(); err != nil {
		beego.Error(err)
		return err
	} else {
		defer db.Close()
		if err = db.Begin().Error; err == nil {
			if err = db.Exec("delete from users where id in (" + strings.Join(ids, ",") + ")").Error; err != nil {
				return err
			}
		} else {
			return err
		}
	}
	return nil
}

func GetUser(id int64) (*User, error) {
	user := &User{}
	user.ID = id
	if db, err := DB(); err != nil {
		beego.Error(err)
		return nil, err
	} else {
		defer db.Close()
		db.First(user)
		return user, nil
	}
}

func GetUserInfo(id int64) (*User, error) {
	user := &User{}
	user.ID = id
	if db, err := DB(); err != nil {
		beego.Error(err)
		return nil, err
	} else {
		defer db.Close()
		db.Select("id,name,qq").First(user)
		db.Model(user).Select("id,name,platform").Association("Roles").Find(&user.Roles)
		for i, per := range user.Roles {
			db.Model(per).Select("id,name,path").Association("Permissions").Find(&user.Roles[i].Permissions)
		}
		db.Model(user).Select("id,name").Association("UserGroups").Find(&user.UserGroups)
		return user, nil
	}
}

func GetUserByAccount(account string) (*User, error) {
	user := &User{}
	user.Account = account
	if db, err := DB(); err != nil {
		beego.Error(err)
		return nil, err
	} else {
		defer db.Close()
		db.Select("account,name,password").Where("account = ?", account).First(user)
		return user, nil
	}
}

// GetUser
// 获取用户列表，并取得其分组、角色及角色所有的权限
func GetUsers() ([]User, error) {
	var users []User
	if db, err := DB(); err != nil {
		beego.Error(err)
		return nil, err
	} else {
		defer db.Close()
		db.Select("id,name").Find(&users)
		for i, user := range users {
			db.Model(&user).Select("id,name,platform").Association("Roles").Find(&users[i].Roles)
			for j, per := range users[i].Roles {
				db.Model(per).Select("id,name,path").Association("Permissions").Find(&users[i].Roles[j].Permissions)
			}
			db.Model(&user).Select("id,name").Association("UserGroups").Find(&users[i].UserGroups)
		}
		return users, nil
	}
}

func GetUserInfoByQQ(qq string) (*User, error) {
	user := &User{}
	user.QQ = qq
	if db, err := DB(); err != nil {
		beego.Error(err)
		return nil, err
	} else {
		defer db.Close()
		db.Select("id,name,qq").Where("qq = ?", qq).First(user)
		db.Model(user).Select("id,name,platform").Association("Roles").Find(&user.Roles)
		for i, per := range user.Roles {
			db.Model(per).Select("id,name,path").Association("Permissions").Find(&user.Roles[i].Permissions)
		}
		db.Model(user).Select("id,name").Association("UserGroups").Find(&user.UserGroups)
		return user, nil
	}
}
