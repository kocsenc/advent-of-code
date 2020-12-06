possible_rows = list(range(128))
possible_cols = list(range(8))


def main(filename='input.txt'):
    with open(filename, 'r') as f:
        seat_assignments = f.readlines()

    print(max(map(seat_id_for_seat, seat_assignments)))


def seat_id_for_seat(seat_code):
    seat_code = list(seat_code.strip())
    seat_row = binary_seat_search(possible_rows, seat_code[:7])
    seat_col = binary_seat_search(possible_cols, seat_code[-3:])

    return seat_row * 8 + seat_col


def binary_seat_search(seat_list, direction_list):
    seat_list_count = len(seat_list)
    if seat_list_count <= 1 or len(direction_list) == 0:
        return seat_list[0]
    
    direction = direction_list.pop(0)
    if direction in ['F', 'L']:
        return binary_seat_search(seat_list[:seat_list_count//2], direction_list)
    elif direction in ['B', 'R']:
        return binary_seat_search(seat_list[seat_list_count//2:], direction_list)


if __name__ == '__main__':
    main()