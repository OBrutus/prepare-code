'''
Author  :   AV 
At      :   @Home 
Date    :   21.5.
ver.    :   0.0 
link    :   
file    :   cf/
'''

# import numpy as np
import sys
import math
import bisect
from sys import stdin,stdout
from math import gcd,floor,sqrt,log,ceil
from collections import defaultdict as dd
from bisect import bisect_left as bl,bisect_right as br

sys.setrecursionlimit(100000000)

ceildiv=lambda x,d: x//d if(x%d==0) else x//d+1

flush  =lambda: stdout.flush()
stdstr =lambda: stdin.readline()
stdint =lambda: int(stdin.readline())
stdpr  =lambda x: stdout.write(str(x))

MOD = 1000000007 # 10**9 + 7

def code(a, n):
	res = None
	return res

#main code
if __name__ == '__main__':
	T = 1
	T = stdint()
	output = ''
	for test_case in range(1, T+1):
		# n = stdint()
		# s = stdstr()
		# a = list(map(int, stdstr().split()))
		# res = code(a, n)
		# print('Case #{}: {}'.format(test_case, res))
		# output += str(res) + '\n'
	print(output, end = '')

