package freshsales

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// the Freshsales analytics interface
type AnalyticsInterface interface {
	Identify(identifier string, properties map[string]interface{}) (err error)
	TrackEvent(identifier string, properties map[string]interface{}) (err error)
	TrackPageView(identifier string, url string) (err error)
}

type Analytics struct {
	Domain string
	Token string
}

// validate a Freshsales request to ensure all required criteria are met
func (analytics *Analytics) validate(action string, data map[string]interface{}) (err error) {
	if strings.TrimSpace(analytics.Domain) == "" || strings.TrimSpace(analytics.Token) == "" {
		err = errors.New("either the Freshsales domain or appToken you've provided during the initialization " +
			"of Freshsales Analytics is empty")
	} else if data == nil {
		err = errors.New("data passed is empty")
	} else if strings.TrimSpace(action) == "" {
		err = errors.New("action is empty")
	} else if data[IdentifierKey] == nil {
		// customer's unique identifier not found
		err = errors.New("identifier passed is empty")
	} else if action == PageViewsEndpoint {
		if data[PageViewKey] == nil {
			// customer's page_view payload missing
			err = errors.New("page_view payload is empty")
		} else {
			if pageViews, ok := data[PageViewKey].(map[string]interface{}); !ok || pageViews == nil ||
				pageViews[UrlKey] == nil {

				// customer viewed URL is missing
				err = errors.New("page URL passed to trackPageView method is empty")
			}
		}
	}
	return
}

// make a POST request to the Freshsales tracking endpoint
func (analytics *Analytics) post(action string, data map[string]interface{}) (err error){
	if err = analytics.validate(action, data); err == nil {
		data[TokenKey] = analytics.Token
		data[SdkKey] = "golang"
		path := fmt.Sprintf("%s/track/%s", analytics.Domain, action)
		err = request(http.MethodPost, path, data)
	}
	return
}

// identify new customer's
func (analytics *Analytics) Identify(identifier string, properties map[string]interface{}) (err error) {
	data := map[string]interface{}{
		IdentifierKey: identifier,
		VisitorKey: properties,
	}
	err = analytics.post(IdentifyEndpoint, data)
	return
}

// track a customer's events
func (analytics *Analytics) TrackEvent(identifier string, properties map[string]interface{}) (err error) {
	data := map[string]interface{}{
		IdentifierKey: identifier,
		EventKey: properties,
	}
	err = analytics.post(EventsEndpoint, data)
	return
}

// track a customer's viewed pages
func (analytics *Analytics) TrackPageView(identifier string, url string) (err error) {
	data := map[string]interface{}{
		IdentifierKey: identifier,
		PageViewKey: map[string]interface{}{UrlKey:url},
	}
	err = analytics.post(PageViewsEndpoint, data)
	return
}