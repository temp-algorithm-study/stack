# 10773. 제로
# S4
# 구현, 자료구조, 스택

from sys import stdin

# 입력 줄 개수
K = int(input())

# 장부
ledger = []
for _ in range(K):
    
    # 금액
    amount = int(stdin.readline())
    
    if amount == 0: ledger.pop()
    else: ledger.append(amount)

print(sum(ledger))


# 정수가 0일 경우에 지울 수 있는 수가 있음을 보장할 수 있다는 조건이 있으므로 별도의 처리 없이 .pop() 사용
# 빈 리스트일 때 .pop() 사용 시 IndexError 발생