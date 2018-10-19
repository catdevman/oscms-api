package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

/* Test Helpers */
func expect(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Errorf("Expected %v (type %v) - Got %v (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
	}
}

func TestJSONDecodeError(t *testing.T) {
	res := httptest.NewRecorder()
	e := errors.New("error")

	JSONDecodeError(res, e)
	expect(t, res.Code, http.StatusInternalServerError)
	er := err{}
	json.NewDecoder(res.Body).Decode(&er)
	expect(t, er.Status, JSONDecodeErrorString)
	expect(t, er.Details, e.Error())
}

func TestDBError(t *testing.T) {
	res := httptest.NewRecorder()
	e := errors.New("error")

	DBError(res, e)
	expect(t, res.Code, http.StatusInternalServerError)
	er := err{}
	json.NewDecoder(res.Body).Decode(&er)
	expect(t, er.Status, DBErrorString)
	expect(t, er.Details, e.Error())
}
