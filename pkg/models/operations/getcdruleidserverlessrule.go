// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	"net/http"
)

type GetCdRuleIDServerlessRuleRequest struct {
	RuleID string `pathParam:"style=simple,explode=false,name=ruleId"`
}

func (o *GetCdRuleIDServerlessRuleRequest) GetRuleID() string {
	if o == nil {
		return ""
	}
	return o.RuleID
}

type GetCdRuleIDServerlessRuleResponse struct {
	// OK
	CdServerlessRule *shared.CdServerlessRule
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetCdRuleIDServerlessRuleResponse) GetCdServerlessRule() *shared.CdServerlessRule {
	if o == nil {
		return nil
	}
	return o.CdServerlessRule
}

func (o *GetCdRuleIDServerlessRuleResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetCdRuleIDServerlessRuleResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetCdRuleIDServerlessRuleResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
