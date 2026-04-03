#include <iostream>

using namespace std;

struct Node {
    int value;
    Node* left;
    Node* right;
};

Node* createNode(int value) {
    return new Node{value, nullptr, nullptr};
}

void addNode(int value, Node*& root) {
    if (root == nullptr) {
        root = createNode(value);
        return;
    }

    if (value < root->value) {
        addNode(value, root->left);
    } else if (value > root->value) {
        addNode(value, root->right);
    }
}

int closestValue(Node* root, int x) {
    int best = root->value;
    Node* current = root;

    while (current != nullptr) {
        int curDiff = abs(current->value - x);
        int bestDiff = abs(best - x);

        if (curDiff < bestDiff || (curDiff == bestDiff && current->value < best)) {
            best = current->value;
        }

        if (x < current->value) {
            current = current->left;
        } else if (x > current->value) {
            current = current->right;
        } else {
            return current->value;
        }
    }

    return best;
}

void printTree(Node* root) {
    if (root == nullptr) return;
    printTree(root->left);
    cout << root->value << " ";
    printTree(root->right);
}

void clearTree(Node*& root) {
    if (root == nullptr) return;
    clearTree(root->left);
    clearTree(root->right);
    delete root;
    root = nullptr;
}

int main() {
    Node* root = nullptr;

    int n;
    cin >> n;

    for (int i = 0; i < n; i++) {
        int x;
        cin >> x;
        addNode(x, root);
    }

    int target;
    cin >> target;

    cout << closestValue(root, target) << endl;

    clearTree(root);
    return 0;
}