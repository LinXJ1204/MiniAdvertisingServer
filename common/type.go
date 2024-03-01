package common

type Ad struct {
	Title      string `json:"title"`
	StartAt    string `json:"startAt"`
	EndAt      string `json:"endAt"`
	*Condition `json:"condition"`
}

type Condition struct {
	AgeStart uint     `json:"ageStart"`
	AgeEnd   uint     `json:"ageEnd"`
	Gender   []string `json:"gender"`
	Country  []string `json:"country"`
	Platform []string `json:"platform"`
}

type SearchCondition struct {
	offset   uint
	limit    uint
	age      uint
	gender   string
	country  string
	platform string
}

type Respond struct {
	Title string `json:"endAt"`
	EndAt string `json:"endAt"`
}
