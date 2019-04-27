# Use an official Debian runtime as a parent image
FROM debian

# Set the working directory
WORKDIR /home/teamspeak

# Install needed packages
RUN apt-get update && \
	apt-get upgrade -y && \
	apt-get install -y \
	wget \
	bzip2 && \
	wget http://dl.4players.de/ts/releases/3.0.13.8/teamspeak3-server_linux_amd64-3.0.13.8.tar.bz2 && \
	tar xfv teamspeak3-server_linux_amd64-3.0.13.8.tar.bz2

WORKDIR teamspeak3-server_linux_amd64

# Make port available to the world outside this container
EXPOSE 9987/udp 10011 30033

# Define environment variable
ENV TS3SERVER_LICENSE=accept

# Open teamspeak
#ENTRYPOINT sh ts3server_startscript.sh start
ENTRYPOINT sh ts3server_minimal_runscript.sh

# Follow steps to launch docker container
# Build
# docker build -t ex01 .
# Run
# docker run --rm -p 9987:9987/udp -p 10011:10011 -p 30033:30033 ex01
