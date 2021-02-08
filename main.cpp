#include <iostream>
#include <string>
#include <stack>
#include <math.h>

using namespace std;

const int T_S = 20;

class HashTableEntry {
public:
    string k;
    int v;

    HashTableEntry(string k, int v) {
        this->k = k;
        this->v = v;
    }
};

class HashMapTable {
private:
    HashTableEntry **t;
public:
    HashMapTable() {
        t = new HashTableEntry *[T_S];
        for (int i = 0; i < T_S; i++) {
            t[i] = NULL;
        }
    }

    // changing with string version
    unsigned int hash_func(string key) {
        unsigned int h = 0;
        unsigned int o = 31415;
        const unsigned int t = 27183;
        for (int i = 0; i < key.length(); ++i) {
            h = (o * h + key[i]) % T_S;
            o = o * t % (T_S - 1);
        }
        return h;
    }

    void Insert(string k, int v) {
        int h = hash_func(k);
        while (t[h] != NULL && t[h]->k != k) {
            h = (h + 1) % T_S;
        }
        if (t[h] != NULL)
            delete t[h];
        t[h] = new HashTableEntry(k, v);
    }

    int SearchKey(string k) {
        int h = hash_func(k);
        while (t[h] != NULL && t[h]->k != k) {
            h = (h + 1) % T_S;
        }
        if (t[h] == NULL)
            return -1;
        else
            return t[h]->v;
    }

    void display() {
        for (int i = 0; i < T_S; ++i) {

            cout << "for ";
            cout << i;
            if (t[i] != NULL) {
                cout << "  " + t[i]->k + "   ";
                cout << t[i]->v << endl;
            } else
                cout << "\n";
        }
    }

    ~HashMapTable() {
        for (int i = 0; i < T_S; i++) {
            if (t[i] != NULL)
                delete t[i];
            delete[] t;
        }
    }
};

//////////////////////////////////////////////////// Linked list implementation :
// Each link object is a node that we define in this class
class Node {
public:
    string key;
    Node *next;
    Node *prev;

    Node(string key);
};

Node *cursor = NULL;

// Node class constructor
Node::Node(string newKey) {
    key = newKey;
}

// Create a node in memory and return the address
Node *create_node(string newkey) {
    Node *temp = new Node(newkey);
    temp->next = NULL;
    temp->prev = NULL;
    return temp;
}

class List {
public:
    Node *head;
    Node *tail;
    string string1;

    List();

    void insert_at_head(string newKey);

    void insert_at_tail(string newKey);

    void remove_from_head();

    void remove_from_tail();

    void insert(string newKey);

    void remove();

    void print_list();

    string getString();

    void deleteNode(Node *tmp);
};

List::List() {
    head = NULL;
    tail = NULL;
    string1 = "";
}

void List::deleteNode(Node *tmp) {
    if (tmp != NULL) {
        if (tmp == head) {
            remove_from_head();
        } else if (tmp == tail) {
            remove_from_tail();
        } else {
            Node *tmp2 = cursor;
            tmp->prev->next = tmp->next;
            tmp->next->prev = tmp->prev;
            delete tmp2;
        }
    }
}

string List::getString() {
    Node *tmp = head;
    string res;
    while (tmp != NULL) {
        res += tmp->key;
        tmp = tmp->next;
    }
    return res;
}

// Inserts a new node at the head
void List::insert_at_head(string newKey) {
    Node *temp = create_node(newKey);
    if (head == NULL) {
        head = temp;
        tail = temp;
    } else {
        temp->next = head;
        temp->next->prev = temp;
        head = temp;
        temp->prev = NULL;
    }
}

// Adds a new node at tail
void List::insert_at_tail(string newKey) {
    Node *temp = create_node(newKey);
    if (tail == NULL) {
        tail = temp;
        head = temp;
    } else {
        temp->prev = tail;
        temp->prev->next = temp;
        tail = temp;
        temp->next = NULL;
    }

}

// Removes the node at the head
void List::remove_from_head() {
    Node *temp = head;
    head = temp->next;
    head->prev = NULL;
    delete temp;
}

// Deletes the node at the tail
void List::remove_from_tail() {
    Node *temp = tail;
    tail = temp->prev;
    tail->next = NULL;
    delete temp;
}

// Adds a new node anywhere in the list
void List::insert(string newKey) {
    if (cursor == NULL) {
        insert_at_head(newKey);
        cursor = head;
    } else if (cursor == tail) {
        insert_at_tail(newKey);
        cursor = tail;
    } else {
        Node *newNode = create_node(newKey);
        newNode->next = cursor->next;
        cursor->next->prev = newNode;
        newNode->prev = cursor;
        cursor->next = newNode;
        cursor = newNode;
    }
}

// Removes a node at any given place from the list
void List::remove() {
    if (cursor != NULL) {
        if (cursor == head) {
            remove_from_head();
            cursor = NULL;
        } else if (cursor == tail) {
            remove_from_tail();
            cursor = tail;
        } else {
            Node *tmp = cursor;
            cursor->prev->next = cursor->next;
            cursor->next->prev = cursor->prev;
            cursor = cursor->prev;
            delete tmp;
        }
    }
}

int mode(string num) {
    int newNum = 0;
    int max = (int) pow(10, 9) + 7;
    for (int i = 0; i < num.length(); ++i) {
        if (num[i] != '+' && num[i] != '-')
            newNum = (newNum * 10 + (int) num[i] - '0') % max;
    }
    return newNum;
}


