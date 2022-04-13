# Calendar app (HTTP server/client)

HTTP server and client for event planning application. Server and client implemented using the standard library.

## Class diagram

![](../../../doc/other_tasks/calendar/calendar_class.png?raw=true)

## API handlers

### Create event

Create new event in calendar for specified user.

```http request
POST /create_event
```

Parameters

|  Name   |             Type              |  In  | Required |                 Description                  |
|:-------:|:-----------------------------:|:----:|:--------:|:--------------------------------------------:|
| user_id |            integer            | Body |   Yes    | user ID for which the event is being created |
|  name   |            string             | Body |   Yes    |                name of event                 |
|  date   | string in format "YYYY-MM-DD" | Body |   Yes    |                date of event                 |


Request example:

```http request
POST /create_event
Content-Type: application/x-www-form-urlencoded

user_id=1&name=event1&date=2019-12-20
```

Response

```
Status: 201 Created
```

```json
{
  "result": {
    "event": {
      "name": "event1",
      "date": "2019-12-20"
    }
  }
}
```

Status codes
* 201 Created
* 400 Bad Request

### Update event

Updates data of specified event.

```http request
POST /update_event
```

Parameters

|   Name   |             Type              |  In  | Required |                 Description                  |
|:--------:|:-----------------------------:|:----:|:--------:|:--------------------------------------------:|
| user_id  |            integer            | Body |   Yes    | user ID for which the event is being created |
| old_name |            string             | Body |   Yes    |              old name of event               |
| old_date | string in format "YYYY-MM-DD" | Body |   Yes    |              old date of event               |
| new_name |       new name of event       | Body |    No    |              new name of event               |
| new_date | string in format "YYYY-MM-DD" | Body |    No    |              new date of event               |

Request example:

```http request
POST /update_event
Content-Type: application/x-www-form-urlencoded

user_id=1&old_name=event1&old_date=2019-12-20&new_name=event2&new_date=2019-11-20
```

Response

```
Status: 200 OK
```

```json
{
  "result": {
    "event": {
      "name": "event2",
      "date": "2019-11-20"
    }
  }
}
```

Status codes
* 200 OK
* 400 Bad Request
* 404 Not Found


### Delete event

Delete specified event.

```http request
POST /delete_event
```

Parameters

|  Name   |             Type              |  In  | Required |            Description             |
|:-------:|:-----------------------------:|:----:|:--------:|:----------------------------------:|
| user_id |            integer            | Body |   Yes    | user ID of the event to be deleted |
|  name   |            string             | Body |   Yes    |    name of event to be deleted     |
|  date   | string in format "YYYY-MM-DD" | Body |   Yes    |    date of event to be deleted     |


Request example:

```http request
POST /delete_event
Content-Type: application/x-www-form-urlencoded

user_id=1&name=event1&date=2019-12-20
```

Response

```
Status: 204 No Content
```

```json
{
  "result": {
    "event": {
      "name": "event1",
      "date": "2019-12-20"
    }
  }
}
```

Status codes
* 204 No Content
* 400 Bad Request
* 404 Not Found


### Events for day

Get all events for the specified day.

```http request
GET /events_for_week
```

Parameters

|  Name   |             Type              |  In   | Required |               Description                |
|:-------:|:-----------------------------:|:-----:|:--------:|:----------------------------------------:|
| user_id |            integer            | Query |   Yes    | ID of user for whom events are displayed |
|  date   | string in format "YYYY-MM-DD" | Query |   Yes    |                   Date                   |


Request example:

```http request
GET /events_for_week
Content-Type: application/x-www-form-urlencoded

user_id=1&date=2019-12-20
```

Response

```
Status: 200 OK
```

```json
{
  "result": {
    "events": [
      {
      "name": "event1",
      "date": "2019-12-20"
      },
      {
        "name": "event2",
        "date": "2019-12-20"
      },
      {
        "name": "event3",
        "date": "2019-12-20"
      }
    ]
  }
}
```

Status codes
* 200 OK
* 400 Bad Request

### Events for week

Get all events for the week. The week is counted from the date specified in the parameters.
For example, if the date 2019-12-13 is indicated in the request, then the response will
include all events from 2019-12-13 to 2019-12-19 inclusive.

```http request
GET /events_for_week
```

Parameters

|  Name   |             Type              |  In   | Required |                 Description                 |
|:-------:|:-----------------------------:|:-----:|:--------:|:-------------------------------------------:|
| user_id |            integer            | Query |   Yes    |  ID of user for whom events are displayed   |
|  date   | string in format "YYYY-MM-DD" | Query |   Yes    |   The date that is the start of the week    |


Request example:

```http request
GET /events_for_week
Content-Type: application/x-www-form-urlencoded

user_id=1&date=2019-12-20
```

Response

```
Status: 200 OK
```

```json
{
  "result": {
    "events": [
      {
      "name": "event1",
      "date": "2019-12-20"
      },
      {
        "name": "event2",
        "date": "2019-12-20"
      },
      {
        "name": "event3",
        "date": "2019-12-26"
      }
    ]
  }
}
```

Status codes
* 200 OK
* 400 Bad Request


### Events for month

Get all events for the month. The month is determined by the date specified in the parameters. 
For example, if the date 2019-12-13 is indicated in the request, then the response will 
include all events from 2019-12-01 to 2019-12-31 inclusive.

```http request
GET /events_for_month
```

Parameters

|  Name   |             Type              |  In   | Required |                Description                |
|:-------:|:-----------------------------:|:-----:|:--------:|:-----------------------------------------:|
| user_id |            integer            | Query |   Yes    | ID of user for whom events are displayed  |
|  date   | string in format "YYYY-MM-DD" | Query |   Yes    | The date by which the month is determined |


Request example:

```http request
GET /events_for_month
Content-Type: application/x-www-form-urlencoded

user_id=1&date=2019-12-20
```

Response

```
Status: 200 OK
```

```json
{
  "result": {
    "events": [
      {
      "name": "event1",
      "date": "2019-12-01"
      },
      {
        "name": "event2",
        "date": "2019-12-20"
      },
      {
        "name": "event3",
        "date": "2019-12-31"
      }
    ]
  }
}
```

Status codes
* 200 OK
* 400 Bad Request
