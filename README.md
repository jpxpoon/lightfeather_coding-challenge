# Jonathan Poon

# LightFeather Codeing Challenge

Created a application with Home Page and Notification Services Page with flexibility to extend additional pages


## Cleint - Frontend

- Using [vue](https://vuejs.org/guide/introduction.html) to build reactive UI
- Using [axios](https://www.npmjs.com/package/axios) API as http client


## Server - Backend

- Using [mux](https://pkg.go.dev/github.com/gorilla/mux) to build backend API, handling http url request
- Function `saveNotification` `/server/server.go` can be extend by adding database architecture 

## Simple Run 

    # Run server.
    cd server && go run .

    # Run client
    cd client && npm start
You can reach your client at `http://localhost:3000/` and your server at `http://localhost:3001/`

## Docker environment

`docker-compose.yml` for local environment

❯ docker-compose up