FROM rust:1.41.0 as build-env

EXPOSE 8000
WORKDIR /usr/src/hello-rocket

COPY . .

RUN rustup default nightly
RUN cargo build --release
RUN cargo install --path .

CMD ["/usr/local/cargo/bin/hello-rocket"]