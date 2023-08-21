# task 1:
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
    "A Y",
    "B X",
    "C Z",
]
guide_from_file = []

file = open("input.txt","r")
for lines in file:
    guide_from_file.append(lines)
file.close()


def task1(guide):
    condition = [
            "A X", "A Y", "A Z",
            "B X", "B Y", "B Z",
            "C X", "C Y", "C Z",
            ]
    drawed = [0,4,8]
    win = [6,1,5]
    score = 0

    for i in guide:
        condtion_satisfied = condition.index(i[:3])
        bonus = (condtion_satisfied % 3) + 1
        if condtion_satisfied in drawed:
            score += bonus + 3
            pass
        elif condtion_satisfied in win:
            score += 6 + bonus
            pass
        else:
            score += bonus

    print("Solution to task1:", score)
def task2(guide):
    condition = [
            "B X", "C X", "A X", #lose
            "A Y", "B Y", "C Y", #draw
            "C Z", "A Z", "B Z", #win
            ]
    score = 0

    for i in guide:
        condtion_satisfied = condition.index(i[:3])

        bonus = (condtion_satisfied % 3) + 1
        draw = 2 < condtion_satisfied <= 5
        win = 5 < condtion_satisfied <= 8

        if draw:
            score += bonus + 3
            pass
        elif win:
            score += 6 + bonus
            pass
        else:
            score += bonus

    print("Solution to task2:", score)

task1(guide_from_file)
task2(guide_from_file)
