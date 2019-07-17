# h1 kubernetes-gosample 
# h4 this application is related to kubentes with golang sample

#h2 How can I use ko with golang sample?

# ko is a little unknown gem, lost in the GitHub maze.

+ build Go binaries
+ containerize them and publish to a registry : I used docker hub because I have windows machine  and hate docker on windows also hyper-v demolishing virtualbox concurency. 
+ automatically update Kubernetes manifests to references the correct container image
