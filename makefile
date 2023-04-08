app: 
	@cd app-chart && helm dependency build && helm install app app-chart -f values.yaml

collection:
	@newman run ./postman-collection/collection.json

install-nginx-ingress:
	@helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx && helm repo update && helm install ingress-nginx ingress-nginx/ingress-nginx

remove-app:
	@helm uninstall app