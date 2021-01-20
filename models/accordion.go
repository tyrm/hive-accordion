package models

type Accordion struct {
	Accordion []Header `yaml:"accordion"`
}

type Header struct {
	Title string `yaml:"title"`
	Links []Link `yaml:"links"`
}

type Link struct {
	Title string `yaml:"title"`
	Link  string `yaml:"link"`
}
