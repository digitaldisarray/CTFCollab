# CTFCollab
Collaborative tool for CTF (capture the flag) teams to keep track of and collaborate on challenges during competitions. 

# Project Values
The goal of this project is to provide an open source, developer friendly, pickup and go solution to collaberative development while solving cybersecurity puzzles in CTF competitions.

- Pick up and go: Our website is meant to be as simple as possible. Users don’t need accounts, they just enter a PIN and get access to an instance similar to Kahoot.
- Open source: We value transparency and community. By choosing to make our software open source we are enabling others to contribute to and audit the code.
- Industry-Standard Tech Stack: Built with Go (backend), SvelteKit (frontend), and MySQL (database), our stack prioritizes performance, scalability, and developer familiarity. 
- Developer-Friendly Codebase: We prioritize code organization and helpful comments to make the project accessible to new contributors.

# Team Members and Roles
- Kevin Rossel: Resident CTF’er. Backend and database schema. Guided UI/UX
- Teo Liam Ippolito: UX Designer and Frontend Development Lead
- Brayden Aldrich: Frontend, backend, and API development
- Danny T Nguyen: Implemented frontend features
- Maximiliano Maldonado: Backend development, worked on the database, API and implementing frontend
- Sory Diagouraga: Backend development and API documentation.

## Status   
Development of CTFNote is ongoing. Use at your own risk.  

# Table of contents
- [CTFCollab](#ctfcollab)
- [Project Values](#project-values)
- [Team Members and Roles](#team-members-and-roles)
  - [Status](#status)
- [Table of contents](#table-of-contents)
  - [Features](#features)
  - [UI](#ui)
  - [Running in docker](#running-in-docker)
  - [Running for frontend development](#running-for-frontend-development)
  - [Running locally](#running-locally)
    - [Just the frontend](#just-the-frontend)
    - [Just the backend](#just-the-backend)
  - [Running the backend test suite](#running-the-backend-test-suite)
- [End User Documentation](#end-user-documentation)

## Features
- Self-hostable  
- Live document editing  
- Challenge tracking  
- CTF tracking  
- CTF member tracking  
- Free and open source  

## UI  

![UI flow graph](https://github.com/digitaldisarray/CTFCollab/tree/main/src/common/images/CTFNote-UI.png "CTFNote UI")


## Running in docker
Create a `.env` file in the root directory of CTFCollab. Here is a template:
```
MYSQL_ROOT_PASSWORD=CHANGE_ME
MYSQL_DATABASE=ctfcollab
MYSQL_USER=ctfuser
MYSQL_PASSWORD=CHANGE_ME
MYSQL_URL=${MYSQL_USER}:${MYSQL_PASSWORD}@tcp(db:3306)/${MYSQL_DATABASE}?multiStatements=true&parseTime=true
CMD_DB_URL=mysql://hedgedoc:hedgedoc_pass@db:3306/hedgedoc
CMD_DOMAIN=localhost:3001
HEDGEDOC_URL=http://hedgedoc:3000
CMD_PROTOCOL_USESSL=false
CMD_URL_ADDPORT=false
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
docker compose -f compose.backend_only.yaml up --build

# Run the frontend outside of docker (make sure to do npm install)
npm run dev
```


## Running locally

### Just the frontend
If for whatever reason you want to run the frontend outside of docker, first run `npm install` (or `pnpm install` or `yarn`) in the repository directory. Then to run the frontend, use one of these commands:

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

# End-User Documentation

## Getting Started

### Admin Log In
- Go to the website domain
- Choose **Log In**
- Provide:
  - **Username**
  - **Password**

### Joining a CTF Event
- Navigate to the **Landing Page**
- Enter a valid **Join Phrase** provided by an Admin
- If joining as a guest, select a **nickname**
- Get to taking notes!

---

## User Guide

### Landing Page
- Join new events through the landing page by inputing a phrase provided by the Admin

### CTF Page
- View CTF info: name, description, flag

### Shared Notes
- Open a challenge to write or edit shared notes
- Supports **Markdown formatting**
- Collaboration is real-time and team-wide

---

## Admin Guide

### Creating a New CTF
1. Go to **Admin Panel → Create CTF**
2. Fill in:
   - Name, Description
   - Start/End Time
   - Description

### Adding a Challenge
- Navigate to your CTF
- Click **Add Challenge**
- Input:
  - Name
  - Description

---

## Onboarding Process

### For New Users
- Share the **Join Phrase**
- Recommend users pick unique nicknames
- Provide a quick tour:
  - Landing Page → CTF Page → Notes

### For Admins
- Use admin token to create first account
- Set up your CTFs in the Admin Panel
- Share join phrases with users or teams

---

## Troubleshooting

| Issue                             | Solution                                                              |
|----------------------------------|-----------------------------------------------------------------------|
| Invalid Join Phrase              | Double-check spelling or confirm with event organizer                |
| Nickname already taken           | Choose a different nickname unique to the current CTF                |
| Page won't load                  | Check your internet connection, refresh, or clear your browser cache |
| Can't add notes or challenges    | Ensure you’re logged in and have appropriate permissions              |
| 500 Internal Server Error        | Check console logs (if self-hosted) or contact Admin                |

---

## FAQ

### Can I join without an account?
Yes, guests can join using a nickname and the join phrase.

### How do I become an admin?
Admins are set manually or by using a special token during registration.

### Can I leave and rejoin a CTF?
Yes, as long as the CTF is still active and your nickname hasn't been taken.

