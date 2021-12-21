#include "linkedlist.h"
#include <stdlib.h>

struct Node *head = NULL;
struct Node *current = NULL;

int length() {
     int count = 0;
     struct Node* node = head;
     while (node != NULL) {
         node = node -> next;
         count = count + 1;
     }
     return count;
}

void put(int key, int value) {
    if (head == NULL) {
        head = (struct Node*)malloc(sizeof(struct Node));
        head -> key = key;
        head -> value = value;
        current = head;
    }
    else {
        struct Node *next = (struct Node*)malloc(sizeof(struct Node));
        next -> key = key;
        next -> value = value;
        current -> next = next;
        current = next;
    }
}

int get(int key) {
     struct Node* node = head;
     while (node->key != key && node != NULL) {
         node = node -> next;
     }
     if (node == NULL) {
        return -1;
     }
     return node -> value;
}

int* getAllValues() {
     int len = length();
     int* values = malloc(len * sizeof(int));
     struct Node* node = head;
     int index = 0;

     while (node != NULL) {
         values[index] = node -> value;
         node = node -> next;
         index = index + 1;
     }
     return values;
}

void close() {
  struct Node* node;
  while (head != NULL) {
      node = head;
      head = head->next;
      free(node);
  }
}