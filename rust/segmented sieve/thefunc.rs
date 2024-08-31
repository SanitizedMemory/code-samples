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
    // insert position is upper
    let insertion = vec![0; source.len()];
    target.extend(insertion);
    let mut i = target.len() - 1;
    while i - source.len() >= upper {
        target[i] = target[i - source.len()];
        i -= 1;
    }
    for j in 0..source.len() {
        target[upper + j] = source[j];
    }
}

fn main() {
    let mut target = vec![0, 1, 2, 3, 100, 200, 300];
    let source = vec![20, 30, 40];
    sorted_insert(&mut target, source);
    println!("{:?}", target);
}
