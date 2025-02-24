def Quicksort(array):

    if len(array)< 2:
        return array
    
    else: 
        pivot = array[0]
        less = [i for i in array[1:] if i <= pivot]
        grater = [i for i in array[1:] if i > pivot]

        return Quicksort(less) + [pivot] + Quicksort(grater)
    

x = [1,2,5,3,5,7,4,7,4,34,234,6,56,4,3,23,2,5,7,6,7,78,7,5,545,43,453,3,52,35,23,5]  
print(Quicksort(x))    