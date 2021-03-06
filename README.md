# Vorto Coffeeshop

## Deploy on Docker environment.

Clone git repository.

```sh
git clone https://github.com/anaxdev/vorto-coffeeshop.git
```

Build api docker image.

```sh
cd vorto-coffeeshop
docker-compose build api
```

Run system

```sh
docker-compose up -d
```

If everything is ok, `http://localhost:8080/` or `http://server-address:8080/` should show the API welcome page.


## Deploy on Kubernetes cluster

Push API docker image to dockerhub (or any other docker registry).

```sh
docker login
docker push anadk/vorto-coffeeshop-api:1.0
```

> Note: docker registry name for api image should be changed according to the own registry.


Create a ConfigMap for database initialization script

```sh
kubectl create configmap pg-init-script \
--from-literal=coffee.sql="$(curl -fsSL https://raw.githubusercontent.com/anaxdev/vorto-coffeeshop/main/sql/coffee.sql)"
```

Deploy all applications into Kubernetes cluster
```sh
git clone https://github.com/anaxdev/vorto-coffeeshop.git
cd vorto-coffeeshop
kubectl apply -f ./k8s
```


## Test API endpoints

There is a postman file [grpc_test.postman_collection.json](./postman/grpc_test.postman_collection.json) to test API endpoints. Import this file to Postman tool.

API has two endpoints: `delivery` and `statistics`

| Endpoint | Method | Parameter:Description | Response |
| --- | --- | --- | --- |
| delivery | Post | bean_type_id : CoffeeBean Id <br/> carrierIdStr : Carrier Id <br/> supplierIdStr : Supplier Id | Success (also new delivery is inserted to the table) <br/> or <br/> Fail with invalid delivery message |
| statistics | Get |  | Return the probability of valid delivery. |


## Import the invalid deliveries

There is a sql file [invalid-deliveries.sql](./sql/invalid-deliveries.sql) that produces the invalid deliveries. (Invalid deliveries are deliveries that a carrier cannot perform due to carrier bean constraints.)

```sh
docker exec -i <postgres-container> psql -U postgres postgres < invalid-deliveries.sql
```

or

Run sql query of this file in postgres client tool.
