package gorm

type User struct {
	ID   uint64 `gorm:"primaryKey"`
	Name string `gorm:"uniqueIndex" json:"name"`
	Age  int    `json:"age"`
}
