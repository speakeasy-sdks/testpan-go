// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	"net/http"
)

type PutCdRuleIDServerlessRuleRequest struct {
	CdServerlessRule shared.CdServerlessRule `request:"mediaType=application/json"`
	RuleID           string                  `pathParam:"style=simple,explode=false,name=ruleId"`
}

func (o *PutCdRuleIDServerlessRuleRequest) GetCdServerlessRule() shared.CdServerlessRule {
	if o == nil {
		return shared.CdServerlessRule{}
	}
	return o.CdServerlessRule
}

func (o *PutCdRuleIDServerlessRuleRequest) GetRuleID() string {
	if o == nil {
		return ""
	}
	return o.RuleID
}

type PutCdRuleIDServerlessRuleResponse struct {
	// updated
	CdServerlessRule *shared.CdServerlessRule
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutCdRuleIDServerlessRuleResponse) GetCdServerlessRule() *shared.CdServerlessRule {
	if o == nil {
		return nil
	}
	return o.CdServerlessRule
}

func (o *PutCdRuleIDServerlessRuleResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutCdRuleIDServerlessRuleResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutCdRuleIDServerlessRuleResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
