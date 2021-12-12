FROM golang:1.17.5

WORKDIR /app

ENV USER=gouser

RUN useradd --create-home $USER \
  && chown -R $USER:$USER /app

USER $USER

COPY --chown=$USER:$USER . .

CMD ["./test.sh"]