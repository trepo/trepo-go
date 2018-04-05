---
id: format
title: Format
---
### Log Entry
- `id` - The id of the entry (`required`, incrementing integer)
- `prev` - The id of the previous entry (`optional`, integer, used for CAS)
- `author` - The entry author (`optional`, `Name <email>` format)
- `message` - The entry message (`optional`, string)
- `nodes` - The entry nodes (`required`, array)
- `edges` - The entry edges (`required`, array)

### Node
- `id` - The node id (`required`, ID, immutable)
- `shard` - The shard identifier (`required if props is null`)
- `props` - A map of properties (`required if shard is null`)

### Edge
- `id` - The node id (`required`, ID, immutable)
- `from` - The id of the from node
- `to` - The id of the to node
- `props` - A map of properties (`required if shard is null`)

### ID
The concatenation of the `type` and a `uuidv4`, separated by a `:` (i.e. `Person:c38d39c5-b67c-4f1c-b480-0498c6648d45`).
The uuidv4 is required to be unique across all nodes, so it would be disallowed to re-use a uuidv4 for any reason. In materialized views that are graph based, the `uuidv4` may be used as the element id and the `type` may be used as a label. The `uuidv4` may be stored as an integer internally for performance reasons.

### Type
They type can be any string matching `[A-Za-z][A-Za-z0-9_]{0,217}`

### Props
`key` can be any string matching `[A-Za-z][A-Za-z0-9_]{0,254}`
`value` must be one of: `string`, `integer (32 bit)`, `boolean`, or a `homogenous array` of one of the previous types


### Notes
- Does an entry have to contain the from/to nodes for every edge? Is it optional? Answer: Yes. The edge from/to ID contains enough information.
- The amount of information per node is kept small on purpose, as overrides are on an element by element basis
- To delete a Node, `shard` and `props` are set to Null. To delete an Edge, `props` is set to Null.
