package web

import (
	"github.com/gorilla/mux"
	"github.com/juju/loggo"
	"hive-accordian/config"
	"github.com/markbates/pkger"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	logger     *loggo.Logger
	templates  *template.Template
)

func Init(c *config.Config) error {
	newLogger := loggo.GetLogger("web")
	logger = &newLogger

	// Load Templates
	templateDir := pkger.Include("/web/templates")
	t, err := compileTemplates(templateDir)
	if err != nil {
		return err
	}
	templates = t

	// Setup Router
	r := mux.NewRouter()
	r.Use(Middleware)

	// Static Files
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(pkger.Dir("/web/static"))))

	// Display Accordion
	r.HandleFunc("/", GetAccordion).Methods("GET")

	go func() {
		srv := &http.Server{
			Handler:      r,
			Addr:         ":5000",
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		}
		err := srv.ListenAndServe()
		if err != nil {
			logger.Errorf("Could not start web server %s", err.Error())
		}
	}()

	return nil

}

func compileTemplates(dir string) (*template.Template, error) {
	tpl := template.New("")

	tpl.Funcs(template.FuncMap{
		"dec": func(i int) int {
			i--
			return i
		},
		"htmlSafe": func(html string) template.HTML {
			return template.HTML(html)
		},
	})

	err := pkger.Walk(dir, func(path string, info os.FileInfo, _ error) error {
		if info.IsDir() || !strings.HasSuffix(path, ".gohtml") {
			return nil
		}
		f, err := pkger.Open(path)
		if err != nil {
			logger.Errorf("could not open pkger path %s: %s", path, err.Error())
			return err
		}
		// Now read it.
		sl, err := ioutil.ReadAll(f)
		if err != nil {
			logger.Errorf("could not read pkger file %s: %s", path, err.Error())
		}

		// It can now be parsed as a string.
		_, err = tpl.Parse(string(sl))
		if err != nil {
			logger.Errorf("could not open parse template %s: %s", path, err.Error())
			return err
		}

		return nil
	})

	return tpl, err
}