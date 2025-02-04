def insertionSort(arr):
    n = len(arr)
    for i in range (1, n):
        current = arr[i]
        j = i-1
        while j >= 0 and current < arr[j]:
            arr[j+1] = arr[j]
            j -= 1
        arr[j+1] = current

    return arr


arr= [1,2,5,3,5,7,4,7,4,34,234,6,56,4,3,23,2,5,7,6,7,78,7,5,545,43,453,3,52,35,23,5]

print(insertionSort(arr))
