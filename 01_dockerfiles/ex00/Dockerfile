# Use an official Alpine runtime as a parent image
FROM alpine

# Install needed packages
RUN apk update && apk upgrade && apk add vim

# Open vim
ENTRYPOINT ["vim"]

# Follow steps to launch docker container
# Build
# docker built -t ex00 .
# Run
# docker run --rm -ti ex00
