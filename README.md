# kubernetes-gosample 
 this application is related to kubentes with golang sample

 **Description**: How can I use ko with golang sample?

ko is  a little unknown gem, lost in the GitHub maze.

+ build Go binaries
+ containerize them and publish to a registry : I used docker hub because I have windows machine  and hate docker on windows also hyper-v demolishing virtualbox concurency. 
+ automatically update Kubernetes manifests to references the correct container image

# Install
+ You have install "ko" first.
~~~python
go get github.com/google/go-containerregistry/cmd/ko
~~~

# Configure a Registry

You have tell to KO_DOCKER_REPO about your repository's place.

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

Then to start the build, containerization and deployment a single ko command is necessary.

~~~python
ko apply -f config/
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