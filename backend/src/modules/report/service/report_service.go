package service

import (
	"twistagram/src/modules/report/dao"
	"twistagram/src/modules/report/domain"
)

func PostReport(report *domain.Report) (*domain.Report, error) {
	return dao.PostReport(report)
}

func GetReportCounts(PostID uint64) (int, error) {
	return dao.GetReportCounts(PostID)
}
