# n-labels

[![Go Report Card](https://goreportcard.com/badge/github.com/nimesh-mittal/n_labels)](https://goreportcard.com/report/github.com/nimesh-mittal/n_labels)

# Background
Aim of this service is to provide ability to manage set of entities by associating entities with tags, also called as labels. Labels are attached to an entity in order to achive following benefits

1. Labels provides a way of organising collection of entities
2. Helps retriving one or more related entities by label
3. Allow special treatment to certain entities by categorising them using labels
4. Labels can be use to assign entities to users i.e., to create fluid workflows
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
  // create new label
  Create(labelName string, namespace string) string    
  // delete exisiting label
  Delete(labelID uuid, namespace string) bool          
  // attach label to an entity
  Attach(labelID uuid, namespace string, entityID uuid) bool    
  // search labels by keyword
  Search(query string) []string    
  // list entities of a label
  ListEntities(labelID uuid) []uuid 
}

```

## Data Model
| Table Name | Description | Columns |
| ------- | ---- | ---- |
| label | Represents label | (namespace, name, enable, *who...*)
| label_entity | Represents mapping between label and entity | (namespace, label_id, entity_id, attached_by, *who...*)

Note: expand who... as created_by, created_at, updated_by, updated_at

## Database choice
A close look at the API interface reveals equal amount to read and write requests. While data consistency is required, need for strong transaction semantics is not required hence any database that provides eventual consistency can also be consider. 

Given Mongo DB provides high availability and horizontal scalability it can be use to achive hight throughput and low latency.

# Scalability and Fault tolerance
Inorder to survive host failure, multiple instances of the service needs to be deployed behind load balancer. Load balance should detect host failure and transfer any incoming request to only healthy node. One choice is to use ngnix server to perform load balancing.

Given one instance of service can serve not more than 2000 request per second, one must deploy more than one instance to achive required throughput from the system

Load balancer should also rate limiting incoming requests to avoid single user/client penalising all other user/client due to heavy load.

# Functional and Load testing
Service should implement good code coverage and write functional and load test cases to maintain high engineering quality standards.

# Logging, Alerting and Monitoring
Service should also expose health end-point to accuretly monitor the health of the service. 

Service should also integrate with alerting framework like new relic to accuretly alert developers on unexpected failured and downtime

Service should also integrate with logging library like zap and distributed tracing library like Jager for easy debugging

# Security
Service should validated the request using Oauth token to ensure request is coming from authentic source

# Documentation
This README file provides complete documentation. Link to any other documentation will be provided in the Reference section of this document.

# Local Development Setup
- Start service
```go run main.go```
- Run testcases with coverage
```go test ./... -cover```