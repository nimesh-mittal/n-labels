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

# Implementation
## API Interface
## Data Model
## Database choice

# Scalability and Fault tolerance
## Functional and Load testing

# Alerting and Monitoring

# Security
