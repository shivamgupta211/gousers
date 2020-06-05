# Kube-Safe Assignment

## Front-end
1. Preact(see: https://preactjs.com/) as the usecase is very small module, so build with Preact only.
    Advantages: 
    - minimal size alternative to reactjs but with same API
    - Blazing Fast front-end loads
    
2. Redux
    - State management

## Back-end
1. Golang
    - Better performance than python
2. Postgres as DB (Why SQL DB?)
    - As we have fix schema and no json data and MongoDB is ideal for unstructured data, still psql can also work with the JSON data but not as efficiently than Mongodb
    - Further in future as data and tables grows might need more and more relations between the tables

# -------------- Running in Development --------------
1. Start the Front-end
    - Run `cd ./frontend && yarn`
    - Run `yarn run start`
    - Server will spin up and all the front-end code will be visible on "localhost:8080" i.e. 8080 port

2. Start Backend
    - cd `./backend` & run `go mod init gousers`
    - Run `go run server.go`
    - Server will satrt on the port "2369"


# -------------- Running in Production --------------
As we have already separated the front-end and back part so we can deploy them independently to each other in a K8s cluster.
For backend we will deploy it in the k8s cluster

### DEPLOY back-end
## Step 1: Setup a PSQL database
-   Create new PSQL cluster/instance which can connected from inside the k8s cluster
-   Create a new database in PSQL
-   Create new tables as per defined in `db-schema.sql` file
-   Once created, please populate the database related ENV varibales present in 
        `./backend/export-envs.sh` file in terminal

## Step 2: Build back-end docker image
1. it has `Dockerfile`
2. - To build new images
    - Change the `DOCKER_HUB_USERNAME` ENV value to your public docker hub username
    - Run `bash ./backend/build-and-push-images.sh` in terminal

## Step 3: Kubernetes deployment
1. Make sure you have defined the proper ENVs in the `./export-envs.sh` file
2. Run `bash ./kube.sh` (could've used HELM but to save time, just did the templates)
3. Once deployed the service will be exposed as the "Loadbalancer" URL


### DEPLOY Front-end
## Step 1: Make client build
-   Replace the value of "API" to the above generated "Loadbalancer" URL
-   Run `cd frontend && yarn && yarn run build`
-   Copy the `./frontend/build/` files to the S3 storage and host that as a website