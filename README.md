# banking
##### Run `./start.sh` to download the dependencies and run the the application

To run the application, you have to define the environment variables, default values of the variables are defined inside `start.sh`

- SERVER_ADDRESS    `[IP Address of the machine]`
- SERVER_PORT       `[Port of the machine]`
- DB_USER           `[Database username]`
- DB_PASSWORD       `[Database password]`
- DB_ADDRESS        `[IP address of the database]`
- DB_PORT           `[Port of the database]`
- DB_NAME           `[Name of the database]`

# mysql database
To start the docker container, run the `docker-compose up` inside the `resources/docker` folder
