import random

from numba import jit
from tqdm import tqdm


def main():
    num_trials = 10_000_000
    simulate_monte_hall(num_trials)


def play_monte_hall(should_switch: bool) -> bool:
    doors = [False] * 3
    car_index = random.randint(0, len(doors) - 1)
    doors[car_index] = True

    selected_door_index = random.randint(0, len(doors) - 1)
    revealed_door = next(
        iter(set(range(len(doors))) - {selected_door_index, car_index})
    )

    if should_switch:
        selected_door_index = next(
            iter(set(range(len(doors))) - {selected_door_index, revealed_door})
        )

    return selected_door_index == car_index


def simulate_monte_hall(num_trials: int) -> float:
    num_wins = 0
    for _ in tqdm(range(num_trials)):
        if play_monte_hall(True):
            num_wins += 1
    print(f"Win rate: {round(num_wins / num_trials * 100, 2)}%")


if __name__ == "__main__":
    main()
