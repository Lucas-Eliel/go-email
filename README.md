# go-email
Aplicação em GO para disparo de email

# PostgreSQL
docker run --name postgres -e POSTGRES_PASSWORD=d4#rt6 -e POSTGRES_USER=email_dev -p 5432:5432 -d postgres

# Keycloak
docker run --name keycloak -p 8080:8080 -e KEYCLOAK_ADMIN=admin -e KEYCLOAK_ADMIN_PASSWORD=admin quay.io/keycloak/keycloak:21.1.1 start-dev