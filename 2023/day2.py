
red_count = 12
green_count = 13
blue_count = 14

map_count = {"red":red_count,"green":green_count, "blue": blue_count}

id_sum = 0
with open('day2Input.txt') as input:
    for line in input:
        game_id = line.split(':')[0].split(' ')[1]
        pull_tries = line.split(':')[1].split(';')
        bool_valid = True
        for pull_try in pull_tries:
            pull = pull_try.split(', ')
            for part in pull:
                count = int(part.strip().split(' ')[0])
                color = part.strip().split(' ')[1]
                if count > map_count[color]:
                    bool_valid = False
                    break
        if bool_valid:
            id_sum += int(game_id)

print(id_sum)