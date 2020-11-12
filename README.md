# Docker-Kubernetes Web Application   

**ABSTRACT**   

This is the application which returns the largest prime factor of the given input in URL.
And I've packaged the web application in a Docker container image, and run it on GKE(Google Kubernetes Engine) cluster.
Then, deployed the web application. 
All these are all based on _[THIS TUTORIAL](https://cloud.google.com/kubernetes-engine/docs/tutorials/hello-app)_.

## 1. Result      


## 2. Files Contained   
This directory contains:

- `main.go`    
  It contains the HTTP server implementation. It responds to HTTP
  requests that contains `?input=${inputValue}` with responses as follows:
  ```
    The input is ${inputValue}   
    The largest prime factor is ${maxPrime}
  ```   
  
- `Dockerfile`   
  It is used to build the Docker image for the application.


- `hello-deployment.yaml`   

- `hello-service.yaml`   

