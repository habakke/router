FROM busybox:musl
COPY router /
COPY conf /confx
EXPOSE 9888
CMD ["/router"]
