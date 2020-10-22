## SubjectAccessReview(SAR) Checks Operator

A controller which does SAR checks for a username and a Kubernetes Resource.

### Running this locally

```
$ cd sar-checks-operator
$ kubectl apply -f deploy/crds/example.sbose.com_examples_crd.yaml
$ operator-sdk run --local
```

I happen to use operator-sdk 1.18.


```
apiVersion: example.sbose.com/v1alpha1
kind: Example
metadata:
  name: example-example
spec:
  username: consoledeveloper
  testSubject:
    group: ""
    version: "v1"
    resource: "pods"
    name: "mypod" # Deploy something prior to this.
```

The status would be updated with `.status.allowed: true`.

The logs should have

```
{true false RBAC: allowed by ClusterRoleBinding "consoledeveloper_role2" of ClusterRole "view" to User "consoledeveloper" }
```