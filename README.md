# Docker-Kubernetes Web Application   

**ABSTRACT**   

This is the application which returns the largest prime factor of the given input in URL.
And I've packaged the web application in a Docker container image, and run it on GKE(Google Kubernetes Engine) cluster.
Then, deployed the web application. 
All these are all based on _[THIS TUTORIAL](https://cloud.google.com/kubernetes-engine/docs/tutorials/hello-app)_.

## 1. Result      
You can access to the application with following URL:   
   
[http://34.64.245.122/?input=13195](http://34.64.245.122/?input=13195)   

<img width=1000 src="https://github.com/SeiwonPark/Docker-Kubernetes/blob/main/images/example1.png" >   

This shows result like the image above.   
And also, you can check if it works by changing the inputs as follows:   

### Example 1: `input = 100`   


### Example 2: `input = 12345678`   


### Example 3: `input = 0`   


### Example 4: `input = 1000000`   



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
(You can )

- `hello-service.yaml`   

