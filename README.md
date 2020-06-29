## Kubernetes(K8S) Pod Scheduler ![Go](https://github.com/ahmagdy/k8s-pod-scheduler/workflows/Go/badge.svg?branch=master) [![Total alerts](https://img.shields.io/lgtm/alerts/g/ahmagdy/k8s-pod-scheduler.svg?logo=lgtm&logoWidth=18)](https://lgtm.com/projects/g/ahmagdy/k8s-pod-scheduler/alerts/)
A simple and lightweight go tool that can run on top of K8s to schedule and run long-running tasks.

It's something smaller and simpler than Apache Airflow for small or simple tasks. Instead of using a workflow management tool with a lot of overhead, this tool can be more fit for simple and small jobs.
 

### Use cases of this package
When a user wants to read jobs from a remote-store and schedule them accordingly.




### TODO:
- Write more docs.
- Store registered jobs state in a persistent store instead of storing them in memory.
- Provide API to update, and remove tasks. 
 
