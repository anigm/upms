package models

import "github.com/astaxie/beego"

type Permission struct {
	Model
	Path             string
	Name             string            `gorm:"unique_index:idx_name_platform"`
	Platform         string            `gorm:"unique_index:idx_name_platform"`
	Roles            []Role            `gorm:"many2many:role_permissions;" json:",omitempty"`
	PermissionGroups []PermissionGroup `gorm:"many2many:group_permissions;" json:",omitempty"`
}

type PermissionGroup struct {
	Model
	Name             string            `gorm:"unique_index:idx_name_parent"`
	ParentID         int64             `gorm:"unique_index:idx_name_parent"`
	PermissionGroups []PermissionGroup `gorm:"ForeignKey:ParentID;AssociationForeignKey:ID" json:",omitempty"`
	Permissions      []Permission      `gorm:"many2many:group_permissions;" json:",omitempty"`
}

func SavePermission(permission *Permission) error {
	if db, err := DB(); err != nil {
		beego.Error(err)
		return err
	} else {
		defer db.Close()
		if err := db.Save(permission).Error; err != nil {
			return err
		}
	}
	return nil
}
