FROM busybox:musl
COPY router /
COPY conf /conf
EXPOSE 9888
CMD ["/router"]
