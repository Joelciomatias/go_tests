import datetime
start_time = datetime.datetime.now()
# start ------------------------

arr1 = [0,1,2,3,4,5,6,7,8,9]
arr2 = [9,8,7,6,5,4,3,2,1,0]
arr3 = []

for k in range(1000000):
    for i in range(10):
        arr3.append(arr2[i]^i)
        arr3.append((arr1[i]+arr2[i])*i)

arr3.sort()
print(len(arr3))

end_time = datetime.datetime.now()
# end ------------------------
print(end_time - start_time)