## Requirement
### Install nginx ingress controller (classic)
```bash
helm upgrade --install ingress-nginx ingress-nginx \
  --repo https://kubernetes.github.io/ingress-nginx \
  --namespace ingress-nginx --create-namespace

or

kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.1.0/deploy/static/provider/cloud/deploy.yaml
```

## Deployment
```bash
helm install --debug  messages . --namespace messages --create-namespace \
--set prometheus-pushgateway.serviceMonitor.namespace=messages \
--set grafana."grafana\.ini".server.domain=http://ingress-nginx-controller.ingress-nginx.svc.cluster.local \
--set emqx-ee.emqxConfig.EMQX_PROMETHEUS__PUSH__GATEWAY__SERVER=http://messages-prometheus-pushgateway.messages.svc.cluster.local:9091
```

## Upgrade
```helm upgrade --debug  messages . --namespace messages --create-namespace \
--set prometheus-pushgateway.serviceMonitor.namespace=messages \
--set grafana."grafana\.ini".server.domain=http://ingress-nginx-controller.ingress-nginx.svc.cluster.local \
--set emqx-ee.emqxConfig.EMQX_PROMETHEUS__PUSH__GATEWAY__SERVER=http://messages-prometheus-pushgateway.messages.svc.cluster.local:9091
```


## Testing
```bash
helm test messages --logs -n messages
```

## Uninstall
```bash
helm uninstall messages --namespace messages

kubectl delete crd alertmanagerconfigs.monitoring.coreos.com
kubectl delete crd alertmanagers.monitoring.coreos.com
kubectl delete crd podmonitors.monitoring.coreos.com
kubectl delete crd probes.monitoring.coreos.com
kubectl delete crd prometheuses.monitoring.coreos.com
kubectl delete crd prometheusrules.monitoring.coreos.com
kubectl delete crd servicemonitors.monitoring.coreos.com
kubectl delete crd thanosrulers.monitoring.coreos.com
```
**Note**: Don't forget to delete pvc if you need.

## Config prometheus and grafana
### Install ingress for prometheus and grafana
```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  namespace: messages
  name: prom-ingress
  annotations:
    kubernetes.io/ingress.class: "nginx"
    nginx.ingress.kubernetes.io/rewrite-target: /$2
spec:
  rules:
  - http:
      paths:
      - path: /prom(/|$)(.*)
        pathType: Prefix
        backend:
          service:
            name: ${promethues-service-name}
            port:
              number: 9090
      - path: /grafana(/|$)(.*)
        pathType: Prefix
        backend:
          service:
            name: ${grafana-service-name}
            port:
              number: 80
```

### Url
> http://${nginx-ingress}/prom/graph  
http://${nginx-ingress}/grafana/login

### Add Data Sources
> ${prometheus-service-name}.messages.svc.cluster.local

## EMQX rule engine Configuration
### Postgres
> db: mqtt  
user: chaos  
passwd: public

### Kafka
> user:  
passwd:  
topic: testTopic
