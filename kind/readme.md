# Kind
```shell
kind create cluster --name hsc-dev --config kind-config.yaml
```

```shell
helm repo add metrics-server https://kubernetes-sigs.github.io/metrics-server/
helm repo update
helm upgrade --install --set args={--kubelet-insecure-tls} metrics-server metrics-server/metrics-server --namespace kube-system


```
```shell

kubectl port-forward svc/xnassimulator -n opentp 50061

```

### Deploy nginx
```shell
kubectl create ns demo-nginx
kubectl create deploy nginx --image nginx -n demo-nginx
kubectl expose -n demo-nginx deploy nginx --port 80 --type NodePort

```


### Setup Kong ingress controller (KIC) Kong manager
```shell
kubectl apply -f kic.yaml 

```

### App nginx example
```shell

kubectl patch deployment -n kong proxy-kong -p '{"spec":{"replicas":1,"template":{"spec":{"containers":[{"name":"proxy","ports":[{"containerPort":8e3,"hostPort":80,"name":"proxy-tcp","protocol":"TCP"},{"containerPort":8443,"hostPort":443,"name":"proxy-ssl","protocol":"TCP"}]}],"nodeSelector":{"ingress-ready":"true"},"tolerations":[{"key":"node-role.kubernetes.io/control-plane","operator":"Equal","effect":"NoSchedule"},{"key":"node-role.kubernetes.io/master","operator":"Equal","effect":"NoSchedule"}]}}}}'
kubectl patch service -n kong kong-proxy -p '{"spec":{"type":"NodePort"}}'
kubectl apply -f example/app.yaml 
kubectl -n my-example port-forward svc/my-app 20001

```

```shell
istioctl install --set profile=minimal -y
```

### Kong add istio
```shell

kubectl label namespace kong istio-injection=enabled
kubectl -n kong get deploy ingress-kong -o yaml | istioctl kube-inject -f - | kubectl apply -f -

```


### App bookinfo
```shell

kubectl create ns bookinfo
kubectl label namespace bookinfo istio-injection=enabled
kubectl -n bookinfo apply -f  istio-1.10.2/samples/bookinfo/platform/kube/bookinfo.yaml


kubectl -n bookinfo apply -f - <<EOF
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: productpage
  namespace: bookinfo
spec:
  ingressClassName: kong
  rules:
  - http:
      paths:
      - path: /
        pathType: ImplementationSpecific
        backend:
          service:
            name: productpage
            port:
              number: 9080

EOF

```



### Start addons Kiali grafana
```shell

kubectl apply -f istio-1.10.2/samples/addons
kubectl apply -f istio-1.10.2/samples/addons/extras

kubectl -n istio-system port-forward svc/grafana 3000
kubectl -n istio-system port-forward svc/kiali 20001


```



### Delete cluster
```shell

kind delete cluster --name hsc-dev
```
