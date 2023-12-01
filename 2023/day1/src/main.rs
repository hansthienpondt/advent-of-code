use getopts::Options;

use std::collections::HashMap;
use std::env;
use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;

fn main() {
    let args: Vec<String> = env::args().collect();
    let program = args[0].clone();

    let mut opts = Options::new();
    opts.optopt("f", "file", "input file", "default is input.txt");
    opts.optflag("h", "help", "print this help menu");
    let matches = opts.parse(&args).unwrap();
    // if matches.opt_present("h") || !matches.opt_present("f") {
    if matches.opt_present("h") {
        let brief = format!("Usage: {} [options]", program);
        print!("{}", opts.usage(&brief));
        return;
    }

    let filename = matches.opt_str("f").unwrap_or("src/input.txt".to_owned());

    let day1a: u32 = day1a(filename.clone());
    println!("Day1a SUM is {}", day1a);
    //let testname: String = "src/testinput1b.txt".to_string();
    let day1b: u32 = day1b(filename.clone());
    println!("Day1b SUM is {}", day1b);
}

// The output is wrapped in a Result to allow matching on errors
// Returns an Iterator to the Reader of the lines of the file.
fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>>
where
    P: AsRef<Path>,
{
    let file = File::open(filename)?;
    Ok(io::BufReader::new(file).lines())
}

fn day1a(filename: String) -> u32 {
    let mut sum: u32 = 0;

    // File input.txt must exist in the current path
    if let Ok(lines) = read_lines(filename) {
        // Consumes the iterator, returns an (Optional) String
        for line in lines {
            if let Ok(l) = line {
                let mut vec = Vec::new();
                //for (i, c) in l.chars().enumerate() {
                for c in l.chars() {
                    if c.is_numeric() {
                        let number: u32 = c.to_string().parse().unwrap();
                        vec.push(number)
                    }
                }
                let linenumber: u32 = vec[0] * 10 + vec[vec.len() - 1];
                sum += linenumber;
            }
        }
    }
    return sum;
}
fn day1b(filename: String) -> u32 {
    let mut sum: u32 = 0;
    let mut data: HashMap<u32, String> = HashMap::new();

    data.insert(1, String::from("one"));
    data.insert(2, String::from("two"));
    data.insert(3, String::from("three"));
    data.insert(4, String::from("four"));
    data.insert(5, String::from("five"));
    data.insert(6, String::from("six"));
    data.insert(7, String::from("seven"));
    data.insert(8, String::from("eight"));
    data.insert(9, String::from("nine"));
    //println!("HashMap = {:?}", data);

    // File input.txt must exist in the current path
    if let Ok(lines) = read_lines(filename) {
        // Consumes the iterator, returns an (Optional) String
        for line in lines {
            if let Ok(l) = line {
                let mut result: HashMap<usize, u32> = HashMap::new();
                //println!("{}", l);
                for (word, value) in &data {
                    for (index, _) in l.to_lowercase().match_indices(&value.to_lowercase()) {
                        result.insert(index, *word);
                        //println!("{} {} found at {}", &value, &word, index);
                    }
                }
                for (index, c) in l.chars().enumerate() {
                    if c.is_numeric() {
                        let number: u32 = c.to_string().parse().unwrap();
                        result.insert(index, number);
                    }
                }
                let Some(min) = result.keys().min() else {
                    todo!()
                };
                let Some(max) = result.keys().max() else {
                    todo!()
                };

                let linenumber: u32 = result[min] * 10 + result[max];
                sum += linenumber;
            }
        }
    }
    return sum;
}
