---
id: access
title: Authentication & Authorization
---
### Authentication/Authorization
Authentication/Authorization is handled via the `Authorization` header. The contents are passed to the specified Authorization Endpoint, and the resulting error code is used on failure (`403` returned from the Auth URL passes through along with the failure message).

Request Body
```
{
  auth: '', // The contents of the Authorization header, including the type
  action: '', // The action being performed
  shard: '', // The shard, or null if not applicable
}
```

The response body must be `200` to continue

Response Body
```
{
  error: '', // On error, this will be returned as the error message to the original request
  message: '', // TODO ? An additional message to append to the passed in message
  author: 'name <email>',
}
```
