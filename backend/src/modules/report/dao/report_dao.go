package dao

import (
	"twistagram/src/modules/report/domain"
	"twistagram/src/modules/report/domain/api"
	"twistagram/src/orm"
)

func isReported(UserID uint64, PostID uint64) bool {
	var report domain.Report
	var exist bool

	if err := orm.Engine.Where("post_id = ? AND user_id = ?", PostID, UserID).First(&report).Error; err != nil {
		exist = false
	} else {
		exist = true
	}

	return exist
}

func PostReport(report *domain.Report) (*domain.Report, error) {
	var newReport domain.Report
	newReport = *report

	reportExist := isReported(uint64(newReport.UserID), uint64(newReport.PostID))
	if reportExist == true {
		return nil, nil
	}

	res := orm.Engine.Create(&newReport)

	if res.Error != nil {
		return nil, res.Error
	}

	return report, nil
}

func GetReportCounts(PostID uint64) (*api.ReportAPI, error) {
	var report domain.Report
	var reportCounts *api.ReportAPI

	res := orm.Engine.Table("reports").Select("user_id").Where("post_id = ?", PostID).Find(&report)

	if res.Error != nil {
		reportCounts.Report = -1
		return reportCounts, res.Error
	}

	reportCounts.Report = int(res.RowsAffected)
	return reportCounts, nil

}
