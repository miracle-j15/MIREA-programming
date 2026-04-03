#include <iostream>
#include <vector>
#include <string>

using namespace std;

struct Node {
    string name;
    vector<string> phones;
    Node* left;
    Node* right;
};

Node* createNode(string name, string phone) {
    Node* node = new Node{name, {phone}, nullptr, nullptr};
    return node;
}

void addContact(string name, string phone, Node*& root) {
    if (root == nullptr) {
        root = createNode(name, phone);
        return;
    }

    if (name < root->name) {
        addContact(name, phone, root->left);
    } else if (name > root->name) {
        addContact(name, phone, root->right);
    } else {
        for (string currentPhone : root->phones) {
            if (currentPhone == phone) return;
        }
        root->phones.push_back(phone);
    }
}

Node* searchContact(string name, Node* root) {
    if (root == nullptr) return nullptr;
    if (name == root->name) return root;
    if (name < root->name) return searchContact(name, root->left);
    return searchContact(name, root->right);
}

Node* findMin(Node* root) {
    while (root != nullptr && root->left != nullptr) {
        root = root->left;
    }
    return root;
}

void deleteContact(string name, Node*& root) {
    if (root == nullptr) return;

    if (name < root->name) {
        deleteContact(name, root->left);
        return;
    }

    if (name > root->name) {
        deleteContact(name, root->right);
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

    Node* minRight = findMin(root->right);
    root->name = minRight->name;
    root->phones = minRight->phones;
    deleteContact(minRight->name, root->right);
}

void printBook(Node* root) {
    if (root == nullptr) return;
    printBook(root->left);
    cout << root->name << ": ";
    for (string phone : root->phones) {
        cout << phone << " ";
    }
    cout << endl;
    printBook(root->right);
}

void clearTree(Node*& root) {
    if (root == nullptr) return;
    clearTree(root->left);
    clearTree(root->right);
    delete root;
    root = nullptr;
}

// main снова в нейронке сделал по тем же причинам
int main() {
    Node* root = nullptr;

    int command;
    do {
        cout << "\n1 - Добавить контакт" << endl;
        cout << "2 - Найти контакт" << endl;
        cout << "3 - Удалить контакт" << endl;
        cout << "4 - Показать всю телефонную книгу" << endl;
        cout << "0 - Выход" << endl;
        cout << "Введите команду: ";
        cin >> command;

        if (command == 1) {
            string name, phone;
            cout << "Введите имя: ";
            cin >> name;
            cout << "Введите телефон: ";
            cin >> phone;
            addContact(name, phone, root);
        } else if (command == 2) {
            string name;
            cout << "Введите имя для поиска: ";
            cin >> name;
            Node* found = searchContact(name, root);
            if (found == nullptr) {
                cout << "Контакт не найден" << endl;
            } else {
                cout << found->name << ": ";
                for (string phone : found->phones) {
                    cout << phone << " ";
                }
                cout << endl;
            }
        } else if (command == 3) {
            string name;
            cout << "Введите имя для удаления: ";
            cin >> name;
            deleteContact(name, root);
        } else if (command == 4) {
            if (root == nullptr) {
                cout << "Телефонная книга пуста" << endl;
            } else {
                printBook(root);
            }
        } else if (command != 0) {
            cout << "Неизвестная команда" << endl;
        }

    } while (command != 0);

    clearTree(root);
    return 0;
}