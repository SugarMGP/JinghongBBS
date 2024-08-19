package reportService

import (
	"BBS/app/models"
	"BBS/config/database"
)

func NewReport(report models.Report) error {
	result := database.DB.Create(&report)
	return result.Error
}

func GetReports(user uint) ([]models.Report, error) {
	var reports []models.Report
	result := database.DB.Where("user = ?", user).Find(&reports)
	if result.Error != nil {
		return nil, result.Error
	}
	return reports, nil
}

func GetAllReportsUnhandled() ([]models.Report, error) {
	var reports []models.Report
	result := database.DB.Where("status = 0").Find(&reports)
	if result.Error != nil {
		return nil, result.Error
	}
	return reports, nil
}

func GetReportByID(user uint, post uint) (*models.Report, error) {
	var report *models.Report
	result := database.DB.Where("user = ?", user).Where("post = ?", post).First(&report)
	if result.Error != nil {
		return nil, result.Error
	}
	return report, nil
}
