// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/testpan-go/pkg/utils"
	"net/http"
)

type GetImagesImagesHashRequest struct {
	// image hash to search for ( as prefix and suffix )
	ImageHash *string `queryParam:"style=form,explode=true,name=imageHash"`
	// The number of entries to return (pagination)
	MaxResults *float64 `default:"100" queryParam:"style=form,explode=true,name=maxResults"`
}

func (g GetImagesImagesHashRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetImagesImagesHashRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetImagesImagesHashRequest) GetImageHash() *string {
	if o == nil {
		return nil
	}
	return o.ImageHash
}

func (o *GetImagesImagesHashRequest) GetMaxResults() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxResults
}

type GetImagesImagesHashResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Success
	Strings []string
}

func (o *GetImagesImagesHashResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetImagesImagesHashResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetImagesImagesHashResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetImagesImagesHashResponse) GetStrings() []string {
	if o == nil {
		return nil
	}
	return o.Strings
}
