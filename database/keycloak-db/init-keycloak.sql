CREATE DATABASE keycloak_db;
CREATE USER keycloak_user WITH ENCRYPTED PASSWORD 'keycloak123';
GRANT ALL PRIVILEGES ON DATABASE keycloak_db TO keycloak_user;