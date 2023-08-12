# A,X = rock = 1 point
# B,Y = paper = 2 point
# C,Z = Scissors 3 point
# Win = 6 point, Draw = 3 point, Lose = 0 point

strategy_guide = {
    "A":"Z", 
    "B":"X", 
    "C":"X", 
}

win_condtion = {
    "A":["C","Z"],
    "C":["B","Y"],
    "B":["A","X"],
}

draw_conditon = {
    "A":"X",
    "B":"Y",
    "C":"Z"
}

point_chart = { "A":1, "X":1, "B":2, "Y":2, "C":3, "Z":3, }

player_points = {
    "p1":{"Name":"John","points":0},
    "p2":{"Name":"Joey","points":0},
}

def play(p1,p2):
    if p2 in win_condtion[p1]:
        print(p1, "wins")
        player_points["p1"]['points'] += 6 + point_chart[p1]
        player_points["p2"]['points'] += 0 + point_chart[p2]
    elif p2 == draw_conditon[p1]:
        print("Draw")
        player_points["p1"]['points'] += 0 + point_chart[p1]
        player_points["p2"]['points'] += 0 + point_chart[p2]
    else:
        print(p2, "wins")
        player_points["p1"]['points'] += 0 + point_chart[p1]
        player_points["p2"]['points'] += 6 + point_chart[p2]

for p1,p2 in strategy_guide.items():
    play(p1,p2)


print(player_points)
