package handler

import (
	"context"
	"encoding/json"

	"github.com/dadrus/heimdall/internal/heimdall"
)

type Authenticator interface {
	Authenticate(context.Context, RequestContext, *heimdall.SubjectContext) error
	WithConfig(config json.RawMessage) (Authenticator, error)
}
