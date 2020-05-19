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

type User struct { 
    UserId *uint `gorm:"column:user_id;primary_key;auto_increment"`
    UserName string `gorm:"column:user_name;type:varchar(100);"` 
    UserEmail string `gorm:"column:user_email;type:varchar(100);"` 
    UserPassword string `gorm:"column:user_password;type:varchar(100);"` 
    UserCreatedAt zone.Time `gorm:"column:user_created_at;"` 
    UserUpdatedAt zone.Time `gorm:"column:user_updated_at;"` 
    UserDeletedAt zone.Time `gorm:"column:user_deleted_at;"` 
    model.BaseModel 
}

func (user *User) TableName() string{ 
  return user.SetTableName("user") 
}
/*设置 UserId 值*/
func (user *User) SetUserIdAttribute(value interface{}) interface{} {
 return user.UserId 
}
/**获取 UserId 值
*/
func (user *User) GetUserIdAttribute(value interface{}) interface{} {
 return user.UserId 
}

/*设置 UserName 值*/
func (user *User) SetUserNameAttribute(value interface{}) interface{} {
 return user.UserName 
}
/**获取 UserName 值
*/
func (user *User) GetUserNameAttribute(value interface{}) interface{} {
 return user.UserName 
}

/*设置 UserEmail 值*/
func (user *User) SetUserEmailAttribute(value interface{}) interface{} {
 return user.UserEmail 
}
/**获取 UserEmail 值
*/
func (user *User) GetUserEmailAttribute(value interface{}) interface{} {
 return user.UserEmail 
}

/*设置 UserPassword 值*/
func (user *User) SetUserPasswordAttribute(value interface{}) interface{} {
 return user.UserPassword 
}
/**获取 UserPassword 值
*/
func (user *User) GetUserPasswordAttribute(value interface{}) interface{} {
 return user.UserPassword 
}

/*设置 UserCreatedAt 值*/
func (user *User) SetUserCreatedAtAttribute(value interface{}) interface{} {
 return user.UserCreatedAt 
}
/**获取 UserCreatedAt 值
*/
func (user *User) GetUserCreatedAtAttribute(value interface{}) interface{} {
 return user.UserCreatedAt 
}

/*设置 UserUpdatedAt 值*/
func (user *User) SetUserUpdatedAtAttribute(value interface{}) interface{} {
 return user.UserUpdatedAt 
}
/**获取 UserUpdatedAt 值
*/
func (user *User) GetUserUpdatedAtAttribute(value interface{}) interface{} {
 return user.UserUpdatedAt 
}

/*设置 UserDeletedAt 值*/
func (user *User) SetUserDeletedAtAttribute(value interface{}) interface{} {
 return user.UserDeletedAt 
}
/**获取 UserDeletedAt 值
*/
func (user *User) GetUserDeletedAtAttribute(value interface{}) interface{} {
 return user.UserDeletedAt 
}

