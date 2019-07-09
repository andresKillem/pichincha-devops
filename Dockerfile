FROM node:10-alpine

# Add our configuration files and scripts
WORKDIR /usr/src/app

COPY . .
#ADD . /app
RUN npm update -g
RUN npm install -g npm@3
RUN npm install

EXPOSE 3000

ENTRYPOINT ["npm", "run", "start"]
