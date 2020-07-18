## Kubernetes(K8S) Pod Scheduler ![Go](https://github.com/ahmagdy/k8s-pod-scheduler/workflows/Go/badge.svg?branch=master) [![Total alerts](https://img.shields.io/lgtm/alerts/g/ahmagdy/k8s-pod-scheduler.svg?logo=lgtm&logoWidth=18)](https://lgtm.com/projects/g/ahmagdy/k8s-pod-scheduler/alerts/) [![codecov](https://codecov.io/gh/ahmagdy/k8s-pod-scheduler/branch/master/graph/badge.svg)](https://codecov.io/gh/ahmagdy/k8s-pod-scheduler)
A simple and lightweight go tool that can run on top of K8s to schedule and run long-running tasks with a simple gRPC interface.

It's something smaller and simpler than Apache Airflow for small or simple tasks. Instead of using a workflow management tool with a lot of overhead, this tool can be more fit for simple and small tasks that don't really require a lot of configurations to execute. 

### Use cases of this package
When a user wants to read jobs from a remote-store and schedule them accordingly.

## User guide
Assuming you have a backend that goes doing some process to come with jobs that need to be scheduled. If you don't have a program and you want to test the service you can use any gRPC client like [BloomRPC](https://github.com/uw-labs/bloomrpc) or [gRPCurl](https://github.com/fullstorydev/grpcurl).

As you can see in idl >> job.proto. The service has only one method which is Add to add a new job to the scheduler.
This job will accept these parameters:
- Name string
- Cron expression string. It accepts cron expressions with seconds. Also, "@daily" and "@every <duration>"

The service will add your job to the scheduler, run it as scheduled, and kill the pod when the job is done. 
Note that when the job the scheduler will kill it as well.

The recommended way is to use the proto interface provided above and call it via gRPC.
Feel free to add REST or CLI support if your use case needs it.

#### Supported Fields:
 
|Field|Type|Required|
|---|---|---|
Name | String | ✓|
Cron | String | ✓|
Spec.image | String| ✓|
Spec.Args | List | |
Spec.Commands | List| 

### Start the program
You can run it from your machine directly or use this image as an example from dockerhub `ahmedmagdi/k8s-pod-scheduler`.

```bash
docker pull ahmedmagdi/k8s-pod-scheduler:0.0.1
kubectl apply -f ./example/k8s-deployment.yaml
``` 
This will start a pod listening on port `30007`. You can call it then using BloomRPC or similar tools.

##### Triggering a new job:

```bash
grpcurl -plaintext -import-path ./idl -proto job.proto -d '{"job": {"name":"my-first-job", "cron":"@every 0h0m30s", "spec":{"image":"ahmedmagdi/go-sample-task:1.0.0"}}}' \
    127.0.0.1:30007 job.JobService/Add
```
<img width="1436" alt="ScreenShot of BloomRPC using k8s-pod-scheduler" src="https://user-images.githubusercontent.com/10447926/87860756-533fc700-c940-11ea-97fc-0b50f464fac7.png">

After that if you checked the logs of the running pod or watched the changes in k8s (`kubectl get pods -n podscheduler -w`) you will notice the following:

```bash
NAME                            READY   STATUS    RESTARTS   AGE
podscheduler-58dcffbdd7-zz2hg   1/1     Running   0          53m
```
When triggering a build
```bash
NAME                            READY   STATUS    RESTARTS   AGE
podscheduler-58dcffbdd7-zz2hg   1/1     Running   0          53m
my-first-jobw7qld               0/1     Pending   0          0s
my-first-jobw7qld               0/1     Pending   0          0s
my-first-jobw7qld               0/1     ContainerCreating   0          0s
my-first-jobw7qld               1/1     Running             0          1s
my-first-jobw7qld               0/1     Completed           0          6s
my-first-jobw7qld               0/1     Terminating         0          6s
my-first-jobw7qld               0/1     Terminating         0          6s
```

### Local Development:
```bash
make setup
```
Will setup all the needed tools and dependencies.

### TODO:
- Store registered jobs state in a persistent store instead of storing them in memory.
- Provide API to update, and remove tasks. 
 
