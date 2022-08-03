package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/snowflake"
)

func CreatePost(p *models.Post) error {
	// 生成post id
	p.ID = snowflake.GenID()
	// 保存到数据库
	return mysql.CreatePost(p)
	// 返回
}