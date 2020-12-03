


def main(filename='input.txt'):
    with open(filename, 'r') as f:
        puzzle_lines = f.readlines()

    position1 = 0
    position2 = 0
    position3 = 0
    position4 = 0
    position5 = 0

    tree_count1 = 0
    tree_count2 = 0
    tree_count3 = 0
    tree_count4 = 0
    tree_count5 = 0

    mod = len(puzzle_lines[0]) - 1

    i = 0
    for line in puzzle_lines:
        if line[position1 % mod] == '#':
            tree_count1 += 1
        if line[position2 % mod] == '#':
            tree_count2 += 1
        if line[position3 % mod] == '#':
            tree_count3 += 1
        if line[position4 % mod] == '#':
            tree_count4 += 1

        if i % 2 == 0:
            if line[position5 % mod] == '#':
                tree_count5 += 1
            position5 += 1 # DOWN 2?

        position1 += 1
        position2 += 3
        position3 += 5
        position4 += 7

        i += 1

    print('part 1: {}'.format(tree_count2))
    print('trees: {}'.format(tree_count1 * tree_count2 * tree_count3 * tree_count4 * tree_count5))

'''
Right 1, down 1.
Right 3, down 1. (This is the slope you already checked.)
Right 5, down 1.
Right 7, down 1.
Right 1, down 2.

'''

if __name__ == '__main__':
    main()
