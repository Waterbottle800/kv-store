Simple KV-Store using Go

OVERVIEW
This project is a simple in-memory key-value store written in Go, exposing an HTTP API with basic operations such as creating, retrieving, and searching key-value pairs. It also supports prefix and suffix-based search functionality.

FILES IN THE PROJECT
1. go.mod - This file defines the module name (mykv) and Go version (1.23.1) required for the project.

2. sample.go - Contains the main logic for the key-value store, which includes the following functions:

- Set a key-value pair
- Stores key-value pairs in memory.
- Get a value by key
= Retrieves the value associated with a given key.
- Search functionality
- Allows searching for keys based on prefix and suffix.
- HTTP API endpoints
        - /set (POST): Set a key-value pair by passing key and value in the request body.
        - /get/{key} (GET): Retrieve the value for a given key.
        - /search (GET): Search keys based on a provided prefix or suffix query parameters.

3. my_test.go
Contains unit tests to verify the functionalities of the service:

        - TestSet(): Tests setting key-value pairs in the store.
        - TestGet(): Tests retrieving values, checking both existing and non-existing keys.
        - TestSearch(): Tests prefix and suffix-based key search functionality.

RUNNING THE SERVICE IN DOCKER

1. Run the container pratheep10/kv-store:1.0 from Dockerhub exposing port 8080 and mapping the same to port 8080 in local machine

        - $ docker run -d --name kv-store -p 8080:8080 pratheep10/kv-store:1.0

2. Check the service at http://localhost:8080, it should return "404 page not found error"

        - curl http://localhost:8080


RUNNING THE SERVICE IN KUBERNETES

1. Use the manifest files provided in kv-store/kv-store-k8s to deploy the service in default namespace

        - $ kubectl apply -f kv-store/kv-store-k8s/kv-store-k8s.yaml

2. The manifest deploys single replica of container kvs (image: pratheep10/kv-store:1.0) exposing port 8080 of the container and a NodePort service with target port 8080 and nodePort 30001. Make sure port 30001 is not mapped with any other service in the local machine.

3. To access the service in local Kind based Kubernetes setup, use the IP of the worker node the service is running in. Use docker inspect to obtain the IP address of the node the service is running in and access the service at http://<worker node container ip>:30001

       - $ kubectl get pods -o wide         // To get the node name in which the service is running
       - $ docker ps                        // To get the container ID/Name of the node
       - $ docker inspect <container ID>

4. Another way to access the service is by using port-forwarding and access the service at http://localhost:8080

        - $ kubectl port-forward svc/kvs-service 8080:80

USE THE API SERVICE (USING CURL)

1. Set a key-value pair

        - $ curl -X POST http://localhost:8080/set -d '{"key": "abc-1", "value": "value1"}' -H "Content-Type: application/json"

2. Get a value by key

        - $ curl http://localhost:8080/get/abc-1

3. Search for keys by prefix

        - $ curl http://localhost:8080/search?prefix=abc

4. Search for keys by suffix

        - $ curl http://localhost:8080/search?suffix=-1




   

       


        
