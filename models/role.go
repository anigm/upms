package models

import "github.com/astaxie/beego"

type Role struct {
	Model
	Name        string       `gorm:"unique_index:idx_name_platform"`
	Platform    string       `gorm:"unique_index:idx_name_platform"`
	Permissions []Permission `gorm:"many2many:role_permissions;" json:",omitempty"`
	Users       []User       `gorm:"many2many:user_roles;" json:",omitempty"`
	RoleGroups  []RoleGroup  `gorm:"many2many:group_roles;" json:",omitempty"`
}

type RoleGroup struct {
	Model
	Name       string      `gorm:"unique_index:idx_name_parent"`
	ParentID   int64       `gorm:"unique_index:idx_name_parent"`
	RoleGroups []RoleGroup `gorm:"ForeignKey:ParentID;AssociationForeignKey:ID" json:",omitempty"`
	Roles      []Role      `gorm:"many2many:group_roles;" json:",omitempty"`
}

func SaveRole(role *Role) error {
	if db, err := DB(); err != nil {
		beego.Error(err)
		return err
	} else {
		defer db.Close()
		if err := db.Save(role).Error; err != nil {
			return err
		}
	}
	return nil
}

func SaveRoleGroup(roleGroup *RoleGroup) error {
	if db, err := DB(); err != nil {
		beego.Error(err)
		return err
	} else {
		defer db.Close()
		if err := db.Save(roleGroup).Error; err != nil {
			return err
		}
	}
	return nil
}
