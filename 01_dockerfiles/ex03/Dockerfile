# Use an official Ubuntu runtime as a parent image
FROM ubuntu

# Install needed packages
RUN apt-get update && \
	apt-get upgrade -y && \
	apt-get install -y \
	curl \
	openssh-server \
	ca-certificates \
	postfix
RUN curl https://packages.gitlab.com/install/repositories/gitlab/gitlab-ce/script.deb.sh | bash && \
	apt-get install -y gitlab-ce

# Make port available to the world outside this container
EXPOSE 443 80 22

ENTRYPOINT (/opt/gitlab/embedded/bin/runsvdir-start &) && gitlab-ctl reconfigure && gitlab-ctl tail

# Follow steps to launch docker container
# Build
# docker build -t ex03 .
# Run
# docker run -it -p 8443:443 -p 8080:80 -p 8022:22 --privileged ex03

