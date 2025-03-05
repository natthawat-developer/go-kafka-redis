
```
go-kafka-redis
├─ README.md
├─ cmd
│  ├─ consumer
│  │  └─ main.go
│  └─ producer
│     └─ main.go
├─ config
│  ├─ config.go
│  └─ config.yaml
├─ deployments
│  ├─ Dockerfile.base
│  ├─ Dockerfile.consumer
│  ├─ Dockerfile.producer
│  ├─ docker-compose.yml
│  └─ k8s.yaml
├─ go.mod
├─ go.sum
└─ pkg
   ├─ kafka
   │  ├─ consumer.go
   │  └─ producer.go
   ├─ logger
   │  └─ logger.go
   ├─ redis
   │  └─ client.go
   └─ utils

```