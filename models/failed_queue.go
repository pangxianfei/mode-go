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

type FailedQueue struct { 
    FailedQueueId *uint `gorm:"column:failed_queue_id;primary_key;auto_increment"`
    FailedQueueHash string `gorm:"column:failed_queue_hash;type:varchar(100);"` 
    FailedQueueTopicName string `gorm:"column:failed_queue_topic_name;type:varchar(100);"` 
    FailedQueueChannelName string `gorm:"column:failed_queue_channel_name;type:varchar(100);"` 
    FailedQueueData string `gorm:"column:failed_queue_data;type:varbinary(2048);"` 
    FailedQueuePushedAt zone.Time `gorm:"column:failed_queue_pushed_at;"` 
    FailedQueueDelay string `gorm:"column:failed_queue_delay;type:bigint(20) unsigned;"` 
    FailedQueueRetries *uint `gorm:"column:failed_queue_retries;type:int(10) unsigned;"` 
    FailedQueueTried *uint `gorm:"column:failed_queue_tried;type:int(10) unsigned;"` 
    FailedQueueErr string `gorm:"column:failed_queue_err;type:longtext;"` 
    FailedQueueCreatedAt zone.Time `gorm:"column:failed_queue_created_at;"` 
    FailedQueueUpdatedAt zone.Time `gorm:"column:failed_queue_updated_at;"` 
    FailedQueueDeletedAt zone.Time `gorm:"column:failed_queue_deleted_at;"` 
    model.BaseModel 
}

func (failed_queue *FailedQueue) TableName() string{ 
     return failed_queue.SetTableName("failed_queue") 
}
/*设置 FailedQueueId 值*/
func (failed_queue *FailedQueue) SetFailedQueueIdAttribute(value interface{}) interface{} {
 return failed_queue.FailedQueueId 
}
/**获取 FailedQueueId 值
*/
func (failed_queue *FailedQueue) GetFailedQueueIdAttribute(value interface{}) interface{} {
 return failed_queue.FailedQueueId 
}

/*设置 FailedQueueHash 值*/
func (failed_queue *FailedQueue) SetFailedQueueHashAttribute(value interface{}) interface{} {
 return failed_queue.FailedQueueHash 
}
/**获取 FailedQueueHash 值
*/
func (failed_queue *FailedQueue) GetFailedQueueHashAttribute(value interface{}) interface{} {
 return failed_queue.FailedQueueHash 
}

/*设置 FailedQueueTopicName 值*/
func (failed_queue *FailedQueue) SetFailedQueueTopicNameAttribute(value interface{}) interface{} {
 return failed_queue.FailedQueueTopicName 
}
/**获取 FailedQueueTopicName 值
*/
func (failed_queue *FailedQueue) GetFailedQueueTopicNameAttribute(value interface{}) interface{} {
 return failed_queue.FailedQueueTopicName 
}

/*设置 FailedQueueChannelName 值*/
func (failed_queue *FailedQueue) SetFailedQueueChannelNameAttribute(value interface{}) interface{} {
 return failed_queue.FailedQueueChannelName 
}
/**获取 FailedQueueChannelName 值
*/
func (failed_queue *FailedQueue) GetFailedQueueChannelNameAttribute(value interface{}) interface{} {
 return failed_queue.FailedQueueChannelName 
}

/*设置 FailedQueueData 值*/
func (failed_queue *FailedQueue) SetFailedQueueDataAttribute(value interface{}) interface{} {
 return failed_queue.FailedQueueData 
}
/**获取 FailedQueueData 值
*/
func (failed_queue *FailedQueue) GetFailedQueueDataAttribute(value interface{}) interface{} {
 return failed_queue.FailedQueueData 
}

/*设置 FailedQueuePushedAt 值*/
func (failed_queue *FailedQueue) SetFailedQueuePushedAtAttribute(value interface{}) interface{} {
 return failed_queue.FailedQueuePushedAt 
}
/**获取 FailedQueuePushedAt 值
*/
func (failed_queue *FailedQueue) GetFailedQueuePushedAtAttribute(value interface{}) interface{} {
 return failed_queue.FailedQueuePushedAt 
}

/*设置 FailedQueueDelay 值*/
func (failed_queue *FailedQueue) SetFailedQueueDelayAttribute(value interface{}) interface{} {
 return failed_queue.FailedQueueDelay 
}
/**获取 FailedQueueDelay 值
*/
func (failed_queue *FailedQueue) GetFailedQueueDelayAttribute(value interface{}) interface{} {
 return failed_queue.FailedQueueDelay 
}

/*设置 FailedQueueRetries 值*/
func (failed_queue *FailedQueue) SetFailedQueueRetriesAttribute(value interface{}) interface{} {
 return failed_queue.FailedQueueRetries 
}
/**获取 FailedQueueRetries 值
*/
func (failed_queue *FailedQueue) GetFailedQueueRetriesAttribute(value interface{}) interface{} {
 return failed_queue.FailedQueueRetries 
}

/*设置 FailedQueueTried 值*/
func (failed_queue *FailedQueue) SetFailedQueueTriedAttribute(value interface{}) interface{} {
 return failed_queue.FailedQueueTried 
}
/**获取 FailedQueueTried 值
*/
func (failed_queue *FailedQueue) GetFailedQueueTriedAttribute(value interface{}) interface{} {
 return failed_queue.FailedQueueTried 
}

/*设置 FailedQueueErr 值*/
func (failed_queue *FailedQueue) SetFailedQueueErrAttribute(value interface{}) interface{} {
 return failed_queue.FailedQueueErr 
}
/**获取 FailedQueueErr 值
*/
func (failed_queue *FailedQueue) GetFailedQueueErrAttribute(value interface{}) interface{} {
 return failed_queue.FailedQueueErr 
}

/*设置 FailedQueueCreatedAt 值*/
func (failed_queue *FailedQueue) SetFailedQueueCreatedAtAttribute(value interface{}) interface{} {
 return failed_queue.FailedQueueCreatedAt 
}
/**获取 FailedQueueCreatedAt 值
*/
func (failed_queue *FailedQueue) GetFailedQueueCreatedAtAttribute(value interface{}) interface{} {
 return failed_queue.FailedQueueCreatedAt 
}

/*设置 FailedQueueUpdatedAt 值*/
func (failed_queue *FailedQueue) SetFailedQueueUpdatedAtAttribute(value interface{}) interface{} {
 return failed_queue.FailedQueueUpdatedAt 
}
/**获取 FailedQueueUpdatedAt 值
*/
func (failed_queue *FailedQueue) GetFailedQueueUpdatedAtAttribute(value interface{}) interface{} {
 return failed_queue.FailedQueueUpdatedAt 
}

/*设置 FailedQueueDeletedAt 值*/
func (failed_queue *FailedQueue) SetFailedQueueDeletedAtAttribute(value interface{}) interface{} {
 return failed_queue.FailedQueueDeletedAt 
}
/**获取 FailedQueueDeletedAt 值
*/
func (failed_queue *FailedQueue) GetFailedQueueDeletedAtAttribute(value interface{}) interface{} {
 return failed_queue.FailedQueueDeletedAt 
}

