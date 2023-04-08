postgres:
	@helm install postgres bitnami/postgresql -f values_postgresql.yaml
	
app: 
	@helm install app app-chart -f app-chart/values.yaml

remove-postgres:
	@helm uninstall postgres

remove-app:
	@helm uninstall app