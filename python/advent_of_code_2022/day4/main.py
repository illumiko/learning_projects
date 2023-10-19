import re
input = []
file = open("input.txt","r")
for lines in file:
    lines = lines[:-1].strip("-").split(",")
    input.append(lines)
file.close()

def task1(input:list):
    pair = 0
    for i in input:
        isContained = False
        pattern = r'(\d+)-(\d+)'
        match1 = re.search(pattern, i[0])
        match2 = re.search(pattern, i[1])

        if match1 and match2:
            range1 = - int(match1.group(1)) + int(match1.group(2))
            range2 = - int(match2.group(1)) + int(match2.group(2))

            # print(match1.group(1),"-",match1.group(2), range1, "|" ,match2.group(1),"-",match2.group(2),":", range2)

            if range2 > range1:
                isContained = int(match2.group(1)) < int(match1.group(1)) and int(match2.group(2)) >= int(match1.group(1))
            elif range2 < range1:
                isContained = int(match2.group(1)) > int(match1.group(1)) and int(match2.group(2)) < int(match1.group(2))
            elif range2 == range1:
                isContained = int(match2.group(1)) == int(match1.group(1)) and int(match2.group(2)) == int(match1.group(2))

            if isContained:
                pair+=1

    print(pair)

task1(input)
