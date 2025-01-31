# CTFCollab
Collaborative tool for CTF teams


## Running in docker
Create a `.env` file in the root directory of CTFCollab. Here is a template:
```
MYSQL_ROOT_PASSWORD=CHANGE_ME
MYSQL_DATABASE=ctfcollab
MYSQL_USER=ctfuser
MYSQL_PASSWORD=CHANGE_ME
MYSQL_URL=${MYSQL_USER}:${MYSQL_PASSWORD}@tcp(db:3306)/${MYSQL_DATABASE}?multiStatements=true&parseTime=true
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
```

## Running locally

### Frontend
Once you've created a project and installed dependencies with `npm install` (or `pnpm install` or `yarn`), start a development server:

```bash
npm run dev

# or start the server and open the app in a new browser tab
npm run dev -- --open
```

It is also recommended to get the svelte extension for vscode. If you use another editor there is likely a svelte extension/plugin for it.

To create a production version of your app:

```bash
npm run build
```

You can preview the production build with `npm run preview`.

### Backend
Golang is required to compile and run the backend code on your computer

For working on the database schema, sqlc is required https://docs.sqlc.dev/en/stable/overview/install.html

Running code: `go run .` while in the backend directory

You can modify the queries and schema in `backend/db/`. Once you modify them, use `sqlc generate` to generate the corresponding go code. You will need to [install sqlc](https://docs.sqlc.dev/en/stable/overview/install.html) to do so. 


### Testing

To run the tests, make sure your database is up to date. Run:   
```
docker volume rm ctfcollab_db-data   
```
Relaunch the server with Docker. Then, run the commands:   
```
cd backend   
go test ./tests    
```
