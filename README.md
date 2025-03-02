# CTFCollab
Collaborative tool for CTF (capture the flag) teams to keep track of and collaborate on challenges during competitions


## Table of contents
- [Running in docker](#running-in-docker)
- [Running for frontend development](#running-for-frontend-development)
- [Running locally](#running-locally)
- [Running backend test suite](#running-the-backend-test-suite)



## Running in docker
Create a `.env` file in the root directory of CTFCollab. Here is a template:
```
MYSQL_ROOT_PASSWORD=CHANGE_ME
MYSQL_DATABASE=ctfcollab
MYSQL_USER=ctfuser
MYSQL_PASSWORD=CHANGE_ME
MYSQL_URL=${MYSQL_USER}:${MYSQL_PASSWORD}@tcp(db:3306)/${MYSQL_DATABASE}?multiStatements=true&parseTime=true
TEST_MODE=False
ADMIN_USERNAME=Admin
ADMIN_PASSWORD=CHANGE_ME
JWT_SECRET=CHANGE_ME
```

Starting
```
docker compose up --build
```

Stopping
```
docker compose down
```

Removing database
```
docker volume rm ctfcollab_db-data
# or, to remove the database while spinning the containers down
docker compose down -v
```

## Running for frontend development
For frontend development, it helps to not run the frontend in docker so it auto updated with each change to the source code.
```bash
# Run just the backed in docker
docker compose up -f compose.backend_only.yaml

# Run the frontend outside of docker (make sure to do npm install)
npm run dev
```


## Running locally

### Just the frontend
If for whatever reason you want to run the frontend outside of docker, fisrt run `npm install` (or `pnpm install` or `yarn`) in the repository directory. Then to run the frontend, use one of these commnds:

```bash
npm run dev

# or start the server and open the app in a new browser tab
npm run dev -- --open
```

It is also recommended to get the svelte extension for vscode. If you use another editor there is likely a svelte extension/plugin for it.

To create a production version of the frontend:
```bash
npm run build
```

You can preview the production build with `npm run preview`.

### Just the backend
Golang is required to compile and run the backend code on your computer

For working on the database schema, sqlc is required https://docs.sqlc.dev/en/stable/overview/install.html

Running code: `go run .` while in the backend directory

You can modify the queries and schema in `backend/db/`. Once you modify them, use `sqlc generate` to generate the corresponding go code. You will need to [install sqlc](https://docs.sqlc.dev/en/stable/overview/install.html) to do so. 


Also, we recommend using the mysql docker container to run the database locally

## Running the backend test suite

To run the tests, make sure the backend is running, and then run the commands:   
```
cd backend   
go test ./tests    
```

Sometimes the database can get messed up, in that case you can one of run the following commands:
```
# Spin down the backend and remove the database volume at the same time
docker compose down -v

# Remove the database volume
docker volume rm ctfcollab_db-data
```
