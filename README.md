# kubernetes-gosample 
 this application is related to kubernetes with golang sample

 **Description**: How can I bind github project with dockerhub to create docker images includes with golang sample?



+ build Go binaries
+ containerize them and publish to a registry : I used docker hub because I have a windows machine and hate using docker as containerized my project  on windows. also hyper-v demolishing VirtualBox concurrency. It is another handicap for my advanture of my containered world.
+ automatically update Kubernetes manifests to references the correct container image

# Install
+ Actually, you only need to install minikube and kubectl for handling our pods. Because we never use any docker to create images locally. We will handle them on dockerhub by binding github and dockerhub.

# Configure a Registry

You have told to KO_DOCKER_REPO about your repository's place.

~~~python
export KO_DOCKER_REPO=yusufkaratoprak/kubernetes-gosample
~~~

#Usage

By writing Kubernetes manifest with docker images, you can draw a roadmap for your minikube.

~~~python
apiVersion: apps/v1beta1
kind: Deployment
metadata:
  name: hello-world
spec:
  replicas: 1
  selector:
    matchLabels:
      app: hello
  template:
    metadata:
      name: hello
      labels:
        app: hello
    spec:
      containers:
      - name: hello-world
        # Image from dockerhub, This is the import path for the Go binary to build and run.
        image: yusufkaratoprak/kubernetes-gosample:latest
        ports:
        - containerPort: 8090

~~~

#Build

Then to start the build, containerization, and deployment a single ko command is necessary.

~~~python
kubectl apply -f config/
~~~

# Check Pods!

~~~python
C:\cygwin64>kubectl get pods
NAME      READY     STATUS    RESTARTS   AGE
hello     1/1       Running   0          8s
~~~

# Create Service

~~~python
C:\cygwin64>kubectl port-forward pod/hello-world-74fdccf99-7zn7g 8090:8090
~~~

#Check your golang service
~~~python
$ curl localhost:8090
Handling connection for 8090
Hello world. My name is YUSUF!
~~~
