package dto

type WeekPlan struct {
	Plan []Plan `json:"week_plan"`
}

type UserPlan struct {
	UserId string `json:"user_id" bson:"user_id"`
	Plan   []Plan `json:"plan" bson:"plan"`
}

type Plan struct {
	Day       string     `json:"day" bson:"day"`
	Dishes    Dishes     `json:"dishes" bson:"dishes"`
	Exercises []Exercise `json:"exercise" bson:"exercise"`
}

type Exercise struct {
	Name        string `json:"name" bson:"name"`
	Approaches  string `json:"approaches" bson:"approaches"`
	Repetitions string `json:"repetitions" bson:"repetitions"`
}

type Dishes struct {
	Breakfast []Dish `json:"breakfast" bson:"breakfast"`
	Dinner    []Dish `json:"dinner" bson:"dinner"`
	Lunch     []Dish `json:"lunch" bson:"lunch"`
}

type Dish struct {
	Name string `json:"name" bson:"name"`
	Gram int    `json:"gram" bson:"gram"`
}
