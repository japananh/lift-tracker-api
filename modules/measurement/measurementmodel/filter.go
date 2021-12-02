package measurementmodel

type Filter struct {
	FakeUserId string `json:"user_id" form:"user_id"`
	UserId     int
}
