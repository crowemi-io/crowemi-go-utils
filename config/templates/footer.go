package config

type Footer struct {
	Copyright Copyright `json:"copyright" omitempty:"true"`
	Socials   []Social  `json:"socials" omitempty:"true"`
}
type Copyright struct {
	Year    int    `json:"year"`
	Company string `json:"company"`
	Tag     string `json:"tag"`
}
type Social struct {
	Name string `json:"name"`
	Icon string `json:"icon"`
	URL  string `json:"url"`
}
