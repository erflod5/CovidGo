FROM rust:1.41.0 as build-env

EXPOSE 8000
#WORKDIR /usr/src/hello-rocket

#COPY . .

#RUN rustup default nightly
#RUN cargo build --release
#RUN cargo install --path .

#CMD ["/usr/local/cargo/bin/hello-rocket"]
WORKDIR /app
ADD . /app

RUN rustup default nightly
RUN cargo build --release
#RUN cargo install --path .

FROM gcr.io/distroless/cc
COPY --from=build-env /app/target/release/hello-rocket /
COPY --from=build-env /app/static/ /static/

CMD ["./hello-rocket"]