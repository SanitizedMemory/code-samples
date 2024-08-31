use num::{cast, pow::Pow, NumCast};

#[derive(Debug)]
struct Point<T, U>
where
    T: NumCast + Copy,
    U: NumCast + Copy,
{
    x: T,
    y: U,
}

trait HasLength {
    fn length<T: NumCast>(&self) -> Option<T>;
}

impl<T, U> HasLength for Point<T, U>
where
    T: NumCast + Copy,
    U: NumCast + Copy,
{
    fn length<V: NumCast>(&self) -> Option<V> {
        let x: f32 = cast(self.x)?;
        let y: f32 = cast(self.y)?;
        let a: f32 = x.pow(2) + y.pow(2);
        cast(a.sqrt())
    }
}

#[derive(Debug)]
struct Rectangle<T: NumCast + Copy> {
    perimeter: Option<T>,
    area: Option<T>,
}

trait Shape<T: NumCast> {
    fn perimeter(&self) -> Option<T>;
    fn area(&self) -> Option<T>;
}

impl<T: NumCast + Copy> Shape<T> for Rectangle<T> {
    fn perimeter(&self) -> Option<T> {
        None
    }

    fn area(&self) -> Option<T> {
        None
    }
}

fn main() {
    let a = Point { x: 10, y: 20.0 }; // Point::new(10, 20.0);
    println!("(x, y) = ({}, {})", a.x, a.y);
    println!("Length of line segment: {}", a.length::<f32>().unwrap())
}
