make up: install-nginx-ingress
	kubectl apply -f postgres-db-pv.yaml
	kubectl apply -f postgres-db-pvc.yaml
	kubectl apply -f postgres-secret.yaml
	kubectl apply -f postgres-db-deployment.yaml
	kubectl apply -f postgres-db-service.yaml
	kubectl apply -f job.yaml
	kubectl apply -f app-configmap.yaml
	kubectl apply -f app-deployment.yaml
	kubectl apply -f app-service.yaml
	kubectl apply -f app-ingress.yaml

install-nginx-ingress:
	@helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx && helm repo update && helm install ingress-nginx ingress-nginx/ingress-nginx