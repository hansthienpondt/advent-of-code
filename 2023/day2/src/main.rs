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
    let maxred: u32 = 12;
    let maxgreen: u32 = 13;
    let maxblue: u32 = 14;

    // File input.txt must exist in the current path
    if let Ok(lines) = read_lines(filename) {
        for line in lines {
            if let Ok(l) = line {
                let lineSlice: Vec<&str> = l.split(":").collect();
                let gameStr: Vec<&str> = lineSlice[0].split(" ").collect();
                let gameNum: u32 = gameStr[1].parse().unwrap();
                let mut possible = true;

                // println!("{:?} {:?} {:?}", lineSlice, gameStr, gameNum);
                for colors in lineSlice[1].split(";") {
                    let mut red: u32 = 0;
                    let mut green: u32 = 0;
                    let mut blue: u32 = 0;

                    for color in colors.split(",") {
                        let colorSlice: Vec<&str> = color.split(" ").collect();
                        let colorValue: u32 = colorSlice[1].parse::<u32>().unwrap();

                        match colorSlice[2] {
                            "blue" => blue += colorValue,
                            "red" => red += colorValue,
                            "green" => green += colorValue,
                            _ => println!("AI! No color match"),
                        }
                    }

                    //println!("{} {} {}", blue, red, green);
                    if red > maxred || green > maxgreen || blue > maxblue {
                        possible = false;
                    }
                }
                if possible {
                    sum += gameNum;
                }
            }
        }
    }
    return sum;
}
fn day1b(filename: String) -> u32 {
    let mut sum: u32 = 0;

    // File input.txt must exist in the current path
    if let Ok(lines) = read_lines(filename) {
        for line in lines {
            if let Ok(l) = line {
                let lineSlice: Vec<&str> = l.split(":").collect();
                let gameStr: Vec<&str> = lineSlice[0].split(" ").collect();

                let mut red: u32 = 0;
                let mut green: u32 = 0;
                let mut blue: u32 = 0;
                // println!("{:?} {:?} {:?}", lineSlice, gameStr, gameNum);
                for colors in lineSlice[1].split(";") {
                    for color in colors.split(",") {
                        let colorSlice: Vec<&str> = color.split(" ").collect();
                        let colorValue: u32 = colorSlice[1].parse::<u32>().unwrap();

                        match colorSlice[2] {
                            "blue" => {
                                if colorValue > blue {
                                    blue = colorValue
                                }
                            }
                            "red" => {
                                if colorValue > red {
                                    red = colorValue
                                }
                            }
                            "green" => {
                                if colorValue > green {
                                    green = colorValue
                                }
                            }
                            _ => println!("AI! No color match"),
                        }
                    }

                    //println!("{} {} {}", blue, red, green);
                }
                sum += red * green * blue;
                //println!("{}", sum);
            }
        }
    }
    return sum;
}
