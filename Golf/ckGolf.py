field = []
start = 0, 0
finish = 4, 7
silent = True


def forward_passage(x, y, path):
    value = field[x][y]
    path.append([(x, y), value])

    if (not silent):
        print(path)

    if ((x, y) == finish):
        print('result =', path)
        return

    move_forward(x + value, y, path, 'right')
    move_forward(x - value, y, path, 'left')
    move_forward(x, y + value, path, 'down')
    move_forward(x, y - value, path, 'up')

    path.pop()


def move_forward(x, y, path, side):
    if (x < 0 or x >= len(field)):
        return
    if (y < 0 or y >= len(field[0])):
        return

    for p in path:
        if ((x, y) == p[0]):
            return

    value = field[x][y]
    if (value < 0):
        path.append([(x, y), value])
        if (side == 'right'):
            move_forward(x + value, y, path, 'left')
        elif (side == 'left'):
            move_forward(x - value, y, path, 'right')
        elif (side == 'down'):
            move_forward(x, y + value, path, 'up')
        elif (side == 'up'):
            move_forward(x, y - value, path, 'down')
        path.pop()
    else:
        forward_passage(x, y, path)

for line in open("inputGolf.txt"):
    field.append([int(x) for x in line.split()])


forward_passage(start[0], start[1], [])
