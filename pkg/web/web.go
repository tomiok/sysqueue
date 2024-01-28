package web

import (
	"bytes"
	"errors"
	"html/template"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{
	"Cut":      nil,
	"MakeSlug": nil,
}

var templates = make(map[string]*template.Template)

type TemplateData struct {
}

func TemplateRender(w http.ResponseWriter, tmpl string, td *TemplateData, isCached bool) error {
	var (
		t *template.Template
	)

	if isCached {
		var (
			err error
			ok  bool
		)

		if len(templates) == 0 {
			if templates, err = TemplateRenderCache(); err != nil {
				return err
			}
		}

		if t, ok = templates[tmpl]; !ok {
			return errors.New("cache is not available, turn flag to false")
		}
	} else {
		cache, err := TemplateRenderCache()
		if err != nil {
			return err
		}

		t = cache[tmpl]
	}

	buf := new(bytes.Buffer)
	err := t.Execute(buf, td)

	if err != nil {
		return err
	}

	_, err = buf.WriteTo(w)

	if err != nil {
		return err
	}

	return nil
}

func TemplateRenderCache() (map[string]*template.Template, error) {
	pages, err := filepath.Glob("./pkg/templates/*.page.tmpl")
	var templateCache = make(map[string]*template.Template)

	if err != nil {
		return templateCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)

		if err != nil {
			return templateCache, err
		}

		matches, err := filepath.Glob("./pkg/templates/*.layout.tmpl")

		if err != nil {
			return templateCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./pkg/templates/*.layout.tmpl")
			if err != nil {
				return templateCache, err
			}
			templateCache[name] = ts
		}
	}

	return templateCache, nil
}
