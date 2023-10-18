alphabets = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
list = []

file = open("input.txt","r")
for lines in file:
    list.append(lines)
file.close()

def processTask1(i):
    score = 0
    commonItems = []
    for i in i:
        divider = round(len(i[:-1])/2)
        compartment1 = i[:divider]
        compartment2 = i[divider:]
        y =  [i for i in compartment1 if i in compartment2][0]
        commonItems.append(y)
    
    for i in commonItems:
        score += alphabets.index(i) + 1

    print("Solution to task1 = ", score)

def processTask2(list):
    badge = []
    priorityNo = 0
    for i in range(0,len(list),3):
        a = list[i]
        b = list[i+1]
        c = list[i+2]
        y = [i for i in a if i in b if i in c][0]
        badge.append(y)

    for i in badge:
        priorityNo += alphabets.index(i) + 1

    print("Solution to task2 = ", priorityNo)

processTask1(list)
processTask2(list)



