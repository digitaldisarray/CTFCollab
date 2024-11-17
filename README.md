# CTFCollab
Collaborative tool for CTF teams

## Frontend

### Developing

Once you've created a project and installed dependencies with `npm install` (or `pnpm install` or `yarn`), start a development server:

```bash
npm run dev

# or start the server and open the app in a new browser tab
npm run dev -- --open
```

It is also recommended to get the svelte extension for vscode. If you use another editor there is likely a svelte extension/plugin for it.

### Building

To create a production version of your app:

```bash
npm run build
```

You can preview the production build with `npm run preview`.

> To deploy your app, you may need to install an [adapter](https://svelte.dev/docs/kit/adapters) for your target environment.

## Backend

### Development
Golang is required to compile and run the backend code on your computer

For working on the database schema, sqlc is required https://docs.sqlc.dev/en/stable/overview/install.html

Running code: `go run .` while in the backend directory

#### Modifying queries or schema
You can modify the queries and schema in `backend/db/`. Once you modify them, use `sqlc generate` to generate the corresponding go code. You will need to [install sqlc](https://docs.sqlc.dev/en/stable/overview/install.html) to do so. 
