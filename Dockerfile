FROM bitnami/minideb:latest

ENV GITHUB_TOKEN="" GITHUB_USER_EMAIL="" GITHUB_USER_NAME="" V3SWAGGER=""

# The place to store bind a dir that will contain
# the swagger files to update:

VOLUME /data

# A script to create a PR with the updated files
COPY bin/create_update_pr /create_update_pr
COPY bin/gh /usr/bin

RUN install_packages ca-certificates curl git openssh-client

# RUN apt-get update
# RUN apt-get install -y gnupg
# RUN apt-get install -y software-properties-common
# RUN apt-key adv --keyserver keyserver.ubuntu.com --recv-key C99B11DEB97541F0
# RUN apt-add-repository https://cli.github.com/packages
# RUN apt-get update
# RUN apt-get install -y gh

CMD ["bash", "/create_update_pr"]
