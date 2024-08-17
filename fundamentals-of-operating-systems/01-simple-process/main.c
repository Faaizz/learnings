#include <stdio.h>

int sum(int a, int b) {
    return a + b;
}

int main() {
    int a = 101;
    int b = 202;
    int c = sum(a, b);
    printf("a + b = %d", c);

    return 0;
}
