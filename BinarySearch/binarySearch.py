import time

def binarySearch(arr: list, target: int) -> int:

    left = 0
    right = len(arr) - 1

    def AuxbinarySearch(left, right, arr, target, count):
        
        mid = (right-left)//2
        time.sleep(2)
        print(left, right, mid, arr, target, count, arr[mid])

        if arr[mid] == target:

            count += mid
            print('i enter ')
            return count
        
        elif arr[mid] < target:

            count +=mid
            print('mid < target')
            return AuxbinarySearch(mid, right, arr[mid:], target, count)
        else:
            print('mid > target')
            return AuxbinarySearch(left, left+mid, arr[:mid+1], target, count)

    a =  AuxbinarySearch(left, right, arr, target, count=0)

    return a, arr[a]
    

arr = [11, 22, 34, 45, 54, 66, 77, 88, 99, 101, 122, 134, 145, 156, 163, 174, 186, 198, 205]

print (binarySearch(arr, 66))

