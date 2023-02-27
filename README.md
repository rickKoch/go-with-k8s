# go-with-k8s
Simple go app test with k8s

## Usage
Replace **repository** with your repository.
```
docker build -t go-with-k8s:latest ./app/
docker tag go-with-k8s:latest <repository>/go-with-k8s:latest
docker push <repository>/go-with-k8s:latest
kubectl applay -f ./k8s/go-app.yaml
```

`LoadBalancer` is used, so you can get the address of the app.
```
kubectl get svc go-app -o jsonpath='http://{.status.loadBalancer.ingress[0].*}:5555'
```

## Logs
```
kubectl logs -l app=go-app
```
