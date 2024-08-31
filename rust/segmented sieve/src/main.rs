use std::cmp::min;
use crossbeam_channel::{unbounded, Receiver, Sender};
use std::thread;

const THREADS: i64 = 8;

fn sieve(n: usize) -> Vec<usize> {
    let mut x: Vec<bool> = vec![true; n];
    let mut i = 2;
    let n_root = (n as f64).sqrt() as usize;
    while i < n_root {
        let mut j = 2 * i;
        while j < x.len() {
            x[j] = false;
            j += i;
        }
        i += 1;
    }
    (2..x.len())
        .filter(|y| x[*y])
        .collect()
}

fn segment(offset: usize, chunk_size: usize, seed_primes: &Vec<usize>) -> Vec<usize> {
    let mut x: Vec<bool> = vec![true; chunk_size];
    for i in seed_primes {
        let mut j = {
            if (offset % i) == 0 {
               0 
            } else {
                i - (offset % i)
            }
        };
        while j < x.len() {
            x[j] = false;
            j += i;
        }
    }
    (0..x.len())
        .filter(|y| x[*y])
        .map(|y| y + offset)
        .collect()
}

fn worker(receiver: Receiver<(usize, usize)>, seed_primes: &Vec<usize>, sender: Sender<Vec<usize>>) {
    loop {
        let (offset, length) = match receiver.recv() {
            Ok((offset, length)) => (offset, length),
            Err(_) => break,
        };
        let primes = segment(offset, length, seed_primes);
        sender.send(primes).unwrap();
    }
}

fn sorted_insert(target: &mut Vec<usize>, source: Vec<usize>) {
    let first_number = source[0];
    let mut upper = target.len();
    let mut lower = 0;
    while lower < upper {
        let mid = (upper + lower) / 2;
        if first_number < target[mid] {
            upper = mid;
        } else if first_number > target[mid] {
            lower = mid;
        }
        if upper - lower == 1 {
            break;
        }
    }
     insert position is upper
    let insertion = vec![0; source.len()];
    target.extend(insertion);
    let mut i = target.len() - 1;
    let mut j = 0;
    while i - source.len() >= upper {
        target[i] = target[i - source.len()];
        i -= 1;
    }
    for j in 0..source.len() {
        target[upper + j] = source[j];
    }
}

fn main() {
    let n = 18_446_744_073;
    let n = 100;
    let chunk_size = (n as f64).sqrt().ceil() as usize;
    let seed_primes = sieve(chunk_size);
    let mut total_chunks = 0;

    let (input_sender, input_receiver) = unbounded();
    let (output_sender, output_receiver) = unbounded();
    thread::scope(|s| {
        for _ in 0..THREADS {
            let receiver = input_receiver.clone();
            let sender = output_sender.clone();
            let seed_ref = &seed_primes;
            s.spawn(move || {
                worker(receiver, seed_ref, sender);
            });
        }
        let mut start = chunk_size + 1;
        let mut end = start + chunk_size;
        while start < n {
            input_sender.send((start, min(end, n) - start));
            start += chunk_size;
            end += chunk_size;
            total_chunks += 1;
        }
        drop(input_sender);
    });

    let mut primes = seed_primes;
    for _ in 0..total_chunks {
        let chunk = output_receiver.recv().unwrap();
        sorted_insert(&mut primes, chunk);
    }
    println!("{:?}", primes)
}
