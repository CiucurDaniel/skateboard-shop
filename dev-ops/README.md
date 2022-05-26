# dev-ops

Here the Kubernetes manifests and any potential scripts will be present.

# Logging

Apply all objects in the folder logging

```
kubectl apply f logging/
```

In order to access Kibana you need to do a port-forward:

```
kubectl port-forward kibana-5749b5778b-vwn47  5601:5601 --namespace=kube-logging
```

Visit the following web URL:

```
http://localhost:5601
```

Create the index and you are good to go.