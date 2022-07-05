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

## Frontend

Minikube makes it a bit tricky to expose your app so for now just:

```
k port-forward frontend-689b64c8ff-2pz9v 9001:9000 -n skateshop
```

## Auth

Test auth deployment:

```
k port-forward auth-679fbf6d75-wbxgr 8071:8070 -n skateshop
```