struct Node {
    int key;
    int value;
    struct Node *next;
};

int length();
void put(int key, int value);
int get(int key);
int* getAllValues();
void close();