# 10828. 스택
# S4
# 구현, 자료구조, 스택

from sys import stdin

# 입력 줄 개수
N = int(input())

stack = []
for _ in range(N):
    order = stdin.readline().strip();
    if " " in order: action, num = order.split()
    else: action, num = order, None

    try:
      match action:
          case "push": stack.append(num)
          case "pop": print(stack.pop())
          case "size": print(len(stack))
          case "empty": print(1 if len(stack) == 0 else 0)
          case "top": print(stack[-1])
        
    except IndexError:
       print(-1)


# match-case는 break가 없다..!