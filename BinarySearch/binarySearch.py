import time

def binarySearch(arr: list, target: int) -> int:

    left = 0
    right = len(arr) - 1

    def AuxbinarySearch(left, right, arr, target):
        
        mid = (right+left)//2
        print(left, right, mid, arr, target, arr[mid])

        if arr[mid] == target:
            print('i enter ')
            return mid 
        
        elif arr[mid] < target:
            print('mid < target')
            return AuxbinarySearch(mid+1, right, arr, target)
        else:
            print('mid > target')
            return AuxbinarySearch(left, mid-1, arr, target)

    a =  AuxbinarySearch(left, right, arr, target)

    return a, arr[a]
    

arr = [11, 22, 34, 45, 54, 66, 77, 88, 99, 101, 122, 134, 145, 156, 163, 174, 186, 198, 205]

print (binarySearch(arr, 66))

