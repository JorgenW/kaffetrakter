from threading import *

i=0
lokk = Lock()
def adder1():

	global i

	for x in range(1000000):
		lokk.acquire(True)
		i=i+1
		#print("1")
		lokk.release()
def adder2():
	global i
	for x in range(1000000):
		lokk.acquire(True)
		i=i-1
		#print("2")
		lokk.release()

def main():
	adder_thr= Thread(target = adder1)
	adder_thr.start()
	adder_thr2=Thread(target=adder2)
	adder_thr2.start()
	adder_thr.join()
	adder_thr2.join()
	print("Done: " + str(i))
  
main()
