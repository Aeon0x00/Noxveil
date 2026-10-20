package domain

import "time"

type User struct {
	ID 			int64		`db:"id"`	
	Balance 	float64		`db:"balance"`
	TelegramID	int64		`db:"telegram_id"`
	CreatedAt 	time.Time	`db:"created_at"`
}

type Subsriptions struct {
	ID				string		`db:"id"`
	UserID			int64		`db:"user_id"`
	SubID			string		`db:"sub_id"`
	IsActive		bool		`db:"is_active"`
	LimitDevice		int32		`db:"limit_device"`	
	TotalBytes		int64		`db:"total_bytes"`
	ExpiryTime		*time.Time	`db:"expiry_time"`
	CreatedAt		time.Time	`db:"created_at"`
}