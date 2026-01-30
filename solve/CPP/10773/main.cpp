#include <iostream>
using namespace std;

int main() {
    ios::sync_with_stdio(false); // 입출력 빨라짐 c와 동기화 비활성화 로 인해
    cin.tie(nullptr); // cin 과 cout 묶음을 풀어서 입출력 빨라짐
    
    int K=0;

    cin >> K;
    int zero[100000];
    int zcount=0;
    int N=0;
    int sum=0;

    for(int i=0; i < K; i++){
        cin >> N;
        if(N == 0){
            if (zcount > 0) zcount--;
        }
        else{
            zero[zcount] = N;
            zcount++;
        }
    }

    for(int i=0; i<zcount; i++){
        sum += zero[i];
    }
    cout << sum << '\n';

    return 0;
}
