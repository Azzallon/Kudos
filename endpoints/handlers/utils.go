package handlers

import (
	"kudos/tests"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func RunGenericGetRequestTest(t *testing.T, expectedCode int, fHandler func(ctx echo.Context) error,
	target string, ctxParams *tests.ContextParams, expectedJson string) {

	e := echo.New()

	request := httptest.NewRequest(http.MethodGet, target, nil)
	recorder := httptest.NewRecorder()
	c := e.NewContext(request, recorder)

	if ctxParams != nil {
		c.SetParamNames(ctxParams.Keys...)
		c.SetParamValues(ctxParams.Values...)
	}

	if expectedCode == http.StatusOK {
		if assert.NoError(t, fHandler(c)) {
			assert.Equal(t, expectedCode, recorder.Code)
			assert.JSONEq(t, expectedJson, recorder.Body.String())
		}
	} else {
		_ = fHandler(c)
		assert.Equal(t, expectedCode, recorder.Code)
		assert.JSONEq(t, expectedJson, recorder.Body.String())
	}
}

func RunGenericPostRequestTest(
	t *testing.T, expectedCode int, fHandler func(ctx echo.Context) error,
	target string, ctxParams *tests.ContextParams, expectedJson, bodyJson string) {
	ast := assert.New(t)

	request := httptest.NewRequest(
		http.MethodPost,
		target,
		strings.NewReader(bodyJson),
	)

	e := echo.New()
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recorder := httptest.NewRecorder()
	c := e.NewContext(request, recorder)
	if ctxParams != nil {
		c.SetParamNames(ctxParams.Keys...)
		c.SetParamValues(ctxParams.Values...)
	}
	if ast.NoError(fHandler(c)) {
		ast.Equal(expectedCode, recorder.Code)
		if expectedJson != "" {
			ast.JSONEq(expectedJson, recorder.Body.String())
		}
	}
}

func RunGenericPutRequestTest(t *testing.T, expectedCode int, fHandler func(ctx echo.Context) error,
	target string, ctxParams *tests.ContextParams, expectedJson string) {

	e := echo.New()

	request := httptest.NewRequest(http.MethodPut, target, nil)
	recorder := httptest.NewRecorder()
	c := e.NewContext(request, recorder)

	if ctxParams != nil {
		c.SetParamNames(ctxParams.Keys...)
		c.SetParamValues(ctxParams.Values...)
	}

	if expectedCode == http.StatusOK {
		if assert.NoError(t, fHandler(c)) {
			assert.Equal(t, expectedCode, recorder.Code)
			assert.JSONEq(t, expectedJson, recorder.Body.String())
		}
	} else {
		_ = fHandler(c)
		assert.Equal(t, expectedCode, recorder.Code)
		assert.JSONEq(t, expectedJson, recorder.Body.String())
	}
}
