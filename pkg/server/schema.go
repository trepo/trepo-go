package server

var Schema = `
	schema {
		query: Query
	}

	type Query {
		shards: [Shard!]!
	}

  # A Trepo Shard
  type Shard {
    # The Shard ID
    id: String!
    # The Shard Name
    name: String!
  }
`
type Resolver struct{}

func (r *Resolver) Shards() []*shardResolver {
  var s []*shardResolver

  return s
}

type shardResolver struct {}

func (r *shardResolver) ID() string {
	return "1"
}

func (r *shardResolver) Name() string {
	return "Name"
}
