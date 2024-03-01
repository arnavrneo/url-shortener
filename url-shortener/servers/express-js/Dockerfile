# TODO: error building; possible: node version
FROM node:21-alpine
LABEL authors="arnavrneo"

WORKDIR /app

COPY package*.json ./

RUN npm install

EXPOSE 5678

CMD [ "node", "index.js" ]