FROM golang:1.18beta1-bullseye

WORKDIR /app

ENV USER=gopher

RUN useradd --create-home $USER \
  && chown -R $USER:$USER /app

USER $USER

COPY --chown=$USER:$USER . .

CMD ["./test.sh"]