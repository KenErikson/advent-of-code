power_sum = 0
with open('day2Input.txt') as input:
    for line in input:
        game_id = line.split(':')[0].split(' ')[1]
        pull_tries = line.split(':')[1].split(';')
        min_blue = 0
        min_red = 0
        min_green = 0
        for pull_try in pull_tries:
            pull = pull_try.split(', ')
            for part in pull:
                count = int(part.strip().split(' ')[0])
                color = part.strip().split(' ')[1]
                match color:
                    case "green":
                        min_green = max(min_green,count)
                    case "blue":
                        min_blue = max(min_blue,count)
                    case "red": 
                        min_red = max(min_red,count)
        power_sum += min_blue * min_green * min_red             
        

print(power_sum)