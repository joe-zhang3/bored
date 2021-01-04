FROM node:14.8.0-alpine

WORKDIR /app

COPY package.json yarn.lock ./

RUN yarn 

COPY . .

CMD ls -l
CMD ["yarn", "start"]

