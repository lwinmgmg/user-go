FROM "node:20-alpine3.19"

RUN apk update && apk upgrade && apk add bash

ENV WORK_DIR="/frontend"
ENV NODE_ENV=production

WORKDIR $WORK_DIR

COPY "*.json" ${WORK_DIR}/.
COPY "*.ts" ${WORK_DIR}/.
COPY "*.mjs" ${WORK_DIR}/.

RUN npm install

COPY app/ ${WORK_DIR}/app/
COPY public/ ${WORK_DIR}/public/

RUN npm run build

COPY ./entrypoint.sh /

EXPOSE 80

ENTRYPOINT [ "/entrypoint.sh" ]

CMD [ "npm" ]
