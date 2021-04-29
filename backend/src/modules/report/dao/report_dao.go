package dao

import (
	"twistagram/src/modules/report/domain"
	"twistagram/src/orm"
)

func isReported(UserID uint64, PostID uint64) bool {
	var report domain.Report
	var exist bool

	if err := orm.Engine.Where("post + id ? AND user_id = ?", PostID, UserID).First(&report).Error; err != nil {
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

func GetReportCounts(PostID uint64) (int, error) {
	var report domain.Report
	var reportCounts int

	res := orm.Engine.Table("reports").Select("user_id").Where("post_id = ?", PostID).Find(&report)

	if res.Error != nil {
		return -1, res.Error
	}

	reportCounts = int(res.RowsAffected)
	return reportCounts, nil

}
