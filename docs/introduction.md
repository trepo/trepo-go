---
id: introduction
title: Introduction
---
Under the covers Trepo is a distributed graph database broken up into commit logs called shards. Each shard, or sub graph, is hosted by an individual or corporation. Each node or edge in the graph has a globally unique id (uuid) and an immutable type, or label associated with it. You can create your own unique "view" of Trepo by determining the override order of the shards.
