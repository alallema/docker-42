FROM debian:9.8

EXPOSE 443 80 22

RUN apt-get update && DEBIAN_FRONTEND=noninteractive apt-get upgrade -y && \
    apt-get install -y ca-certificates openssh-server wget postfix
RUN sed -i 's/#PermitRootLogin prohibit-password/PermitRootLogin yes/' /etc/ssh/sshd_config

RUN wget https://packages.gitlab.com/install/repositories/gitlab/gitlab-ce/script.deb.sh && \
    chmod +x script.deb.sh && ./script.deb.sh && \
    apt-get install -y gitlab-ce

RUN apt-get clean && rm -rf /var/lib/apt/lists/*


RUN mkdir -p /etc/gitlab/ssl
RUN chmod +x /etc/gitlab/ssl
RUN openssl req -x509 -nodes -days 365 -newkey rsa:2048 \
    -subj /C=FR/ST=Paris/L=Paris/O=42/OU=0/CN=127.0.0.1/emailAddress=alallema@student.42.fr \
    -keyout /etc/gitlab/ssl/selfsigned.key -out /etc/gitlab/ssl/selfsigned.crt

RUN echo "external_url \"https://127.0.0.1\"" >> /etc/gitlab/gitlab.rb
RUN echo "nginx['redirect_http_to_https'] = true" >> /etc/gitlab/gitlab.rb
RUN echo "nginx['ssl_certificate'] = \"/etc/gitlab/ssl/selfsigned.crt\"" >> /etc/gitlab/gitlab.rb
RUN echo "nginx['ssl_certificate_key'] = \"/etc/gitlab/ssl/selfsigned.key\"" >> /etc/gitlab/gitlab.rb
RUN echo "gitlab_rails['gitlab_shell_ssh_port'] = 8022" >> /etc/gitlab/gitlab.rb

ENTRYPOINT service ssh start && (/opt/gitlab/embedded/bin/runsvdir-start &) && gitlab-ctl reconfigure && gitlab-ctl tail

# docker build -t ex03 .
# docker run -it --rm -p 80:80 -p 8022:22 -p 443:443 --privileged ex03
# accept selfsigned certif:  export GIT_SSL_NO_VERIFY=true
