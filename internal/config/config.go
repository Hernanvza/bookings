package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

// APPconfig holds the application config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	Infolog       *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
}
