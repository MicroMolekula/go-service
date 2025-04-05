package models

type User struct {
	ID              uint    `gorm:"primaryKey" json:"id"`
	Name            string  `json:"name"`
	Email           string  `json:"email"`
	Password        string  `json:"-"`
	YandexID        int     `json:"-"`
	Gender          string  `json:"gender"`
	LevelOfTraining string  `json:"level_of_training"`
	Inventory       string  `json:"inventory"`
	Target          string  `json:"target"`
	Weight          int     `json:"weight"`
	Age             int     `json:"age"`
	Height          float64 `json:"height"`
	DesiredWeight   int     `json:"desired_weight"`
	FilledInData    bool    `json:"filled_in_data"`
	Details         string  `json:"details"`
}
