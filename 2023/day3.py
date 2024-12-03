result = 0

nums = {"0","1","2","3","4","5","6","7","8","9"}

non_sign = {"0","1","2","3","4","5","6","7","8","9", ".", '\n'}


inputs = []

line_length = 0
with open('day3Input.txt') as input:
    for line in input:
        inputs.append(line)
        line_length = len(line)


for i,line in enumerate(inputs):
    #print(f"line: {i}")
    #print(inputs[max(0,i-1)])
    #print(line)
    #print(inputs[min(line_length-1,i+1)])
    potential_numbers_start_indices = []
    potential_numbers_stop_indices = []
    inside_number = False
    for j in range(len(line)):
        if inside_number and line[j] in nums:
            continue
        elif line[j] in nums:
            potential_numbers_start_indices.append(j)
            inside_number = True
        elif inside_number:
            inside_number = False
            potential_numbers_stop_indices.append(j - 1)

    if len(potential_numbers_start_indices) > len(potential_numbers_stop_indices):
        potential_numbers_stop_indices.append(line_length - 1)
        print("adding end")
    if len(potential_numbers_start_indices) != len(potential_numbers_stop_indices):
        raise Exception("Exception")

    for j in range(len(potential_numbers_start_indices)):
        number = int(line[potential_numbers_start_indices[j]:potential_numbers_stop_indices[j]+1])
        if number == 76:
            print(number)

        min_x = max(0,potential_numbers_start_indices[j] - 1)
        max_x = min(line_length,potential_numbers_stop_indices[j] + 1)
        min_i = max(0, i - 1)
        max_i = min(len(inputs)-1, i + 1)
        added = False
        for xi in range(min_i,max_i+1):
            if added:
                 break
            for x in range(min_x,max_x+1):
                if inputs[xi][x] not in non_sign:
                    added = True
                    result += number
                    if number < 1:
                        raise Exception("test")
                    break
        if added and number == 76:
            print("YES")
        elif number == 76:
            print("NO")


print(result)