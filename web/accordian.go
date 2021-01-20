package web

import (
	"gopkg.in/yaml.v2"
	"hive-accordian/models"
	"io/ioutil"
	"log"
	"net/http"
)

type AccordionTemplate struct {
	Accordion models.Accordion
}

func GetAccordion(w http.ResponseWriter, r *http.Request) {
	// Init template variables
	tmplVars := &AccordionTemplate{}

	// Read File
	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &tmplVars.Accordion)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	err = templates.ExecuteTemplate(w, "accordion", tmplVars)
	if err != nil {
		logger.Errorf("could not render home template: %s", err.Error())
	}

}