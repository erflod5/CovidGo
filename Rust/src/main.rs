#![feature(proc_macro_hygiene, decl_macro)]
#[macro_use] extern crate rocket;
extern crate rocket_contrib;
extern crate cpu_monitor;
extern crate redis;

use rocket::response::content;
use rocket_contrib::serve::StaticFiles;

use std::io;
use std::time::Duration;
use std::thread;
use cpu_monitor::CpuInstant;
use redis::Commands;

#[get("/cpu")]
fn world() -> String {
    let mut i = 0.0;
    cputest(&mut i).expect("Error");
    format!("{:.0}", i)
}

fn rocket() -> rocket::Rocket {
    rocket::ignite()
    .mount("/", StaticFiles::from("static"))
    .mount("/api", routes![world])
}

fn main(){
    thread::spawn(|| {
        while true {
            let mut i = 3.3;
            cputest(&mut i).expect("Error");
            thread::sleep(Duration::from_millis(5000));
        }
    });

    rocket().launch();
}

fn cputest(e: &mut f64) -> Result<(),io::Error>{
    let start = CpuInstant::now()?;
    std::thread::sleep(Duration::from_millis(100));
    let end = CpuInstant::now()?;
    let duration = end - start;
    *e = duration.non_idle() * 100.;
    do_something(duration.non_idle() * 100.);
    Ok(())
}


fn do_something(val : f64) -> redis::RedisResult<()> {
    let client = redis::Client::open("redis://192.168.1.187:7001")?;
    let mut con = client.get_connection()?;
    let _ : () = con.lpush("cpu",val)?;
    //let _ : () = con.set("my_counter", val)?;
    /*let count : f64 = con.get("my_counter")?;
    println!("Data: {:.0}",count);*/
    Ok(())
}