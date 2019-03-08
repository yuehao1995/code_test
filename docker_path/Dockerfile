FROM alpine

# The glibc is not default installed in alphine, in alphine is musl libc
# we need make use of follow symbol link to work around
RUN mkdir -p /opt/eechains/dockerpath/bin

COPY ./bin /opt/eechains/dockerpath/bin
# COPY cfg /opt/eechains/scaring/cfg

ENTRYPOINT ["/opt/eechains/dockerpath/bin/dockerpath"]
