# Use an official Ruby runtime as a parent image
FROM ruby:2.6.3

# Install needed packages
RUN apt-get update -qq && apt-get install -y nodejs postgresql-client sqlite3
RUN gem install bundler -v 2.0.1

# Set the working directory
COPY app /opt/app
WORKDIR /opt/app

# An ONBUILD command executes after the current Dockerfile build completes when the parent image is called
ONBUILD RUN bundler install
ONBUILD RUN bundle exec rake db:migrate
ONBUILD RUN bundle exec rake db:seed

# Follow steps to launch docker container
# Download or clone an application compatible with ruby 2.6.3 rails 5.1.6 bundle 2.0.1 in app directory
# I use this app:
# git clone https://github.com/gitpod-io/rails_sample_app.git app
# Build
# docker build -t ft-rail:on-build .
# Copy the Dockerfile from the subjet in Dockerfile_ex02 and build it
# docker build -t ex02 -f Dockerfile_ex02 .
# docker run --rm -it -p 3000:3000 ex02
