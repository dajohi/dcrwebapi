// Copyright (c) 2017-2018 The Decred developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strings"
)

const (
	// semanticAlphabet defines the allowed characters for the pre-release
	// portion of a semantic version string.
	semanticAlphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
)

// normalizeSemString returns the passed string stripped of all characters
// which are not valid according to the provided semantic versioning alphabet.
func normalizeSemString(str, alphabet string) string {
	var result bytes.Buffer
	for _, r := range str {
		if strings.ContainsRune(alphabet, r) {
			result.WriteRune(r)
		}
	}
	return result.String()
}

// NormalizePreRelString returns the passed string stripped of all characters
// which are not valid according to the semantic versioning guidelines for
// pre-release strings.  In particular they MUST only contain characters in
// semanticAlphabet.
func NormalizePreRelString(str string) string {
	return normalizeSemString(str, semanticAlphabet)
}

// writeJSONResponse convenience func for writing json responses
func writeJSONResponse(writer *http.ResponseWriter, code int,
	respJSON *[]byte) {
	(*writer).Header().Set("Content-Type", "application/json")
	(*writer).Header().Set("Strict-Transport-Security", "max-age=15552001")
	(*writer).Header().Set("Vary", "Accept-Encoding")
	(*writer).WriteHeader(code)
	(*writer).Write(*respJSON)
}

// writeSVGResponse convenience func for writing svg responses
func writeSVGResponse(writer *http.ResponseWriter, code int,
	svg *string) {
	(*writer).Header().Set("Strict-Transport-Security", "max-age=15552001")
	(*writer).Header().Set("Vary", "Accept-Encoding")
	(*writer).Header().Set("Content-Type", "image/svg+xml")
	(*writer).WriteHeader(code)
	fmt.Fprint(*writer, *svg)
}

// writeJSONErrorResponse convenience func for writing json error responses
func writeJSONErrorResponse(writer *http.ResponseWriter, code int, err error) {
	errorBody := map[string]interface{}{}
	errorBody["error"] = err.Error()
	errorJSON, _ := json.Marshal(errorBody)
	(*writer).Header().Set("Content-Type", "application/json")
	(*writer).Header().Set("Strict-Transport-Security", "max-age=15552001")
	(*writer).Header().Set("Vary", "Accept-Encoding")
	(*writer).WriteHeader(code)
	(*writer).Write(errorJSON)
}

// round rounding func
func round(f float64, places uint) float64 {
	shift := math.Pow(10, float64(places))
	return math.Floor(f*shift+.5) / shift
}
