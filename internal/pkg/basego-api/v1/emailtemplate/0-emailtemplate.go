package emailtemplate

import (
	"bytes"
	"fmt"
	"html/template"
	"sync"

	"github.com/jonylim/basego/internal/pkg/common/logger"
)

type parsedTemplate struct {
	tpl *template.Template
	err error
}

var mapTemplates = make(map[string]*parsedTemplate)
var mutexTemplate sync.Mutex

func getByFilename(filename string) (*template.Template, error) {
	mutexTemplate.Lock()
	defer mutexTemplate.Unlock()

	t, ok := mapTemplates[filename]
	if t == nil || !ok {
		tpl, err := template.ParseFiles("assets/templates/email/" + filename)
		t = &parsedTemplate{tpl, err}
		mapTemplates[filename] = t
	}
	return t.tpl, t.err
}

func generateFromTemplate(templateFilename string, data interface{}) (body string, err error) {
	var t *template.Template
	t, err = getByFilename(templateFilename)
	if err != nil {
		logger.Fatal("emailtemplate", fmt.Sprintf("generateFromTemplate: %v", logger.FromError(err)))
		return
	}
	var buf bytes.Buffer
	if err = t.Execute(&buf, data); err != nil {
		logger.Fatal("emailtemplate", fmt.Sprintf("generateFromTemplate: %v", logger.FromError(err)))
		return
	}
	body = string(buf.Bytes())
	return
}
