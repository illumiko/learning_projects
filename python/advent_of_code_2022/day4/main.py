input = []
file = open("input.txt","r")
for lines in file:
    lines = lines[:-1].strip("-").split(",")
    input.append(lines)
file.close()

def task1(input:list):
    for i in input:
        seperatorIndex1 = round((len(i[0])-1)/2)
        seperatorIndex2 = round((len(i[1])-1)/2)


task1(input)
