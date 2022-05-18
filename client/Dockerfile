FROM node:lts-alpine
RUN npm install -g webpack-dev-server
WORKDIR /app/client
COPY package*.json ./
RUN npm install
COPY . .
EXPOSE 3000
CMD [ "npm", "start"]
