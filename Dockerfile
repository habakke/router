FROM busybox:musl
ADD router /
ADD conf /conf
EXPOSE 9888
CMD ["/router"]
