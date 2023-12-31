// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	"net/http"
)

type GetCdRuleIDConnectionsRuleRequest struct {
	RuleID string `pathParam:"style=simple,explode=false,name=ruleId"`
}

func (o *GetCdRuleIDConnectionsRuleRequest) GetRuleID() string {
	if o == nil {
		return ""
	}
	return o.RuleID
}

type GetCdRuleIDConnectionsRuleResponse struct {
	// OK
	CdConnectionRule *shared.CdConnectionRule
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetCdRuleIDConnectionsRuleResponse) GetCdConnectionRule() *shared.CdConnectionRule {
	if o == nil {
		return nil
	}
	return o.CdConnectionRule
}

func (o *GetCdRuleIDConnectionsRuleResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetCdRuleIDConnectionsRuleResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetCdRuleIDConnectionsRuleResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
