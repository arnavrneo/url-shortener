FROM node:18
LABEL authors="arnavrneo"

WORKDIR /app

COPY package*.json ./

RUN npm install

COPY . .

EXPOSE 5678

CMD [ "node", "--watch", "index.js" ]