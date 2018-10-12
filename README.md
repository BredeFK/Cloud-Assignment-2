Note:
# Assignment 2: Persistence, Live feed, Notifications
## Root url:
https://oblig2-klausen.herokuapp.com/

## Note:
I tried to build without 'gopkg.in' in vendor folder, but I couldn't do it.

## Instructions
Develop a service, that will allow a user to monitor a currency ticker, and notify a webhook upon certain conditions are met, such as the price falling below or going above a given threshold. The API must allow the user to specify the base currency (for simplicity, if the base is not EURO, your service should respond that this is not yet implemented), the target currency (arbitrary, from all currencies supported by fixer), and the min and max price for the event to trigger the notification. The notification will be provided via a webhook specified by the user, and multiple webhooks should be provided (predefined types).

In addition, the service will be able to monitor the currencies (all from the http://api.fixer.io/latest?base=EUR query) at regular time intervals (once a day) and store the results in a MongoDB database. The system will allow the user to query for the "latest" ticker of given currency pair between EUR/xxx, and also, to query for the "running average" of the last 3 days.

**Clarification:** The service will only need to support the use of EUR as base currency. Other currencies could be supported, if you want, but are not required. If the user specifies base currency other than EURO, your service can respond: "not implemented", along with the corresponding status code.

## Service specification
### Registration of new webhook
New webhooks can be registered using POST requests with the following schema. Note we will use /root as a placeholder for the root path of your web service (i.e. the path you will submit to the submission system). For example, if your web service runs on https://localhost:8080/exchange, then this is the root path you would submit.

#### Request
Path: /root

Payload specification:
```
{
    "webhookURL": {
      "type": "string"
    },
    "baseCurrency": {
      "type": "string"
    },
    "targetCurrency": {
      "type": "string"
    },
    "minTriggerValue": {
      "type": "number"
    },
    "maxTriggerValue": {
      "type": "number"
    }
}
```

Example:

```
{
    "webhookURL": "http://remoteUrl:8080/randomWebhookPath",
    "baseCurrency": "EUR",
    "targetCurrency": "NOK",
    "minTriggerValue": 1.50,
    "maxTriggerValue": 2.55
}
```

#### Response
The response body should contain the id of the created resource, as string. Note, the response body will contain only the created id, as string, not the entire path; no json encoding. Response code upon success should be 200 or 201.

### Invoking a registered webhook
When invoking a registered webhook, use the following payload specification:

#### Request
Path: /webhookUrl

Payload Specification:
```
{
    "baseCurrency": {
      "type": "string"
    },
    "targetCurrency": {
      "type": "string"
    },
    "currentRate": {
      "type": "number"
    },
    "minTriggerValue": {
      "type": "number"
    },
    "maxTriggerValue": {
      "type": "number"
    }
}
```

Example:

```
{
    "baseCurrency": "EUR",
    "targetCurrency": "NOK",
    "currentRate": 2.75,
    "minTriggerValue": 1.50,
    "maxTriggerValue": 2.55
}
```
#### Response
Upon successful notification you will receive either status code 200 (for trigger) or 204 (when no trigger).

### Accessing registered webhooks
Registered webhooks should be accessible using the GET method and the webhook id generated during registration.

#### Request GET
Path: /root/{id}

#### Response
Body: (see POST request)

### Deleting registered webhooks
Registered webhooks can further be deleted using the DELETE method and the webhook id.

Path: /root/{id}

### Retrieving the latest currency exchange rates

#### Request
Path:  /root/latest

Body:
```
{
    "baseCurrency": "EUR",
    "targetCurrency": "NOK"
}
```
#### Response
The response should contain only the latest exchange rate value (no json tags).

Body: value

Example: 1.56

### Retrieving the running average over the past seven days

#### Request
Path: /root/average

Body:
```
{
    "baseCurrency": "EUR",
    "targetCurrency": "NOK"
}
```

#### Response
The response should contain only the running average value (no json tags).

Body: value

Example: 2.75

### Addendum: Trigger webhooks for testing purposes
This trigger (Method: GET) invokes all webhooks (i.e. bypasses the timed trigger) and sends the payload as specified under 'Invoking a registered webhook'. This functionality is meant for testing and evaluation purposes.

#### Request
Path: /root/evaluationtrigger

Body: empty

#### Response
Reasonable status code

Important: For all requests, ensure that you use appropriate status codes and semantics (see [IETF RFC 7231](https://tools.ietf.org/html/rfc7231) for details).
