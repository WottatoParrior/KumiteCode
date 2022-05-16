# KumiteCode

Conduct Interviews by Code Reviewing with the candidate

**Prereqs:**

-Docker
-Go

---

If you have Make installed run **make run dev**

If you don't have Make installed, run **docker compose up -d** and then **air**

## To Add

-~~CI/CD~~

-Docker/Kubernetes

-Protobuff for communication

-Postgress

-AWS

-gRPC

-~~Jenkins~~

-Ansible for DB

-RabbitMQ for event streaming

-Terraform

-Sentry for logging

-~~Air for golang~~

## To Expand

-CI/CD for Golang

## Journey

Day 1.Added some basic github actions to get familiar with CI/CD processes, will need to add specific ones
for runing tests and separate ones for deploying the app.

Day 2.Did some research on Docker and how to create multi stage builds using distroless container images.
Packaged everything in one container using docker compose. Theoretically the app and the db should be on different containers, I think. Also setup a Postgres container, need to check if it can actually communicate with the app.

Day 3.Added air for live reload and a Makefile to run all the commands needed at once. This runs the docker compose file and then
runs air to enable live-reload.
