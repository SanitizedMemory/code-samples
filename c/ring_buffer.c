#include <stdio.h>
#include <stdlib.h>

#define BUFFER_SIZE 10

typedef struct Rbuffer {
    int size;
    int start;
    int end;
    char *buffer;
} Rbuffer;

void read(Rbuffer *buffer, int num_of_values) {
    int start = buffer->start;
    int size = buffer->size;
    for (int _ = 0; _ < num_of_values; _++) {
        printf("%c", buffer->buffer[start]);
        start = (start + 1) % size;
    }
    printf("\n");
    buffer->start = start;
}

void write(Rbuffer *buffer, char *data) {
    int data_idx = 0;
    while (data[data_idx] != '\0') {
        data_idx++;
    }

    int i = 0;
    int write = buffer->end;
    while (data[i] != '\0') {
        buffer->buffer[write] = data[i];
        i++;
        write = (write + 1) % buffer->size;
    }
    buffer->end = write;
}

int main() {
    Rbuffer ring_buffer = {
        .size = BUFFER_SIZE,
        .start = 0,
        .end = 0,
        .buffer = (char *)malloc(sizeof(char) * BUFFER_SIZE),
    };

    char *data = "HelloWorld";
    write(&ring_buffer, data);
    read(&ring_buffer, 5);
    read(&ring_buffer, 4);
    read(&ring_buffer, 7);
    return 0;
}
