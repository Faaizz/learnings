# Introduction to Kubernetes

## Content
- [Kubernetes Architecture](#kuberentes-architecture)
- [Minikube](#minikube)
- [Kubernetes Object Model](#kubernetes-object-model)
- [Authentication, Authorization, & Access Control](#authentication-authorization--access-control)
- [Services](#services)
- [Accessing Deployed Application on `minikube`](#accessing-deployed-application-on-minikube)
- [Readiness & Liveness Probes](#readiness--liveness-probes)
- [Volume Management](#volume-management)
- [ConfigMaps & Secrets](#configmaps-and-secrets)
- [Ingress](#ingress)
- [Advanced Topics](#advanced-topics)
  
## Kubernetes Architecture
- [Master Node Components](#master-node-components)
- [Worker Node](#worker-node)
- [Networking](#netwroking)
- [Configuration](#configuration)

### Master Node Components
- Control Plane Components:
  - **API Server (`kube-apiserver`)**: Is the only service that can directly read-from or write-to the data store, thus it serves as a middle layer between all other control plane services and the data store (etcd). Can scale horizontally, and can be configured with secondary API servers (in this case, the primary API server becomes a reverse proxy to the secondary API servers).
  - **Scheduler (`kube-scheduler`)**: Assigns workload objects (e.g. Pods) to Nodes. During scheduling, the workload object's requirements (obtained from the API server) and the Kubernetes cluster state (obtained from the data store via the API server) are taken into account. The result of the decision process is communicated to the API server which then relays it to some toher control plane agaent for deployment.
  - **Controller Managers**: Are responsible for managing controllers. Controllers are watch-loops that constantly compare the desired state (obtained from objects' configuration data) and the current state of the cluster (obtained from the data store via the API server). If a mismatch is detected, corrective action is taken until the desired and current states match. `kube-controller-manager` is responsible for ensuring pod counts as expected, creating endpoints & service accounts, & managing API access tokens. `cloud-controller-manager` is responsible for managing storage volumes provided by a cloud service, load balancing, & routing.
  - **Data Store (`etcd`)**: Is a distributed key-value data store used to persist a Kubernetes cluster's state. New data is added by appending the existing data, data is never replaced in the data store. `stacked topology` is when the data store is run on the Master Node (and all its High-Availability replicas) along with other control place services, while `external topology` is when the data store is run on its own dedicated host (requires seperate HA replication, but provides isolation from other control plane services). Besides the cluster's state, etcd is also used to store configuration details such as subnets, ConfigMaps, Secrets, and so on.
- Others:
  - Container Runtime
  - Node Agent
  - Proxy

### Worker Node
Client applications which are containerized microservices are encapsulated in Pods and run on worker nodes. A Pod is a logical collection of one or more containers which can be started, stopped, or rescheduled as a unit of work. Worker nodes provide Pods with their compute, storage, & memory reaources, and also networking to communicate with one another and with the outside world. In a multi-node cluster, communication between client users and Pods are handled by the worker nodes (and not routed through the master node).

- **Container Runtime**
- **Node Agent (`kubelet`)**: Runtime agent that runs on each node and is in constant communication with the control plane. Recieves pod descriptions from the API server and interacts with the container runtime on the node to run containers associated with the pods. `kubelet` relates with the container runtime via the Container Runtime Interface (CRI). The CRI has 2 services: "ImageService" which is responseible for image-related operations, and "RuntimeService" which is responsible for all Pod and container-related operations.
- **Proxy (`kube-proxy`)**: Network agent that runs on every node. It is responsible for maintaning all networks rules of the node. It handles the details of pods network and forwards connection request to pods.
- **Addons** for DNS, dashboard user interface, cluster-level monitoring and logging

### Netwroking
- **Container-to-Container**: Each Pod is configured with an extra *Pause* container which is spawned for the sole purpose of creating a 'network namespace' for the Pod. Each container in a Pod share the same network namespace and thus are able to communicate with one another via *localhost*.
- **Pod-to-Pod**: Kubernetes uses an *IP-per-Pod* model to facilitate communication between pods in the cluster. It treates Pods like VMs and assign a unique IP address to each Pod.
- **External-to-Pod**: To enable Pod communications with the external world, Kubernetes uses Services. Services implement complex encapsulation of network routing rules in *iptables* on cluster nodes, these are then used by `kube-proxy` to expose applications to user outside of the cluster over virtual IP addresses.

### Configuration
- All-in-One Single Node
- Single-Master (stacked etcd) and Multi-Worker
- Single-Master with Single-Node etcd and Multi-Worker
- Multi-Master (stacked etcd) and Multi-Worker
- Multi-Master with Multi-Node etcd and Multi-Worker


## Minikube
- [Accessing Minikube](#accessing-minikube)
- [Enable Registry Addon](#enable-registry-addon)

### Accessing Minikube
- Start: `minikube start`
- Status: `minikube status`
- Stop: `minikube stop`

- `kubectl`
- Kubernetes Dashboard
- `curl`

### Enable Registry Addon
For some weird reason (or lack of patience) I couold not successfully setup and use the minikube container registry so I just stuck with Microk8s.
```shell
# Enable insecure registry on minikube start
minikube start --insecure-registry "192.168.49.0/24"
# Replace "192.168.49.0/24" with appropraite values according to 'minikube ip'
minikube addons enable registry

# Build image with tag
docker build --tag $(minikube ip):5000/<img-name>
# Push image to registry
docker push $(minikube ip):5000/<img-name>
```
Enable minikube IP as insecure registry in docker daemon
```json
# /etc/docker/daemon.json
{

}
```
```shell
# Run socat to forward traffic from localhost to minikube on port 5000
docker run --rm -it --network=host alpine ash -c "apk add socat && socat TCP-LISTEN:5000,reuseaddr,fork TCP:$(minikube ip):5000"
```

#### API Server
HTTP API directory tree of Kubernetes:
```shell
/ 
  -|/healthz
  -|/metrics
  -|/logs
  -|/ui
  -|/api/v1
    -|/pods
    -|/nodes
    -|/services
    -|...
  -|/apis/$NAME/$VERSION
  -|...
```

#### `kubectl`
```
kubectl [command] [TYPE] [NAME] [flags]
```
- `command`: Operation to perform. Commands can have subcommands, for example `kubectl config view`. `view` here is a subcommand of `config`.
- `TYPE`: Resource type. Case-insensitive, and can be specified in the singular, plural, or abbrevated form.
- `NAME`: Resource name. Combinations of resource type and name can be used to specify multiple resource for a given command. E.g.
  ```shell
  kubectl get pod pod1 pod2 pod3
  kubectl delete pod/demo-pod1 deploy/demo-deploy1 svc/demo-service1
  ```

Some common operations include:
- `annotate`: Add or update annotations of resource 
- `apply`: Apply a configuration change to a resource (creates the resouce if it doesn't already exist)
- `attach`: Attach to a running container 
- `cluster-info`: Display information about the cluster 
- `cp`: Copy files and directories to and fro containers
- `delete`: Delete one or more resources 
- `describe`: Display detailed information about one or more resources 
- `edit`: Edit the configuration of one or more resources on the server using a text editor
- `exec`: Execute a command against a container running in a pod 
- `explain`: Get documentation on various resources e.g. pods, deployments, services, nodes, etc 
- `expose`: Expose a replication controller, service, or pod 
- `get`: List resources 
- `label`: Add/update resource labels 
- `logs`: Print the logs for a container in a pod
- `port-forward`: Forward local ports to a pod 
- `proxy`: STarts a server which proxies requests to the API server
- `run`: Run a specified image on the cluster 
- `scale`: Update the size of specified replication controller 
- `top`: Display resource usage
 
 #### Kubernetes Dashboard
 `minikube dashboard` or via kubectl by running `kubectl proxy` which then makes the page available at [http://localhost:8001/api/v1/namespaces/kubernetes-dashboard/services/http:kubernetes-dashboard:/proxy/](http://localhost:8001/api/v1/namespaces/kubernetes-dashboard/services/http:kubernetes-dashboard:/proxy/). 
 We can enable the metrics-sever addon to get access to cluster metrics:
 ```shell
 minikube addons list
 minikube addons enable metrics-server
 ```


## Kubernetes Object Model
- [Object Model](#object-model)
- [Pods](#pods)
- [Labels](#labels)
- [Label Selectors](#label-selectors)
- [ReplicationSets](#replicationsets)
- [Deployments](#deployments)
- [Namespaces](#namespaces)

### Object Model
Kubernetes object model represents different persistent entities in the cluster describing:
- What containerized applications are running
- The nodes where they are running
- The resource consumption of those applications
- The policies attached to those applications. E.g. restart policy, fault tolerance policy, etc

The *spec* section of an object model holds the desired state of a resource, this is defined by the administrator. The *status* section is managed by Kubernetes, it holds the actual state of the resource. The control plane constantly tries to bring the actual state to match the desired state. A sample configuration manifest for a Deployment object is given:
```yaml
apiVersion: /apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
  labels:
    app: nginx
spec:
  replicas: 3
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app:nginx
    containers:
    - name: nginx
      image: nginx:1.15
      ports:
      - containerPort: 80
```
The *apiVersiom* field specifies the API endpoint which we want to connect to.
Once the deployment has been created,configured the Kubernetes system attaches a *status* section to hold the current state of the resource.

### Pods
A Pod is the smallest resource unit in Kubernetes. It is a logical collection of containers that:
- Are scheduled to run on the same Node
- Share the same network namespace, hence have the same IP address (the Pod's IP address)
- Have access to mount the same external storages (volumes)

Pods are ephemeral and thus are often used with controllers which manage their replication, fault tolerance, etc. Such controllers include Deployments, ReplicaSets, etc. Pods are often attached to the controller's spec using Pod Template, albeit leaving out the 'apiVersion' and 'kind' fields. Sample Pod Template:
```yaml
apiVersion: v1
kind: Pod
metadata: 
  name: nginx-pod
  labels:
    app: nginx
spec:
  containers:
  - name: nginx
    image: nginx:latest
    ports:
    - containerPort: 80
```
The 'containerPort' field specifies the port on the container that's exposed by Kubernetes resources for access from other applications or external clients.
**TIP** _The 'spec' field gives the desired state of the Pod, also named the **PodSpec**_

### Labels
Labels are *key-value* pairs attached to Kubernetes objects (e.g. Namespaces, Deployments, Services, etc). They can be used to organize or select objects matching some specified requirements. For instance, controllers can logically group pods together according to their labels.
Some commonly used label commands are shown below:
```shell
# List pods along with their atached labels. If the pod has the label specified,
# its value will be displayed, otherwise, the column is left empty.
kubectl get pods -L label1,label2
# Select Pods with a given label
kubectl get pods -l k8s-app=web-app
```

### Label Selectors
- **Equality-Based:** Objects are filtered based on label keys and values using the equality operator *=*/*==* and/or the inequality operator *!=*.
- **Set-Based:** Objects are filtered using sets. Operators such as *in* and *notin* can be used to check if the value of a label is present in a set- e.g. **env** notin (dev,qa)- while operators such as *exists*/*does not exist* are used to check the presence of a key.

### ReplicationSets
Pods cannot restart themselveconfiguredhe practice that controllers are used to manage such pods. ReplicationSet is a controller that can by used to scale how many pods should run an application container, either manually or via an autoscaler. They support both equality-based and set-based selectors. If the number of running pods exceed the specified amount, the controller randomly kills pods to match the values. If on the other hand, the number of running pods is smaller than the requested quantity, then the controller requests deployment of additional pods to match the values.

### Deployments
A Deployment is apable of creating, deleting,a nd updating Pods. It automatically creates a ReplicaSet which in turn creates Pods. In addition to being able to maintain replicas, a Deployment makes applcation updates and rollbacks seamless to perform. The DeploymentController is a part of the comtroller-manager on the control-plane.
Considering an illustration where a Deployment is deployed to create 3 pods running the nginx:1.13 image. The deployment creates a ReplicaSet (say ReplicaSet A) which then creates 3 pods (say Pod-1, Pod-2, Pod-3) running containers spwaned from the nginx-1.13 image. This setup is tagged as a Revision (say Revision 1). Let's say the image tag on the template field of the deplyment then gets bumped to 1.91 (i.e. update to nginx:1.91). Then the deployment creates a new ReplicaSet (say ReplicaSet B) which in turn creates new Pods, yielding a new Revision (say Revision 2). When the new revision is up and running, the deployment reduces the number of running pods on Revision 1 (ReplicaSet A) to zero, but skill keeps the ReplicaSet. This makes it possible to rollback to a previous configuration if required.
```shell
# Create deployment
kubectl create deployment mynginx --image=nginx:1.13-alpine
# List resources tagged with app:mynginx
kubectl get deply,rs,po -t app-mynginx
# Scale deployment to 3 pods
kubectl scale deploy mynginx --replicas=3
# List resources and note ReplicaSet name
kubectl get deply,rs,po -t app-mynginx
# View rollout history
kubectl rollout history deploy mynginx
# Show detailed info about revision 1
kubectl rollout history deploy mynginx --revision=1
# Upgrade image to ngnix:1.16
kubectl set image deployment mynginx nginx=nginx:1.16-alpine
# Show detailed info about revision 2
kubectl rollout history deploy mynginx --revision=2
# Rollback to revision 1
kubectl rollout undo deployment mynginx --to-revision=1
```
A sample Deployment config is given below:
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: webserver
  labels:
    app: nginx
spec:
  replicas: 3
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: nginx:alpine
        ports:
        - containerPort: 80
```


### Namespaces
Namespaces help to create virtual sub-clusters within a Kubernetes cluster. This yields many advantages since the names of resources need only be unique within a particular namespace. It makes isolation of users, teams, and applications possible. Resource quotas can be used to limit the overall consumption of resources by namespaces, while LimitRanges help limit resource consumption by containers or pods in a namespace.


## Authentication, Authorization, & Access Control
- [Authentication](#authentication)
- [Authorization](#authorization)

### Authentication
In Kubernetes, there are 2 kinds of users: *Normal Users* which are managed outside of the cluster, and *Service Users* which allow in-cluster processes to communicate with the API server. Such processes mount their credentials as secrets when attempting to communicate with the API server.
Kubernetes doesn't keep user access information in it's data store, so it resolves to external modules for user identification and authentication. These methods include: CLient Certificates, Static Token File, Service Account Tokens, Webhook Token Authentication, etc. 
Below is a sample implementation of Client Certificaate Authentication:
```shell
# Create namespace to use
kubectl create namespace lfs158
# Create certificate private key
openssl genrsa -out student.key 2048
# Create certificate signing request
openssl req -new -key student.key -out student.csr -subj "/CN=student/O=learner"
# Output the signing request encoded in base64
cat student.csr | base64 | tr -d '\n','%'
```
```yaml
# signing-request.yaml
apiVersion: certificates.k8s.io/v1
kind: CertificateSigningRequest
metadata:
  name: student-csr
spec:
  groups:
  - system:authenticated
  request: <insert base64 encoded student.csr>
  signerName: kubernetes.io/kube-apiserver-client
  usages:
  - digital signature
  - key encipherment
  - client auth
```
```shell
# Create signing request in Kubernetes
kubectl create -f signing-request.yaml
# Approve signature
kubectl certificate approve student-csr
# Save certificate to file
kubectl get csr student-csr -o jsonpath='{.status.certifacate}' | base64 --decode > student.crt
# Create user entry in kubectl
kubectl config set-credentials student --client-certificate=student.crt --client-key=student.key
# Create context in kubectl that uses the user and set namespace
kubectl config set-context student-context --cluster=minikube --namespace=lfs158 --user=student
# Use context
kubectl --context=student-context get pods
```

### Authorization
Upon successful authentication, users can send requests to the API server. Here, the attributes of these requests are inspected using various authorization modules. Multiple modules can be configured and each is checked in sequence. Some availaable authorization modules include: Node Authorization, Attribute-Based Acces Control (ABAC), Webhook, Role-Based Access Control (RBAC).
In RBAC, **Role**s let us specify permissions for users in a namespace (while **ClusterRole**s allow cluster-wide permissions to be set). We can assign users to roles using **RoleBinding**s.
A sample implementation of **Role** & **RoleBinding** is given below:
```yaml
# role.yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: pod-reader
  namespace: lfs158
rules:
- apiGroups: [""] # "" apiGroup signifies the core API
  resources: ["pods"]
  verbs: ["get", "watch", "list"] # verbs indicate which actions are allowed
```
```shell
# Create role
kubectl create -f role.yaml
kubectl -n lfs158 get roles
```
```yaml
# rolebiding.yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: pod-read-access
  namespace: lfs158
subjects:
- kind: User
  name: student
  apiGroup: rbac.authorization.kubernetes.io
roleRef:
  kind: Role
  name: pod-reader
  apiGroup: rbac.authorization.kubernetes.io
```
```shell
# Create role binding for user student to role pod-reader
kubectl create -f rolebinding.yaml
# View rolebindings
kubectl -n lfs158 get rolebindings
```

## Services
- [Service](#service)
- [`kube-proxy`](#kube-proxy)
- [Service Discovery](#service-discovery)
- [ServiceType](#servicetype)

### Service
A **Service** helps to provide a single DNS entry for a containerized application that is managed by the Kubernetes cluster by providing a single load-balacing access point to a set of pods logically grouped together & managed by a controller (e.g. Deployment, ReplicaSet, etc).
Services helps us abstract communication between internal cluster microservices or with the external world.
**Services** recieve an IP address that is routable only inside of the cluster called *ClusterIP* upon creation. A **Service** performs load balancing by default. A sample configuration file is provided below:
```yaml
apiVersion: v1
kind: Service
metadata:
  name: frontend-svc
spec:
  selector:
    app: frontend # Selects resources with label app==frontend
  ports:
  - protocol: TCP
    port: 80  # Port through which service recieves traffic
    targetPort: 5000 # Port on Pod which recieves traffic, if not defined, 'port' is used. This value should match 'containerPort' on the Pod specification
```
*Service endpoints* are a logical collection of pod IPs and corresponding `targetPort` for a Service. For example, if the `frontend-svc` matches Pods with IPs `1.0.0.2` & `1.0.0.4`, then the service endpoints would be: [`1.0.0.2:5000`, `1.0.0.4:5000`]. To get a list of Service endpoints, we can use: `kubectl describe svc <service-name>`

### `kube-proxy`
`kube-proxy` is a deamon that runs on all nodes which constantly watches the API server for addition, updates, and removal of Services & endpoints. For each Service, it sets up an `iptables` entry that captures traffic directed towards the Service's ClusterIP and forwards it to one of the Services endpoints.

### Service Discovery
Kubernetes supports 2 methods of service discovery at runtime:
- **Environment Variables:** Upon startup of a Pod, `kubelet` daemon running on the cluster node supplies the Pod with environment variables corresponding to info about running services. For example, if we have a service *redis-master* that exposes port *6379* and is assigned CLusterIP *172.17.0.6*, then for every Pod created after spinning up the Service, the following environment variables get passed:
  ```shell
  REDIS_MASTER_SERVICE_HOST=172.17.0.6
  REDIS_MASTER_SERVICE_PORT=6379
  REDIS_MASTER_PORT=tcp://172.17.0.6:6379
  REDIS_MASTER_PORT_6379_TCP=tcp://172.17.0.6:6379
  REDIS_MASTER_PORT_6379_TCP_PROTO=tcp
  REDIS_MASTER_PORT_6379_TCP_PORT=6379
  REDIS_MASTER_PORT_6379_TCP_ADDR=172.17.0.6
  ```
It is noteworthy to remember that with this approach, the Service **MUST** be created before the Pods are deployed, otherwise this won't work.

- **DNS:** Kubernetes uses a DNS addon that creates DNS records for each service using the following format: **service-name.namespace-name.svc.cluster.local**. Pods (and Services) in the same namespace as the service can find the service just by its name. Pods from other namespaces can find the service by specifying **service-name.namespace-name** or the FQDN of the service which is **service-name.namespace-name.svc.cluster.local**.

### ServiceType
ServiceType determines the access scope of a Service. We candecide whether the Service:
- Is only accessible from within the cluster
- Is accessible from within the cluster and the external world
- Maps to an entity which resides inside or outside of the cluaster.

The ServiceType of a Service can be specified when creating the Service. The default is `ClusterIP`. 

The `NodePort` ServiceType assigns a high-port (from 30000-32767) to the respective Service in addition to its ClusterIP. WHen a client request is made to any worker node on the assigned port, the requests gets forwarded to one of the Service's endpoints. This allows access to the Service from the external world. If a specific NodePort is preferred, it can be assigned from the specified range while creating the Service.

Other ServiceTypes such as `LoadBalancer` & `ExternalIP` require clud providers. The `ExternalName` ServiceType is a special ServiceType that has no selectors and does not define any endpoints. Rather it is returns a CNAME record of an externally configured Service. Its primary use case is to make available externally configured Services in the cluster.

## Accessing Deployed Application on `minikube`
To get the IP address of the minikube VM, run `minikube ip`. The running application can then be accessed using `<minikubeIP>:<NodePort>`. The same can also be acheived by running `minikube service <service-name>`, which opens the service up in the default browser.

## Readiness & Liveness Probes
- [Liveness Probes](#liveness-probes)
- [Readiness Probes](#readiness-probes)

### Liveness Probes
 Liveness Probes help us heck if an application in a container is functioning as expected, if not, `kubelet` restarts the container. Liveness Probes can be set up as:
 - Liveness command
 - Liveness HTTP request
 - TCP Liveness probe

Example Liveness command:
```yaml
apiVersion: v1
kind: Pod
metadata:
  labels:
    name: liveness-exec
    test: liveness
spec:
  containers:
  - name: liveness
    image: k8s.gcr.io/busybox
    args:
    - /bin/sh
    - -c
    - touch /tmp/healthy; sleep 30; rm -rf /tmp/healthy; sleep 600
    livenessProbe:
      exec:
        command:
        - cat
        - /tmp/healthy
      initialDelaySeconds: 3  # Delay before first probe, should be enough time for container to spin up application
      failureThreshold: 1 # Number of failures before `kubelet` renders container as unhealthy and perform restart 
      periodSeconds: 5  # Liveness Probe is run every 5secs
```

Example Liveness HTTP request:
```yaml
...
    livenessProbe:
      httpGet:
        path: /healthz  # application endpoint to probe
        port: 8080
        httpHeaders:
        - name: X-Custom-Header
          value: Awesome
      initialDelaySeconds: 3
      periodSeconds: 3
...
```

Example TCP Liveness probe. The `kubelet` attempts to open TCP socket to the container on specified port, on the event of failure, the container is deemed unhealthy and restarted.
```yaml
...
    livenessProbe:
      tcpSocket:
        port: 8080
      initialDelaySeconds: 13
      periodSeconds: 30
...
```


### Readiness Probes
We use Readiness Probes to tell Kubernetes when an application is ready to recieve traffic. A Pod with containers that do not report ready status will not receive traffic from Kubernetes Services. Readiness Probes are configured similar to Liveness Probes. An example is given below:
```yaml
...
    readinessProbe:
      exec:
        command:
        - cat
        - /tmp/healthy
      initialDelaySeconds: 5 
      periodSeconds: 5
...
```

## Volume Management
- [Volume Types](#volume-types)
- [Persistent Volumes](#persistent-volumes)
- [Persistent Volume Claims](#persistent-volume-claims)
- [Container Storage Interface (CSI)](#container-storage-interface-csi)
- [Illustration of How to Consume a Volume inside a Pod](#illustration-of-how-to-consume-a-volume-in-containers-inside-a-pod)

### Volume Types
Containers running in Pods are ephemeral, they lose all data in the if they fail & have to be restarted by `kubelet`. Volumes can be attached to Pods, allowing containers in a Pod to persist data via the Volume. The Volume outlives the Containers but not the Pod. Kubernetes offer different Volume Types for use with Pods, these include: 
- *emptyDir:* Empty Volume which loses its data permanently when the Pod dies.
- *hostDir:* Shared directory from the Host to the Pod.
- *gcePersistentDisk*, *awsElasticBlockStore*, *azureDisk*, *azeureFile*.
- *secret:* To pass sensitive information to Pods (eg. passwords).
- *configMap:* To provide configuration data or shell command arguments to Pods.
- *persistentVolumeClaim:* Attach a PersistentVolume to a Pod.

### Persistent Volumes
A PersistentVolume is a storage abstraction in Kubernetes which is backed by various underlying storage technologies, ranging from local storage on the Pod's host, to network attached storage, cloud storage, or a distributed storage solution.
PersistentVolumes can be dynamically provisioned using the StorageClass resource. They can be configured using the PersistentVolume API and consumed using the PersistentVolumeClaim API.

### Persistent Volume Claims
A PersistentVolumeClaim is made by users to request a PersisentVolume for use. Users request for volumes based on type, access mode, and size. There are 3 access modes:
- *ReadOnlyMany:* Read-only to by multiple nodes.
- *ReadWriteOnce:* Read-write by a single node.
- *ReadWriteMany:* Read-write by multiple nodes.

Once a suitabe PersistentVOlume is found, it is boud to the PersistentVolumeClaim and made available to the user, who in turn mounts it into the Pod(s) where it is needed. When the user is done with the PersistentVolume, the volume is released. Depending on the `persistentVolumeeclaimPolicy`, one of three things occur. The PersistentVOlume can be:
- *reclaimed:* For an admin to verify and/or aggregate data.
- *deleted:* Both data and volume deleted.
- *recycled:* Data deleted but volume kept for future usage.

### Container Storage Interface (CSI)
CSI allows third-party storage providers to develop storage solutions without need to add them into the Kubernetes core.

### Illustration of How to Consume a Volume in Containers Inside a Pod
```shell
# SSH into Minikube VM
minikube ssh

mkdir pod-volume
cd ./pod-volume
pwd
```
```yaml
#share-pod.yaml
apiVersion: v1
kind: Pod
metadata:
  name: share-pod
  labels:
    app: share-pod
spec:
  volumes:
  - name: host-volume
    hostPath:
      path: /home/docker/pod-volume
  containers:
  - image: nginx
    name: nginx
    ports:
    - containerPort: 80
    volumeMounts:
    - mountPath: /usr/share/nginx/html
      name: host-volume
  - image: debian
    name: debian
    volumeMounts:
    - mountPath: /host-vol
      name: host-volume
    command: ["/bin/sh", "-c", "ech0 Introduction to Kubenetes > /host-vol/index.html; sleep 3600"]
```
```shell
# Create Pod from configuration
kubectl create -f share-pod.yaml
# List Pods
kubectl get pods
# Expose Pod via a NodePort Service
kubectl expose pod share-pod --type=NodePort --port=80
# List services and endpoints
kubectl get svc,endpoints
# Open service in web browser
minikube service share-pod
```

## ConfigMaps and Secrets
- [ConfigMaps](#configmaps)
- [Secrets](#secrets)

The ConfigMap API allows runtime parameters to be passed to Pods in Kubernetes. The Secret API does the same, but provides additional security for sensitive information.

### ConfigMaps
Allow us to decouple configuration details from container images. They provide configuration information as key-value pairs which can be consumed by Pods, or any other system components  or controllers as environment variables, sets of commands and arguments, or as volumes. ConfigMaps can be created from literal values, configuration files, or from one ore more files or directories. A sample configuration file for a ConfigMap is provided below:
```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: web-configmap
data:
  MARIA_BD_HOST: 192.168.0.13
  MARIA_DB_PORT: 3066
  APACHE_LOG_DIR: /var/log/apache2
```

We can create config maps from .env files. This is illustrated below:
```shell
#.env
MARIA_BD_HOST=192.168.0.13
MARIA_DB_PORT=3066
APACHE_LOG_DIR=/var/log/apache2
```
```shell
kubectl create configmap web-config --from-file=.env
kubectl get configmaps web-config -o yaml
```

#### Using ConfigMaps Inside Pods
To retrieve all key-value pairs specified in a ConfigMap in a Container, we use:
```yaml
...
  containers:
  - name: <app-name>
    image: <app-image>
    envFrom:
    - configMapRef:
      name: <config-map-name>
...
```
To retrieve specific key-value pairs from ConfigMaps, we can use:
```yaml
...
  containers:
  - name: <app-name>
    image: <app-image>
    env:
    - name: MARIA_DB_HOST
      valueFrom:
        configMapKeyRef:
          name: <config-map-name>
          key: MARIA_DB_HOST
    - name: MARIA_DB_PORT
      valueFrom:
        configMapKeyRef:
          name: <config-map-name>
          key: MARIA_DB_PORT
...
```
We can also mount ConfigMaps as Volumes. With this approach, a file is created for each entry in the ConfigMap with the entry's key as the file name and the value as the content of the file.
```yaml
...
  containers:
    - name: <app-name>
      image: <app-image>
      volumeMounts:
      - name: config-volume
        mountPath: /etc/config
        readOnly: true
  volumes:
  - name: config-volume
    configMap:
      name: <config-name>
...
```

More details can be found on the [ConfigMap](https://kubernetes.io/docs/tasks/configure-pod-container/configure-pod-configmap/) entry of the Kubernetes documentation.

### Secrets
Just as ConfigMaps, data can be stored in Secrets as key-value pairs. However, unlike ConfigMaps, when Secret data are used, the Secret object is referenced without actually exposing its contents. It is worthy to note that Secrets are stored as plain texts in the data store (etcd) by default, thus access to the API server and etcd must be controlled. Encryption of Secrets can also be enabled at an API server level.
Secrets can be created and used similar to ConfigMaps. A sample configuration for Secrets is shown below:
```yaml
apiVersion: v1
kind: Secret
metadate:
  name: passwords
type: Opaque
data:
  MARIA_DB_PASS: MariaDBPassword
```
Asample usage of Secrets inside a Pod is given below:
```yaml
...
  containers:
  - name: <app-name>
    image: <app-image>
    env:
    - name: MARIA_DB_PASS
      valueFrom:
        secretKeyRef:
          name: <secret-name>
          key: MARIA_DB_PASS
...
```

**Generally, you should NOT commit Secret configuration files to version control.**


## Ingress
- [Ingress Resource](#ingress-resource)
- [Ingress Controller](#ingress-controller)


### Ingress Resource
Ingress configures Layer 7 HTTP/HTTPS load balancer for Services and provides:
- Transport Layer Security
- Name-based virtual hosting
- Fanout routing
- Load Balancing
- Custom rules

With Ingress, users do not have to worry about connecting ti the right Service, they just connect to the Ingress endpoint which in turn routes the traffic to the appropriate Service. Sample Name-Based Virtual Hosting Ingress configuration is given below:
```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: virtual-host-ingress
  namespace: default
spec:
  rules: 
  - host: blue.example.com
    http:
      paths:
      - path: /
        backend:
          service:
            name: webserver-blue-svc
            port:
              number: 80
        pathType: ImplementationSpecific
  - host: green.example.com
    http:
      paths:
      - path: /
        backend:
          service:
            name: webserver-green-svc
            port:
              number: 80
        pathType: ImplementationSpecific
```

With **Fanout** Ingress rules, we can forward traffic to different paths on the same domain to different services. Illustration:
```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: fanout-ingress
  namespace: default
spec:
  rules: 
  - host: example.com
    http:
      paths:
      - path: /blue
        backend:
          service:
            name: webserver-blue-svc
            port:
              number: 80
        pathType: ImplementationSpecific
      - path: /green
        backend:
          service:
            name: webserver-green-svc
            port:
              number: 80
        pathType: ImplementationSpecific
```

The Ingress resource only accepts a definition of forwarding rules, the Ingress Controller is a reverse proxy responsible for fulfilling the traffic routing.

### Ingress Controller
An Ingress Controller constantly watches the API server for cahnges in the Ingress esources and updates the Layer 7 Load Balancing accordingly. Kubernetes supports numerous Ingress Controllers which include: `GCE L7 Load Balancer`, `Nginx Ingress Controller`, `Contour`, `HAProxy Ingress`, etc.
Minikube ships with NGINX Ingress Controller asan addon which we can enable with: `minikube addons enable ingress`.


## Advanced Topics
- [Annotations](#annotations)
- [Quota & Limits Management](#quota--limits-management)
- [Autoscaling](#autoscaling)
- [Jobs & CronJobs](#jobs-and-cronjobs)
- [DaemonSets](#daemonsets)
- [StatefulSet](#statefulset)
- [Network Policies](#network-policies)
- [Monitoring & Logging](#monitoring-and-logging)
- [Helm](#helm)

### Annotations
Unlike Lables, Annotations aren't used to select & identify resources, rather they are used to: store build/release information, PR numbers, git brance, etc; info of people responsible; pointers to logging, monitoring, analytics, etc; Ingress controller info; deployment state and revision information.

### Quota & Limits Management
The ResourceQuota API allows setting the following types of quotas per Namespace:
- Compute Resource Quota (limit CPU, memory, etc.)
- Storage Resoure Quota (limit PersistentVolumeClaims, etc.)
- Object Count Quota (limi no of Pods, Containers, etc.).

The LimitRange resource helps limit resource allocation within a namespace, it allows us to:
- Set compute usage limits per Po/Container
- Set storage request limits per PersistentVolumeClaim
- Set default requests and limits and automatically inject them into Containers'environment at runtime
- etc.

### Autoscaling
Autoscaling can be configured via controllers which peridically adjust the number of running objects based on single, multiple, or custom metrics.

- Horizontal Pod Autoscaler: Adjust replicas in a ReplicaSet, Replication Controller, or a Deployment based on
CPU utilization.
- Vertical Pod Autoscaler: Set Container compute requirements in a Pod and dynamically adjust them in runtime.
- Cluster Autoscaler: Dynamically adjust the size of the cluster when there's insufficient resources for scheduled tasks.

### Jobs and CronJobs
Allow one-off or periodic execution of some "task(s)". A Job manages Pods to fulfil its requirements, while a CronJob manages Jobs (which in turn manage Pods) to fulfil its requirements.

### DaemonSets
DeamonSets make it possible to spin up a Pod on all nodes. `kube-proxy` is an example of DaemonSet managed Pod. It is also possible for a DaemonSet to filter which nodes to deploy Pods to by configuring `nodeSelectors`and node `affinity` rules. DaemonSets support rolling updates and rollbacks.

### StatefulSet
Used for staeful applications which require unique identities such as name, network identifications, or strict ordering. Also supports rolling updates and rollbacks.
  
### Network Policies
Behave like firewalls and help us control how Pods are allowed to talk to other Pods.
  
### Monitoring and Logging
For metrics collection, Metrics Server or Prometheus is a popular candidate. For cluster-wide logging, [Elasticsearch](https://v1-18.docs.kubernetes.io/docs/tasks/debug-application-cluster/logging-elasticsearch-kibana/) ccan be used in conjunction with [fluentd](http://www.fluentd.org/).

### Helm
Is an `apt`or `yum` style package manager for K8s. K8s manifests can be bundled into Charts which can be served via repositories and managed via Helm.


## References
- [Introduction to Kubernetes- edX](https://learning.edx.org/course/course-v1:LinuxFoundationX+LFS158x+3T2020/home)
