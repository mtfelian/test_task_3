package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	e "github.com/mtfelian/error"
	"github.com/mtfelian/validation"
)

const (
	// ErrorInvalidInput is an error code for invalid input
	ErrorInvalidInput uint = iota + 1
	// ErrorValidation is an error code for validation error
	ErrorValidation
	// ErrorStorage is an error code for storage error
	ErrorStorage
)

func extractBody(c *gin.Context, to interface{}) error {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, to)
}

// GetParamBody represents request body for the GetParam API handler
type GetParamBody struct {
	Type string
	Data string
}

// Validate the params. Returns:
// - bool if validation has errors
// - string errors description
func (p *GetParamBody) Validate() (bool, string) {
	p.Type, p.Data = strings.TrimSpace(p.Type), strings.TrimSpace(p.Data)
	v := &validation.Validation{}
	v.Check(p.Type, validation.Required{}).Message("Body field is required")
	v.Check(p.Data, validation.Required{}).Message("Data field is required")
	return v.HasErrors(), v.String()
}

// GetParam is an API handler to get configuration entry
func GetParam(c *gin.Context) {
	var body GetParamBody
	if err := extractBody(c, &body); err != nil {
		c.JSON(http.StatusUnprocessableEntity, e.NewError(ErrorInvalidInput, err))
		return
	}

	hasErrors, details := body.Validate()
	if hasErrors {
		c.JSON(http.StatusUnprocessableEntity, e.NewErrorf(ErrorValidation, details))
		return
	}

}
