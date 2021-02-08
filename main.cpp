#include <iostream>
#include <string>
#include <stack>
#include <math.h>
#include <unordered_map>

using namespace std;

struct stringRes {
    string infExpression;
    string tableExpression;
};


//////////////////////////////////////////////////// doubly Linked list with tail implementation :
class Node {
public:
    string key;
    Node *next;
    Node *prev;

    Node(string key);
};

Node *cursor = NULL;

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

    stringRes getString();
};

List::List() {
    head = NULL;
    tail = NULL;
    string1 = "";
}


stringRes List::getString() {
    Node *tmp = head;
    stringRes res;
    string prev = "";
    while (tmp != NULL) {
        if (((prev == ")" || isdigit(prev[0])) && tmp->key == "(") || (prev == ")" && isdigit(tmp->key[0]))) {
            res.infExpression += "*";
        }
        res.infExpression += tmp->key;
        res.tableExpression += tmp->key;
        prev = tmp->key;
        tmp = tmp->next;
    }
    return res;
}

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

void List::remove_from_head() {
    Node *temp = head;
    head = temp->next;
    head->prev = NULL;
    delete temp;
}

void List::remove_from_tail() {
    Node *temp = tail;
    tail = temp->prev;
    tail->next = NULL;
    delete temp;
}

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


void List::print_list() {
    bool printed = false;
    Node *temp = head;
    if (cursor == NULL) {
        cout << "|";
        printed = true;
    }
    while (temp != NULL) {
        //cout << "@";
        cout << temp->key;
        //cout << "@";
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

List infixToPostfix(string expression) {
    std::stack<string> st;
    st.push("N");
    List list2;
    bool isNegetive = false;
    string prev = "";
    int counter = 0;
    for (int i = 0; i < expression.length(); ++i) {
        if (isdigit(expression[i])) {
            string num;
            if (isNegetive) {
                num += "-";
                isNegetive = false;
            }
            num += expression[i];
            while (i < expression.length() && isdigit(expression[i + 1])) {
                // cout << "tss";
                num += expression[i + 1];
                i++;
            }
            if (num.length() >= 8) {
                std::string s = std::to_string(mode(num));
                num = s;
            }
            list2.insert_at_tail(num);
        } else if (expression[i] == '(') {
            st.push("(");
        } else if (expression[i] == ')') {
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
                if (expression[i] == '-') {
                    isNegetive = true;
                }
            } else {
                while (st.top() != "N" && prec(expression[i]) <= prec(st.top()[0])) {
                    string c = st.top();
                    st.pop();
                    list2.insert_at_tail(c);
                }
                std::string s{expression[i]};
                //   cout << s << endl;
                st.push(s);
            }
        }
        prev = expression[i];
    }
    // Pop all the remaining elements from the stack
    while (st.top() != "N") {
        string c = st.top();
        st.pop();
        list2.insert_at_tail(c);
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
            long long int val1 = stoll(st.top(), nullptr, 10);
            st.pop();
            long long int val2 = stoi(st.top(), nullptr, 10);
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
    unordered_map<string, long long int> table;
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
                break;
                /*case '@': {
                    for (auto itr = table.begin(); itr != table.end(); itr++) {
                        cout << (*itr).first << ": " << (*itr).second << endl;
                    }
                }
                    break;*/
            case '!': {
                stringRes s = list.getString();
                unordered_map<std::string, long long int>::const_iterator got = table.find(s.tableExpression);
                if (got != table.end()) {
                    //  cout << "from table :";
                    cout << got->second << endl;
                } else {
                    //  cout << s.tableExpression << endl;
                    // cout << s.infExpression << endl;
                    List post = infixToPostfix(s.infExpression);
                    //  post.print_list();
                    string res = evaluatePostfix(post);
                    long long int val = stoll(res, nullptr, 10);
                    while (val < 0) {
                        val += pow(10, 9) + 7;
                    }
                    cout << val << endl;
                    table.emplace(s.tableExpression, val);
                }
            }
                break;
        }
    }
    return 0;
}

/*
// hash table implementation which I did not use it :
const int T_S = 10;

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

    // hash function with O (n)
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
*/