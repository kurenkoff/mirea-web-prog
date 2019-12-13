FROM ubuntu

COPY server /opt
COPY players.json /opt
WORKDIR /opt
RUN ls
EXPOSE 8080
ENTRYPOINT ["./server"]