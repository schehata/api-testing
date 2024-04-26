// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// CreateAnnotation implements createAnnotation operation.
	//
	// Create a new annotation.
	//
	// POST /annotations
	CreateAnnotation(ctx context.Context, req *NewAnnotation) (CreateAnnotationRes, error)
	// DeleteAnnotation implements deleteAnnotation operation.
	//
	// Delete an annotation by id.
	//
	// DELETE /annotations/{id}
	DeleteAnnotation(ctx context.Context, params DeleteAnnotationParams) (DeleteAnnotationRes, error)
	// GetAnnotation implements getAnnotation operation.
	//
	// Get a single annotation by id.
	//
	// GET /annotations/{id}
	GetAnnotation(ctx context.Context, params GetAnnotationParams) (GetAnnotationRes, error)
	// GetAnnotations implements getAnnotations operation.
	//
	// Get all annotations and filter by datasets or timerange.
	//
	// GET /annotations
	GetAnnotations(ctx context.Context, params GetAnnotationsParams) (GetAnnotationsRes, error)
	// UpdateAnnotation implements updateAnnotation operation.
	//
	// Update an annotation.
	//
	// PUT /annotations/{id}
	UpdateAnnotation(ctx context.Context, req *UpdateAnnotation, params UpdateAnnotationParams) (UpdateAnnotationRes, error)
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h   Handler
	sec SecurityHandler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, sec SecurityHandler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		sec:        sec,
		baseServer: s,
	}, nil
}