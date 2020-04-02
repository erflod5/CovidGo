#![feature(proc_macro_hygiene, decl_macro)]
#[macro_use] extern crate rocket;
extern crate rocket_contrib;
extern crate cpu_monitor;

use rocket::response::content;
use rocket_contrib::serve::StaticFiles;

use std::io;
use std::time::Duration;
use std::thread;
use cpu_monitor::CpuInstant;

#[get("/rt1")]
fn world() -> &'static str {
    "Hello, world!"
}

#[get("/rt2")]
fn json() -> content::Json<&'static str> {
    content::Json("{ 'hi': 'world' }")
}

fn rocket() -> rocket::Rocket {
    rocket::ignite()
    .mount("/", StaticFiles::from("static"))
    .mount("/api", routes![json])
    .mount("/api", routes![world])
}

fn main(){
    thread::spawn(|| {
        while true {
            cputest().expect("Error");
            thread::sleep(Duration::from_millis(5000));
        }
    });

    rocket().launch();
}

fn cputest() -> Result<(),io::Error>{
    let start = CpuInstant::now()?;
    std::thread::sleep(Duration::from_millis(100));
    let end = CpuInstant::now()?;
    let duration = end - start;
    println!("cpu: {:.0}%", duration.non_idle() * 100.);
    Ok(())
}