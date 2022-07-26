package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
)

func GetCommunityList() (community []*models.Community, err error) {
	return mysql.GetCommunityList()
}

func GetCommunityDetail(id int64) (*models.CommunityDetail, error) {
	return mysql.GetCommunityDetailByID(id)
}