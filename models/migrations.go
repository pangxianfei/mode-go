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

type Migrations struct { 
    Id *uint `gorm:"column:id;primary_key;auto_increment"`
    Migration string `gorm:"column:migration;type:varchar(255);"` 
    Batch *uint `gorm:"column:batch;type:int(10) unsigned;"` 
    model.BaseModel 
}

func (migrations *Migrations) TableName() string{ 
     return migrations.SetTableName("migrations") 
}
/*设置 Id 值*/
func (migrations *Migrations) SetIdAttribute(value interface{}) interface{} {
 return migrations.Id 
}
/**获取 Id 值
*/
func (migrations *Migrations) GetIdAttribute(value interface{}) interface{} {
 return migrations.Id 
}

/*设置 Migration 值*/
func (migrations *Migrations) SetMigrationAttribute(value interface{}) interface{} {
 return migrations.Migration 
}
/**获取 Migration 值
*/
func (migrations *Migrations) GetMigrationAttribute(value interface{}) interface{} {
 return migrations.Migration 
}

/*设置 Batch 值*/
func (migrations *Migrations) SetBatchAttribute(value interface{}) interface{} {
 return migrations.Batch 
}
/**获取 Batch 值
*/
func (migrations *Migrations) GetBatchAttribute(value interface{}) interface{} {
 return migrations.Batch 
}

