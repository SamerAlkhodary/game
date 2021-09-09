import json
import numpy as np

file = open('data.json')
data = json.load(file)
file.close()

width = data['width']
height = data['height']
n = data['n'][0]
location_gateway= (height//2,width//2)
locations = set()
while len(locations)< n:
    x = np.random.randint(height)
    y = np.random.randint(width)
    if location_gateway != (x,y):
        locations.add((x,y))

f = open("config.txt", "w")
f.write(str(width)+','+str(height)+'\n')

f.write(str(data['Tp'])+'\n')
f.write(str(data['Ts'])+'\n')
f.write(str(data['r'])+'\n')

f.write(str(location_gateway[0])+','+str(location_gateway[1])+'\n')
for elem in locations:
    f.write(str(elem[0])+','+str(elem[1])+'\n')
f.close()


