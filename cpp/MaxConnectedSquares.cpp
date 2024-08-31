#include <string>
#include <vector>
#include <set>
#include <iostream>

using namespace std;

//Because it explores
int dora(int square[], int size, set<int> &explored, int index) {
    if (index < 0 || index >= size * size) return 0;
    if (square[index] == 0) return 0;
    if (explored.find(index) != explored.end()) return 0;
    
    int branchSum = 1;
    explored.insert(index);
    if (index % size != 0) {
        branchSum += dora(square, size, explored, index - 1);
    }
    if ((index + 1) % size != 0) {
        branchSum += dora(square, size, explored, index + 1);
    }

    branchSum += dora(square, size, explored, index + size);
    branchSum += dora(square, size, explored, index - size);
    return branchSum;
}
int maximumConnectedArea(int square[], int size) {
    set<int> explored;
    int largestCluster = 0;
    for (int i = 0; i < size * size; i++) {
        int clusterSize = dora(square, size, explored, i);
        if (clusterSize > largestCluster) {
            largestCluster = clusterSize;
        }
    }
    return largestCluster;
}

int main() {
    int square[] = {
        1, 1, 1, 1,
        0, 0, 1, 1,
        0, 1, 0, 1,
        0, 0, 0, 1
    };
    int size = 4;
    cout << maximumConnectedArea(square, size) << endl;
    return 0;
}
