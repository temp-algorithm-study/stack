# 9012. 괄호
# S4
# 자료구조, 문자열, 스택

from sys import stdin

# 입력 줄 개수
T = int(input())

stack = []
for _ in range(T):
    parens = stdin.readline();
    cnt_open = 0
    cnt_close = 0

    for paren in parens:
        match paren:
            case "(": cnt_open += 1     
            case ")": cnt_close += 1
        if cnt_close > cnt_open: break

    else:
        if cnt_open == cnt_close:
            print("YES")
            continue

    print("NO")