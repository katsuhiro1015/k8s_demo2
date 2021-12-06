// Package generate provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.0 DO NOT EDIT.
package generate

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"3d8CAAD//xmGt4NVEgAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}