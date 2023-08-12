import csv

inventory_from_file = dict()
with open("input.csv") as csvfile:
    spamreader = csv.reader(csvfile, delimiter="\n")
    i = 0
    id = ""
    for row in spamreader:
        if row == []:
            i+=1
            id = "elf"+str(i)
            inventory_from_file[id] = []
        elif row != []:
            inventory_from_file[id].append(''.join(row))


def get_total(list_calories:list) -> int: 
    current = 0
    for i in list_calories:
        current += int(i)

    return current

def find_elf(data, hook):
    for elf,calorie in data.items():
        if hook == get_total(calorie):
            return [elf, get_total(calorie)]

def get_elf_with_most_calories(data:dict):
    temp = []
    for elf in data:
        temp.append(get_total(data[elf]))

    top_3 = []
    for i in range(3):
        elf_with_highest_calorie = max(temp)
        found = find_elf(data,elf_with_highest_calorie)
        top_3.append(found)
        temp.remove(elf_with_highest_calorie)

    sum = 0
    for i in top_3:
        sum += i[1]

    return sum




print(get_elf_with_most_calories(inventory_from_file))
