FROM golang:1.21-bookworm AS compiler

ARG OUTPUT_DIR=/bin
ARG OUTPUT_BINARY_NAME=user-go-http
ARG DB_MIGRATE_BINARY=user-go-migrator
ARG WORKDIR_NAME=/build

WORKDIR ${WORKDIR_NAME}

COPY . ${WORKDIR_NAME}/.

WORKDIR ${WORKDIR_NAME}/cmd/http-server

RUN go get && go mod vendor

RUN go build -o "${OUTPUT_DIR}/${OUTPUT_BINARY_NAME}"

RUN cd ${WORKDIR_NAME}/cmd/migrate-database && go build -o "${OUTPUT_DIR}/${DB_MIGRATE_BINARY}"

FROM debian:bookworm

ARG OUTPUT_DIR=/bin
ARG OUTPUT_BINARY_NAME=user-go-http
ARG DB_MIGRATE_BINARY=user-go-migrator
ARG WORKDIR_NAME=/build

ENV USER_GOLANG_VERSION=1.21
ENV USER_USER=gouser
ENV USER_USER_HOME_DIR="/home/${USER_USER}"
ENV USER_USER_UID=2023
ENV USER_WORKDIR="/etc/user/config"
ENV USER_BINARY_DIR="/usr/local/user/bin"
ENV PATH=${PATH}:${USER_BINARY_DIR}
ENV USER_SETTING_PATH="${USER_WORKDIR}/settings.yaml"
ENV GIN_MODE=release

RUN groupadd --gid ${USER_USER_UID} ${USER_USER} \
    && useradd --uid ${USER_USER_UID} --gid ${USER_USER_UID} \
    --home-dir ${USER_USER_HOME_DIR} --shell /bin/bash ${USER_USER}

RUN apt update
RUN apt install -y ca-certificates

WORKDIR ${USER_WORKDIR}

# Copy compiled binary
COPY --from=compiler --chown=${USER_USER}:${USER_USER} "${OUTPUT_DIR}/${OUTPUT_BINARY_NAME}" "${USER_BINARY_DIR}/${OUTPUT_BINARY_NAME}"

# Copy migrator binary
COPY --from=compiler --chown=${USER_USER}:${USER_USER} "${OUTPUT_DIR}/${DB_MIGRATE_BINARY}" "${USER_BINARY_DIR}/${DB_MIGRATE_BINARY}"

# To save vendor package if they go unsupported
COPY --from=compiler --chown=${USER_USER}:${USER_USER} "${WORKDIR_NAME}/vendor" "${USER_USER_HOME_DIR}/vendor"

EXPOSE 3001

CMD [ "user-go-http" ]
