# n-labels

# Background
Aim of this service is to provide ability to manage set of entities by associating entities with tags, also called as labels. Labels are attached to an entity in order to achive following benefits

1. Labels provides a way to organise collection of entities
2. Helps retriving one or more related entities by label
3. Allow special treatment to certain entities by categorising them using labels
4. Labels can be use to assign entities to users
5. Labels can be use to control the visibility of an entity

While this list covers some common usecases it is no way exhaustive.

# Requirments
Label service should provide following abilities
- Ability to create and delete label
- Ability to attach label to an entity
- Ability to search labels by name or substring
- Ability to list all the labels of an entity
- Ability to list all the entities attached to a label

# Service SLA
- Availability
Label service should target 99.99% uptime
- Latency
Label service should aim for less than 100 ms P95 latency
- Throughput
Label service should provide QPS of 2000 per node
- Freshness
Label service should provide newly created labels in search results immediately 

# Architecture
![image](https://user-images.githubusercontent.com/10060860/115158869-02ffbc00-a0ae-11eb-9c77-8a7fcb58ba39.png)

# Implementation
## API Interface

```go

type Interface{
  Create(labelName string) string    // create new label
  Delete(labelID uuid) bool          // delete exisiting label
  Attach(labelID uuid, entityID uuid) bool    // attach label to an entity
  Search(keyword string) []string    // search labels by keyword
  ListLabels(entityID uuid) []uuid // list labels of an entity
  ListEntities(labelID uuid) []uuid // list entities of a label
}

```

## Data Model
| Table Name | Description | Columns |
| ------- | ---- | ---- |
| label | Represents label | (namespace, name, enable, *who...*)
| label_entity | Represents mapping between label and entity | (namespace, label_id, entity_id, *who...*)

Note: expand who... as created_by, created_at, updated_by, updated_at

## Database choice
A close look at the API interface reveals equal amount to read and write. While data consistency is required need for strong transaction semantics is not required hence any database that provides eventual consistency can also be consider. 

Given Mongo DB provides high availability and horizontal scalability it can be use to achive hight throughput and low latency.

# Scalability and Fault tolerance

## Functional and Load testing

# Alerting and Monitoring

# Security
