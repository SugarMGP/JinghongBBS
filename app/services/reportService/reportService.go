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
