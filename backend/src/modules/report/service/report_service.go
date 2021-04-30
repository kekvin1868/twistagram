package service

import (
	"twistagram/src/modules/report/dao"
	"twistagram/src/modules/report/domain"
	"twistagram/src/modules/report/domain/api"
)

func PostReport(report *domain.Report) (*domain.Report, error) {
	return dao.PostReport(report)
}

func GetReportCounts(PostID uint64) (*api.ReportAPI, error) {
	return dao.GetReportCounts(PostID)
}
