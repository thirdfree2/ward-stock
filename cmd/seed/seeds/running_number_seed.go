package seeds

import (
	"fmt"
	"ward-stock-backend/internal/domain"

	"gorm.io/gorm"
)

func SetRunningNumberAfterUserSeed(db *gorm.DB) {
	var lastUser domain.User
	err := db.Order("code desc").First(&lastUser).Error
	if err != nil {
		return // ไม่มี user ไม่ต้องตั้งค่า
	}

	// ตัดเลขจาก Code เช่น USR00015 → 15
	var lastNo int
	_, err = fmt.Sscanf(lastUser.Code, "USR%05d", &lastNo)
	if err != nil {
		return
	}

	db.FirstOrCreate(&domain.RunningNumber{
		Key:    "USER",
		LastNo: lastNo,
	}, domain.RunningNumber{
		Key: "USER",
	})
}