// Prints the linked list objects
void List::print_list() {
    bool printed = false;
    Node *temp = head;
    if (cursor == NULL) {
        cout << "|";
        printed = true;
    }
    while (temp != NULL) {
        cout << temp->key;
        if (temp == cursor && !printed) {
            cout << "|";
            printed = true;
        }
        temp = temp->next;
    }
    cout << endl;
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////
int prec(char c) {
    if (c == '*' || c == '/')
        return 2;
    else if (c == '+' || c == '-')
        return 1;
    else
        return -1;
}

List infixToPostfix(List list) {
    std::stack<string> st;
    st.push("N");
    List list2;
    bool isNegetive = false;
    string prev = "";
    Node **addedNodes;
    int counter = 0;
    Node *current = list.head;
    string listString; // will be used for hash table
    while (current != NULL) {
        listString += current->key;
        if (isdigit(current->key[0])) {
            if (prev == ")") {
                Node *newNode = create_node("*");
                addedNodes[counter] = newNode;
                counter++;
                newNode->next = current;
                current->prev->next = newNode;
                newNode->prev = current->prev;
                current->prev = newNode;
                current = current->prev->prev;
            } else {
                string num;
                if (isNegetive) {
                    num += "-";
                    isNegetive = false;
                }
                num += current->key;
                while (current->next != NULL && isdigit(current->next->key[0])) {
                    num += current->next->key;
                    current = current->next;
                }
                if (num.length() >= 8) {
                    std::string s = std::to_string(mode(num));
                    num = s;
                }
                list2.insert_at_tail(num);
            }
        } else if (current->key[0] == '(') {
            if (isdigit(prev[0]) || prev == ")") {
                Node *newNode = create_node("*");
                addedNodes[counter] = newNode;
                counter++;
                newNode->next = current;
                current->prev->next = newNode;
                newNode->prev = current->prev;
                current->prev = newNode;
                current = current->prev->prev;
            } else {
                st.push("(");
            }
        } else if (current->key[0] == ')') {
            while (st.top() != "N" && st.top() != "(") {
                string c = st.top();
                st.pop();
                list2.insert_at_tail(c);
            }
            if (st.top() == "(") {
                string c = st.top();
                st.pop();
            }
        } else {
            if (prev == "+" || prev == "-" || prev == "*" || prev == "/" || prev == "" || prev == "(") {
                if (current->key == "-") {
                    isNegetive = true;
                }
            } else {
                while (st.top() != "N" && prec(current->key[0]) <= prec(st.top()[0])) {
                    string c = st.top();
                    st.pop();
                    list2.insert_at_tail(c);
                }
                st.push(current->key);
            }
        }
        prev = current->key;
        current = current->next;
    }

    // Pop all the remaining elements from the stack
    while (st.top() != "N") {
        string c = st.top();
        st.pop();
        list2.insert_at_tail(c);
    }
    list2.string1 = listString; // yadet bashe in ham beham mirizeeeeee
    for (int i = 0; i < counter; ++i) {
        Node *tmp = addedNodes[i];
        list.deleteNode(tmp);
    }
    return list2;
}

////////////////////////////////////////////////////////////////////////////////////////////////////
string evaluatePostfix(List list) {
    // Create a stack of capacity equal to expression size
    std::stack<string> st;
    //st.push("N");
    Node *current = list.head;
    while (current != NULL) {
        //  cout << current->key;
        // If the scanned character is an operand (number here),
        // push it to the stack.
        if (isdigit(current->key[0]))
            st.push(current->key);
        else if (current->key.length() > 1) {
            if (isdigit(current->key[1]))
                st.push(current->key);
        } else {
            //   cout << "tss";
            long long int val1 = stoll(st.top());
            st.pop();
            long long int val2 = stoll(st.top());
            st.pop();
            long long int res = 0;
            switch (current->key[0]) {
                case '+':
                    res = val2 + val1;
                    break;
                case '-':
                    res = val2 - val1;
                    break;
                case '*':
                    res = val2 * val1;
                    break;
                case '/':
                    res = val2 / val1;
                    break;
            }
            //cout << res << endl;
            std::string s = std::to_string(res);
            if (s.length() >= 8) {
                std::string s2 = std::to_string(mode(s));
            }
            st.push(s);
        }
        current = current->next;
    }
    return st.top();
}

int main() {
    string output = "";
    //HashMapTable table ;
    int numberOfRequests;
    cin >> numberOfRequests;
    string input;
    cin >> input;
    List list;
    for (int i = 0; i < input.length(); i++) {
        std::string s1{input[i]};
        list.insert_at_tail(s1);
    }
    //list.print_list();
    char inp1, inp2;
    for (int i = 0; i < numberOfRequests; ++i) {
        cin >> inp1;
        switch (inp1) {
            case '>':
                if (cursor == NULL) {
                    cursor = list.head;
                } else if (cursor != list.tail) {
                    cursor = cursor->next;
                }
                break;
            case '<':
                if (cursor != NULL) {
                    cursor = cursor->prev;
                }
                break;
            case '+': {
                cin >> inp2;
                std::string s2{inp2};
                list.insert(s2);
            }
                break;
            case '-':
                list.remove();
                break;
            case '?':
                list.print_list();
                output += list.getString() + "\n\n\n";
                break;
            case '!': {
                List post = infixToPostfix(list);
                //cout << post.string1;
                //cout << post.head->next->next->next;
                //post.print_list();
                string res = evaluatePostfix(post);
                long long int val = stoll(res);
                while (val < 0) {
                    val += pow(10, 9) + 7;
                }
                cout << val << endl;
                std::string s2 = std::to_string(val);
                output += s2 + "\n\n\n";
            }
                break;
        }
    }
    cout << output << endl;
    return 0;
}
