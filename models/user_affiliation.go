/*
Author:pangxianfei
*DATE:20200-51-05 01:30:51 
*contact:421339244@qq.com
*/
package models 

 import (
     "github.com/totoval/framework/model"
     "github.com/totoval/framework/helpers/zone"
  )

type UserAffiliation struct { 
    UserId *uint `gorm:"column:user_id;type:int(10) unsigned;"` 
    UaffCode string `gorm:"column:uaff_code;type:varchar(32);"` 
    UaffFromCode string `gorm:"column:uaff_from_code;type:varchar(32);"` 
    UaffRootId *uint `gorm:"column:uaff_root_id;type:int(10) unsigned;"` 
    UaffParentId *uint `gorm:"column:uaff_parent_id;type:int(10) unsigned;"` 
    UaffLeftId *uint `gorm:"column:uaff_left_id;type:int(10) unsigned;"` 
    UaffRightId *uint `gorm:"column:uaff_right_id;type:int(10) unsigned;"` 
    UaffLevel *uint `gorm:"column:uaff_level;type:int(10) unsigned;"` 
    UserCreatedAt zone.Time `gorm:"column:user_created_at;"` 
    UserUpdatedAt zone.Time `gorm:"column:user_updated_at;"` 
    UserDeletedAt zone.Time `gorm:"column:user_deleted_at;"` 
    model.BaseModel 
}

func (user_affiliation *UserAffiliation) TableName() string{ 
     return user_affiliation.SetTableName("user_affiliation") 
}
/*设置 UserId 值*/
func (user_affiliation *UserAffiliation) SetUserIdAttribute(value interface{}) interface{} {
 return user_affiliation.UserId 
}
/**获取 UserId 值
*/
func (user_affiliation *UserAffiliation) GetUserIdAttribute(value interface{}) interface{} {
 return user_affiliation.UserId 
}

/*设置 UaffCode 值*/
func (user_affiliation *UserAffiliation) SetUaffCodeAttribute(value interface{}) interface{} {
 return user_affiliation.UaffCode 
}
/**获取 UaffCode 值
*/
func (user_affiliation *UserAffiliation) GetUaffCodeAttribute(value interface{}) interface{} {
 return user_affiliation.UaffCode 
}

/*设置 UaffFromCode 值*/
func (user_affiliation *UserAffiliation) SetUaffFromCodeAttribute(value interface{}) interface{} {
 return user_affiliation.UaffFromCode 
}
/**获取 UaffFromCode 值
*/
func (user_affiliation *UserAffiliation) GetUaffFromCodeAttribute(value interface{}) interface{} {
 return user_affiliation.UaffFromCode 
}

/*设置 UaffRootId 值*/
func (user_affiliation *UserAffiliation) SetUaffRootIdAttribute(value interface{}) interface{} {
 return user_affiliation.UaffRootId 
}
/**获取 UaffRootId 值
*/
func (user_affiliation *UserAffiliation) GetUaffRootIdAttribute(value interface{}) interface{} {
 return user_affiliation.UaffRootId 
}

/*设置 UaffParentId 值*/
func (user_affiliation *UserAffiliation) SetUaffParentIdAttribute(value interface{}) interface{} {
 return user_affiliation.UaffParentId 
}
/**获取 UaffParentId 值
*/
func (user_affiliation *UserAffiliation) GetUaffParentIdAttribute(value interface{}) interface{} {
 return user_affiliation.UaffParentId 
}

/*设置 UaffLeftId 值*/
func (user_affiliation *UserAffiliation) SetUaffLeftIdAttribute(value interface{}) interface{} {
 return user_affiliation.UaffLeftId 
}
/**获取 UaffLeftId 值
*/
func (user_affiliation *UserAffiliation) GetUaffLeftIdAttribute(value interface{}) interface{} {
 return user_affiliation.UaffLeftId 
}

/*设置 UaffRightId 值*/
func (user_affiliation *UserAffiliation) SetUaffRightIdAttribute(value interface{}) interface{} {
 return user_affiliation.UaffRightId 
}
/**获取 UaffRightId 值
*/
func (user_affiliation *UserAffiliation) GetUaffRightIdAttribute(value interface{}) interface{} {
 return user_affiliation.UaffRightId 
}

/*设置 UaffLevel 值*/
func (user_affiliation *UserAffiliation) SetUaffLevelAttribute(value interface{}) interface{} {
 return user_affiliation.UaffLevel 
}
/**获取 UaffLevel 值
*/
func (user_affiliation *UserAffiliation) GetUaffLevelAttribute(value interface{}) interface{} {
 return user_affiliation.UaffLevel 
}

/*设置 UserCreatedAt 值*/
func (user_affiliation *UserAffiliation) SetUserCreatedAtAttribute(value interface{}) interface{} {
 return user_affiliation.UserCreatedAt 
}
/**获取 UserCreatedAt 值
*/
func (user_affiliation *UserAffiliation) GetUserCreatedAtAttribute(value interface{}) interface{} {
 return user_affiliation.UserCreatedAt 
}

/*设置 UserUpdatedAt 值*/
func (user_affiliation *UserAffiliation) SetUserUpdatedAtAttribute(value interface{}) interface{} {
 return user_affiliation.UserUpdatedAt 
}
/**获取 UserUpdatedAt 值
*/
func (user_affiliation *UserAffiliation) GetUserUpdatedAtAttribute(value interface{}) interface{} {
 return user_affiliation.UserUpdatedAt 
}

/*设置 UserDeletedAt 值*/
func (user_affiliation *UserAffiliation) SetUserDeletedAtAttribute(value interface{}) interface{} {
 return user_affiliation.UserDeletedAt 
}
/**获取 UserDeletedAt 值
*/
func (user_affiliation *UserAffiliation) GetUserDeletedAtAttribute(value interface{}) interface{} {
 return user_affiliation.UserDeletedAt 
}

