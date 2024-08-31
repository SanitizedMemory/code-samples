import subprocess
import sys


if __name__ == '__main__':
    old = "/Users/michael/Desktop/rusty/old"
    new = "/Users/michael/Desktop/rusty/new"
    x = 0
    for i in range(100):
        x += float([i for i in subprocess.run(["/usr/bin/time", old], capture_output=True, text=True).stderr.split(" ") if len(i) != 0][0])
    print(x/100)

    y = 0
    for i in range(100):
        y += float([i for i in subprocess.run(["/usr/bin/time", new], capture_output=True, text=True).stderr.split(" ") if len(i) != 0][0])
    print(y/100)

    print("result:", y / x * 100)
