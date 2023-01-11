import time

list1=(12,12,11,11,10,10,9,9,8,8,7,7,9,8,8,7,7,6,6,5,5,7,7,6,6,5,5)
list2=(12,11,11,10,10,9,9,8,8,11,11,10,7,7,6,6,5,5,8,8,7,4,3,3,2,2,1)

time.sleep(1)
print('你个水性杨花、\n')
time.sleep(1)
print('两面三刀的家伙\n\n')
time.sleep(1)
print('我受够你了!\n\n')
time.sleep(3)
for j in range(3):
    print('\n\n你就活吧！')
    time.sleep(1)

for i in range(0,27):
    n1=list1[i]
    n2=list2[i]
    print(" "*(int(n1)),end="")
    print("*"*(int(n2)))
time.sleep(1000)
