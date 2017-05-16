import os

for name in os.listdir("./Content"):
	f = open("./Content/" + name, 'r')
	f1 = open("./Content1/" + name + ".html", 'w')
	f1.write(f.read())
	f1.close()
	f.close()
