package utils

import (
	"fmt"
	"github.com/MicroMolekula/gpt-service/internal/models"
)

func GetUserGender(genderCode string) string {
	switch genderCode {
	case "male":
		return "Мужчина"
	case "female":
		return "Женщина"
	}
	return ""
}

func GetUserLevel(levelCode string) string {
	switch levelCode {
	case "nothing":
		return "Никогда не было физических нагрузок"
	case "walk":
		return "Много двигался пешком"
	case "fit":
		return "Поддерживал форму"
	case "active":
		return "Активно занимался спортом"
	}
	return ""
}

func GetUserEquipment(eqCode string) string {
	switch eqCode {
	case "minimal":
		return "Нет инвентаря для занятий"
	case "home":
		return "Есть домашний зал с небольшим количеством инвенторя"
	case "gym":
		return "Хожу в тренажерный зал"
	}
	return ""
}

func GetUserTarget(targetCode string) string {
	switch targetCode {
	case "strength":
		return "Хочу набрать мышечную массу"
	case "fit":
		return "Хочу оставаться в хорошей форме"
	case "thick":
		return "Хочу сбросить вес"
	}
	return ""
}

func GenerateQueryByUserData(user *models.User) string {
	target := fmt.Sprintf("Цель: %s", GetUserTarget(user.Target))
	gender := fmt.Sprintf("Пол: %s", GetUserGender(user.Gender))
	age := fmt.Sprintf("Возраст: %d", user.Age)
	height := fmt.Sprintf("Рост: %f", user.Height)
	weight := fmt.Sprintf("Вес: %d", user.Weight)
	targetWeight := fmt.Sprintf("Целевой вес: %d", user.DesiredWeight)
	equipment := fmt.Sprintf("Доступ к спортивному инвентарю: %s", GetUserEquipment(user.Inventory))
	level := fmt.Sprintf("Уровень физической подготовки: %s", GetUserLevel(user.LevelOfTraining))
	other := fmt.Sprintf("Дополнительная информация: %s", user.Details)
	return fmt.Sprintf("%s \n%s \n%s \n%s \n%s \n%s \n%s \n%s \n%s",
		target,
		gender,
		age,
		height,
		weight,
		targetWeight,
		equipment,
		level,
		other,
	)
}
