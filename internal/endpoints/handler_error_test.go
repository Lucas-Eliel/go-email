package endpoints

import (
	"emailn/internal/internalerrors"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_HalderError_when_enpoint_returns_internal_error(t *testing.T) {
	assert := assert.New(t)

	endpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
		return nil, 0, internalerrors.ErrInternalError
	}

	handlerFunc := HandlerError(endpoint)

	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	handlerFunc.ServeHTTP(res, req)

	assert.Equal(http.StatusInternalServerError, res.Code)
	assert.Contains(res.Body.String(), internalerrors.ErrInternalError.Error())
}

func Test_HalderError_when_enpoint_returns_domain_error(t *testing.T) {
	assert := assert.New(t)

	endpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
		return nil, 0, errors.New("Domain error")
	}

	handlerFunc := HandlerError(endpoint)

	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	handlerFunc.ServeHTTP(res, req)

	assert.Equal(http.StatusBadRequest, res.Code)
	assert.Contains(res.Body.String(), "Domain error")
}

func Test_HalderError_when_enpoint_returns_object_and_status(t *testing.T) {
	assert := assert.New(t)
	type BodyForTest struct {
		Id int
	}
	objetcExpect := BodyForTest{Id: 2}

	endpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
		return objetcExpect, http.StatusCreated, nil
	}

	handlerFunc := HandlerError(endpoint)

	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	handlerFunc.ServeHTTP(res, req)

	assert.Equal(http.StatusCreated, res.Code)
	objReturned := BodyForTest{}
	json.Unmarshal(res.Body.Bytes(), &objReturned)
	assert.Equal(objetcExpect, objReturned)
}
