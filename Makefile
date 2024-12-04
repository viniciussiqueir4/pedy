database:
	docker run --name pedy-postgres -e POSTGRES_PASSWORD=pedy -e POSTGRES_USER=pedy -e POSTGRES_DB=pedy -p "5432:5432" -d postgres