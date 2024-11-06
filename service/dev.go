package service

import (
	"context"
	"io"

	"github.com/a-h/templ"
	"github.com/datasektionen/logout/pkg/config"
	"github.com/datasektionen/logout/templates"
)

func (s *Service) DevLoginForm() templ.Component {
	if !config.Config.Dev {
		return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			return nil
		})
	}
	return templates.DevLoginForm()
}
