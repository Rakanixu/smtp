# root CA certificates required for secure connection to SMTP server
# Using ubutu image for simplicity
FROM ubuntu:14.04.1
RUN apt-get update
RUN apt-get install -y ca-certificates
ADD srv /
ENTRYPOINT [ "/srv" ]
