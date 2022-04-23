#include <iostream>
#include <string>
#include <vector>
using namespace std;

template<class T> void chmin(T& a, T b) {
    if (a > b) {
        a = b;
    }
}

void printTable(vector<vector<int>> td) {
    for (int i = 0; i < td.size(); i++) {
        for (int j = 0; j < td.size(); j++) {
            cout << td[i][j];
            cout << " ";
        }
        cout << endl;
    }
}

const int INF = 1 << 29;

// if [ -e "bar" ];then rm bar;fi;g++ -std=c++11 -o bar bar.cpp && ./bar
int main() {

    string S = "logistic";
    string T = "algorithm";

    vector<vector<int>> dp(S.size()+1, vector<int>(T.size()+1, INF));

    dp[0][0] = 0;

    printTable(dp);
    // DPループ
    for (int i = 0; i <= S.size(); ++i) {
        for (int j = 0; j <= T.size(); ++j) {
            if (i > 0 && j > 0) {
                if (S[i - 1] == T[j - 1]) {
                    chmin(dp[i][j], dp[i-1][j-1]);
                }
                else {
                    chmin(dp[i][j], dp[i-1][j-1]+1);
                }
            }
            if (i > 0) chmin(dp[i][j], dp[i-1][j]+1);
            if (j > 0) chmin(dp[i][j], dp[i][j-1]+1);
        }
    }
    printTable(dp);

    cout << dp[S.size()][T.size()] << endl;
}
