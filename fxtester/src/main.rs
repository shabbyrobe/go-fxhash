extern crate fxhash;

use std::io;
use std::io::Read;

fn main() {
    let mut buf = [0u8; 4];
    let mut rdr = io::stdin();
    
    loop {
        // FIXME: work out how to do this properly
        match rdr.read_exact(&mut buf) {
            Ok(_) => 0,
            Err(_) => break,
        };

        let sz = u32::from_le_bytes(buf);

        let mut data = vec![0u8; sz as usize];
        match rdr.read_exact(&mut data) {
            Ok(_) => 0,
            Err(_) => break,
        };

        let hash = fxhash::hash64(&data);
        println!("{}", hash);
    }
}
