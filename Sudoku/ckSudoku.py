EMPTY_VALUE = '.'


def print_sudoku(sudoku):
    if sudoku is None:
        print 'Sudoku is None'
        return
    for i in xrange(9):
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
        if len(s) == 9:
            return s
    return s


def search_next_empty(sudoku, row):
    for i in xrange(row, 9):
        for j in xrange(0, 9):
            if sudoku[i][j] == EMPTY_VALUE:
                return (i, j)
    return None, None


def check_sudoku(sudoku, row, col):
    checked = set()
    for i in xrange(9):
        value = sudoku[row][i]
        if value == EMPTY_VALUE:
            continue
        if value in checked:
            return False
        checked.add(value)

    checked = set()
    for i in xrange(9):
        value = sudoku[i][col]
        if value == EMPTY_VALUE:
            continue
        if value in checked:
            return False
        checked.add(value)

    def check_33(r, c, sudoku):
        checked = set()
        row_start = r * 3
        col_start = c * 3
        for i in xrange(3):
            for j in xrange(3):
                value = sudoku[row_start + i][col_start + j]
                if value == EMPTY_VALUE:
                    continue
                if value in checked:
                    return False
                checked.add(value)
        return True

    ch_res = check_33(row / 3, col / 3, sudoku)
    if not ch_res:
        return False
    return True


def solve(sudoku, row, col):
    r, c = search_next_empty(sudoku, row)
    if r is None:
        return sudoku

    for v in xrange(1, 10):
        sudoku[r][c] = str(v)
        if r < 2:
            print_sudoku(sudoku)
        if check_sudoku(sudoku, r, c):
            res = solve(sudoku, r, c)
            if res is not None:
                return res
        sudoku[r][c] = EMPTY_VALUE
    return None

sudoku = load_sukoku()
sudoku = solve(sudoku, 0, 0)
print_sudoku(sudoku)
