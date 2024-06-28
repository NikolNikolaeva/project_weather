package http

import (
	"net/url"
	"path"
	"strings"

	"github.com/NikolNikolaeva/project_weather/it/internal/gomisc/lang"
	arrays "github.com/NikolNikolaeva/project_weather/it/internal/gomisc/lang/array"
	"github.com/NikolNikolaeva/project_weather/it/internal/gomisc/lang/maps"
)

type URLBuilder interface {
	WithHost(host string) URLBuilder
	WithScheme(scheme string) URLBuilder
	WithPath(parts ...string) URLBuilder
	AppendPath(parts ...string) URLBuilder
	WithQuery(query map[string][]string) URLBuilder

	Build() *url.URL
}

func NewURLBuilder(urls ...*url.URL) URLBuilder {
	urls = append(urls, &url.URL{})

	return &_URLBuilder{
		host:   urls[0].Host,
		scheme: urls[0].Scheme,
		query:  urls[0].Query(),
		path: arrays.Map(
			strings.Split(urls[0].EscapedPath(), "/"),
			func(_ int, fragment string) string {
				return lang.First(url.PathUnescape(fragment))
			},
		),
	}
}

type _URLBuilder struct {
	host   string
	scheme string
	path   []string
	query  map[string][]string
}

func (self *_URLBuilder) WithHost(host string) URLBuilder {
	self.host = host
	return self
}

func (self *_URLBuilder) WithScheme(scheme string) URLBuilder {
	self.scheme = scheme
	return self
}

func (self *_URLBuilder) WithPath(parts ...string) URLBuilder {
	self.path = append(make([]string, 0), parts...)
	return self
}

func (self *_URLBuilder) AppendPath(parts ...string) URLBuilder {
	self.path = append(self.path, parts...)
	return self
}

func (self *_URLBuilder) WithQuery(query map[string][]string) URLBuilder {
	self.query = maps.Merge(query)
	return self
}

func (self *_URLBuilder) Build() *url.URL {
	return &url.URL{
		Host:     self.host,
		Scheme:   self.scheme,
		Path:     path.Join(self.path...),
		RawQuery: url.Values(self.query).Encode(),
		RawPath: path.Join(
			arrays.Map(self.path, func(_ int, segment string) string {
				return url.PathEscape(segment)
			})...,
		),
	}
}
