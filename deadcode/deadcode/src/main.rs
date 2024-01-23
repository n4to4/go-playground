fn main() {
    let g: &dyn Greeter;
    g = &Helloer;
    g.greet();
}

trait Greeter {
    fn greet(&self);
}

struct Helloer;
struct Goodbyer;

impl Greeter for Helloer {
    fn greet(&self) {
        hello();
    }
}

impl Greeter for Goodbyer {
    fn greet(&self) {
        goodbye();
    }
}

fn hello() {
    println!("hello");
}

fn goodbye() {
    println!("goodbye");
}
