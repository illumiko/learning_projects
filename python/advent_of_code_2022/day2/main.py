# A,X = rock = 1 point
# B,Y = paper = 2 point
# C,Z = Scissors 3 point
# Win = 6 point, Draw = 3 point, Lose = 0 point

# strategy_guide = {
#     "A":"Y", 
#     "B":"X", 
#     "C":"Z", 
# }
guide = [
    "B Y",
    "C X",
    "A Z",
]
guide_from_file = []

file = open("input.txt","r")
for lines in file:
    guide_from_file.append(lines)
file.close()


point_chart = { "A":1, "X":1, "B":2, "Y":2, "C":3, "Z":3, }

player_points = {
    "p1":{"Name":"John","points":0},
    "p2":{"Name":"YOU","points":0},
}

win_condtion = {
    "A":["C","Z"],
    "C":["B","Y"],
    "B":["A","X"],
}

draw_conditon = {
    "A":"X",
    "B":"Y",
    "C":"Z",
}

def play(p1,p2):
    if p2 in win_condtion[p1]:
        player_points["p1"]['points'] += 6 + point_chart[p1]
        player_points["p2"]['points'] += 0 + point_chart[p2]
    elif p2 == draw_conditon[p1]:
        player_points["p1"]['points'] += 3 + point_chart[p1]
        player_points["p2"]['points'] += 3 + point_chart[p2]
    else:
        player_points["p1"]['points'] += 0 + point_chart[p1]
        player_points["p2"]['points'] += 6 + point_chart[p2]

def task1(input:list):

    for i in input:
        p1_moves = i[0] 
        p2_moves = i[2]
        play(p1_moves,p2_moves)
    __import__('pprint').pprint(player_points)

def task2(input:list):
    p2 = ""
    for x in input:
        if x[2] == "Y":
            p2 = draw_conditon[x[0]]
            print("I draw")
        if x[2] == "X":
            p2 = win_condtion[x[0]][0]
            print("I lost")
        if x[2]=="Z":
            p2 = ("A" not in win_condtion[x[0]] and "A") or ("B" not in win_condtion[x[0]] and "B") or ("C" not in win_condtion[x[0]] and "C")
            print("I won")
        play(x[0], p2)


task2(guide_from_file)
print(player_points)
# Z is 588
# Y is 1723
# X is 189
