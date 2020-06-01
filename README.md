## Kubernetes(K8S) Pod Scheduler
A simple and lightweight go tool that can run on K8s to schedule and run long-running tasks.

It's something smaller and simpler than Apache Airflow for a small or simple tasks. Instead of using workflow management tool with a lot of overhead, this tool can do the job faster.
 

### Use cases of this package
When a user wants to read jobs from a remote-store and schedule them accordingly.




### TODO:
- Store the registered jobs state in a persistent store instead of storing them in memory.
- Provide API to add, update and remove tasks. 
 