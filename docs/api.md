---
id: api
title: API
---

|URL|Method|Description|
|---|------|-----------|
|/v1/|GET|Get Server health and version information|
|/v1/shard|GET|Get a list of shards|
|/v1/shard|POST|Create a shard|
|/v1/shard/:id|GET|Get a shard|
|/v1/shard/:id|PUT|Create a shard (only if shard does not exist)|
|/v1/shard/:id|PATCH|Update a shard|
|/v1/shard/:id|DELETE|Delete a shard (and all entries)|
|/v1/shard/:id/entry|GET|Get entries|
|/v1/shard/:id/entry|POST|Append an entry (optionally with a CAS check)|
|/v1/shard/:id/entry/:id|GET|Get an entry|
|/v1/shard/:id/snapshot|GET|Get a snapshot|
