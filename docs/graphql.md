---
id: graphql
title: GraphQL
---
GraphQL Schema

```graphql
schema {
  query: Query
  mutation: Mutation
  subscription: Subscription
}

type Query {
  shards: Shards
  shard(id: ShardID!): Shard
  commits(shard: ShardID!): Commits
  commit(shard: ShardID!, id: CommitID): Commit
}

type Mutation {
  createShard(input: CreateShardRequest!): CreateShardResponse
  createCommit(input: CreateCommitRequest!): CreateCommitResponse
}

type Subscription {
  commitAdded(shards: [ShardID!]): Commit
}

type Shard {
  id: ShardID!
  name: String!
}

type Commit {
  id: CommitID!
  shardID: ShardID!
  author: String
  message: String
  nodes: [Node!]!
  edges: [Edge!]!
}

type Node {
  id: ElementID!
  shard: ShardID
  props: Properties
}

type Edge {
  id: ElementID!
  from: ElementID!
  to: ElementID!
  props: Properties
}

type Error {
  code: Int!
  message: String!
}

input CreateShardRequest {
  name: String!
}

type CreateShardResponse {
  shard: Shard
}

input CreateCommitRequest {
  shardID: ShardID!
  prev: CommitID
  author: String
  message: String
  nodes: [Node!]!
  edges: [Edge!]!
}

type CreateCommitResponse {
  error: Error
  commit: Commit
}

scalar ShardID
scalar CommitID
scalar ElementID
scalar Properties
```
