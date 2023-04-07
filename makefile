postgres:
	@cd manifests/postgres && kubectl apply -f . && cd ../..

job:
	@cd manifests && kubectl apply -f . && cd ../..

install-ingress:
	@helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx && helm repo update && helm install ingress-nginx ingress-nginx/ingress-nginx

app:
	@cd manifests/app && kubectl apply -f . && cd ../..