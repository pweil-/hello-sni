# hello-sni

Demonstrates a TLS tcp client connecting through the OpenShift router to a TLS tcp server.

1.  Build the docker image with `build.sh`
1.  Start OpenShift
1.  Install the router
1.  Create the pod, service, and route
1.  Run the client

```
[vagrant@openshiftdev paul_temp]$ osc get pods && osc get services && osc get routes
POD                 IP                  CONTAINER(S)                   IMAGE(S)                          HOST                           LABELS              STATUS
router              172.17.0.2          origin-haproxy-router-router   openshift/origin-haproxy-router   openshiftdev.local/127.0.0.1   <none>              Running
tls-server          172.17.0.4          tls-server                     pweil/tls-server                  openshiftdev.local/127.0.0.1   name=tls-server     Running
NAME                LABELS                                    SELECTOR            IP                  PORT
kubernetes          component=apiserver,provider=kubernetes   <none>              172.30.17.2         443
kubernetes-ro       component=apiserver,provider=kubernetes   <none>              172.30.17.1         80
tls-service         <none>                                    name=tls-server     172.30.17.20        2999
NAME                HOST/PORT           PATH                SERVICE             LABELS
route-passthrough   my-tls-server                           tls-service   

[vagrant@openshiftdev paul_temp]$ go run client.go 
Hello TLS
[vagrant@openshiftdev paul_temp]$
```
