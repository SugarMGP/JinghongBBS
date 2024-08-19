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

func GetReportByID(post uint) (*models.Report, error) {
	var report models.Report
	result := database.DB.Where("post = ?", post).First(&report)
	if result.Error != nil {
		return nil, result.Error
	}
	return &report, nil
}

func SetReportStatus(post uint, status uint) error {
	result := database.DB.Where("post = ?", post).First(&models.Report{}).Update("status", status)
	return result.Error
}
