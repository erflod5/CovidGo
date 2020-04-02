FROM rust:1.31

WORKDIR /usr/src/app
COPY . .

RUN rustup default nightly
RUN cargo build

CMD ["cargo","run"]