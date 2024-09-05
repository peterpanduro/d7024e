Deploy the containers:

`docker compose up -d`

To test ping between containers:

`docker exec lab_0-node-1 ping -c 4 lab_0-node-2`
