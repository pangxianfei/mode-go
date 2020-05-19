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

type UserToken struct { 
    Id *uint `gorm:"column:id;primary_key;auto_increment"`
    Uid *uint `gorm:"column:uid;type:int(11);"` 
    ClientId string `gorm:"column:client_id;type:varchar(255);"` 
    Name string `gorm:"column:name;type:varchar(100);"` 
    Token string `gorm:"column:token;type:varchar(255);"` 
    CreatedAt zone.Time `gorm:"column:created_at;"` 
    UpdatedAt zone.Time `gorm:"column:updated_at;"` 
    PastedAt zone.Time `gorm:"column:pasted_at;"` 
    DeletedAt zone.Time `gorm:"column:deleted_at;"` 
    model.BaseModel 
}

func (user_token *UserToken) TableName() string{ 
     return user_token.SetTableName("user_token") 
}
/*设置 Id 值*/
func (user_token *UserToken) SetIdAttribute(value interface{}) interface{} {
 return user_token.Id 
}
/**获取 Id 值
*/
func (user_token *UserToken) GetIdAttribute(value interface{}) interface{} {
 return user_token.Id 
}

/*设置 Uid 值*/
func (user_token *UserToken) SetUidAttribute(value interface{}) interface{} {
 return user_token.Uid 
}
/**获取 Uid 值
*/
func (user_token *UserToken) GetUidAttribute(value interface{}) interface{} {
 return user_token.Uid 
}

/*设置 ClientId 值*/
func (user_token *UserToken) SetClientIdAttribute(value interface{}) interface{} {
 return user_token.ClientId 
}
/**获取 ClientId 值
*/
func (user_token *UserToken) GetClientIdAttribute(value interface{}) interface{} {
 return user_token.ClientId 
}

/*设置 Name 值*/
func (user_token *UserToken) SetNameAttribute(value interface{}) interface{} {
 return user_token.Name 
}
/**获取 Name 值
*/
func (user_token *UserToken) GetNameAttribute(value interface{}) interface{} {
 return user_token.Name 
}

/*设置 Token 值*/
func (user_token *UserToken) SetTokenAttribute(value interface{}) interface{} {
 return user_token.Token 
}
/**获取 Token 值
*/
func (user_token *UserToken) GetTokenAttribute(value interface{}) interface{} {
 return user_token.Token 
}

/*设置 CreatedAt 值*/
func (user_token *UserToken) SetCreatedAtAttribute(value interface{}) interface{} {
 return user_token.CreatedAt 
}
/**获取 CreatedAt 值
*/
func (user_token *UserToken) GetCreatedAtAttribute(value interface{}) interface{} {
 return user_token.CreatedAt 
}

/*设置 UpdatedAt 值*/
func (user_token *UserToken) SetUpdatedAtAttribute(value interface{}) interface{} {
 return user_token.UpdatedAt 
}
/**获取 UpdatedAt 值
*/
func (user_token *UserToken) GetUpdatedAtAttribute(value interface{}) interface{} {
 return user_token.UpdatedAt 
}

/*设置 PastedAt 值*/
func (user_token *UserToken) SetPastedAtAttribute(value interface{}) interface{} {
 return user_token.PastedAt 
}
/**获取 PastedAt 值
*/
func (user_token *UserToken) GetPastedAtAttribute(value interface{}) interface{} {
 return user_token.PastedAt 
}

/*设置 DeletedAt 值*/
func (user_token *UserToken) SetDeletedAtAttribute(value interface{}) interface{} {
 return user_token.DeletedAt 
}
/**获取 DeletedAt 值
*/
func (user_token *UserToken) GetDeletedAtAttribute(value interface{}) interface{} {
 return user_token.DeletedAt 
}

