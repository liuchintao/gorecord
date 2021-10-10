import random
import time

name = "".join(random.choices("abcdefghijklmnopqrstuvwxyz", k=8))
with open("gen-"+name, "w+") as f:
  f.write("hello")
  time.sleep(5)
  f.write("world")
  f.close()