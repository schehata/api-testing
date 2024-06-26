// Code generated by ogen, DO NOT EDIT.

package api

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

// DeleteAnnotationParams is parameters of deleteAnnotation operation.
type DeleteAnnotationParams struct {
	// The id of the annotation.
	ID string
}

func unpackDeleteAnnotationParams(packed middleware.Parameters) (params DeleteAnnotationParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(string)
	}
	return params
}

func decodeDeleteAnnotationParams(args [1]string, argsEscaped bool, r *http.Request) (params DeleteAnnotationParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
			if err := func() error {
				if err := (validate.String{
					MinLength:    0,
					MinLengthSet: false,
					MaxLength:    0,
					MaxLengthSet: false,
					Email:        false,
					Hostname:     false,
					Regex:        regexMap["^ann"],
				}).Validate(string(params.ID)); err != nil {
					return errors.Wrap(err, "string")
				}
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// GetAnnotationParams is parameters of getAnnotation operation.
type GetAnnotationParams struct {
	// The id of the annotation.
	ID string
}

func unpackGetAnnotationParams(packed middleware.Parameters) (params GetAnnotationParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(string)
	}
	return params
}

func decodeGetAnnotationParams(args [1]string, argsEscaped bool, r *http.Request) (params GetAnnotationParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
			if err := func() error {
				if err := (validate.String{
					MinLength:    0,
					MinLengthSet: false,
					MaxLength:    0,
					MaxLengthSet: false,
					Email:        false,
					Hostname:     false,
					Regex:        regexMap["^ann"],
				}).Validate(string(params.ID)); err != nil {
					return errors.Wrap(err, "string")
				}
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// GetAnnotationsParams is parameters of getAnnotations operation.
type GetAnnotationsParams struct {
	// The datasets to filter by.
	Datasets []string
	// If set, will filter to events after this date. Should be in RFC3339.
	Start OptDateTime
	// If set, will filter to events before this date. Should be in RFC3339.
	End OptDateTime
}

func unpackGetAnnotationsParams(packed middleware.Parameters) (params GetAnnotationsParams) {
	{
		key := middleware.ParameterKey{
			Name: "datasets",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Datasets = v.([]string)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "start",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Start = v.(OptDateTime)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "end",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.End = v.(OptDateTime)
		}
	}
	return params
}

func decodeGetAnnotationsParams(args [0]string, argsEscaped bool, r *http.Request) (params GetAnnotationsParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: datasets.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "datasets",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				return d.DecodeArray(func(d uri.Decoder) error {
					var paramsDotDatasetsVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						paramsDotDatasetsVal = c
						return nil
					}(); err != nil {
						return err
					}
					params.Datasets = append(params.Datasets, paramsDotDatasetsVal)
					return nil
				})
			}); err != nil {
				return err
			}
			if err := func() error {
				var failures []validate.FieldError
				for i, elem := range params.Datasets {
					if err := func() error {
						if err := (validate.String{
							MinLength:    1,
							MinLengthSet: true,
							MaxLength:    0,
							MaxLengthSet: false,
							Email:        false,
							Hostname:     false,
							Regex:        nil,
						}).Validate(string(elem)); err != nil {
							return errors.Wrap(err, "string")
						}
						return nil
					}(); err != nil {
						failures = append(failures, validate.FieldError{
							Name:  fmt.Sprintf("[%d]", i),
							Error: err,
						})
					}
				}
				if len(failures) > 0 {
					return &validate.Error{Fields: failures}
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "datasets",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: start.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "start",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotStartVal time.Time
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToDateTime(val)
					if err != nil {
						return err
					}

					paramsDotStartVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Start.SetTo(paramsDotStartVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "start",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: end.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "end",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotEndVal time.Time
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToDateTime(val)
					if err != nil {
						return err
					}

					paramsDotEndVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.End.SetTo(paramsDotEndVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "end",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// UpdateAnnotationParams is parameters of updateAnnotation operation.
type UpdateAnnotationParams struct {
	// The id of the annotation.
	ID string
}

func unpackUpdateAnnotationParams(packed middleware.Parameters) (params UpdateAnnotationParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(string)
	}
	return params
}

func decodeUpdateAnnotationParams(args [1]string, argsEscaped bool, r *http.Request) (params UpdateAnnotationParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
			if err := func() error {
				if err := (validate.String{
					MinLength:    0,
					MinLengthSet: false,
					MaxLength:    0,
					MaxLengthSet: false,
					Email:        false,
					Hostname:     false,
					Regex:        regexMap["^ann"],
				}).Validate(string(params.ID)); err != nil {
					return errors.Wrap(err, "string")
				}
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}
