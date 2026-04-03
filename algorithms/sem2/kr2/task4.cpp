#include <iostream>

using namespace std;

struct Node {
    int leftBorder;
    int rightBorder;
    Node* left;
    Node* right;
};

Node* createNode(int leftBorder, int rightBorder) {
    return new Node{leftBorder, rightBorder, nullptr, nullptr};
}

Node* buildTree(int leftBorder, int rightBorder) {
    if (leftBorder > rightBorder) return nullptr;

    Node* root = createNode(leftBorder, rightBorder);

    if (leftBorder == rightBorder) return root;

    int middle = (leftBorder + rightBorder) / 2;

    root->left = buildTree(leftBorder, middle);
    root->right = buildTree(middle + 1, rightBorder);

    return root;
}

void printTree(Node* root) {
    if (root == nullptr) return;

    cout << "[" << root->leftBorder << ", " << root->rightBorder << "] ";
    printTree(root->left);
    printTree(root->right);
}

int countSegments(Node* root, int x) {
    if (root == nullptr) return 0;

    int count = 0;
    if (root->leftBorder <= x && x <= root->rightBorder) {
        count = 1;
    }

    return count + countSegments(root->left, x) + countSegments(root->right, x);
}

void clearTree(Node*& root) {
    if (root == nullptr) return;

    clearTree(root->left);
    clearTree(root->right);

    delete root;
    root = nullptr;
}

// этот main уже сам делал))
int main() {
    int l, r;
    cout << "Введите начало и конец интервала: ";
    cin >> l >> r;

    if (l > r) {
        cout << "Некорректный интервал" << endl;
        return 0;
    }

    Node* root = buildTree(l, r);

    cout << "Прямой обход дерева: ";
    printTree(root);
    cout << endl;

    int x;
    cout << "Введите точку X: ";
    cin >> x;

    cout << "Количество интервалов, содержащих X: " << countSegments(root, x) << endl;

    clearTree(root);
    return 0;
}