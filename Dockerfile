#
# Harmony Maestro DockerFile
#

FROM ubuntu:14.04.2



#
# Misc Docker Config
#

# set the maintainer
MAINTAINER pmccarren

# Set correct environment variables.
ENV HOME /root

# set the command to run
CMD ["/maestro"]

# enter the container at home
WORKDIR /gocode/src/github.com/dronemill/harmony-maestro



#
# Install packages
#

RUN apt-get update && \
	apt-get install -y wget && \
	wget -q https://storage.googleapis.com/golang/go1.4.2.linux-amd64.tar.gz && \
	tar -C /usr/local -xzf go1.4.2.linux-amd64.tar.gz && \
	rm go1.4.2.linux-amd64.tar.gz


#
# Cleanup & optimize packages
#

RUN find /var/log -type f -delete && \
	apt-get clean && \
	rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*


#
# Go Env
#
ENV PATH $PATH:/usr/local/go/bin
ENV GOPATH /gocode