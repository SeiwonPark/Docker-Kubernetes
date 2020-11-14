# Docker-Kubernetes Web Application   

**ABSTRACT**   

This is the application which returns the largest prime factor of the given input in URL.
And I've packaged the web application in a Docker container image, and run it on GKE(Google Kubernetes Engine) cluster.
Then, deployed the web application. 
All these are all based on _[THIS TUTORIAL](https://cloud.google.com/kubernetes-engine/docs/tutorials/hello-app)_.

## 1. Result      
You can access to the application with following URL:   

<del>   [http://34.64.233.162/?input=13195](http://34.64.233.162/?input=13195)</del>  **(To avoid incurring charges to my GCP account for the resources, I've deleted the service.)**   

This shows the result like the image below(_**Example 1**_).   
And also, you can check if it works by changing the inputs as follows:   

### Example 1: `input = 13195`   

<img width=800 src="https://github.com/SeiwonPark/Docker-Kubernetes/blob/main/images/image1.png">   

### Example 2: `input = 0`   

<img width=800 src="https://github.com/SeiwonPark/Docker-Kubernetes/blob/main/images/image2.png">   

### Example 3: `input = 1000000`   

<img width=800 src="https://github.com/SeiwonPark/Docker-Kubernetes/blob/main/images/image3.png">   

### Example 4: `input = 100`   

<img width=800 src="https://github.com/SeiwonPark/Docker-Kubernetes/blob/main/images/image4.png">   

### Example 5: `input = 123456`   

<img width=800 src="https://github.com/SeiwonPark/Docker-Kubernetes/blob/main/images/image5.png">   


## 2. Files Contained   
This directory contains:

- `main.go`    
  It contains the HTTP server implementation. It responds to HTTP
  requests that contains `?input=${inputValue}` with responses as follows:
  ```
    The input is ${inputValue}   
    The largest prime factor is ${maxPrime}
  ```   
  or if has no prime factor,   
  ```
    The input is ${inputValue} has no prime factor
  ```
  
- `Dockerfile`   
  It is used to build the Docker image for the application.


- `hello-deployment.yaml`   

- `hello-service.yaml`   

