// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type APIServiceFuzzingTest struct {
	Tags        *FuzzingTestTags           `json:"tags,omitempty"`
	TestDetails *APIServiceFuzzingProgress `json:"testDetails,omitempty"`
}

func (o *APIServiceFuzzingTest) GetTags() *FuzzingTestTags {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *APIServiceFuzzingTest) GetTestDetails() *APIServiceFuzzingProgress {
	if o == nil {
		return nil
	}
	return o.TestDetails
}
