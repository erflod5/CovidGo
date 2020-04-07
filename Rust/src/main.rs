#![feature(proc_macro_hygiene, decl_macro)]
#[macro_use] extern crate rocket;
extern crate rocket_contrib;
extern crate redis;
extern crate rocket_cors;

//ROCKET
//use rocket::response::content;
use rocket_contrib::serve::StaticFiles;

//cpu
use std::time::Duration;
use std::thread;

//redis
use redis::Commands;

//CORS
use rocket::http::Method; // 1.
use rocket_cors::{
    AllowedHeaders, AllowedOrigins, // 2.
    Cors, CorsOptions // 3.
};

//HTTP-CLIENT
use futures::{Future, Stream};
use reqwest::r#async::{Client, Decoder};
use std::mem;

#[get("/cpu")]
fn world() -> String {
    let i = get_last();
    println!("{:?}",i);
    format!("{:?}", i)
}

fn rocket() -> rocket::Rocket {
    rocket::ignite()
    .mount("/", StaticFiles::from("static"))
    .mount("/api", routes![world])
    .attach(make_cors()) // 7.
}

fn make_cors() -> Cors {
    let allowed_origins = AllowedOrigins::some_exact(&[ // 4.
        "http://localhost:5001",
        "http://127.0.0.1:5001",
        "http://localhost:5001",
        "http://0.0.0.0:5001",
    ]);

    CorsOptions { // 5.
        allowed_origins,
        allowed_methods: vec![Method::Get].into_iter().map(From::from).collect(), // 1.
        allowed_headers: AllowedHeaders::some(&[
            "Authorization",
            "Accept",
            "Access-Control-Allow-Origin", // 6.
        ]),
        allow_credentials: true,
        ..Default::default()
    }
    .to_cors()
    .expect("error while building CORS")
}

fn main(){
    thread::spawn(|| {
        while true {
            let f = post_greeting();
            tokio::run(f);
            thread::sleep(Duration::from_millis(5000));
        }
    });
    rocket().launch();
}

fn add_to_rust(val : String, val2 : String) -> redis::RedisResult<()> {
    let client = redis::Client::open("redis://192.168.1.187:7001")?;
    let mut con = client.get_connection()?;
    let _ : () = con.lpush("cpu",val)?;
    let _ : () = con.set("last",val2)?;
    Ok(())
}

fn post_greeting() -> impl Future<Item=(), Error=()> {
    Client::new()
        .get("http://localhost:8002/cpu")
        .send()
        .and_then(|mut res| {
            let body = mem::replace(res.body_mut(), Decoder::empty());
            body.concat2()
        })
        .map_err(|err| println!("request error: {}", err))
        .map(|body| {
            let v = body.to_vec();
            let s = String::from_utf8_lossy(&v);
            println!("response: {} ", s);
            add_to_rust(s.to_string(),s.to_string()).expect("Error");
        })
}

fn get_last() -> redis::RedisResult<String> {
    let client = redis::Client::open("redis://192.168.1.187:7001")?;
    let mut con = client.get_connection()?;
    con.get("last")
}