from typing import List


def self_dividing_number(left: int, right: int) -> List[int]:
    def is_self_dividing(n):
        n_copy = n
        while n > 1:
            digit = int(n % 10)
            if digit == 0 or n_copy % digit != 0:
                return False
            n /= 10
        return True

    result = []
    for i in range(left, right + 1):
        if is_self_dividing(i):
            result.append(i)
    return result

print(self_dividing_number(1, 22))
