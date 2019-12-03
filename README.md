### freshsales
a GoLang SDK to help track events for the https://www.freshsales.io CRM platform


## How ?

**import** the library in your go file 
```go
package freshsales

import "github.com/desertjinn/freshsales"
```
**initialise** to start tracking customer events 
```go
package yourpackage

import "github.com/desertjinn/freshsales"
...
...
    ...
    ...

    analytics := &Analytics{
        Domain: <your Freshsales domain host. eg: yourcompany.freshsales.io>, 
        Token: <your Freshsales API token>
    }
```
**identify** a New Customer
```go
    ...
    ...
    
    identifier := "a unique identifier for your customer"
    properties : map[string]interface{}{
        "key": "value", 
        "the customer": "properties",
    }   
    err := analytics.Identify(identifier, properties)
    if err != nil {
        fmt.Println(err)
    }
```
**track** an Event
```go
    ...
    ...
    
    identifier := "your customer's unique identifier"
    properties : map[string]interface{}{
        "key": "value", 
        "the event": "properties",
    }   
    err := analytics.TrackEvent(identifier, properties)
    if err != nil {
        fmt.Println(err)
    }
```
**track** a Page View
```go
    ...
    ...
    
    identifier := "your customer's unique identifier"
    url := "the page viewed"
    err := analytics.TrackPageView(identifier, url)
    if err != nil {
        fmt.Println(err)
    }
```
