# Calendar app (HTTP server/client)

HTTP server and client for event planning application. Server and client implemented using the standard library.

## Class diagram

### Client diagram

![](../../../doc/other_tasks/calendar/calendar_client_class.png?raw=true)

### Server diagram

![](../../../doc/other_tasks/calendar/calendar_server_class.png?raw=true)

### Middleware diagram

![](../../../doc/other_tasks/calendar/calendar_middleware_class.png?raw=true)

## API handlers

### Create event

Create new event in calendar for specified user.

```http request
POST /event
```

Parameters

|  Name   |             Type              |  In  | Required |                 Description                  |
|:-------:|:-----------------------------:|:----:|:--------:|:--------------------------------------------:|
| user_id |            integer            | Body |   Yes    | user ID for which the event is being created |
|  name   |            string             | Body |   Yes    |                name of event                 |
|  date   | string in format "YYYY-MM-DD" | Body |   Yes    |                date of event                 |

Response example

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
PUT /event
```

Parameters

|   Name   |             Type              |  In  | Required |                 Description                  |
|:--------:|:-----------------------------:|:----:|:--------:|:--------------------------------------------:|
| user_id  |            integer            | Body |   Yes    | user ID for which the event is being created |
|   name   |            string             | Body |   Yes    |              old name of event               |
|  date    | string in format "YYYY-MM-DD" | Body |   Yes    |              old date of event               |
| new_name |       new name of event       | Body |    No    |              new name of event               |
| new_date | string in format "YYYY-MM-DD" | Body |    No    |              new date of event               |

Response example

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
DELETE /event
```

Parameters

|  Name   |             Type              |  In  | Required |            Description             |
|:-------:|:-----------------------------:|:----:|:--------:|:----------------------------------:|
| user_id |            integer            | Body |   Yes    | user ID of the event to be deleted |
|  name   |            string             | Body |   Yes    |    name of event to be deleted     |
|  date   | string in format "YYYY-MM-DD" | Body |   Yes    |    date of event to be deleted     |

Response example

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

### Events for period

Get all events for the specified period. End of period not included. For example, 
if in the query start_date=2017-01-15 and end_date=2017-01-20, then the result 
will include the events of 2017-01-15 and will not include the events of 2017-01-20.

```http request
GET /events_for_period
```

Parameters

|    Name    |             Type              |  In   | Required |               Description                |
|:----------:|:-----------------------------:|:-----:|:--------:|:----------------------------------------:|
|  user_id   |            integer            | Query |   Yes    | ID of user for whom events are displayed |
| start_date | string in format "YYYY-MM-DD" | Query |   Yes    |            Period start date             |
|  end_date  | string in format "YYYY-MM-DD" | Query |   Yes    |             Period end date              |

Response example

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
