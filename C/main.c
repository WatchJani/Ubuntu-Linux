#include <stdio.h>
#include <stdlib.h>
#include <time.h>

size_t readFile(char *path, char *buffer, size_t sizeOfBuffer)
{
    FILE *fptr = fopen(path, "r");

    if (fptr == NULL)
    {
        perror("Error opening file");
        return 0;
    }

    size_t read = fread(buffer, 1, sizeOfBuffer, fptr);

    fclose(fptr);

    return read;
}

int main()
{
    size_t sizeOfBuffer = 4096;
    char *buffer = (char *)malloc(sizeOfBuffer);

    if (buffer == NULL)
    {
        perror("Buffer allocation error");
        return 1;
    }

    int numberOfIteration = 184000;
    clock_t start = clock();

    for (int i = 0; i < numberOfIteration; i++)
    {
        size_t read = readFile("AAAHNCYNCP.txt", buffer, sizeOfBuffer);
    }

    clock_t end = clock();

    double elapsedTime = ((double)(end - start)) / CLOCKS_PER_SEC;

    printf("Function is executed %d times in %.6f second.\n", numberOfIteration, elapsedTime);

    free(buffer);

    return 0;
}
