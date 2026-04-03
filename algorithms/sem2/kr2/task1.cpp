#include <iostream>
#include <vector>
#include <random>
#include <ctime>

using namespace std;

struct Node {
    int value;
    int count;
    Node* left;
    Node* right;
};

Node* createNode(int value) {
    Node* node = new Node{value, 1, nullptr, nullptr};
    return node;
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
    } else {
        root->count++;
    }
}

void printTree(Node* root) {
    if (root == nullptr) return;
    printTree(root->left);
    for (int i = 0; i < root->count; i++) {
        cout << root->value << " ";
    }
    printTree(root->right);
}

int depthTree(Node* root) {
    if (root == nullptr) return 0;
    int leftDepth = depthTree(root->left);
    int rightDepth = depthTree(root->right);
    return (leftDepth > rightDepth ? leftDepth : rightDepth) + 1;
}

Node* searchNode(int value, Node* root) {
    if (root == nullptr) return nullptr;
    if (value == root->value) return root;
    if (value < root->value) return searchNode(value, root->left);
    return searchNode(value, root->right);
}

Node* findMax(Node* root) {
    if (root == nullptr) return nullptr;
    while (root->right != nullptr) {
        root = root->right;
    }
    return root;
}

void deleteNode(int value, Node*& root) {
    if (root == nullptr) return;

    if (value < root->value) {
        deleteNode(value, root->left);
        return;
    }

    if (value > root->value) {
        deleteNode(value, root->right);
        return;
    }

    if (root->count > 1) {
        root->count--;
        return;
    }

    if (root->left == nullptr && root->right == nullptr) {
        delete root;
        root = nullptr;
        return;
    }

    if (root->left == nullptr) {
        Node* temp = root;
        root = root->right;
        delete temp;
        return;
    }

    if (root->right == nullptr) {
        Node* temp = root;
        root = root->left;
        delete temp;
        return;
    }

    Node* maxLeft = findMax(root->left);
    root->value = maxLeft->value;
    root->count = maxLeft->count;
    maxLeft->count = 1;
    deleteNode(maxLeft->value, root->left);
}

void clearTree(Node*& root) {
    if (root == nullptr) return;
    clearTree(root->left);
    clearTree(root->right);
    delete root;
    root = nullptr;
}

// main делал через нейронку для красивого вывода и тестов
int main() {
    Node* root = nullptr;

    int n;
    cout << "Введите количество элементов дерева: ";
    cin >> n;

    mt19937 gen(time(0));
    uniform_int_distribution<int> dist(1, 100);

    vector<int> generated;
    for (int i = 0; i < n; i++) {
        int value = dist(gen);
        generated.push_back(value);
        addNode(value, root);
    }

    cout << "Сгенерированные значения: ";
    for (int x : generated) {
        cout << x << " ";
    }
    cout << endl;

    cout << "Симметричный обход дерева: ";
    printTree(root);
    cout << endl;

    cout << "Высота дерева: " << depthTree(root) << endl;

    int newValue;
    cout << "Введите значение для добавления: ";
    cin >> newValue;
    addNode(newValue, root);

    cout << "Дерево после добавления: ";
    printTree(root);
    cout << endl;

    int searchValue;
    cout << "Введите значение для поиска: ";
    cin >> searchValue;
    Node* found = searchNode(searchValue, root);
    if (found == nullptr) {
        cout << "Элемент не найден" << endl;
    } else {
        cout << "Элемент найден: " << found->value << ", count = " << found->count << endl;
    }

    int deleteValue;
    cout << "Введите значение для удаления: ";
    cin >> deleteValue;
    deleteNode(deleteValue, root);

    cout << "Дерево после удаления: ";
    printTree(root);
    cout << endl;

    clearTree(root);
    return 0;
}