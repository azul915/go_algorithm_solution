#include <iostream>
#include <string>
#include <vector>

using namespace std;

struct Node {
    Node* next;
    string name;
    Node(string name_ = "") : next(NULL), name(name_) { }
};

Node* nil;

void init() {
    nil = new Node();
    nil->next = nil;
}

void printList() {
    Node* cur = nil->next;
    for(; cur!= nil;cur = cur->next) {
        cout << cur->name << " -> ";
    }
    cout << endl;
}

void insert(Node* v, Node* p = nil) {
    cout << p->next << endl;
    v->next = p->next;
    p->next = v;
}

// if [ -e "hoge" ];then rm hoge;fi;g++ -std=c++11 -o hoge hoge.cpp && ./hoge
int main() {
    init();

    vector<string> names = {"yamamoto", "watanabe", "ito", "takahashi","suzuki", "sato"};
    for (int i =0; i < (int)names.size(); ++i) {
        Node* node = new Node(names[i]);

        insert(node);

        cout << "step " << i << ": ";
        printList();
    }
}