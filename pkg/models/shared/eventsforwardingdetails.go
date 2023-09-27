// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type EventsForwardingDetails struct {
	EventsForwardingDetailsType EventsForwardingDetailsTypeEnum `json:"eventsForwardingDetailsType"`
	EventsToForward             []EventsToForward               `json:"eventsToForward"`
	ID                          *string                         `json:"id,omitempty"`
	Name                        string                          `json:"name"`
	URL                         *string                         `json:"url,omitempty"`
}

func (o *EventsForwardingDetails) GetEventsForwardingDetailsType() EventsForwardingDetailsTypeEnum {
	if o == nil {
		return EventsForwardingDetailsTypeEnum("")
	}
	return o.EventsForwardingDetailsType
}

func (o *EventsForwardingDetails) GetEventsToForward() []EventsToForward {
	if o == nil {
		return []EventsToForward{}
	}
	return o.EventsToForward
}

func (o *EventsForwardingDetails) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *EventsForwardingDetails) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *EventsForwardingDetails) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

type EventsForwardingDetailsInput struct {
	EventsForwardingDetailsType EventsForwardingDetailsTypeEnum `json:"eventsForwardingDetailsType"`
	EventsToForward             []EventsToForward               `json:"eventsToForward"`
	Name                        string                          `json:"name"`
	URL                         *string                         `json:"url,omitempty"`
}

func (o *EventsForwardingDetailsInput) GetEventsForwardingDetailsType() EventsForwardingDetailsTypeEnum {
	if o == nil {
		return EventsForwardingDetailsTypeEnum("")
	}
	return o.EventsForwardingDetailsType
}

func (o *EventsForwardingDetailsInput) GetEventsToForward() []EventsToForward {
	if o == nil {
		return []EventsToForward{}
	}
	return o.EventsToForward
}

func (o *EventsForwardingDetailsInput) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *EventsForwardingDetailsInput) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}
