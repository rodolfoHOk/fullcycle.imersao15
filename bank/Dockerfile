FROM node:18-slim

RUN apt update -y && \
    apt install -y procps && \
    npm install -g @nestjs/cli@10.1.18

WORKDIR /home/node/app

USER node

CMD [ "tail", "-f", "/dev/null" ]
