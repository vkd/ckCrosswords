import copy

EMPTY_VALUE = '.'


def print_sudoku(sudoku):
    if sudoku is None:
        print 'Sudoku is None'
        return
    for i in xrange(9):
        # print sudoku[i]
        line = ''
        for j in xrange(9):
            line += sudoku[i][j]
            if j % 3 == 2 and j != 8:
                line += '|'
        print line
        if i % 3 == 2 and i != 8:
            print '--- --- ---'
    print ''


def load_sukoku():
    s = []
    for line in open("inputSudoku.txt"):
        s.append([c for c in line[:-1]])
    return s


def search_next_empty(sudoku):
    for i in xrange(9):
        for j in xrange(9):
            if sudoku[i][j] == EMPTY_VALUE:
                return (i, j)


def check_sudoku(sudoku, row, col):
    checked = set()
    for i in xrange(9):
        value = sudoku[row][i]
        if value == '.':
            continue
        # print 'value =', value
        if value in checked:
            return False
        checked.add(value)
    # print 'rows checked'

    checked = set()
    for i in xrange(9):
        value = sudoku[i][col]
        if value == '.':
            continue
        if value in checked:
            return False
        checked.add(value)
    # print 'columns checked'

    def check_33(r, c, sudoku):
        checked = set()
        for i in xrange(3):
            for j in xrange(3):
                value = sudoku[r * 3 + i][c * 3 + j]
                if value == '.':
                    continue
                if value in checked:
                    return False
                checked.add(value)
        return True

    ch_res = check_33(row / 3, col / 3, sudoku)
    if not ch_res:
        return False
    # print 'cubes checked'
    return True


def solve(sudoku):
    empty_pnt = search_next_empty(sudoku)
    if empty_pnt is None:
        return sudoku

    s = copy.deepcopy(sudoku)
    for v in xrange(1, 10):
        row = empty_pnt[0]
        col = empty_pnt[1]
        s[row][col] = str(v)
        if row < 3:
            print_sudoku(s)
        if check_sudoku(s, row, col):
            # print 'checked'
            # print '------------------------------'
            res = solve(s)
            if res is not None:
                return res
        else:
            # print 'EMPTY_VALUE'
            # print '------------------------------'
            s[row][col] = EMPTY_VALUE
    return None

sudoku = load_sukoku()
sudoku = solve(sudoku)
print_sudoku(sudoku)
