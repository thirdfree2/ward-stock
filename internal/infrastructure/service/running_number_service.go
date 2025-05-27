package service

import (
	"fmt"
	"time"
	"ward-stock-backend/internal/domain"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type RunningNumberService struct {
	db *gorm.DB
}

func NewRunningNumberService(db *gorm.DB) *RunningNumberService {
	return &RunningNumberService{db: db}
}

func (s *RunningNumberService) NextNumber(key string, prefix string, digit int) (string, error) {
	var rn domain.RunningNumber

	err := s.db.Transaction(func(tx *gorm.DB) error {
		// STEP 1: ลองหา key ก่อน
		err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("key = ?", key).
			First(&rn).Error

		if err != nil {
			if err == gorm.ErrRecordNotFound {
				// STEP 2: ถ้ายังไม่มี → สร้างใหม่
				rn = domain.RunningNumber{
					Key:    key,
					LastNo: 1, // เริ่มที่ 1
				}
				return tx.Create(&rn).Error
			}
			return err // error อื่น ๆ
		}

		// STEP 3: ถ้ามีอยู่แล้ว → +1 และ update
		rn.LastNo++
		rn.UpdatedAt = time.Now()

		return tx.Save(&rn).Error
	})

	if err != nil {
		return "", err
	}

	runningCode := fmt.Sprintf("%s%0*d", prefix, digit, rn.LastNo)
	return runningCode, nil
}

var _ domain.RunningNumberRepository = (*RunningNumberService)(nil)
