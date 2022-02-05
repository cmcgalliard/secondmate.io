# SecondMate.io
A tool to help your nonProduction Kubernetes Clusters running clean.
The goal of this tool is to add some features to non production clusters to help developemnt workflows.

Current Feature:
- You can add labels to a namespace and it will be deleted on a specified date

Future Features:
- Remove any workloads that have not been modified in a specified amounts of dates
- Tool to sync secrets from secret tools
- Tool to quickly identify workloads with outdated API configuration (eg. ingress)
- Tooling to Identify who installed what where (think, mutating webhook annotations)

## Namespace Pruning
You can set a few lables on a namespace and the tool will run hourly based on your configuration and delete any old Namespaces/workloads.
here's the config you'll need:

```
apiVersion: v1
kind: Namespace
metadata:
  labels:
    kubernetes.io/metadata.name: test-true
    secondmate.io/purge: "true"
    secondmate.io/purge-date: "2022-03-20"
    secondmate.io/purge-hour: "15"
    secondmate.io/purge-tz: EST
  name: test-true
````

Example Run:
```
--- config ---
dry-run: false 
labelMatcher: secondmate.io/purge=true 
--- config ---

PURGE     Namespace     Created                           PurgeDate                         NamespaceDeleted
---       ---           ---                               ---                               ---
false     test-true     2022-02-05 16:36:48 -0500 EST     2022-03-20 16:00:00 -0400 EDT     false
```