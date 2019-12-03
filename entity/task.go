package entity

// Task ... タスクモデルのプロパティ
type Task struct {
	ID            int    `gorm:"primary_key"`
	TaskName      string `gorm:"not null"`
	Deadline      string `gorm:"varchar(12)"`
	ImportantFlag bool   `gorm:"not null"`
	FinisheFlag   bool   `gorm:"not null"`
	DeleteFlag    bool   `gorm:"not null"`
}
