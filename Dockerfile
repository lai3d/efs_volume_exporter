FROM        busybox
LABEL maintainer="Larry Lai<cgpanda.ca@gmail.com>"

COPY ./efs_volume_exporter /bin/efs_volume_exporter

USER 1001

ENTRYPOINT [ "/bin/efs_volume_exporter" ]
CMD        [ "--volume-dir=bin:/bin", \
             "--web.listen-address=:9888" ]
